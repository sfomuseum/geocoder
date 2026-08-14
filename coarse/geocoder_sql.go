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
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode"

	_ "modernc.org/sqlite"
	_ "modernc.org/sqlite/vec"

	"github.com/aaronland/go-pagination"
	"github.com/aaronland/go-pagination/countable"
	"github.com/bwmarrin/snowflake"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
	geocoder_sql "github.com/sfomuseum/geocoder/coarse/sql"
	"github.com/sfomuseum/geocoder/placeholder"
	x_vfs "github.com/sfomuseum/geocoder/x/vfs"
	"github.com/sfomuseum/go-database/sql"
	"github.com/sfomuseum/go-embeddings"
	"github.com/whosonfirst/go-whosonfirst/v4/hierarchies"
	"modernc.org/sqlite/vfs"
	// sqlite_vec "modernc.org/sqlite/vec"
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
	embedder           embeddings.Embedder[float32]
	embedder_models    []string
	vector_compression string
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

	switch u.Host {
	case "sqlite":

		q := u.Query()

		vfs_enable := false

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
			CreateTablesIfNecessary: true,
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

func (g *SQLGeocoder) RecordExists(ctx context.Context, id int64) (bool, error) {

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

func (g *SQLGeocoder) AddRecord(ctx context.Context, rec *Record) error {

	g.mu.Lock()
	defer g.mu.Unlock()

	g.records = append(g.records, rec)

	if len(g.records) >= g.batch_size {

		err := g.addRecords(ctx, g.records...)

		if err != nil {
			return err
		}

		g.records = make([]*Record, 0)
	}

	return nil
}

func (g *SQLGeocoder) addRecords(ctx context.Context, records ...*Record) error {

	logger := slog.Default()
	logger.Debug("Add bulk records", "count", len(records), "workers", g.bulk_workers)

	tx, err := g.db.BeginTx(ctx, &db_sql.TxOptions{
		Isolation: db_sql.LevelDefault,
		ReadOnly:  false,
	})

	if err != nil {
		return fmt.Errorf("Failed to begin transaction, %w", err)
	}

	t1 := time.Now()

	defer func() {

		tx.Rollback()

		if err != nil && err != db_sql.ErrTxDone {
			logger.Error("Failed to rollback transaction", "error", err)
		}

		logger.Debug("Time to bulk index records", "count", len(records), "time", time.Since(t1))
	}()

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	done_ch := make(chan bool)
	err_ch := make(chan error)

	throttle := make(chan bool, g.bulk_workers)

	for i := 0; i < g.bulk_workers; i++ {
		throttle <- true
	}

	for _, rec := range records {

		<-throttle

		select {
		case <-ctx.Done():
			break
		default:
			// pass
		}

		go func(rec *Record) {

			logger := slog.Default()
			logger = logger.With("id", rec.Id)

			defer func() {
				throttle <- true
				done_ch <- true
			}()

			enc_hierarchies, err := json.Marshal(rec.Hierarchies)

			if err != nil {
				err_ch <- fmt.Errorf("Failed to marshal hierarchies, %w", err)
				return
			}

			record_hash, err := rec.Hash()

			if err != nil {
				err_ch <- fmt.Errorf("Failed to hash record, %w", err)
				return
			}

			rec_q := fmt.Sprintf("INSERT OR REPLACE INTO %s (id, parent_id, name, placetype, latitude, longitude, country, inception, cessation, hierarchies, is_current, population_rank, record_hash) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", g.tableName("records"))

			_, err = tx.ExecContext(ctx, rec_q, rec.Id, rec.ParentId, rec.Name, rec.Placetype, rec.Centroid.Lat(), rec.Centroid.Lon(), rec.Country, rec.Inception, rec.Cessation, string(enc_hierarchies), rec.IsCurrent, rec.PopulationRank, record_hash)

			if err != nil {
				err_ch <- fmt.Errorf("Failed to insert in to records, %w", err)
				return
			}

			// Placetypes (alt)

			pt_stq := fmt.Sprintf("INSERT INTO %s (id, placetype) VALUES(?, ?)", g.tableName("placetypes_alt"))

			pt_st, err := tx.Prepare(pt_stq)

			if err != nil {
				err_ch <- fmt.Errorf("Failed to prepare placetypes statement, %w", err)
				return
			}

			defer pt_st.Close()

			for _, pt := range rec.PlacetypeAlt {

				_, err = pt_st.ExecContext(ctx, rec.Id, pt)

				if err != nil {
					err_ch <- fmt.Errorf("Failed to execute placetype statement, %w", err)
					return
				}
			}

			// Ancestors

			anc_stq := fmt.Sprintf("INSERT INTO %s (id, ancestor_id) VALUES(?, ?)", g.tableName("ancestors"))
			anc_st, err := tx.Prepare(anc_stq)

			if err != nil {
				err_ch <- fmt.Errorf("Failed to prepare ancestors statement, %w", err)
				return
			}

			defer anc_st.Close()

			ancestors := make([]int64, 0)

			for _, hier := range rec.Hierarchies {

				for _, id := range hier {

					if !slices.Contains(ancestors, id) {
						ancestors = append(ancestors, id)
					}
				}
			}

			for _, id := range ancestors {

				_, err = anc_st.ExecContext(ctx, rec.Id, id)

				if err != nil {
					err_ch <- fmt.Errorf("Failed to execute ancestors statement, %w", err)
					return
				}
			}

			// Bounds

			bounds_stq := fmt.Sprintf("INSERT INTO %s (minx, miny, maxx, maxy, wofid) VALUES(?, ?, ?, ?, ?)", g.tableName("bounds"))
			bounds_st, err := tx.Prepare(bounds_stq)

			if err != nil {
				err_ch <- fmt.Errorf("Failed to prepare bounds statement, %w", err)
				return
			}

			defer bounds_st.Close()

			for _, b := range rec.Bounds {

				_, err = bounds_st.ExecContext(ctx, b.Min.X(), b.Min.Y(), b.Max.X(), b.Max.Y(), rec.Id)

				if err != nil {
					err_ch <- fmt.Errorf("Failed to execute bounds statement, %w", err)
					return
				}
			}

			// Dates

			start_outer, start_inner, end_inner, end_outer := rec.DateRanges()

			date_fields := []string{
				"id",
			}

			date_args := []any{
				rec.Id,
			}

			if start_outer != nil {
				date_fields = append(date_fields, "start_outer")
				date_args = append(date_args, start_outer.Unix())
			}

			if start_inner != nil {
				date_fields = append(date_fields, "start_inner")
				date_args = append(date_args, start_inner.Unix())
			}

			if end_inner != nil {
				date_fields = append(date_fields, "end_inner")
				date_args = append(date_args, end_inner.Unix())
			}

			if end_outer != nil {
				date_fields = append(date_fields, "end_outer")
				date_args = append(date_args, end_outer.Unix())
			}

			if len(date_fields) > 1 {

				date_placeholders := make([]string, len(date_fields))

				for i, _ := range date_fields {
					date_placeholders[i] = "?"
				}

				date_q := fmt.Sprintf("INSERT OR REPLACE INTO %s (%s) VALUES (%s)", g.tableName("dates"), strings.Join(date_fields, ","), strings.Join(date_placeholders, ","))

				_, err := tx.ExecContext(ctx, date_q, date_args...)

				if err != nil {
					err_ch <- fmt.Errorf("Failed to insert dates, %v", err)
					return
				}
			}

			// Tokens

			tok_stq := fmt.Sprintf("INSERT INTO %s (id, token, lang, tag) VALUES(?, ?, ?, ?)", g.tableName("tokens"))
			tok_st, err := tx.Prepare(tok_stq)

			if err != nil {
				err_ch <- fmt.Errorf("Failed to prepare token statement, %w", err)
				return
			}

			defer tok_st.Close()

			for lang, tag_tokens := range rec.Tokens {

				for tag, tokens := range tag_tokens {
					_, err = tok_st.ExecContext(ctx, rec.Id, strings.Join(tokens, " "), lang, tag)

					if err != nil {
						err_ch <- fmt.Errorf("Failed to execute token statement, %w", err)
						return
					}
				}
			}

			// Vectors

			if rec.VectorEmbeddings != nil && len(rec.VectorEmbeddings) > 0 {

				slog.Info("Add vector embeddings", "id", rec.Id, "count", len(rec.VectorEmbeddings))

				emb_table := g.tableName("embeddings")

				var vec_q string

				switch g.vector_compression {
				case geocoder_sql.SQLiteVecQuantizeCompression:
					vec_q = fmt.Sprintf("INSERT OR REPLACE INTO %s (rowid, embedding) VALUES (?, vec_quantize_binary(?))", emb_table)
				case geocoder_sql.SQLiteVecMatroyshkaCompression:
					vec_q = fmt.Sprintf("INSERT OR REPLACE INTO %s (rowid, embedding) VALUES (?, vec_normalize(vec_slice(?, 0, %d)))", emb_table, geocoder_sql.SQLiteMatroyshkaDimensions)
				case geocoder_sql.SQLiteVecDefaultCompression:
					vec_q = fmt.Sprintf("INSERT OR REPLACE INTO %s (rowid, embedding) VALUES (?, ?)", emb_table)
				default:
					err_ch <- fmt.Errorf("Invalid or unsupported compression, '%s'", g.vector_compression)
					return
				}

				vec_st, err := tx.Prepare(vec_q)

				if err != nil {
					err_ch <- fmt.Errorf("Failed to prepare vector statement, %w", err)
					return
				}

				defer vec_st.Close()

				vrec_q := fmt.Sprintf("INSERT OR REPLACE INTO %s (id, record_id, model, language, tag) VALUES (?, ?, ?, ?, ?)", g.tableName("embeddings_records"))

				vrec_st, err := tx.Prepare(vrec_q)

				if err != nil {
					err_ch <- fmt.Errorf("Failed to prepare vector record statement, %w", err)
					return
				}

				defer vrec_st.Close()

				for _, v := range rec.VectorEmbeddings {

					for _, e := range v.Embeddings {

						logger.Info("Add embeddings")

						vrec_id, err := g.uidForVectorRecord(ctx, rec.Id, v.Model, e.Language, e.Tag)

						if err != nil {
							err_ch <- err
							return
						}

						logger.Info("Add embeddings with ID", vrec_id)

						enc_e, err := geocoder_sql.SerializeFloat32(e.Embeddings)

						if err != nil {
							err_ch <- err
							return
						}

						_, err = vec_st.ExecContext(ctx, vec_q, vrec_id, enc_e)

						if err != nil {
							err_ch <- err
							return
						}

						_, err = vrec_st.ExecContext(ctx, vrec_id, rec.Id, v.Model, e.Language, e.Tag)

						if err != nil {
							err_ch <- err
							return
						}
					}
				}
			}

			// All done...

		}(rec)
	}

	remaining := len(records)

	for remaining > 0 {
		select {
		case <-done_ch:
			remaining -= 1
			// logger.Info("Bulk add", "remaining", remaining)
		case err := <-err_ch:
			return err
		}
	}

	err = tx.Commit()

	if err != nil {
		return fmt.Errorf("Failed to commit transaction, %w", err)
	}

	return nil
}

func (g *SQLGeocoder) RemoveRecord(ctx context.Context, id int64) error {

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

func (g *SQLGeocoder) Query(ctx context.Context, req *QueryRequest, pg_opts pagination.Options) ([]*geojson.Feature, pagination.Results, error) {

	logger := slog.Default()
	t1 := time.Now()

	defer func() {
		logger.Debug("Time to query", "time", time.Since(t1))
	}()

	if len(req.Query) < g.min_query_length {
		return nil, nil, fmt.Errorf("Query below min query length")
	}

	query_str := g.prepareQuery(req.Query)

	if query_str == "" {
		return nil, nil, fmt.Errorf("empty or invalid search term")
	}

	if len(query_str) < g.min_query_length {
		return nil, nil, fmt.Errorf("Query below min query length")
	}

	logger = logger.With("query", query_str)

	sb := strings.Builder{}

	sb.WriteString(`
		SELECT f.rank, COUNT(*) OVER() as total_count, r.id, r.parent_id, r.name, r.placetype, r.country, r.is_current, r.latitude, r.longitude, r.inception, r.cessation, r.hierarchies
		FROM tokens_fts f
		JOIN tokens t ON t.row_id = f.rowid
		JOIN records r ON r.id = t.id
        `)

	if len(req.Placetype) > 0 {
		sb.WriteString(" LEFT JOIN placetypes_alt p ON r.id = p.id")
	}

	if len(req.BelongsTo) > 0 {
		sb.WriteString(" JOIN ancestors a ON r.id = a.id")
	}

	// dates

	if req.DateStarts != nil || req.DateEnds != nil {
		sb.WriteString(" JOIN dates d ON r.id = d.id")
	}

	// bounds

	if req.Bounds != nil {
		sb.WriteString(" JOIN bounds b ON r.id = b.wofid")
	}

	if req.UseEmbeddings {

		emb_req := &embeddings.EmbeddingsRequest{
			Id:    query_str,
			Model: req.UseEmbeddingsModel,
			Body:  []byte(query_str),
		}

		emb_rsp, err := g.embedder.TextEmbeddings(ctx, emb_req)

		if err != nil {
			return nil, nil, fmt.Errorf("Failed to derive text embeddings for query, %w", err)
		}

		slog.Info("GOT EMBEDDINGS", "rsp", emb_rsp)
		return nil, nil, fmt.Errorf("Not implemented")

	} else {
		sb.WriteString(" WHERE f.token MATCH ?")
	}

	args := []any{
		query_str,
	}

	// Dates

	if req.DateStarts != nil {

		sb.WriteString(" AND (? <= d.start_outer AND ? <= d.start_inner)")
		args = append(args, req.DateStarts.Outer.Start)
		args = append(args, req.DateStarts.Inner.Start)
	}

	if req.DateEnds != nil {

		sb.WriteString(" AND (d.end_inner <= ? AND d.end_outer <= ?)")
		args = append(args, req.DateEnds.Inner.End)
		args = append(args, req.DateEnds.Outer.End)
	}

	// Bounds

	if req.Bounds != nil {

		coords := []any{
			req.Bounds.Min.X(),
			req.Bounds.Max.X(),
			req.Bounds.Min.Y(),
			req.Bounds.Max.Y(),
		}

		sb.WriteString(" AND (b.maxx >= ? AND b.minx <= ? AND b.maxy >= ? AND b.miny <= ?)")
		args = append(args, coords...)
	}

	// Placetypes

	if len(req.Placetype) > 0 {

		placeholders := make([]string, len(req.Placetype))

		for i, pt := range req.Placetype {
			placeholders[i] = "?"
			args = append(args, pt)
		}

		for i, pt := range req.Placetype {
			placeholders[i] = "?"
			args = append(args, pt)
		}

		str_placeholders := strings.Join(placeholders, ",")

		sb.WriteString(fmt.Sprintf(" AND (r.placetype IN (%s) OR p.placetype IN (%s))", str_placeholders, str_placeholders))
	}

	// Belongs to (ancestors)

	if len(req.BelongsTo) > 0 {

		placeholders := make([]string, len(req.BelongsTo))

		for i, anc_id := range req.BelongsTo {
			placeholders[i] = "?"
			args = append(args, anc_id)
		}

		sb.WriteString(fmt.Sprintf(" AND a.ancestor_id IN (%s)", strings.Join(placeholders, ",")))
	}

	// Language

	if req.Lang != "" {
		sb.WriteString(" AND r.lang = ?")
		args = append(args, req.Tag)
	}

	// Language (x-) tag

	if req.Tag != "" {
		sb.WriteString(" AND r.tag = ?")
		args = append(args, req.Tag)
	}

	// Countries

	if len(req.Country) > 0 {

		placeholders := make([]string, len(req.Country))

		for i, pt := range req.Country {
			placeholders[i] = "?"
			args = append(args, pt)
		}

		sb.WriteString(fmt.Sprintf(" AND r.country IN (%s)", strings.Join(placeholders, ",")))
	}

	// Is current

	if req.IsCurrent != nil {
		sb.WriteString(" AND r.is_current = ?")
		args = append(args, req.IsCurrent.StringFlag())
	}

	sb.WriteString(" GROUP BY r.id")

	sb.WriteString(`
		ORDER BY (
			CASE r.is_current
				WHEN 1 THEN 0.0
				WHEN -1 THEN 1.0
				ELSE 2.0
                        END
		) ASC,
                MIN(CASE t.tag
				WHEN 'concordance' THEN 0.5
				WHEN 'preferred'    THEN 1.0
				WHEN 'colloquial' THEN 2.0
				WHEN 'variant'    THEN 4.0
				WHEN 'historical'    THEN 5.0
				WHEN 'unknown'   THEN 6.0
				ELSE 10.0
			END) ASC,
		 r.population_rank DESC,
			(CASE r.placetype
				WHEN 'microhood' THEN 1.0
				WHEN 'neighbourhood' THEN 1.0
				WHEN 'borough' THEN 1.5
				WHEN 'locality' THEN 2.0
				WHEN 'localadmin' THEN 2.25
				WHEN 'campus' THEN 2.5
				WHEN 'postalcode' THEN 2.9
				WHEN 'county' THEN 3.0
				WHEN 'region' THEN 4.0
				WHEN 'country' THEN 5.0	
				ELSE 10.0
                        END) ASC,
                 MIN(f.rank) ASC
	`)

	page := countable.PageFromOptions(pg_opts)
	per_page := pg_opts.PerPage()

	sb.WriteString(fmt.Sprintf(" LIMIT %d", per_page))

	if page > 1 {
		offset := (page - 1) * per_page
		sb.WriteString(fmt.Sprintf(" OFFSET %d", offset))
	}

	q := sb.String()
	//slog.Info(q, "args", args)

	rows, err := g.db.QueryContext(ctx, q, args...)

	if err != nil {
		return nil, nil, err
	}

	defer rows.Close()

	var features []*geojson.Feature
	var total_count int64

	for rows.Next() {

		var rank float64
		var id int64
		var parent_id int64
		var name string
		var country string
		var placetype string
		var is_current int
		var latitude float64
		var longitude float64
		var inception string
		var cessation string
		var enc_hierarchies string

		err := rows.Scan(&rank, &total_count, &id, &parent_id, &name, &placetype, &country, &is_current, &latitude, &longitude, &inception, &cessation, &enc_hierarchies)

		if err != nil {
			return nil, nil, err
		}

		props := map[string]any{
			"wof:id":          id,
			"wof:parent_id":   parent_id,
			"wof:name":        name,
			"wof:country":     country,
			"wof:placetype":   placetype,
			"mz:is_current":   is_current,
			"edtf:inception":  inception,
			"edtf:cessation":  cessation,
			"wof:hierarchies": enc_hierarchies,
			"geocoder:rank":   rank,
		}

		pt := orb.Point([2]float64{longitude, latitude})
		f := geojson.NewFeature(pt)

		f.ID = id
		f.Properties = props
		features = append(features, f)
	}

	err = rows.Err()

	if err != nil {
		return nil, nil, err
	}

	logger = logger.With("total", total_count)

	wg := new(sync.WaitGroup)

	for _, f := range features {

		wg.Go(func() {
			g.assignExtra(ctx, f)
		})
	}

	wg.Wait()

	pg_rsp, err := countable.NewResultsFromCountWithOptions(pg_opts, total_count)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to create pagination response, %w", err)
	}

	return features, pg_rsp, nil
}

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

	logger := slog.Default()
	logger = logger.With("id", f.ID)

	str_hiers := f.Properties.MustString("wof:hierarchies", "")

	if str_hiers == "" {
		f.Properties["wof:hierarchies"] = make([]map[string]int64, 0)
		return nil
	}

	var hiers []map[string]int64

	err := json.Unmarshal([]byte(str_hiers), &hiers)

	if err != nil {
		return fmt.Errorf("Failed to unmarshal hierarchies, %w", err)
	}

	f.Properties["wof:hierarchies"] = hiers

	if len(hiers) == 0 {
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
	case int64:
		parent_id, _ = f_pid.(int64)
	default:
		parent_id = -1
	}

	label_opts := &hierarchies.AncestorIdsForLabelOptions{
		Hierarchies: hiers,
		Placetype:   str_pt,
		ParentId:    parent_id,
	}

	name_ids := hierarchies.AncestorIdsForLabel(label_opts)

	names_q := fmt.Sprintf("SELECT name, placetype, country from %s WHERE id = ?", g.tableName("records"))

	for _, id := range name_ids {

		var id_name string
		var id_placetype string
		var id_country string

		row := g.db.QueryRowContext(ctx, names_q, id)
		err := row.Scan(&id_name, &id_placetype, &id_country)

		switch {
		case err == db_sql.ErrNoRows:
			continue
		case err != nil:
			logger.Warn("Failed to query ID for name", "name id", id)
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

func (g *SQLGeocoder) uidForVectorRecord(ctx context.Context, record_id int64, model string, language string, tag string) (int64, error) {

	q := fmt.Sprintf("SELECT id FROM %s WHERE record_id = ? AND model = ? AND language = ? AND tag = ?", g.tableName("embeddings_records"))

	row := g.db.QueryRowContext(ctx, q, record_id, model, language, tag)

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
