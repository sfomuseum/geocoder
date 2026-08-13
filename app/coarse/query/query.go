package query

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
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
)

func Run(ctx context.Context) error {
	fs := DefaultFlagSet()
	return RunWithFlagSet(ctx, fs)
}

func RunWithFlagSet(ctx context.Context, fs *flag.FlagSet) error {

	opts, err := OptionsFromFlagSet(ctx, fs)

	if err != nil {
		return err
	}

	return RunWithOptions(ctx, opts)
}

func RunWithOptions(ctx context.Context, opts *Options) error {

	if opts.Verbose {
		slog.SetLogLoggerLevel(slog.LevelDebug)
		slog.Debug("Verbose logging enabled")
	}

	if opts.Geocoder == nil {
		return fmt.Errorf("Missing geocoder")
	}

	defer opts.Geocoder.Close()

	if opts.Query == "" {
		return fmt.Errorf("Missing query.")
	}

	req := &coarse.QueryRequest{
		Query: opts.Query,
	}

	if len(opts.Placetypes) > 0 {
		req.Placetype = opts.Placetypes
	}

	if len(opts.Countries) > 0 {
		req.Country = opts.Countries
	}

	if len(opts.BelongsTo) > 0 {
		req.BelongsTo = opts.BelongsTo
	}

	if opts.Lang != "" {
		req.Lang = opts.Lang
	}

	if opts.LangTag != "" {
		req.Tag = opts.LangTag
	}

	if opts.Bounds != "" {

		bounds, err := geocoder.StringBoundsToOrbBounds(opts.Bounds)

		if err != nil {
			return fmt.Errorf("Failed to create bounds, %w", err)
		}

		req.Bounds = bounds
	}

	if opts.DateStarts != "" {

		ok, ranges, err := unix.DeriveRanges(opts.DateStarts)

		if err != nil {
			return fmt.Errorf("Failed to derive ranges for starts date, %w", err)
		}

		if !ok {
			return fmt.Errorf("Unable to derive a range for start date.")
		}

		req.DateStarts = ranges
	}

	if opts.DateEnds != "" {

		ok, ranges, err := unix.DeriveRanges(opts.DateEnds)

		if err != nil {
			return fmt.Errorf("Failed to derive ranges for ends date, %w", err)
		}

		if !ok {
			return fmt.Errorf("Unable to derive a range for start date.")
		}

		req.DateEnds = ranges
	}

	ctx, cancel := context.WithTimeout(ctx, time.Duration(opts.QueryTimeout)*time.Second)
	defer cancel()

	// To do: derive options from URI constructor...
	pg_opts, err := countable.NewCountableOptions()

	if err != nil {
		return fmt.Errorf("Failed to create pagination options, %w", err)
	}

	pg_opts.PerPage(opts.PerPage)
	pg_opts.Pointer(opts.Page)

	rsp, pg_rsp, err := opts.Geocoder.Query(ctx, req, pg_opts)

	if err != nil {
		return fmt.Errorf("Failed to query database, %w", err)
	}

	slog.Info("Query results", "total", pg_rsp.Total(), "page", pg_rsp.Page(), "pages", pg_rsp.Pages())

	switch opts.Mode {
	case "geojson":

		fc := geojson.NewFeatureCollection()
		fc.Features = rsp

		enc := json.NewEncoder(os.Stdout)
		err = enc.Encode(fc)

		if err != nil {
			return fmt.Errorf("Failed to encode response, %w", err)
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

	return nil
}
