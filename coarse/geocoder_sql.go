package coarse

import (
	"context"
	db_sql "database/sql"
	"fmt"
	"io/fs"
	"log/slog"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	_ "modernc.org/sqlite"
	_ "modernc.org/sqlite/vec"

	"github.com/bwmarrin/snowflake"
	geocoder_sql "github.com/sfomuseum/geocoder/coarse/sql"
	x_vfs "github.com/sfomuseum/geocoder/x/vfs"
	"github.com/sfomuseum/go-database/sql"
	"modernc.org/sqlite/vfs"
)

// To do: Support wildcard machine tags
var re_machinetag = regexp.MustCompile(`^[A-Za-z0-9_-]+:[A-Za-z0-9_-]+=[^\s]+$`)

var snowflake_node *snowflake.Node

func init() {

	ctx := context.Background()
	MustRegisterGeocoder(ctx, "sql", NewSQLGeocoder)

	n, err := snowflake.NewNode(1)

	if err != nil {
		panic(err)
	}

	snowflake_node = n
}

type SQLGeocoder struct {
	Geocoder
	db                 *db_sql.DB
	tables             map[string]sql.Table
	vfs                *vfs.FS
	vector_compression string
	vector_query_k     int
	mu                 *sync.RWMutex
	min_query_length   int
	records            []*Record
	batch_size         int
	bulk_workers       int
	identifier_cache   *sync.Map
	identifier_counter int64
	identifier_mu      *sync.RWMutex
}

type NewSQLGeocoderOptions struct {
	Database *db_sql.DB
	VFS      *vfs.FS
}

func NewSQLGeocoder(ctx context.Context, uri string) (Geocoder, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse URI, %w", err)
	}

	vfs_enable := false

	switch u.Host {
	case "sqlite":

		q := u.Query()

		if q.Has("vfs-enable") {

			v, err := strconv.ParseBool(q.Get("vfs-enable"))

			if err != nil {
				return nil, err
			}

			vfs_enable = v
		}

		if vfs_enable {

			vfs_base := q.Get("vfs-base")

			if vfs_base == "" {
				return nil, fmt.Errorf("Missing or empty ?vfs-base paramter")
			}

			_, err := url.Parse(vfs_base)

			if err != nil {
				return nil, fmt.Errorf("Failed to parse ?vfs-base parameter, %w", err)
			}

			vfs_dbname := q.Get("vfs-dbname")

			if vfs_base == "" {
				return nil, fmt.Errorf("Missing or empty ?vfs-dbname paramter")
			}

			vfs_timeout := 5

			if q.Has("vfs-timeout") {

				v, err := strconv.Atoi(q.Get("vfs-timeout"))

				if err != nil {
					return nil, fmt.Errorf("Failed to parse ?vfs-timeout= parameter, %v", err)
				}

				vfs_timeout = v
			}

			vfs_fs := &x_vfs.RemoteHTTPFS{
				BaseURL: vfs_base,
				Client: &http.Client{
					Timeout: time.Duration(vfs_timeout) * time.Second,
				},
			}

			vfs_name, _, err := vfs.New(vfs_fs)

			if err != nil {
				return nil, fmt.Errorf("Failed to derive VFS name, %w", err)
			}

			dsn := fmt.Sprintf("file:%s?vfs=%s&mode=ro", vfs_dbname, vfs_name)
			enc_dsn := url.QueryEscape(dsn)
			uri = fmt.Sprintf("sql://sqlite?dsn=%s", enc_dsn)

			slog.Info("Rewrite geocoder URI to enable VFS", "uri", uri)
		}

	default:
		// pass
	}

	db, err := sql.OpenWithURI(ctx, uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to open database connection, %w", err)
	}

	db.SetMaxOpenConns(1)

	driver := sql.Driver(db)

	switch driver {
	case sql.SQLITE_DRIVER:

		db_tables, err := geocoder_sql.SQLiteTables(ctx)

		if err != nil {
			return nil, fmt.Errorf("Failed to instantiate SQLite tables, %w", err)
		}

		to_create := make([]sql.Table, 0)

		for _, t := range db_tables {
			to_create = append(to_create, t)
		}

		create_tables := true

		if vfs_enable {
			create_tables = false
		}

		db_opts := &sql.ConfigureDatabaseOptions{
			// https://github.com/pelias/placeholder/blob/master/lib/Database.js
			Pragma: []string{
				"PRAGMA JOURNAL_MODE=OFF",
				"PRAGMA SYNCHRONOUS=OFF",
				"PRAGMA FOREIGN_KEYS=OFF",
				"PRAGMA PAGE_SIZE=4096",
				"PRAGMA CACHE_SIZE=-40000",
				"PRAGMA JOURNAL_MODE=MEMORY",
				"PRAGMA TEMP_STORE=MEMORY",
			},
			CreateTablesIfNecessary: create_tables,
			Tables:                  to_create,
		}

		err = sql.ConfigureDatabase(ctx, db, db_opts)

		if err != nil {
			return nil, fmt.Errorf("Failed to configure database, %w", err)
		}

	default:
		return nil, fmt.Errorf("Unsupported SQL driver, %s", driver)
	}

	opts := &NewSQLGeocoderOptions{
		Database: db,
	}

	return NewSQLGeocoderWithOptions(ctx, opts)
}

