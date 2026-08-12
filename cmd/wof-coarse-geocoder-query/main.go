package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"os"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/aaronland/go-pagination/countable"
	"github.com/paulmach/orb/geojson"
	"github.com/sfomuseum/geocoder"
	"github.com/sfomuseum/geocoder/coarse"
	"github.com/sfomuseum/go-edtf/unix"
	"github.com/sfomuseum/go-flags/flagset"
	"github.com/sfomuseum/go-flags/multi"
)

func main() {

	var geocoder_uri string
	var query string
	var lang string
	var lang_tag string

	var placetypes multi.MultiString
	var countries multi.MultiString
	var belongsto multi.MultiInt64
	var str_bounds string
	var date_starts string
	var date_ends string

	var page int64
	var per_page int64

	var query_timeout int

	var mode string
	var verbose bool

	fs := flagset.NewFlagSet("query")

	fs.StringVar(&geocoder_uri, "geocoder-uri", "", "A registered sfomuseum/geocoder/coarse.Geocoder URI.")
	fs.StringVar(&query, "query", "", "The term to query for. Required.")
	fs.Var(&placetypes, "placetype", "Zero or more placetypes to filter results by.")
	fs.Var(&countries, "country", "Zero or more 2-letter country codes to filter results by.")
	fs.Var(&belongsto, "belongs-to", "Zero or more Who's On First ancestor IDs to filter results by.")
	fs.StringVar(&lang, "lang", "", "An optional (3-letter) language code to filter results by,")
	fs.StringVar(&lang_tag, "tag", "", "An option WOF language tag to filter results by.")
	fs.StringVar(&str_bounds, "bounds", "", "Optional bounding box (in the form of 'minx,miny,maxx,mayx') to filter results by.")
	fs.StringVar(&date_starts, "date-starts", "", "Optional ETDF starting date string to filter results by.")
	fs.StringVar(&date_ends, "date-ends", "", "Optional ETDF ending date string to filter results by.")

	fs.IntVar(&query_timeout, "query-timeout", 5, "The maximum allowable time in seconds for a query to complete.")

	fs.Int64Var(&page, "page", 1, "The specific page number to query for paginated result sets.")
	fs.Int64Var(&per_page, "per-page", 100, "The number of results to include for paginated result sets.")

	fs.StringVar(&mode, "mode", "tab", "Output mode for results. Valid options are: geojson, tab.")
	fs.BoolVar(&verbose, "verbose", false, "Enable verbose (debug) logging.")

	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "Query a Who's On First (coarse) geocoding database.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t%s [options]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Valid options are:\n")
		fs.PrintDefaults()
	}

	flagset.Parse(fs)

	if verbose {
		slog.SetLogLoggerLevel(slog.LevelDebug)
		slog.Debug("Verbose logging enabled")
	}

	ctx := context.Background()

	gc, err := coarse.NewSQLGeocoder(ctx, geocoder_uri)

	if err != nil {
		log.Fatalf("Failed to create geocoder, %v", err)
	}

	defer gc.Close()

	if query == "" {
		log.Fatalf("Missing query.")
	}

	req := &coarse.QueryRequest{
		Query: query,
	}

	if len(placetypes) > 0 {
		req.Placetype = placetypes
	}

	if len(countries) > 0 {
		req.Country = countries
	}

	if len(belongsto) > 0 {
		req.BelongsTo = belongsto
	}

	if lang != "" {
		req.Lang = lang
	}

	if lang_tag != "" {
		req.Tag = lang_tag
	}

	if str_bounds != "" {

		bounds, err := geocoder.StringBoundsToOrbBounds(str_bounds)

		if err != nil {
			log.Fatalf("Failed to create bounds, %v", err)
		}

		req.Bounds = bounds
	}

	if date_starts != "" {

		ok, ranges, err := unix.DeriveRanges(date_starts)

		if err != nil {
			log.Fatalf("Failed to derive ranges for starts date, %v", err)
		}

		if !ok {
			log.Fatalf("Unable to derive a range for start date.")
		}

		req.DateStarts = ranges
	}

	if date_ends != "" {

		ok, ranges, err := unix.DeriveRanges(date_ends)

		if err != nil {
			log.Fatalf("Failed to derive ranges for ends date, %v", err)
		}

		if !ok {
			log.Fatalf("Unable to derive a range for start date.")
		}

		req.DateEnds = ranges
	}

	ctx, cancel := context.WithTimeout(ctx, time.Duration(query_timeout)*time.Second)
	defer cancel()

	// To do: derive options from URI constructor...
	pg_opts, err := countable.NewCountableOptions()

	if err != nil {
		log.Fatalf("Failed to create pagination options, %v", err)
	}

	pg_opts.PerPage(per_page)
	pg_opts.Pointer(page)

	rsp, pg_rsp, err := gc.Query(ctx, req, pg_opts)

	if err != nil {
		log.Fatalf("Failed to query database, %v", err)
	}

	slog.Info("Query results", "total", pg_rsp.Total(), "page", pg_rsp.Page(), "pages", pg_rsp.Pages())

	switch mode {
	case "geojson":

		fc := geojson.NewFeatureCollection()
		fc.Features = rsp

		enc := json.NewEncoder(os.Stdout)
		err = enc.Encode(fc)

		if err != nil {
			log.Fatalf("Failed to encode response, %v", err)
		}

	default:

		tw := new(tabwriter.Writer)
		tw.Init(os.Stdout, 1, 8, 1, '\t', 0)

		cols := []string{
			"id",
			"name",
			"placetype",
			"is current",
			"inception",
			"cessation",
			"label",
		}

		fmt.Println("")
		fmt.Fprintln(tw, strings.Join(cols, "\t"))

		for _, f := range rsp {

			vals := []string{
				fmt.Sprintf("%d", f.ID),
				f.Properties["wof:name"].(string),
				f.Properties["wof:placetype"].(string),
				fmt.Sprintf("%v", f.Properties["mz:is_current"]),
				f.Properties["edtf:inception"].(string),
				f.Properties["edtf:cessation"].(string),
				f.Properties["wof:label"].(string),
			}

			fmt.Fprintln(tw, strings.Join(vals, "\t"))
		}

		tw.Flush()
	}
}
