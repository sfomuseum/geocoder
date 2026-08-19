package sql

import (
	"context"
	"fmt"
	"sync"

	sfom_sql "github.com/sfomuseum/go-database/sql"
	"github.com/sfomuseum/geocoder/x/vec"
)

var sqliteTables = sync.OnceValues(func() (map[string]sfom_sql.Table, error) {

	ctx := context.Background()

	records_table, err := NewRecordsTable(ctx)

	if err != nil {
		return nil, fmt.Errorf("Failed to instantiate records table, %w", err)
	}

	tokens_table, err := NewTokensTable(ctx)

	if err != nil {
		return nil, fmt.Errorf("Failed to instantiate tokens table, %w", err)
	}

	dates_table, err := NewDatesTable(ctx)

	if err != nil {
		return nil, fmt.Errorf("Failed to instantiate dates table, %w", err)
	}

	ancestors_table, err := NewAncestorsTable(ctx)

	if err != nil {
		return nil, fmt.Errorf("Failed to instantiate ancestors table, %w", err)
	}

	placetypes_alt_table, err := NewPlacetypesAltTable(ctx)

	if err != nil {
		return nil, fmt.Errorf("Failed to instantiate placetypes table, %w", err)
	}

	bounds_table, err := NewBoundsTable(ctx)

	if err != nil {
		return nil, fmt.Errorf("Failed to instantiate bounds table, %w", err)
	}

	emb_opts, err := DefaultEmbeddingsTableOptions()

	if err != nil {
		return nil, fmt.Errorf("Failed to create embeddings table options, %w", err)
	}

	// To do... pull this in dynamically from... wut?
	emb_opts.Dimensions = vec.DEFAULT_EMBEDDINGS_DIMENSIONS

	embeddings_table, err := NewEmbeddingsTableWithOptions(ctx, emb_opts)

	if err != nil {
		return nil, fmt.Errorf("Failed to instantiate embeddings table, %w", err)
	}

	embeddings_records_table, err := NewEmbeddingsRecordsTable(ctx)

	if err != nil {
		return nil, fmt.Errorf("Failed to instantiate embeddings records table, %w", err)
	}

	db_tables := map[string]sfom_sql.Table{
		"records":            records_table,
		"tokens":             tokens_table,
		"dates":              dates_table,
		"ancestors":          ancestors_table,
		"placetypes_alt":     placetypes_alt_table,
		"bounds":             bounds_table,
		"embeddings":         embeddings_table,
		"embeddings_records": embeddings_records_table,
	}

	return db_tables, nil
})

// SQLiteTables returns a slice of Table objects that describe
// all tables used by the SQLite backend.  The slice can be
// passed to the database configuration code to create the
// required schema on demand.
func SQLiteTables(ctx context.Context) (map[string]sfom_sql.Table, error) {
	return sqliteTables()
}