func NewSQLGeocoderWithFS(ctx context.Context, db_fs fs.FS, db_name string) (Geocoder, error) {

	vfs_name, vfs_fs, err := vfs.New(db_fs)

	if err != nil {
		return nil, fmt.Errorf("Failed to create VFS, %w", err)
	}

	dsn := fmt.Sprintf("file:%s?vfs=%s", db_name, vfs_name)
	db, err := db_sql.Open("sqlite", dsn)

	if err != nil {
		return nil, fmt.Errorf("Failed to open database, %w", err)
	}

	opts := &NewSQLGeocoderOptions{
		Database: db,
		VFS:      vfs_fs,
	}

	return NewSQLGeocoderWithOptions(ctx, opts)
}

func NewSQLGeocoderWithOptions(ctx context.Context, opts *NewSQLGeocoderOptions) (Geocoder, error) {

	err := opts.Database.PingContext(ctx)

	if err != nil {
		return nil, fmt.Errorf("Failed to ping database, %w", err)
	}

	// To do: Account for other SQL databases
	db_tables, err := geocoder_sql.SQLiteTables(ctx)

	if err != nil {
		return nil, fmt.Errorf("Failed to derive tables, %w", err)
	}

	mu := new(sync.RWMutex)

	identifier_cache := new(sync.Map)
	identifier_mu := new(sync.RWMutex)
	identifier_counter := int64(0)

	g := &SQLGeocoder{
		db:                 opts.Database,
		tables:             db_tables,
		vfs:                opts.VFS,
		mu:                 mu,
		vector_compression: geocoder_sql.SQLiteVecDefaultCompression,
		vector_query_k:     50,
		min_query_length:   2,
		records:            make([]*Record, 0),
		batch_size:         1000,
		bulk_workers:       50,
		identifier_cache:   identifier_cache,
		identifier_counter: identifier_counter,
		identifier_mu:      identifier_mu,
	}

	return g, nil
}

func (g *SQLGeocoder) PreIndex(ctx context.Context) error {

	driver := sql.Driver(g.db)

	switch driver {
	case sql.SQLITE_DRIVER:

		q, err := fs.ReadFile(geocoder_sql.FS, "pre_index.sqlite.schema")

		if err != nil {
			return err
		}

		_, err = g.db.ExecContext(ctx, string(q))
		return err

	default:
		return nil
	}
}

func (g *SQLGeocoder) PostIndex(ctx context.Context) error {

	driver := sql.Driver(g.db)

	switch driver {
	case sql.SQLITE_DRIVER:

		err := g.Flush(ctx)

		if err != nil {
			return err
		}

		q, err := fs.ReadFile(geocoder_sql.FS, "post_index.sqlite.schema")

		if err != nil {
			return err
		}

		_, err = g.db.ExecContext(ctx, string(q))
		return err

	default:
		return nil
	}
}

func (g *SQLGeocoder) HasRecordHashChanged(ctx context.Context, rec *Record) (bool, bool, error) {

	compare_hash, err := rec.Hash()

	if err != nil {
		return false, false, err
	}

	q := fmt.Sprintf("SELECT record_hash FROM %s WHERE id = ?", g.tableName("records"))
	row := g.db.QueryRowContext(ctx, q, rec.Id)

	var record_hash string
	err = row.Scan(&record_hash)

	switch {
	case err == db_sql.ErrNoRows:
		return false, true, nil
	case err != nil:
		return false, false, err
	default:
		if record_hash == compare_hash {
			return true, false, nil
		}

		return true, true, nil
	}
}

func (g *SQLGeocoder) RecordExists(ctx context.Context, id string) (bool, error) {

	q := fmt.Sprintf("SELECT 1 FROM %s WHERE id = ?", g.tableName("records"))
	row := g.db.QueryRowContext(ctx, q, id)

	var stub int
	err := row.Scan(&stub)

	switch {
	case err == db_sql.ErrNoRows:
		return false, nil
	case err != nil:
		return false, err
	default:
		return true, nil
	}
}

// AddRecord has been moved in to geocoder_sql_add.go

