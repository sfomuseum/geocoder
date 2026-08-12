package main

import (
	"archive/zip"
	"context"
	db_sql "database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"os"
	"slices"
	"strings"
	"sync"
	"time"

	_ "modernc.org/sqlite"

	"github.com/netascode/xmldot"
	"github.com/paulmach/orb"
	"github.com/sfomuseum/geocoder/coarse"
	"github.com/sfomuseum/geocoder/placeholder"
	"github.com/sfomuseum/geocoder/x/tgn"
	"github.com/sfomuseum/go-database/sql"
	"github.com/sfomuseum/go-edtf/parser"
	"github.com/sfomuseum/go-flags/flagset"
)

func main() {

	var tgn_data string
	var do_index bool

	var index_db_uri string
	var geocoder_uri string

	var verbose bool

	fs := flagset.NewFlagSet("tgn")

	fs.StringVar(&tgn_data, "tgn-data", "", "The path to the compressed (zip) TGN XML records.")
	fs.BoolVar(&do_index, "do-index", false, "...")

	fs.StringVar(&index_db_uri, "index-db-uri", "sql://sqlite?dsn=tgn_records.db", "...")
	fs.StringVar(&geocoder_uri, "geocoder-uri", "sql://sqlite?dsn=tgn2.db", "...")
	fs.BoolVar(&verbose, "verbose", false, "Enable verbose (debug) logging.")

	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "Index one or more Who's On First data sources in a (coarse) geocoding database.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t%s [options] uri(N) uri(N) uri(N)\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Valid options are:\n")
		fs.PrintDefaults()
	}

	flagset.Parse(fs)

	ctx := context.Background()

	if verbose {
		slog.SetLogLoggerLevel(slog.LevelDebug)
		slog.Debug("Verbose logging enabled")
	}

	//

	records_count := 0
	records_seen := 0

	placetype_map := new(sync.Map)
	country_map := new(sync.Map)

	now := time.Now()
	yyyy := now.Format("2006")

	e_now, err := parser.ParseString(yyyy)

	if err != nil {
		log.Fatal(err)
	}

	// Set up "index" database

	if index_db_uri == "" {

		if !do_index {
			log.Fatal("-do-index flag is false but -index-db-uri flag is empty")
		}

		f, err := os.CreateTemp("", "tgn-index.*.db")

		if err != nil {
			log.Fatalf("Failed to create temp file for index database, %v", err)
		}

		f.Close()

		fname := f.Name()
		defer os.Remove(fname)

		index_db_uri = fmt.Sprintf("sql://sqlite?dsn=%s", fname)
	}

	db, err := sql.OpenWithURI(ctx, index_db_uri)

	if err != nil {
		log.Fatalf("Failed to open indexing database, %v", err)
	}

	defer db.Close()

	//

	if do_index {

		_, err = db.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS records (id INTEGER PRIMARY KEY, name TEXT, parent_id INTEGER, placetype)`)

		if err != nil {
			log.Fatalf("Failed to create indexing database 'records' table, %v", err)
		}

		reader, err := zip.OpenReader(tgn_data)

		if err != nil {
			log.Fatalf("failed to open TGN data, %v", err)
		}

		defer reader.Close()

		for _, f := range reader.File {

			if f.FileInfo().IsDir() {
				continue
			}

			fname := f.Name
			records_count += 1

			r, err := f.Open()

			if err != nil {
				log.Fatalf("Failed to open TGN %s for reading, %v", fname, err)
			}

			defer r.Close()

			body, err := io.ReadAll(r)

			r.Close()

			if err != nil {
				log.Fatalf("Failed to read TGN %s, %v", fname, err)
			}

			id_rsp := xmldot.GetBytes(body, "Vocabulary.Subject.@Subject_ID")

			if !id_rsp.Exists() {
				continue
			}

			id := id_rsp.Int()

			name_rsp := xmldot.GetBytes(body, "Vocabulary.Subject.Terms.Preferred_Term.Term_Text")
			name := name_rsp.String()

			parent_rsp := xmldot.GetBytes(body, "Vocabulary.Subject.Parent_Relationships.Preferred_Parent.Parent_Subject_ID")
			parent_id := parent_rsp.Int()

			pt_rsp := xmldot.GetBytes(body, "Vocabulary.Subject.Place_Types.Preferred_Place_Type.Place_Type_ID")
			pt := pt_rsp.String()

			_, err = db.ExecContext(ctx, "INSERT OR REPLACE INTO records (id, name, parent_id, placetype) VALUES (?, ?, ?, ?)", id, name, parent_id, pt)

			if err != nil {
				log.Fatalf("Failed to add %s to indexing database, %v", fname, err)
			}
		}
	} else {

		q := "SELECT COUNT(id) FROM records"
		row := db.QueryRowContext(ctx, q)

		err := row.Scan(&records_count)

		if err != nil {
			log.Fatalf("Failed to determine record count from indexing database, %v", err)
		}
	}

	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	done_ch := make(chan bool)

	defer func() {
		done_ch <- true
	}()

	go func() {

		for {
			select {
			case <-done_ch:
				return
			case <-ticker.C:
				slog.Info("Processing", "seen", records_seen, "total", records_count)
			}
		}
	}()

	// Set up geocoder database

	gc, err := coarse.NewSQLGeocoder(ctx, geocoder_uri)

	if err != nil {
		log.Fatalf("Failed to create geocoder, %v", err)
	}

	defer gc.Close()

	err = gc.PreIndex(ctx)

	if err != nil {
		log.Fatalf("Pre-indexing failed, %v", err)
	}

	// Read TGN data

	reader, err := zip.OpenReader(tgn_data)

	if err != nil {
		log.Fatalf("Failed to open TGN data, %v", err)
	}

	defer reader.Close()

	for _, f := range reader.File {

		if f.FileInfo().IsDir() {
			continue
		}

		fname := f.Name
		records_seen += 1

		r, err := f.Open()

		if err != nil {
			log.Fatalf("Failed to open TGN %s for reading, %v", fname, err)
		}

		defer r.Close()

		body, err := io.ReadAll(r)

		r.Close()

		if err != nil {
			log.Fatalf("Failed to read TGN %s, %v", fname, err)
		}

		id_rsp := xmldot.GetBytes(body, "Vocabulary.Subject.@Subject_ID")

		if !id_rsp.Exists() {
			continue
		}

		id := id_rsp.Int()

		name_rsp := xmldot.GetBytes(body, "Vocabulary.Subject.Terms.Preferred_Term.Term_Text")
		name := name_rsp.String()

		parent_rsp := xmldot.GetBytes(body, "Vocabulary.Subject.Parent_Relationships.Preferred_Parent.Parent_Subject_ID")
		parent_id := parent_rsp.Int()

		pt_rsp := xmldot.GetBytes(body, "Vocabulary.Subject.Place_Types.Preferred_Place_Type.Place_Type_ID")
		tgn_pt := pt_rsp.String()

		start_rsp := xmldot.GetBytes(body, "Vocabulary.Subject.Place_Types.Preferred_Place_Type.PT_Date.Start_Date")
		end_rsp := xmldot.GetBytes(body, "Vocabulary.Subject.Place_Types.Preferred_Place_Type.PT_Date.End_Date")

		start_date := start_rsp.String()
		end_date := end_rsp.String()

		if start_date != "" {

			edtf_start, err := tgn.TgnToEdtfYear(int(start_rsp.Int()))

			if err == nil {
				start_date = edtf_start
			}
		}

		if end_date != "" {

			edtf_end, err := tgn.TgnToEdtfYear(int(end_rsp.Int()))

			if err == nil {
				end_date = edtf_end
			}
		}

		pt := tgn.TgnToWhosOnFirstPlacetype(tgn_pt)

		lat_rsp := xmldot.GetBytes(body, "Vocabulary.Subject.Coordinates.Standard.Latitude.Decimal")
		lon_rsp := xmldot.GetBytes(body, "Vocabulary.Subject.Coordinates.Standard.Longitude.Decimal")
		lat := lat_rsp.Float()
		lon := lon_rsp.Float()

		centroid := orb.Point([2]float64{lon, lat})

		tokens := make(map[string]map[string][]string)

		tokens["eng"] = map[string][]string{
			"preferred": placeholder.Tokenize(name),
		}

		names_rsp := xmldot.GetBytes(body, "Vocabulary.Subject.Terms.Non-Preferred_Term")

		for _, n_rsp := range names_rsp.Array() {

			t_rsp := n_rsp.Get("Term_Text")
			l_rsp := n_rsp.Get("Term_Languages.Term_Language.Language")

			tgn_lang := l_rsp.String()
			wof_lang, wof_tag := tgn.TgnToWhosOnFirstLanguage(tgn_lang)

			lang_map, ok := tokens[wof_lang]

			if !ok {
				lang_map = make(map[string][]string)
				tokens[wof_lang] = lang_map
			}

			lang_toks, ok := tokens[wof_lang][wof_tag]

			if !ok {
				lang_toks = make([]string, 0)
				tokens[wof_lang][wof_tag] = lang_toks
			}

			for _, t := range placeholder.Tokenize(t_rsp.String()) {

				if !slices.Contains(tokens[wof_lang][wof_tag], t) {
					tokens[wof_lang][wof_tag] = append(tokens[wof_lang][wof_tag], t)
				}
			}

			// slog.Debug("Set tokens", "lang", wof_lang, "tag", wof_tag, "tokens", tokens[wof_lang][wof_tag])
		}

		//

		hier := make(map[string]int64)
		wof_co := ""

		ancestor_id := parent_id

		for ancestor_id != -1 {

			q := "SELECT name, placetype, parent_id FROM records WHERE id = ?"
			row := db.QueryRowContext(ctx, q, ancestor_id)

			var anc_name string
			var anc_pt string
			var anc_parent int64

			err := row.Scan(&anc_name, &anc_pt, &anc_parent)

			if err != nil {

				if err != db_sql.ErrNoRows {
					slog.Warn("Failed to query ancestor", "id", ancestor_id, err)
				}
				break
			}

			wof_pt := tgn.TgnToWhosOnFirstPlacetype(anc_pt)

			if wof_pt == "custom" {

				switch anc_pt {
				case "10003/facet":
					// pass
				default:

					_, ok := placetype_map.LoadOrStore(anc_pt, true)

					if !ok {
						slog.Debug("Unregistered placetype", "id", ancestor_id, "pt", anc_pt, "wof pt", wof_pt)
					}
				}

			} else {
				k := fmt.Sprintf("%s_id", wof_pt)
				hier[k] = ancestor_id

				if wof_pt == "country" {

					wof_co = tgn.TgnToWhosOnFirstCountry(anc_name)

					if wof_co == "XY" {
						_, ok := country_map.LoadOrStore(anc_name, true)

						if !ok {
							slog.Debug("Unregistered country", "country", anc_name)
						}
					}
				}
			}

			ancestor_id = anc_parent
		}

		//

		is_current := "-1"

		if end_date == "" {

			// I suppose we should check that start date is not in the future
			// but I am going to go out on a limb and suggest that is not Getty's
			// jam...

			if start_date != "" {
				is_current = "1"
			}

		} else {

			e, err := parser.ParseString(end_date)

			if err != nil {
				slog.Warn("Failed to parse end date", "date", end_date, "error", err)
			} else {

				before, err := e.Before(e_now)

				if err != nil {
					slog.Warn("Failed to determine before-iness", "now", e_now, "then", e, "error", err)
				} else {

					if before {
						is_current = "0"
					}
				}

				after, err := e.After(e_now)

				if err != nil {
					slog.Warn("Failed to determine after-iness", "now", e_now, "then", e, "error", err)
				} else {

					if after {
						is_current = "1"
					}
				}

			}
		}

		//

		rec := &coarse.Record{
			Id:           id,
			ParentId:     parent_id,
			Name:         name,
			Placetype:    pt,
			Country:      wof_co,
			PlacetypeAlt: strings.Split(tgn_pt, "/"),
			Centroid:     &centroid,
			Bounds: []orb.Bound{
				centroid.Bound(),
			},
			IsCurrent: is_current,
			Inception: start_date,
			Cessation: end_date,
			Tokens:    tokens,
			Hierarchies: []map[string]int64{
				hier,
			},
		}

		dump := false

		if dump {
			enc := json.NewEncoder(os.Stdout)
			enc.Encode(rec)
		}

		err = gc.AddRecord(ctx, rec)

		if err != nil {
			log.Fatalf("Failed to index TGN record %s, %v", fname, err)
		}

	}

	err = gc.Flush(ctx)

	if err != nil {
		log.Fatalf("Failed to flush database, %v", err)
	}

	err = gc.PostIndex(ctx)

	if err != nil {
		log.Fatalf("Post-indexing failed, %v", err)
	}

	placetype_map.Range(func(k, v any) bool {
		fmt.Println(k.(string))
		return true
	})

	country_map.Range(func(k, v any) bool {
		fmt.Println(k.(string))
		return true
	})

}
