package sql

import (
	"context"
	db_sql "database/sql"

	sfom_sql "github.com/sfomuseum/go-database/sql"
)

const ANCESTORS_TABLE_NAME string = "ancestors"

type AncestorsTableOptions struct{}

func DefaultAncestorsTableOptions() (*AncestorsTableOptions, error) {

	opts := AncestorsTableOptions{}

	return &opts, nil
}

type AncestorsTable struct {
	sfom_sql.Table
	name    string
	options *AncestorsTableOptions
}

func NewAncestorsTable(ctx context.Context) (sfom_sql.Table, error) {

	opts, err := DefaultAncestorsTableOptions()

	if err != nil {
		return nil, err
	}

	return NewAncestorsTableWithOptions(ctx, opts)
}

func NewAncestorsTableWithOptions(ctx context.Context, opts *AncestorsTableOptions) (sfom_sql.Table, error) {

	t := AncestorsTable{
		name:    ANCESTORS_TABLE_NAME,
		options: opts,
	}

	return &t, nil
}

func NewAncestorsTableWithDatabase(ctx context.Context, db *db_sql.DB) (sfom_sql.Table, error) {

	opts, err := DefaultAncestorsTableOptions()

	if err != nil {
		return nil, err
	}

	return NewAncestorsTableWithDatabaseAndOptions(ctx, db, opts)
}

func NewAncestorsTableWithDatabaseAndOptions(ctx context.Context, db *db_sql.DB, opts *AncestorsTableOptions) (sfom_sql.Table, error) {

	t, err := NewAncestorsTableWithOptions(ctx, opts)

	if err != nil {
		return nil, err
	}

	err = t.InitializeTable(ctx, db)

	if err != nil {
		return nil, err
	}

	return t, nil
}

func (t *AncestorsTable) InitializeTable(ctx context.Context, db *db_sql.DB) error {
	return sfom_sql.CreateTableIfNecessary(ctx, db, t)
}

func (t *AncestorsTable) Name() string {
	return t.name
}

func (t *AncestorsTable) Schema(db *db_sql.DB) (string, error) {
	return sfom_sql.LoadSchema(db, FS, ANCESTORS_TABLE_NAME)
}

func (t *AncestorsTable) IndexRecord(ctx context.Context, db *db_sql.DB, tx *db_sql.Tx, i any) error {
	return NotSupportedError
}
