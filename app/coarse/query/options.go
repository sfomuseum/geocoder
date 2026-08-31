package query

import (
	"context"
	"flag"
	"fmt"

	"github.com/sfomuseum/geocoder/coarse"
	"github.com/sfomuseum/go-embeddings"
	"github.com/sfomuseum/go-flags/flagset"
)

type Options struct {
	Geocoder         coarse.Geocoder
	Embedder         embeddings.Embedder[float32]
	Query            string
	EmbeddingsSearch bool
	EmbeddingsModel  string
	Lang             string
	LangTag          string
	Placetypes       []string
	Countries        []string
	BelongsTo        []string
	Bounds           string
	DateStarts       string
	DateEnds         string
	Source           string
	Page             int64
	PerPage          int64
	QueryTimeout     int
	Mode             string
	Verbose          bool
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

	if embeddings_search {

		embedder, err := embeddings.NewEmbedder32(ctx, embedder_uri)

		if err != nil {
			return nil, fmt.Errorf("Failed to create embedder, %w", err)
		}

		opts.Embedder = embedder
		opts.EmbeddingsModel = embeddings_model
	}

	opts.EmbeddingsSearch = embeddings_search
	opts.Placetypes = placetypes
	opts.Query = query
	opts.Countries = countries
	opts.BelongsTo = belongsto
	opts.Bounds = str_bounds
	opts.DateStarts = date_starts
	opts.DateEnds = date_ends
	opts.Source = source
	opts.Page = page
	opts.PerPage = per_page
	opts.QueryTimeout = query_timeout
	opts.Mode = mode
	opts.Verbose = verbose

	return opts, nil
}
