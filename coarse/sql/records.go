package sql

import (
	"context"
	db_sql "database/sql"

	sfom_sql "github.com/sfomuseum/go-database/sql"
)

const RECORDS_TABLE_NAME string = "records"

type RecordsTableOptions struct {
	IndexAltFiles bool
}

func DefaultRecordsTableOptions() (*RecordsTableOptions, error) {

	opts := RecordsTableOptions{
		IndexAltFiles: false,
	}

	return &opts, nil
}

type RecordsTable struct {
	sfom_sql.Table
	name    string
	options *RecordsTableOptions
}

func NewRecordsTable(ctx context.Context) (sfom_sql.Table, error) {

	opts, err := DefaultRecordsTableOptions()

	if err != nil {
		return nil, err
	}

	return NewRecordsTableWithOptions(ctx, opts)
}

func NewRecordsTableWithOptions(ctx context.Context, opts *RecordsTableOptions) (sfom_sql.Table, error) {

	t := RecordsTable{
		name:    RECORDS_TABLE_NAME,
		options: opts,
	}

	return &t, nil
}

func NewRecordsTableWithDatabase(ctx context.Context, db *db_sql.DB) (sfom_sql.Table, error) {

	opts, err := DefaultRecordsTableOptions()

	if err != nil {
		return nil, err
	}

	return NewRecordsTableWithDatabaseAndOptions(ctx, db, opts)
}

func NewRecordsTableWithDatabaseAndOptions(ctx context.Context, db *db_sql.DB, opts *RecordsTableOptions) (sfom_sql.Table, error) {

	t, err := NewRecordsTableWithOptions(ctx, opts)

	if err != nil {
		return nil, err
	}

	err = t.InitializeTable(ctx, db)

	if err != nil {
		return nil, err
	}

	return t, nil
}

func (t *RecordsTable) InitializeTable(ctx context.Context, db *db_sql.DB) error {
	return sfom_sql.CreateTableIfNecessary(ctx, db, t)
}

func (t *RecordsTable) Name() string {
	return t.name
}

func (t *RecordsTable) Schema(db *db_sql.DB) (string, error) {
	return sfom_sql.LoadSchema(db, FS, RECORDS_TABLE_NAME)
}

func (t *RecordsTable) IndexRecord(ctx context.Context, db *db_sql.DB, tx *db_sql.Tx, i any) error {
	return NotSupportedError
}
