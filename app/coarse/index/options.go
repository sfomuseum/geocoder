package index

import (
	"context"
	"flag"
	"fmt"
	"github.com/aaronland/go-json-query"
	"log/slog"
	"net/url"
	"slices"

	"github.com/sfomuseum/geocoder/coarse"
	"github.com/sfomuseum/go-embeddings"
	"github.com/sfomuseum/go-flags/flagset"
	"github.com/whosonfirst/go-whosonfirst/v4/iterate"
)

type Options struct {
	Geocoder          coarse.Geocoder
	Iterator          iterate.Iterator
	IteratorSources   []string
	Fresh             bool
	Prune             bool
	Offset int64
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

	if exclude_deprecated || exclude_superseded || exclude_funky {

		exclude_deprecated_path := "properties.edtf:deprecated=.*"
		exclude_superseded_path := "properties.wof:superseded_by=.*"
		exclude_funky_path := "propertiees.mz:is_funky=1"

		u, err := url.Parse(iterator_uri)

		if err != nil {
			return nil, fmt.Errorf("Failed to parse iterator URI, %w", err)
		}

		q := u.Query()

		to_exclude := make([]string, 0)

		if exclude_deprecated {
			to_exclude = append(to_exclude, exclude_deprecated_path)
		}

		if exclude_superseded {
			to_exclude = append(to_exclude, exclude_superseded_path)
		}

		if exclude_funky {
			to_exclude = append(to_exclude, exclude_funky_path)
		}

		_, ok := q["exclude"]

		if ok {

			for _, v := range q["exclude"] {

				if !slices.Contains(to_exclude, v) {
					to_exclude = append(to_exclude, v)
				}
			}
		}

		q["exclude"] = to_exclude

		q.Set("exclude_mode", query.QUERYSET_MODE_ANY)
		u.RawQuery = q.Encode()
		iterator_uri = u.String()

		slog.Info("Rewrote iterator URI", "uri", iterator_uri)
	}

	iter, err := iterate.NewIterator(ctx, iterator_uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to create iterator, %w", err)
	}

	opts.Iterator = iter
	opts.IteratorSources = fs.Args()

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
	opts.Offset = offset
	opts.IndexJuggling = index_juggling
	opts.ExcludeDeprecated = exclude_deprecated
	opts.ExcludeSuperseded = exclude_superseded
	opts.ExcludeFunky = exclude_funky
	opts.ExcludeNullIsland = exclude_nullisland
	opts.Verbose = verbose

	return opts, nil
}
