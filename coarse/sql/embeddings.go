package sql

// To do: Update to handle compressions and dimensions

import (
	"context"
	db_sql "database/sql"

	sfom_sql "github.com/sfomuseum/go-database/sql"
)

const EMBEDDINGS_TABLE_NAME string = "embeddings"

type EmbeddingsTableOptions struct {
	Dimensions  int
	Compression string
}

type EmbeddingsTableSchemaVars struct {
	Name        string
	Dimensions  int
	Compression string
}

func DefaultEmbeddingsTableOptions() (*EmbeddingsTableOptions, error) {

	opts := EmbeddingsTableOptions{
		Dimensions:  384,
		Compression: SQLiteVecDefaultCompression,
	}

	return &opts, nil
}

type EmbeddingsTable struct {
	sfom_sql.Table
	name    string
	options *EmbeddingsTableOptions
}

func NewEmbeddingsTable(ctx context.Context) (sfom_sql.Table, error) {

	opts, err := DefaultEmbeddingsTableOptions()

	if err != nil {
		return nil, err
	}

	return NewEmbeddingsTableWithOptions(ctx, opts)
}

func NewEmbeddingsTableWithOptions(ctx context.Context, opts *EmbeddingsTableOptions) (sfom_sql.Table, error) {

	t := EmbeddingsTable{
		name:    EMBEDDINGS_TABLE_NAME,
		options: opts,
	}

	return &t, nil
}

func NewEmbeddingsTableWithDatabase(ctx context.Context, db *db_sql.DB) (sfom_sql.Table, error) {

	opts, err := DefaultEmbeddingsTableOptions()

	if err != nil {
		return nil, err
	}

	return NewEmbeddingsTableWithDatabaseAndOptions(ctx, db, opts)
}

func NewEmbeddingsTableWithDatabaseAndOptions(ctx context.Context, db *db_sql.DB, opts *EmbeddingsTableOptions) (sfom_sql.Table, error) {

	t, err := NewEmbeddingsTableWithOptions(ctx, opts)

	if err != nil {
		return nil, err
	}

	err = t.InitializeTable(ctx, db)

	if err != nil {
		return nil, err
	}

	return t, nil
}

func (t *EmbeddingsTable) InitializeTable(ctx context.Context, db *db_sql.DB) error {
	return sfom_sql.CreateTableIfNecessary(ctx, db, t)
}

func (t *EmbeddingsTable) Name() string {
	return t.name
}

func (t *EmbeddingsTable) Schema(db *db_sql.DB) (string, error) {

	vars := EmbeddingsTableSchemaVars{
		Name:        t.name,
		Dimensions:  t.options.Dimensions,
		Compression: t.options.Compression,
	}

	return sfom_sql.LoadSchemaWithVars(db, FS, t.name, vars)
}

func (t *EmbeddingsTable) IndexRecord(ctx context.Context, db *db_sql.DB, tx *db_sql.Tx, i any) error {
	return NotSupportedError
}
