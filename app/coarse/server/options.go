package server

import (
	"context"
	"flag"
	"fmt"

	"github.com/sfomuseum/geocoder/coarse"
	"github.com/sfomuseum/go-flags/flagset"
)

type Options struct {
	Geocoder             coarse.Geocoder
	ServerURI            string
	Prefix               string
	QueryTimeout         int
	AllowQueryEmbeddings bool
	PaginationPerPage    int64
	Demo                 bool
	Verbose              bool
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

	opts.ServerURI = server_uri
	opts.Prefix = prefix
	opts.QueryTimeout = query_timeout
	opts.PaginationPerPage = per_page
	opts.Demo = demo
	opts.AllowQueryEmbeddings = allow_query_embeddings

	opts.Verbose = verbose

	return opts, nil
}
