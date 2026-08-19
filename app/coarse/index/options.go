package index

import (
	"context"
	"flag"
	"fmt"

	"github.com/sfomuseum/geocoder/coarse"
	"github.com/sfomuseum/go-embeddings"
	"github.com/sfomuseum/go-flags/flagset"
	"github.com/whosonfirst/go-whosonfirst/v4/iterate"
)

type Options struct {
	Geocoder          coarse.Geocoder
	Iterator          iterate.Iterator
	Fresh             bool
	Prune             bool
	IndexJuggling     bool
	ExcludeDeprecated bool
	ExcludeSuperseded bool
	ExcludeFunky      bool
	ExcludeNullIsland bool
	Embedder          embeddings.Embedder[float32]
	EmbeddingsIndex   bool
	EmbeddingsModel   string
	Verbose           bool
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

	iter, err := iterate.NewIterator(ctx, iterator_uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to create iterator, %w", err)
	}

	opts.Iterator = iter

	if embeddings_index {

		embedder, err := embeddings.NewEmbedder32(ctx, embedder_uri)

		if err != nil {
			return nil, fmt.Errorf("Failed to create embedder, %w", err)
		}

		opts.Embedder = embedder
		opts.EmbeddingsModel = embeddings_model
	}

	opts.EmbeddingsIndex = embeddings_index

	opts.Fresh = fresh
	opts.Prune = prune
	opts.IndexJuggling = index_juggling
	opts.ExcludeDeprecated = exclude_deprecated
	opts.ExcludeSuperseded = exclude_superseded
	opts.ExcludeFunky = exclude_funky
	opts.ExcludeNullIsland = exclude_nullisland
	opts.Verbose = verbose

	return opts, nil
}
