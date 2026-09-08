package parquet

import (
	"context"
	"fmt"
	"iter"
	"log/slog"

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

			r, sz, err := OpenURI(uri)

			if err != nil {
				logger.Error("Failed to parse URI", "error", err)
				yield(nil, fmt.Errorf("Failed to parse URI '%s', %w", uri, err))
				return
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
