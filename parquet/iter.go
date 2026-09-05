package parquet

import (
	"context"
	"fmt"
	"iter"
	"log/slog"
	"net/http"
	"net/url"
	"os"

	parquet_go "github.com/parquet-go/parquet-go"
)

// Iterate returns a sequence that yields pointers to records of type `T` read from the
// specified URIs.  The iterator is generic over the record type.
func Iterate[T any](ctx context.Context, uris ...string) iter.Seq2[*T, error] {

	return func(yield func(*T, error) bool) {

		for _, uri := range uris {
			logger := slog.Default()
			logger = logger.With("uri", uri)

			logger.Debug("Read records from URI")

			u, err := url.Parse(uri)
			if err != nil {
				logger.Error("Failed to parse URI", "error", err)
				yield(nil, fmt.Errorf("Failed to parse URI '%s', %w", uri, err))
				return
			}

			var r ReadCloserAt
			var sz int64

			switch u.Scheme {
			case "http", "https":
				rsp, err := http.Get(uri)
				if err != nil {
					logger.Error("Failed to retrieve URI", "error", err)
					yield(nil, fmt.Errorf("Failed to retrieve %s, %w", uri, err))
					return
				}
				r = NewCachedReaderAt(rsp.Body)
				sz = rsp.ContentLength

			default:
				f, err := os.Open(uri)
				if err != nil {
					logger.Error("Failed to open URI for reading", "error", err)
					yield(nil, fmt.Errorf("Failed to open %s for reading, %w", uri, err))
					return
				}
				info, err := f.Stat()
				if err != nil {
					logger.Error("Failed to stat URI", "error", err)
					f.Close()
					yield(nil, fmt.Errorf("Failed to stat %s, %w", uri, err))
					return
				}
				r = f
				sz = info.Size()
			}

			rows, err := parquet_go.Read[*T](r, sz)
			if err != nil {
				logger.Error("Failed to create Parquet reader", "error", err)
				r.Close()
				yield(nil, fmt.Errorf("Failed to create Parquet reader for %s, %w", uri, err))
				return
			}

			for _, rec := range rows {
				if !yield(rec, nil) {
					r.Close()
					return
				}
			}
			r.Close()
		}
	}
}
