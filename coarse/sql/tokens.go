package sql

import (
	"context"
	db_sql "database/sql"

	sfom_sql "github.com/sfomuseum/go-database/sql"
)

const TOKENS_TABLE_NAME string = "tokens"

type TokensTableOptions struct{}

func DefaultTokensTableOptions() (*TokensTableOptions, error) {

	opts := TokensTableOptions{}

	return &opts, nil
}

type TokensTable struct {
	sfom_sql.Table
	name    string
	options *TokensTableOptions
}

func NewTokensTable(ctx context.Context) (sfom_sql.Table, error) {

	opts, err := DefaultTokensTableOptions()

	if err != nil {
		return nil, err
	}

	return NewTokensTableWithOptions(ctx, opts)
}

func NewTokensTableWithOptions(ctx context.Context, opts *TokensTableOptions) (sfom_sql.Table, error) {

	t := TokensTable{
		name:    TOKENS_TABLE_NAME,
		options: opts,
	}

	return &t, nil
}

func NewTokensTableWithDatabase(ctx context.Context, db *db_sql.DB) (sfom_sql.Table, error) {

	opts, err := DefaultTokensTableOptions()

	if err != nil {
		return nil, err
	}

	return NewTokensTableWithDatabaseAndOptions(ctx, db, opts)
}

func NewTokensTableWithDatabaseAndOptions(ctx context.Context, db *db_sql.DB, opts *TokensTableOptions) (sfom_sql.Table, error) {

	t, err := NewTokensTableWithOptions(ctx, opts)

	if err != nil {
		return nil, err
	}

	err = t.InitializeTable(ctx, db)

	if err != nil {
		return nil, err
	}

	return t, nil
}

func (t *TokensTable) InitializeTable(ctx context.Context, db *db_sql.DB) error {
	return sfom_sql.CreateTableIfNecessary(ctx, db, t)
}

func (t *TokensTable) Name() string {
	return t.name
}

func (t *TokensTable) Schema(db *db_sql.DB) (string, error) {
	return sfom_sql.LoadSchema(db, FS, TOKENS_TABLE_NAME)
}

func (t *TokensTable) IndexRecord(ctx context.Context, db *db_sql.DB, tx *db_sql.Tx, i any) error {
	return NotSupportedError
}
