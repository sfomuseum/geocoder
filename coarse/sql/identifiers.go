package sql

import (
	"context"
	db_sql "database/sql"

	sfom_sql "github.com/sfomuseum/go-database/sql"
)

const IDENTIFIERS_TABLE_NAME string = "identifiers"

type IdentifiersTableOptions struct{}

func DefaultIdentifiersTableOptions() (*IdentifiersTableOptions, error) {

	opts := IdentifiersTableOptions{}

	return &opts, nil
}

type IdentifiersTable struct {
	sfom_sql.Table
	name    string
	options *IdentifiersTableOptions
}

func NewIdentifiersTable(ctx context.Context) (sfom_sql.Table, error) {

	opts, err := DefaultIdentifiersTableOptions()

	if err != nil {
		return nil, err
	}

	return NewIdentifiersTableWithOptions(ctx, opts)
}

func NewIdentifiersTableWithOptions(ctx context.Context, opts *IdentifiersTableOptions) (sfom_sql.Table, error) {

	t := IdentifiersTable{
		name:    IDENTIFIERS_TABLE_NAME,
		options: opts,
	}

	return &t, nil
}

func NewIdentifiersTableWithDatabase(ctx context.Context, db *db_sql.DB) (sfom_sql.Table, error) {

	opts, err := DefaultIdentifiersTableOptions()

	if err != nil {
		return nil, err
	}

	return NewIdentifiersTableWithDatabaseAndOptions(ctx, db, opts)
}

func NewIdentifiersTableWithDatabaseAndOptions(ctx context.Context, db *db_sql.DB, opts *IdentifiersTableOptions) (sfom_sql.Table, error) {

	t, err := NewIdentifiersTableWithOptions(ctx, opts)

	if err != nil {
		return nil, err
	}

	err = t.InitializeTable(ctx, db)

	if err != nil {
		return nil, err
	}

	return t, nil
}

func (t *IdentifiersTable) InitializeTable(ctx context.Context, db *db_sql.DB) error {
	return sfom_sql.CreateTableIfNecessary(ctx, db, t)
}

func (t *IdentifiersTable) Name() string {
	return t.name
}

func (t *IdentifiersTable) Schema(db *db_sql.DB) (string, error) {
	return sfom_sql.LoadSchema(db, FS, IDENTIFIERS_TABLE_NAME)
}

func (t *IdentifiersTable) IndexRecord(ctx context.Context, db *db_sql.DB, tx *db_sql.Tx, i any) error {
	return NotSupportedError
}
