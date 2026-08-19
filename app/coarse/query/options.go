package query

import (
	"context"
	"flag"
	"fmt"

	"github.com/sfomuseum/geocoder/coarse"
	"github.com/sfomuseum/go-flags/flagset"
)

type Options struct {
	Geocoder     coarse.Geocoder
	Query        string
	VectorSearch bool
	Lang         string
	LangTag      string
	Placetypes   []string
	Countries    []string
	BelongsTo    []int64
	Bounds       string
	DateStarts   string
	DateEnds     string
	Page         int64
	PerPage      int64
	QueryTimeout int
	Mode         string
	Verbose      bool
}

func OptionsFromFlagSet(ctx context.Context, fs *flag.FlagSet) (*Options, error) {

	flagset.Parse(fs)

	err := flagset.SetFlagsFromEnvVars(fs, "GEOCODER")

	if err != nil {
		return nil, fmt.Errorf("Failed to set flags from environment variables, %w", err)
	}

	opts := new(Options)

	gc, err := coarse.NewGeocoder(ctx, geocoder_uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to create geocoder, %w", err)
	}

	opts.Geocoder = gc

	opts.VectorSearch = vector_search
	opts.Placetypes = placetypes
	opts.Query = query
	opts.Countries = countries
	opts.BelongsTo = belongsto
	opts.Bounds = str_bounds
	opts.DateStarts = date_starts
	opts.DateEnds = date_ends
	opts.Page = page
	opts.PerPage = per_page
	opts.QueryTimeout = query_timeout
	opts.Mode = mode
	opts.Verbose = verbose

	return opts, nil
}
