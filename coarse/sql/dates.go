package sql

import (
	"context"
	db_sql "database/sql"

	sfom_sql "github.com/sfomuseum/go-database/sql"
)

const DATES_TABLE_NAME string = "dates"

type DatesTableOptions struct{}

func DefaultDatesTableOptions() (*DatesTableOptions, error) {

	opts := DatesTableOptions{}

	return &opts, nil
}

type DatesTable struct {
	sfom_sql.Table
	name    string
	options *DatesTableOptions
}

func NewDatesTable(ctx context.Context) (sfom_sql.Table, error) {

	opts, err := DefaultDatesTableOptions()

	if err != nil {
		return nil, err
	}

	return NewDatesTableWithOptions(ctx, opts)
}

func NewDatesTableWithOptions(ctx context.Context, opts *DatesTableOptions) (sfom_sql.Table, error) {

	t := DatesTable{
		name:    DATES_TABLE_NAME,
		options: opts,
	}

	return &t, nil
}

func NewDatesTableWithDatabase(ctx context.Context, db *db_sql.DB) (sfom_sql.Table, error) {

	opts, err := DefaultDatesTableOptions()

	if err != nil {
		return nil, err
	}

	return NewDatesTableWithDatabaseAndOptions(ctx, db, opts)
}

func NewDatesTableWithDatabaseAndOptions(ctx context.Context, db *db_sql.DB, opts *DatesTableOptions) (sfom_sql.Table, error) {

	t, err := NewDatesTableWithOptions(ctx, opts)

	if err != nil {
		return nil, err
	}

	err = t.InitializeTable(ctx, db)

	if err != nil {
		return nil, err
	}

	return t, nil
}

func (t *DatesTable) InitializeTable(ctx context.Context, db *db_sql.DB) error {
	return sfom_sql.CreateTableIfNecessary(ctx, db, t)
}

func (t *DatesTable) Name() string {
	return t.name
}

func (t *DatesTable) Schema(db *db_sql.DB) (string, error) {
	return sfom_sql.LoadSchema(db, FS, DATES_TABLE_NAME)
}

func (t *DatesTable) IndexRecord(ctx context.Context, db *db_sql.DB, tx *db_sql.Tx, i any) error {
	return NotSupportedError
}
