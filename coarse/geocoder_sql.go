package coarse

import (
	"context"
	db_sql "database/sql"
	"encoding/json"
	"fmt"
	"io/fs"
	"log/slog"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode"

	_ "modernc.org/sqlite"
	_ "modernc.org/sqlite/vec"

	"github.com/bwmarrin/snowflake"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
	geocoder_sql "github.com/sfomuseum/geocoder/coarse/sql"
	"github.com/sfomuseum/geocoder/placeholder"
	x_vfs "github.com/sfomuseum/geocoder/x/vfs"
	"github.com/sfomuseum/go-database/sql"
	"github.com/whosonfirst/go-whosonfirst/v4/hierarchies"
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

	g := &SQLGeocoder{
		db:                 opts.Database,
		tables:             db_tables,
		vfs:                opts.VFS,
		mu:                 mu,
		vector_compression: geocoder_sql.SQLiteVecDefaultCompression,
		vector_query_k:     50,
		min_query_length:   2,
		records:            make([]*Record, 0),
		batch_size:         10,
		bulk_workers:       50,
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

		var q string

		switch t {
		case g.tableName("bounds"):
			q = fmt.Sprintf("DELETE FROM %s WHERE wofid = ?", t)
		case g.tableName("embeddings_records"):
			q = fmt.Sprintf("DELETE FROM %s WHERE record_id = ?", t)
		default:
			q = fmt.Sprintf("DELETE FROM %s WHERE id = ?", t)
		}

		_, err := tx.ExecContext(ctx, q, id)

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

func (g *SQLGeocoder) assignExtra(ctx context.Context, f *geojson.Feature) error {

	logger := slog.Default()
	logger = logger.With("id", f.ID)

	err := g.assignHierarchiesAndLabel(ctx, f)

	if err != nil {
		logger.Warn("Failed to assign label", "error", err)
	}

	err = g.assignPlacetypeAlt(ctx, f)

	if err != nil {
		logger.Warn("Failed to assign placetype alt", "error", err)
	}

	err = g.assignBBox(ctx, f)

	if err != nil {
		logger.Warn("Failed to assign bbox", "error", err)
	}

	return nil
}

func (g *SQLGeocoder) assignBBox(ctx context.Context, f *geojson.Feature) error {

	bounds_q := fmt.Sprintf("SELECT MIN(minx), MIN(miny), MAX(maxx), MAX(maxy) FROM %s WHERE wofid = ?", g.tableName("bounds"))

	bounds_row := g.db.QueryRowContext(ctx, bounds_q, f.ID)

	var minx float64
	var miny float64
	var maxx float64
	var maxy float64

	err := bounds_row.Scan(&minx, &miny, &maxx, &maxy)

	if err != nil {
		return fmt.Errorf("Failed to query bounds, %w", err)
	}

	bounds := orb.Bound{
		Min: orb.Point([2]float64{minx, miny}),
		Max: orb.Point([2]float64{maxx, maxy}),
	}

	f.BBox = geojson.NewBBox(bounds)
	return nil
}

func (g *SQLGeocoder) assignHierarchiesAndLabel(ctx context.Context, f *geojson.Feature) error {

	// To do: Update to account for string-based IDs and convert back to WOF int-based IDs
	// where applicable.

	logger := slog.Default()
	logger = logger.With("id", f.ID)

	f_hiers := f.Properties.MustString("wof:hierarchies", "")

	if f_hiers == "" {
		f.Properties["wof:hierarchies"] = make([]map[string]int64, 0)
		return nil
	}

	var str_hiers []map[string]string

	err := json.Unmarshal([]byte(f_hiers), &str_hiers)

	if err != nil {
		return fmt.Errorf("Failed to unmarshal hierarchies, %w", err)
	}

	f.Properties["wof:hierarchies"] = str_hiers

	if len(str_hiers) == 0 {
		return nil
	}

	// Convert back to wof-compliant hierarchies

	wof_hiers := make([]map[string]int64, 0)

	for _, h := range str_hiers {

		wof_h := make(map[string]int64)
		has_wof := false

		for k, v := range h {

			if !strings.HasPrefix(v, "wof:id=") {
				continue
			}

			// To do: proper machine tag parser

			v = strings.Replace(v, "wof:id=", "", 1)
			id, err := strconv.ParseInt(v, 10, 64)

			if err != nil {
				slog.Warn("Failed to parse what looks like a wof:id", "id", h[k], "error", err)
				continue
			}

			wof_h[k] = id
			has_wof = true
		}

		if has_wof {
			wof_hiers = append(wof_hiers, wof_h)
		}
	}

	if len(wof_hiers) == 0 {
		return nil
	}

	// START OF put me in a function

	name := f.Properties.MustString("wof:name")
	str_pt := f.Properties.MustString("wof:placetype", "")

	labels := []string{
		name,
	}

	var parent_id int64
	f_pid := f.Properties["wof:parent_id"]

	switch f_pid.(type) {
	case string:

		str_p := f_pid.(string)

		if strings.HasPrefix(str_p, "wof:id=") {

			// To do: proper machine tag parser

			str_p = strings.Replace(str_p, "wof:id=", "", 1)
			id, err := strconv.ParseInt(str_p, 10, 64)

			if err != nil {
				slog.Warn("Failed to parse what looks like a wof:id", "id", str_p, "error", err)
				parent_id = -1
			} else {
				parent_id = id
			}

		} else {
			parent_id = -1
		}

	default:
		parent_id = -1
	}

	label_opts := &hierarchies.AncestorIdsForLabelOptions{
		Hierarchies: wof_hiers,
		Placetype:   str_pt,
		ParentId:    parent_id,
	}

	name_ids := hierarchies.AncestorIdsForLabel(label_opts)

	names_q := fmt.Sprintf("SELECT name, placetype, country from %s WHERE id = ?", g.tableName("records"))

	for _, id := range name_ids {

		wof_id := fmt.Sprintf("wof:id=%d", id)

		var id_name string
		var id_placetype string
		var id_country string

		row := g.db.QueryRowContext(ctx, names_q, wof_id)
		err := row.Scan(&id_name, &id_placetype, &id_country)

		switch {
		case err == db_sql.ErrNoRows:
			continue
		case err != nil:
			logger.Warn("Failed to query ID for name", "name id", wof_id)
		default:

			switch id_placetype {
			case "country":
				labels = append(labels, id_country)
			default:
				labels = append(labels, id_name)
			}
		}
	}

	f.Properties["wof:label"] = strings.Join(labels, ", ")
	return nil
}

func (g *SQLGeocoder) assignPlacetypeAlt(ctx context.Context, f *geojson.Feature) error {

	pt_q := fmt.Sprintf("SELECT placetype from %s WHERE id = ?", g.tableName("placetypes_alt"))

	pt_rows, err := g.db.QueryContext(ctx, pt_q, f.ID)

	if err != nil {
		return err
	}

	defer pt_rows.Close()

	alt_pt := make([]string, 0)

	for pt_rows.Next() {

		var pt string
		err := pt_rows.Scan(&pt)

		if err != nil {
			return fmt.Errorf("Failed to scan placetypes row, %w", err)
		}

		alt_pt = append(alt_pt, pt)
	}

	err = pt_rows.Err()

	if err != nil {
		return fmt.Errorf("Failed to derive alt placetypes, %w", err)
	}

	if len(alt_pt) > 0 {
		f.Properties["wof:placetype_alt"] = alt_pt
	}

	return nil
}

func (g *SQLGeocoder) prepareQuery(input string) string {

	if re_machinetag.MatchString(input) {
		// To do: Support wildcards
		input = strings.ReplaceAll(input, ":", "_")
		input = strings.ReplaceAll(input, "=", "__")
		return input
	}

	words := strings.Fields(placeholder.Normalize(input))

	if len(words) == 0 {
		return ""
	}

	var sanitized []string

	for _, word := range words {

		// Strip characters that disrupt FTS5 syntax (like raw quotes or dashes)
		clean := strings.Map(func(r rune) rune {
			// unicode.IsLetter handles all global alphabets (Greek, Cyrillic, CJK, Arabic, etc.)
			// unicode.IsNumber handles global digits (0-9, and non-Arabic numeral systems)
			if unicode.IsLetter(r) || unicode.IsNumber(r) {
				return r
			}

			// Keep Unicode letters/numbers if working with global languages
			return -1
		}, word)

		if clean != "" {
			sanitized = append(sanitized, clean)
		}
	}

	if len(sanitized) == 0 {
		return ""
	}

	lastIdx := len(sanitized) - 1
	sanitized[lastIdx] = sanitized[lastIdx] + "*"

	return strings.Join(sanitized, " AND ")
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
