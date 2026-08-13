package main

// This is a sample application demonstrating how to use the
// coarse.NewSQLGeocoderWithFS method to create a geocoding
// tool with an embedded (SQLite) database.

import (
	"context"
	"log"

	"github.com/sfomuseum/geocoder/app/coarse/query"
	"github.com/sfomuseum/geocoder/coarse"
	x_fs "github.com/sfomuseum/geocoder/x/fs"
)

func main() {

	ctx := context.Background()

	fs := query.DefaultFlagSet()

	opts, err := query.OptionsFromFlagSet(ctx, fs)

	if err != nil {
		log.Fatalf("Failed to derive options, %v", err)
	}

	db_fs := x_fs.FS
	db_name := "sfomuseum_arch.db"

	gc, err := coarse.NewSQLGeocoderWithFS(ctx, db_fs, db_name)

	if err != nil {
		log.Fatalf("Failed to create geocoder, %v", err)
	}

	opts.Geocoder = gc

	err = query.RunWithOptions(ctx, opts)

	if err != nil {
		log.Fatalf("Failed to query geocoder, %v", err)
	}
}
