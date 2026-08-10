package parquet

import (
	"bytes"
	"context"
	"fmt"
	"iter"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"sync/atomic"

	parquet_go "github.com/parquet-go/parquet-go"
	"github.com/whosonfirst/go-ioutil"
	"github.com/whosonfirst/go-whosonfirst/v4/iterate"
	"github.com/whosonfirst/go-whosonfirst/v4/iterate/filters"
	parquet_wof "github.com/whosonfirst/go-whosonfirst/v4/parquet"
)

func init() {
	ctx := context.Background()
	iterate.RegisterIterator(ctx, "parquet", NewParquetIterator)
}

// ParquetIterator implements the `Iterator` interface for crawling records in Parquet files produced
// by the `whosonfirst/go-whosonfirst/v4/parquet.ParquetWriter` methods.
type ParquetIterator struct {
	iterate.Iterator
	// The count of documents that have been processed so far.
	seen int64
	// Boolean value indicating whether records are still being iterated.
	iterating *atomic.Bool
	// filters is a `filters.Filters` instance used to include or exclude specific records from being crawled.
	filters filters.Filters
}

// NewParquetIterator() returns a new `ParquetIterator` instance configured by 'uri' in the form of:
//
//	parquet://
//
// Where {PARAMETERS} may be:
// * `?include=` Zero or more `aaronland/go-json-query` query strings containing rules that must match for a document to be considered for further processing.
// * `?exclude=` Zero or more `aaronland/go-json-query`	query strings containing rules that if matched will prevent a document from being considered for further processing.
// * `?include_mode=` A valid `aaronland/go-json-query` query mode string for testing inclusion rules.
// * `?exclude_mode=` A valid `aaronland/go-json-query` query mode string for testing exclusion rules.
func NewParquetIterator(ctx context.Context, uri string) (iterate.Iterator, error) {

	f, err := filters.NewQueryFiltersFromURI(ctx, uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to create filters from query, %w", err)
	}

	it := &ParquetIterator{
		filters:   f,
		seen:      int64(0),
		iterating: new(atomic.Bool),
	}

	return it, nil
}

// Iterate will return an `iter.Seq2[*Record, error]` for each record encountered in 'uris'.
func (it *ParquetIterator) Iterate(ctx context.Context, uris ...string) iter.Seq2[*iterate.Record, error] {

	return func(yield func(rec *iterate.Record, err error) bool) {

		it.iterating.Swap(true)
		defer it.iterating.Swap(false)

		for _, uri := range uris {

			logger := slog.Default()
			logger = logger.With("uri", uri)

			logger.Debug("Read records from URI")

			u, err := url.Parse(uri)

			if err != nil {

				logger.Error("Failed to parse URI", "error", err)

				if !yield(nil, fmt.Errorf("Failed to parse URI '%s', %w", uri, err)) {
					return
				}

				continue
			}

			var r ReadCloserAt
			var sz int64

			switch u.Scheme {
			case "http", "https":

				rsp, err := http.Get(uri)

				if err != nil {

					logger.Error("Failed to retrieve URI", "error", err)

					if !yield(nil, fmt.Errorf("Failed to retrieve %s, %w", uri, err)) {
						return
					}

					continue
				}

				r = NewCachedReaderAt(rsp.Body)
				sz = rsp.ContentLength

			default:

				f, err := os.Open(uri)

				if err != nil {

					logger.Error("Failed to open URI for reading", "error", err)

					if !yield(nil, fmt.Errorf("Failed to open %s for reading, %w", uri, err)) {
						return
					}

					continue
				}

				info, err := f.Stat()

				if err != nil {

					logger.Error("Failed to stat URI", "error", err)
					f.Close()

					if !yield(nil, fmt.Errorf("Failed to stat %s, %w", uri, err)) {
						return
					}

					continue
				}

				r = f
				sz = info.Size()
			}

			rows, err := parquet_go.Read[*parquet_wof.Record](r, sz)

			if err != nil {

				logger.Error("Failed to create Parquet reader", "error", err)
				r.Close()

				if !yield(nil, fmt.Errorf("Failed to create Parquet reader for %s, %w", uri, err)) {
					return
				}

				continue
			}

			for i, p_rec := range rows {

				path := fmt.Sprintf("%s#%d", uri, i)
				atomic.AddInt64(&it.seen, 1)

				body, err := p_rec.AsGeoJSONBytes()

				if err != nil {

					logger.Error("Failed to derive GeoJSON from record", "record", i, "error", err)
					r.Close()

					if !yield(nil, fmt.Errorf("Failed to derive geojson from record for %s, %w", path, i, err)) {
						return
					}

					continue
				}

				br := bytes.NewReader(body)
				rsc, err := ioutil.NewReadSeekCloser(br)

				if err != nil {

					logger.Error("Failed to create ReadSeekCloser for record", "record", i, "error", err)
					r.Close()

					if !yield(nil, fmt.Errorf("Failed to create ReadSeekCloser for %s, %w", path, i, err)) {
						return
					}

					continue
				}

				if it.filters != nil {

					ok, err := iterate.ApplyFilters(ctx, rsc, it.filters)

					if err != nil {
						rsc.Close()
						if !yield(nil, fmt.Errorf("Failed to apply filters for '%s', %w", path, err)) {
							return
						}

						continue
					}

					if !ok {
						rsc.Close()
						continue
					}

				}

				rec := &iterate.Record{
					Path: path,
					Body: rsc,
				}

				if !yield(rec, nil) {
					r.Close()
					return
				}
			}

			r.Close()
			logger.Debug("Finished processing URI", "count seen (total)", it.Seen())
		}
	}
}

// Seen() returns the total number of records processed so far.
func (it *ParquetIterator) Seen() int64 {
	return atomic.LoadInt64(&it.seen)
}

// IsIterating() returns a boolean value indicating whether 'it' is still processing documents.
func (it *ParquetIterator) IsIterating() bool {
	return it.iterating.Load()
}

// Close performs any implementation specific tasks before terminating the iterator.
func (it *ParquetIterator) Close() error {
	return nil
}
