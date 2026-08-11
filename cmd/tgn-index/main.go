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
)

func main() {

	ctx := context.Background()

	db_uri := "sql://sqlite?dsn=tgn_records.db"

	db, err := sql.OpenWithURI(ctx, db_uri)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	geocoder_uri := "sql://sqlite?dsn=tgn2.db"
	gc, err := coarse.NewSQLGeocoder(ctx, geocoder_uri)

	if err != nil {
		log.Fatalf("Failed to create geocoder, %v", err)
	}

	defer gc.Close()

	err = gc.PreIndex(ctx)

	if err != nil {
		log.Fatalf("Pre-indexing failed, %v", err)
	}

	reader, err := zip.OpenReader("/Users/asc/Downloads/tgn_xml_0126.zip")

	if err != nil {
		log.Fatalf("failed to open zip: %v", err)
	}
	defer reader.Close()

	country_map := new(sync.Map)
	placetype_map := new(sync.Map)

	now := time.Now()
	yyyy := now.Format("2006")

	e_now, err := parser.ParseString(yyyy)

	if err != nil {
		log.Fatal(err)
	}

	for _, f := range reader.File {

		// fmt.Println(f.Name)

		if f.FileInfo().IsDir() {
			continue
		}

		r, err := f.Open()

		if err != nil {
			log.Fatal(err)
		}

		defer r.Close()

		body, err := io.ReadAll(r)

		r.Close()

		if err != nil {
			log.Fatal(err)
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

		// fmt.Println(id, lat, lon)

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

			slog.Debug("Set tokens", "lang", wof_lang, "tag", wof_tag, "tokens", tokens[wof_lang][wof_tag])
		}

		//

		hier := make(map[string]int64)

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
						slog.Info("Unregistered placetype", "id", ancestor_id, "pt", anc_pt, "wof pt", wof_pt)
					}
				}

			} else {
				k := fmt.Sprintf("%s_id", wof_pt)
				hier[k] = ancestor_id
				// slog.Info("hier", "pt", wof_pt, "id", ancestor_id)

				if wof_pt == "country" {
					country_map.Store(ancestor_id, anc_name)
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
			}
		}

		//

		rec := &coarse.Record{
			Id:           id,
			ParentId:     parent_id,
			Name:         name,
			Placetype:    pt,
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
			log.Fatalf("Failed to index records, %v", err)
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

	country_map.Range(func(k, v any) bool {
		fmt.Println(v.(string))
		return true
	})

}
