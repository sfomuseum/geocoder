package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"strings"

	_ "github.com/whosonfirst/go-whosonfirst/v4/spatial/pmtiles"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/planar"
	"github.com/sfomuseum/geocoder/coarse"
	"github.com/sfomuseum/geocoder/placeholder"
	"github.com/sfomuseum/go-parquet"
	"github.com/whosonfirst/go-openpois"
	"github.com/whosonfirst/go-openpois/whosonfirst"
	"github.com/whosonfirst/go-reader/v2"
	"github.com/whosonfirst/go-whosonfirst/v4/spatial/database"
)

func main() {

	var geocoder_uri string
	var spatial_database_uri string
	var reader_uri string

	var append_wof bool
	var verbose bool

	flag.StringVar(&geocoder_uri, "geocoder-uri", "sql://sqlite?dsn=:memory:", "A registered sfomuseum/geocoder/coarse.Geocoder URI. Required if -append-whosonfirst-properties=true.")
	flag.StringVar(&spatial_database_uri, "spatial-database-uri", "pmtiles://?tiles=file:///usr/local/data/whosonfirst/whosonfirst-pmtiles&database=whosonfirst-point-in-polygon-z13-20250805&enable-cache=true&zoom=13&layer=whosonfirst", "A registered whosonfirst/go-whosonfirst/v4/spatial/database.SpatialDatabase URI. Required if -append-whosonfirst-properties=true.")
	flag.StringVar(&reader_uri, "reader-uri", "https://data.whosonfirst.org", "A registered whosonfirst/go-reader/v2.Reader URI.")
	flag.BoolVar(&append_wof, "append-whosonfirst-properties", true, "Append Who's On First properties to OpenPOI records before indexing.")
	flag.BoolVar(&verbose, "verbose", false, "Enable verbose (debug) logging.")

	flag.Parse()

	if verbose {
		slog.SetLogLoggerLevel(slog.LevelDebug)
		slog.Debug("Verbose logging enabled")
	}

	ctx := context.Background()
	uris := flag.Args()

	gc, err := coarse.NewGeocoder(ctx, geocoder_uri)

	if err != nil {
		log.Fatal(err)
	}

	defer gc.Close()

	err = gc.PreIndex(ctx)

	if err != nil {
		log.Fatalf("Pre-indexing failed, %v", err)
	}

	var append_opts *whosonfirst.AppendWhosOnFirstPropertiesOptions

	if append_wof {

		wof_r, err := reader.NewReader(ctx, reader_uri)

		if err != nil {
			log.Fatalf("Failed to create WOF reader, %v", err)
		}

		spatial_db, err := database.NewSpatialDatabase(ctx, spatial_database_uri)

		if err != nil {
			log.Fatalf("Failed to create spatial database, %v", err)
		}

		defer spatial_db.Close(ctx)

		append_opts = &whosonfirst.AppendWhosOnFirstPropertiesOptions{
			Database: spatial_db,
			Reader:   wof_r,
		}
	}

	for poi, err := range parquet.Iterate[openpois.Record](ctx, uris...) {

		if err != nil {
			log.Fatalf("Iterator yield an error, %v", err)
		}

		logger := slog.Default()
		logger = logger.With("id", poi.UnifiedID)

		geom, err := poi.ToOrbGeometry()

		if err != nil {
			slog.Warn("Failed to derive geometry", "error", err)
			continue
		}

		centroid, _ := planar.CentroidArea(geom)

		raw := []string{
			poi.Name,
			poi.Brand,
			fmt.Sprintf("osm:id=%s", poi.OsmID),
			poi.OsmName,
			poi.OsmBrand,
			poi.OvertureName,
			poi.OvertureBrand,
		}

		addrs := poi.DistinctAddresses()
		raw = append(raw, addrs...)

		str_raw := strings.Join(raw, " ")

		tokens := make(map[string]map[string][]string)

		tokens["eng"] = map[string][]string{
			"preferred": placeholder.Tokenize(str_raw),
		}

		rec := &coarse.Record{
			Id:        poi.UnifiedID,
			ParentId:  "-1",
			Name:      poi.PrimaryName(),
			Placetype: "venue",
			Country:   "XY",
			Centroid:  &centroid,
			Bounds: []orb.Bound{
				centroid.Bound(),
			},
			IsCurrent:   "1",
			Inception:   "",
			Cessation:   "..",
			Tokens:      tokens,
			Hierarchies: make([]map[string]string, 0),
		}

		if append_wof {

			err := whosonfirst.AppendWhosOnFirstProperties(ctx, poi, append_opts)

			if err != nil {
				logger.Error("Failed to append Who's On First properties", "error", err)
			} else {

				str_hiers := make([]map[string]string, len(poi.WhosOnFirstHierarchies))

				for i, h := range poi.WhosOnFirstHierarchies {

					str_h := make(map[string]string)

					for k, v := range h {
						str_h[k] = fmt.Sprintf("wof:id=%d", v)
					}

					str_hiers[i] = str_h
				}

				rec.Hierarchies = str_hiers
				rec.ParentId = fmt.Sprintf("wof:id=%d", poi.WhosOnFirstParentId)
				rec.Country = poi.WhosOnFirstCountry

				logger.Debug("Add with WOF properties", "parent id", rec.ParentId, "country", rec.Country)
			}
		}

		err = gc.AddRecord(ctx, rec)

		if err != nil {
			log.Fatalf("Failed to index TGN record %s, %v", poi.UnifiedID, err)
		}
	}

	err = gc.PostIndex(ctx)

	if err != nil {
		log.Fatalf("Post-indexing failed, %v", err)
	}

}
