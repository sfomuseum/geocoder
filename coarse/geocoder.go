package coarse

import (
	"context"

	"github.com/aaronland/go-pagination"
	"github.com/paulmach/orb/geojson"
)

// Geocoder defines the interface used by the command‑line utilities
// and HTTP server to interact with a coarse geocoding database.
type Geocoder interface {
	// PreIndex runs any pre‑indexing logic that must happen before records are added.
	// For example, the SQLite implementation this method to drop existing indices and
	// the full‑text search (FTS) table so that bulk inserts can be performed more quickly.
	PreIndex(context.Context) error
	// PostIndex runs any post‑indexing logic that must happen after all records have been added.
	// For example the SQLite implementation uses this method to recreate the indices, deduplicates
	// rows in the auxiliary tables and rebuilds the FTS table.
	PostIndex(context.Context) error
	// AddRecord queues a record for bulk insertion into the database.
	// The record will be written when the internal buffer reaches its
	// batch size or when Flush is called.
	AddRecord(context.Context, *Record) error
	// RemoveRecord deletes a record and all of its dependent rows
	// from the database.
	RemoveRecord(context.Context, int64) error
	// RecordExists reports whether a record with the supplied ID is present in the database.
	RecordExists(context.Context, int64) (bool, error)
	// HasRecordHashChanged checks whether a record already exists  and, if so, whether its
	// content hash has changed. It returns (exists, changed, error).
	HasRecordHashChanged(context.Context, *Record) (bool, bool, error)
	// Query performs a coarse geocoding query using the supplied
	// QueryRequest and pagination options.  It returns a slice of
	// GeoJSON Features and a pagination.Results object.
	Query(context.Context, *QueryRequest, pagination.Options) ([]*geojson.Feature, pagination.Results, error)
	// Flush forces any buffered records to be written to the database.
	Flush(context.Context) error
	// Close closes the underlying database connection.
	Close() error
}
