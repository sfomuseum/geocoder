package sql

import (
	"context"
	db_sql "database/sql"

	sfom_sql "github.com/sfomuseum/go-database/sql"
)

const PLACETYPES_ALT_TABLE_NAME string = "placetypes_alt"

type PlacetypesAltTableOptions struct{}

func DefaultPlacetypesAltTableOptions() (*PlacetypesAltTableOptions, error) {

	opts := PlacetypesAltTableOptions{}

	return &opts, nil
}

type PlacetypesAltTable struct {
	sfom_sql.Table
	name    string
	options *PlacetypesAltTableOptions
}

func NewPlacetypesAltTable(ctx context.Context) (sfom_sql.Table, error) {

	opts, err := DefaultPlacetypesAltTableOptions()

	if err != nil {
		return nil, err
	}

	return NewPlacetypesAltTableWithOptions(ctx, opts)
}

func NewPlacetypesAltTableWithOptions(ctx context.Context, opts *PlacetypesAltTableOptions) (sfom_sql.Table, error) {

	t := PlacetypesAltTable{
		name:    PLACETYPES_ALT_TABLE_NAME,
		options: opts,
	}

	return &t, nil
}

func NewPlacetypesAltTableWithDatabase(ctx context.Context, db *db_sql.DB) (sfom_sql.Table, error) {

	opts, err := DefaultPlacetypesAltTableOptions()

	if err != nil {
		return nil, err
	}

	return NewPlacetypesAltTableWithDatabaseAndOptions(ctx, db, opts)
}

func NewPlacetypesAltTableWithDatabaseAndOptions(ctx context.Context, db *db_sql.DB, opts *PlacetypesAltTableOptions) (sfom_sql.Table, error) {

	t, err := NewPlacetypesAltTableWithOptions(ctx, opts)

	if err != nil {
		return nil, err
	}

	err = t.InitializeTable(ctx, db)

	if err != nil {
		return nil, err
	}

	return t, nil
}

func (t *PlacetypesAltTable) InitializeTable(ctx context.Context, db *db_sql.DB) error {
	return sfom_sql.CreateTableIfNecessary(ctx, db, t)
}

func (t *PlacetypesAltTable) Name() string {
	return t.name
}

func (t *PlacetypesAltTable) Schema(db *db_sql.DB) (string, error) {
	return sfom_sql.LoadSchema(db, FS, PLACETYPES_ALT_TABLE_NAME)
}

func (t *PlacetypesAltTable) IndexRecord(ctx context.Context, db *db_sql.DB, tx *db_sql.Tx, i any) error {
	return NotSupportedError
}
