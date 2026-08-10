package sql

import (
	"context"
	db_sql "database/sql"

	sfom_sql "github.com/sfomuseum/go-database/sql"
)

const BOUNDS_TABLE_NAME string = "bounds"

type BoundsTableOptions struct{}

func DefaultBoundsTableOptions() (*BoundsTableOptions, error) {

	opts := BoundsTableOptions{}
	return &opts, nil
}

type BoundsTable struct {
	sfom_sql.Table
	name    string
	options *BoundsTableOptions
}

func NewBoundsTable(ctx context.Context) (sfom_sql.Table, error) {

	opts, err := DefaultBoundsTableOptions()

	if err != nil {
		return nil, err
	}

	return NewBoundsTableWithOptions(ctx, opts)
}

func NewBoundsTableWithOptions(ctx context.Context, opts *BoundsTableOptions) (sfom_sql.Table, error) {

	t := BoundsTable{
		name:    BOUNDS_TABLE_NAME,
		options: opts,
	}

	return &t, nil
}

func NewBoundsTableWithDatabase(ctx context.Context, db *db_sql.DB) (sfom_sql.Table, error) {

	opts, err := DefaultBoundsTableOptions()

	if err != nil {
		return nil, err
	}

	return NewBoundsTableWithDatabaseAndOptions(ctx, db, opts)
}

func NewBoundsTableWithDatabaseAndOptions(ctx context.Context, db *db_sql.DB, opts *BoundsTableOptions) (sfom_sql.Table, error) {

	t, err := NewBoundsTableWithOptions(ctx, opts)

	if err != nil {
		return nil, err
	}

	err = t.InitializeTable(ctx, db)

	if err != nil {
		return nil, err
	}

	return t, nil
}

func (t *BoundsTable) Name() string {
	return t.name
}

func (t *BoundsTable) Schema(db *db_sql.DB) (string, error) {
	return sfom_sql.LoadSchema(db, FS, BOUNDS_TABLE_NAME)
}

func (t *BoundsTable) InitializeTable(ctx context.Context, db *db_sql.DB) error {
	return sfom_sql.CreateTableIfNecessary(ctx, db, t)
}

func (t *BoundsTable) IndexRecord(ctx context.Context, db *db_sql.DB, tx *db_sql.Tx, i any) error {
	return NotSupportedError
}
