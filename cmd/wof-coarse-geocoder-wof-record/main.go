package main

// Helper tool for generate `coarse.Record` JSON data from Who's On First records
// go run cmd/wof-coarse-geocoder-wof-record/main.go ./fixtures/ca.geojson > ./fixtures/ca-coarse.json

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/sfomuseum/geocoder/whosonfirst"
)

func main() {

	flag.Parse()

	ctx := context.Background()

	for _, path := range flag.Args() {

		body, err := os.ReadFile(path)

		if err != nil {
			log.Fatalf("Failed to read '%s', %v", path, err)
		}

		wof_opts := &whosonfirst.NewCoarseGeocoderRecordOptions{
			Body: body,
		}

		rec, err := whosonfirst.NewCoarseGeocoderRecord(ctx, wof_opts)

		if err != nil {
			log.Fatalf("Failed to create new record, %v", err)
		}

		enc := json.NewEncoder(os.Stdout)
		err = enc.Encode(rec)

		if err != nil {
			log.Fatalf("Failed to encode record, %v", err)
		}
	}
}
