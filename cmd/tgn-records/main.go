package main

import (
	"archive/zip"
	"context"
	"fmt"
	"io"
	"log"
	_ "os"

	_ "modernc.org/sqlite"

	"github.com/netascode/xmldot"
	"github.com/sfomuseum/go-database/sql"
)

func main() {

	ctx := context.Background()

	db_uri := "sql://sqlite?dsn=tgn_records.db"

	db, err := sql.OpenWithURI(ctx, db_uri)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	_, err = db.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS records (id INTEGER PRIMARY KEY, name TEXT, parent_id INTEGER, placetype)`)

	if err != nil {
		log.Fatal(err)
	}

	// Open the zip file
	reader, err := zip.OpenReader("/Users/asc/Downloads/tgn_xml_0126.zip")
	if err != nil {
		log.Fatalf("failed to open zip: %v", err)
	}
	defer reader.Close()

	// Stream file names
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

		// fmt.Println(string(body))
		// break

		id_rsp := xmldot.GetBytes(body, "Vocabulary.Subject.@Subject_ID")

		if !id_rsp.Exists() {
			continue
		}

		id := id_rsp.Int()

		name_rsp := xmldot.GetBytes(body, "Vocabulary.Subject.Terms.PreferredTerm.Term_Text")
		name := name_rsp.String()

		parent_rsp := xmldot.GetBytes(body, "Vocabulary.Subject.Parent_Relationships.Preferred_Parent.Parent_Subject_ID")
		parent_id := parent_rsp.Int()

		pt_rsp := xmldot.GetBytes(body, "Vocabulary.Subject.Place_Types.Preferred_Place_Type.Place_Type_ID")
		pt := pt_rsp.String()

		/*
			lat_rsp := xmldot.GetBytes(body, "Vocabulary.Subject.Coordinates.Standard.Latitude.Decimal")
			lon_rsp := xmldot.GetBytes(body, "Vocabulary.Subject.Coordinates.Standard.Longitude.Decimal")
			lat := lat_rsp.Float()
			lon := lon_rsp.Float()

			fmt.Println(id, lat, lon)
		*/

		_, err = db.ExecContext(ctx, "INSERT OR REPLACE INTO records (id, name, parent_id, placetype) VALUES (?, ?, ?, ?)", id, name, parent_id, pt)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(id, name, parent_id, pt)
	}
}
