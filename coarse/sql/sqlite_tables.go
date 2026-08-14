package sql

import (
	"context"
	"fmt"

	sfom_sql "github.com/sfomuseum/go-database/sql"
)

// SQLiteTables returns a slice of Table objects that describe
// all tables used by the SQLite backend.  The slice can be
// passed to the database configuration code to create the
// required schema on demand.
func SQLiteTables(ctx context.Context) ([]sfom_sql.Table, error) {

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

	embeddings_table, err := NewEmbeddingsTable(ctx)

	if err != nil {
		return nil, fmt.Errorf("Failed to instantiate embeddings table, %w", err)
	}

	embeddings_records_table, err := NewEmbeddingsRecordsTable(ctx)

	if err != nil {
		return nil, fmt.Errorf("Failed to instantiate embeddings records table, %w", err)
	}

	db_tables := []sfom_sql.Table{
		records_table,
		tokens_table,
		dates_table,
		ancestors_table,
		placetypes_alt_table,
		bounds_table,
		embeddings_table,
		embeddings_records_table,
	}

	return db_tables, nil
}
