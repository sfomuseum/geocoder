package sql

// To do: Update to handle compressions and dimensions

import (
	"context"
	db_sql "database/sql"

	sfom_sql "github.com/sfomuseum/go-database/sql"
)

const VEC_TABLE_NAME string = "vec"

type VectorsTableOptions struct {
	Dimensions  int
	Compression string
}

func DefaultVectorsTableOptions() (*VectorsTableOptions, error) {

	opts := VectorsTableOptions{
		Dimensions:  768,
		Compression: SQLiteVecDefaultCompression,
	}

	return &opts, nil
}

type VectorsTable struct {
	sfom_sql.Table
	name    string
	options *VectorsTableOptions
}

func NewVectorsTable(ctx context.Context) (sfom_sql.Table, error) {

	opts, err := DefaultVectorsTableOptions()

	if err != nil {
		return nil, err
	}

	return NewVectorsTableWithOptions(ctx, opts)
}

func NewVectorsTableWithOptions(ctx context.Context, opts *VectorsTableOptions) (sfom_sql.Table, error) {

	t := VectorsTable{
		name:    TOKENS_TABLE_NAME,
		options: opts,
	}

	return &t, nil
}

func NewVectorsTableWithDatabase(ctx context.Context, db *db_sql.DB) (sfom_sql.Table, error) {

	opts, err := DefaultVectorsTableOptions()

	if err != nil {
		return nil, err
	}

	return NewVectorsTableWithDatabaseAndOptions(ctx, db, opts)
}

func NewVectorsTableWithDatabaseAndOptions(ctx context.Context, db *db_sql.DB, opts *VectorsTableOptions) (sfom_sql.Table, error) {

	t, err := NewVectorsTableWithOptions(ctx, opts)

	if err != nil {
		return nil, err
	}

	err = t.InitializeTable(ctx, db)

	if err != nil {
		return nil, err
	}

	return t, nil
}

func (t *VectorsTable) InitializeTable(ctx context.Context, db *db_sql.DB) error {
	return sfom_sql.CreateTableIfNecessary(ctx, db, t)
}

func (t *VectorsTable) Name() string {
	return t.name
}

func (t *VectorsTable) Schema(db *db_sql.DB) (string, error) {
	return sfom_sql.LoadSchema(db, FS, TOKENS_TABLE_NAME)
}

func (t *VectorsTable) IndexRecord(ctx context.Context, db *db_sql.DB, tx *db_sql.Tx, i any) error {
	return NotSupportedError
}
