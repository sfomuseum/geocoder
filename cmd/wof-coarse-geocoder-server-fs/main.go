package main

// This is a sample application demonstrating how to use the
// coarse.NewSQLGeocoderWithFS method to create a geocoding
// tool with an embedded (SQLite) database. Note that it the
// 'sfomuseum_wof.db' database referenced in the code is NOT
// bundled with this package. You will need to do that yourself.
// For example:
//
// $> ./bin/wof-coarse-geocoder-index -geocoder-uri 'sql://sqlite?dsn=x/fs/sfomuseum_wof.db' /usr/local/data/sfomuseum-data-whosonfirst/

import (
	"context"
	"log"

	"github.com/sfomuseum/geocoder/app/coarse/server"
	"github.com/sfomuseum/geocoder/coarse"
	x_fs "github.com/sfomuseum/geocoder/x/fs"
)

func main() {

	ctx := context.Background()

	fs := server.DefaultFlagSet()

	opts, err := server.OptionsFromFlagSet(ctx, fs)

	if err != nil {
		log.Fatalf("Failed to derive options, %v")
	}

	db_fs := x_fs.FS
	db_name := "sfomuseum_wof.db"

	gc, err := coarse.NewSQLGeocoderWithFS(ctx, db_fs, db_name)

	if err != nil {
		log.Fatalf("Failed to create geocoder, %v", err)
	}

	opts.Geocoder = gc

	err = server.RunWithOptions(ctx, opts)

	if err != nil {
		log.Fatalf("Failed to run geocoder server, %v", err)
	}
}