func (g *SQLGeocoder) RemoveRecord(ctx context.Context, id string) error {

	logger := slog.Default()
	logger = logger.With("id", id)

	tx, err := g.db.BeginTx(ctx, &db_sql.TxOptions{
		Isolation: db_sql.LevelDefault,
		ReadOnly:  false,
	})

	if err != nil {
		return fmt.Errorf("Failed to begin transaction, %w", err)
	}

	defer func() {

		tx.Rollback()

		if err != nil && err != db_sql.ErrTxDone {
			logger.Error("Failed to rollback transaction", "error", err)
		}

	}()

	record_id, err := g.getRecordId(ctx, tx, id)

	if err != nil {
		return err
	}

	tables := []string{
		g.tableName("records"),
		g.tableName("placetypes_alt"),
		g.tableName("dates"),
		g.tableName("bounds"),
		g.tableName("ancestors"),
		g.tableName("tokens"),
		g.tableName("embeddings"),
		g.tableName("embeddings_records"),
	}

	for _, t := range tables {

		q := fmt.Sprintf("DELETE FROM %s WHERE record_id = ?", t)

		_, err := tx.ExecContext(ctx, q, record_id)

		if err != nil {
			return fmt.Errorf("Failed to remove %s table, %w", t, err)
		}
	}

	err = tx.Commit()

	if err != nil {
		return fmt.Errorf("Failed to commit transaction, %w", err)
	}

	return nil
}

// Query has been moved in geocoder_sql_query.go

func (g *SQLGeocoder) Flush(ctx context.Context) error {

	g.mu.Lock()
	defer g.mu.Unlock()

	if len(g.records) > 0 {

		err := g.addRecords(ctx, g.records...)

		if err != nil {
			return err
		}

		g.records = make([]*Record, 0)
	}

	return nil
}

func (g *SQLGeocoder) Close() error {

	if g.vfs != nil {
		g.vfs.Close()
	}

	return g.db.Close()
}

func (g *SQLGeocoder) uidForVectorRecord(ctx context.Context, tx *db_sql.Tx, record_id int64, model string, language string, tag string) (int64, error) {

	q := fmt.Sprintf("SELECT id FROM %s WHERE record_id = ? AND model = ? AND language = ? AND tag = ?", g.tableName("embeddings_records"))

	row := tx.QueryRowContext(ctx, q, record_id, model, language, tag)

	var id int64
	err := row.Scan(&id)

	switch {
	case err == db_sql.ErrNoRows:
		new_id := snowflake_node.Generate()
		return new_id.Int64(), nil
	case err != nil:
		return 0, err
	default:
		return id, nil
	}
}

func (g *SQLGeocoder) tableName(label string) string {

	t, ok := g.tables[label]

	if !ok {
		slog.Warn("Failed to retrieve table for label", "label", label)
		return ""
	}

	return t.Name()
}

func (g *SQLGeocoder) storeIdentifier(ctx context.Context, tx *db_sql.Tx, id string) (int64, error) {

	g.identifier_mu.Lock()
	defer g.identifier_mu.Unlock()

	record_id, err := g.retrieveInt64IdentifierTx(ctx, tx, id)

	if err == nil {
		return record_id, nil
	}

	record_id = atomic.AddInt64(&g.identifier_counter, 1)

	q := fmt.Sprintf("INSERT INTO %s (id, identifier) VALUES (?, ?)", g.tableName("identifiers"))

	_, err = tx.ExecContext(ctx, q, record_id, id)

	g.identifier_cache.Store(id, record_id)
	return record_id, nil
}

func (g *SQLGeocoder) retrieveStringIdentifier(ctx context.Context, id int64) (string, error) {

	q := fmt.Sprintf("SELECT identifier FROM %s WHERE id = ?", g.tableName("identifiers"))
	row := g.db.QueryRowContext(ctx, q, id)

	var identifier string
	err := row.Scan(&identifier)

	if err != nil {
		return "", err
	}

	return identifier, nil
}

func (g *SQLGeocoder) retrieveInt64IdentifierTx(ctx context.Context, tx *db_sql.Tx, id string) (int64, error) {

	v, ok := g.identifier_cache.Load(id)

	if ok {
		return v.(int64), nil
	}

	q := fmt.Sprintf("SELECT id FROM %s WHERE identifier = ?", g.tableName("identifiers"))
	row := tx.QueryRowContext(ctx, q, id)

	var record_id int64
	err := row.Scan(&record_id)

	if err != nil {
		return 0, err
	}

	g.identifier_cache.Store(id, record_id)
	return record_id, nil
}

func (g *SQLGeocoder) retrieveInt64IdentifierDb(ctx context.Context, id string) (int64, error) {

	v, ok := g.identifier_cache.Load(id)

	if ok {
		return v.(int64), nil
	}

	q := fmt.Sprintf("SELECT id FROM %s WHERE identifier = ?", g.tableName("identifiers"))
	row := g.db.QueryRowContext(ctx, q, id)

	var record_id int64
	err := row.Scan(&record_id)

	if err != nil {
		return 0, err
	}

	g.identifier_cache.Store(id, record_id)
	return record_id, nil
}

func (g *SQLGeocoder) getRecordId(ctx context.Context, tx *db_sql.Tx, id string) (int64, error) {

	record_q := fmt.Sprintf("SELECT record_id FROM %s WHERE id = ?", g.tableName("records"))
	row := tx.QueryRowContext(ctx, record_q, id)

	var record_id int64

	err := row.Scan(&record_id)

	if err != nil {
		return 0, err
	}

	return record_id, nil
}
