package sql

import (
	"context"
	db_sql "database/sql"

	sfom_sql "github.com/sfomuseum/go-database/sql"
)

const EMBEDDINGS_RECORDS_TABLE_NAME string = "embeddings_records"

type EmbeddingsRecordsTableOptions struct{}

func DefaultEmbeddingsRecordsTableOptions() (*EmbeddingsRecordsTableOptions, error) {

	opts := EmbeddingsRecordsTableOptions{}

	return &opts, nil
}

type EmbeddingsRecordsTable struct {
	sfom_sql.Table
	name    string
	options *EmbeddingsRecordsTableOptions
}

func NewEmbeddingsRecordsTable(ctx context.Context) (sfom_sql.Table, error) {

	opts, err := DefaultEmbeddingsRecordsTableOptions()

	if err != nil {
		return nil, err
	}

	return NewEmbeddingsRecordsTableWithOptions(ctx, opts)
}

func NewEmbeddingsRecordsTableWithOptions(ctx context.Context, opts *EmbeddingsRecordsTableOptions) (sfom_sql.Table, error) {

	t := EmbeddingsRecordsTable{
		name:    EMBEDDINGS_RECORDS_TABLE_NAME,
		options: opts,
	}

	return &t, nil
}

func NewEmbeddingsRecordsTableWithDatabase(ctx context.Context, db *db_sql.DB) (sfom_sql.Table, error) {

	opts, err := DefaultEmbeddingsRecordsTableOptions()

	if err != nil {
		return nil, err
	}

	return NewEmbeddingsRecordsTableWithDatabaseAndOptions(ctx, db, opts)
}

func NewEmbeddingsRecordsTableWithDatabaseAndOptions(ctx context.Context, db *db_sql.DB, opts *EmbeddingsRecordsTableOptions) (sfom_sql.Table, error) {

	t, err := NewEmbeddingsRecordsTableWithOptions(ctx, opts)

	if err != nil {
		return nil, err
	}

	err = t.InitializeTable(ctx, db)

	if err != nil {
		return nil, err
	}

	return t, nil
}

func (t *EmbeddingsRecordsTable) InitializeTable(ctx context.Context, db *db_sql.DB) error {
	return sfom_sql.CreateTableIfNecessary(ctx, db, t)
}

func (t *EmbeddingsRecordsTable) Name() string {
	return t.name
}

func (t *EmbeddingsRecordsTable) Schema(db *db_sql.DB) (string, error) {
	return sfom_sql.LoadSchema(db, FS, EMBEDDINGS_RECORDS_TABLE_NAME)
}

func (t *EmbeddingsRecordsTable) IndexRecord(ctx context.Context, db *db_sql.DB, tx *db_sql.Tx, i any) error {
	return NotSupportedError
}
