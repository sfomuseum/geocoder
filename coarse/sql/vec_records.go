package sql

import (
	"context"
	db_sql "database/sql"

	sfom_sql "github.com/sfomuseum/go-database/sql"
)

const VEC_RECORDS_TABLE_NAME string = "vec_records"

type VectorRecordsTableOptions struct{}

func DefaultVectorRecordsTableOptions() (*VectorRecordsTableOptions, error) {

	opts := VectorRecordsTableOptions{}

	return &opts, nil
}

type VectorRecordsTable struct {
	sfom_sql.Table
	name    string
	options *VectorRecordsTableOptions
}

func NewVectorRecordsTable(ctx context.Context) (sfom_sql.Table, error) {

	opts, err := DefaultVectorRecordsTableOptions()

	if err != nil {
		return nil, err
	}

	return NewVectorRecordsTableWithOptions(ctx, opts)
}

func NewVectorRecordsTableWithOptions(ctx context.Context, opts *VectorRecordsTableOptions) (sfom_sql.Table, error) {

	t := VectorRecordsTable{
		name:    TOKENS_TABLE_NAME,
		options: opts,
	}

	return &t, nil
}

func NewVectorRecordsTableWithDatabase(ctx context.Context, db *db_sql.DB) (sfom_sql.Table, error) {

	opts, err := DefaultVectorRecordsTableOptions()

	if err != nil {
		return nil, err
	}

	return NewVectorRecordsTableWithDatabaseAndOptions(ctx, db, opts)
}

func NewVectorRecordsTableWithDatabaseAndOptions(ctx context.Context, db *db_sql.DB, opts *VectorRecordsTableOptions) (sfom_sql.Table, error) {

	t, err := NewVectorRecordsTableWithOptions(ctx, opts)

	if err != nil {
		return nil, err
	}

	err = t.InitializeTable(ctx, db)

	if err != nil {
		return nil, err
	}

	return t, nil
}

func (t *VectorRecordsTable) InitializeTable(ctx context.Context, db *db_sql.DB) error {
	return sfom_sql.CreateTableIfNecessary(ctx, db, t)
}

func (t *VectorRecordsTable) Name() string {
	return t.name
}

func (t *VectorRecordsTable) Schema(db *db_sql.DB) (string, error) {
	return sfom_sql.LoadSchema(db, FS, TOKENS_TABLE_NAME)
}

func (t *VectorRecordsTable) IndexRecord(ctx context.Context, db *db_sql.DB, tx *db_sql.Tx, i any) error {
	return NotSupportedError
}
