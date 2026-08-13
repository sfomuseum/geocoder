package coarse

import (
	"context"
	"fmt"
	"log/slog"
	"slices"
	"sort"
	"strings"

	"github.com/paulmach/orb"
	"github.com/sfomuseum/geocoder/placeholder"
	"github.com/sfomuseum/go-embeddings"
	"github.com/tidwall/gjson"
	"github.com/whosonfirst/go-rfc-5646/tags"
	"github.com/whosonfirst/go-whosonfirst/v4/feature/geometry"
	"github.com/whosonfirst/go-whosonfirst/v4/feature/properties"
)

type NewWhosOnFirstRecordOptions struct {
	Body           []byte
	Embedder       embeddings.Embedder[float32]
	EmbedderModels []string
}

// NewWhosOnFirstRecord converts a raw Who's On First GeoJSON document
// into a Record struct.  The function parses all required fields,
// normalises text, tokenises names, collects concordances and
// returns a fully populated Record ready for indexing.
func NewWhosOnFirstRecord(ctx context.Context, opts *NewWhosOnFirstRecordOptions) (*Record, error) {

	// Important: Note all the sorting of strings. This is important
	// when generating record hashes.

	id, err := properties.Id(opts.Body)

	if err != nil {
		return nil, fmt.Errorf("Failed to derive ID, %w", err)
	}

	parent_id, err := properties.ParentId(opts.Body)

	if err != nil {
		return nil, fmt.Errorf("Failed to derive parent ID, %w", err)
	}

	name, err := properties.Name(opts.Body)

	if err != nil {
		return nil, fmt.Errorf("Failed to derive name, %w", err)
	}

	pt, err := properties.Placetype(opts.Body)

	if err != nil {
		return nil, fmt.Errorf("Failed to derive placetype, %w", err)
	}

	is_current, err := properties.IsCurrent(opts.Body)

	if err != nil {
		return nil, fmt.Errorf("Failed to derive is current, %w", err)
	}

	centroid, _, err := properties.Centroid(opts.Body)

	if err != nil {
		return nil, fmt.Errorf("Failed to derive centroid, %w", err)
	}

	// START OF put me in go-whosonfirst/v4/feature/geometry

	bounds := make([]orb.Bound, 0)

	geom, err := geometry.Geometry(opts.Body)

	if err != nil {
		return nil, fmt.Errorf("Failed to derive geometry, %w", err)
	}

	orb_geom := geom.Geometry()

	switch geom.Type {
	case "MultiPolygon":

		for _, p := range orb_geom.(orb.MultiPolygon) {
			bounds = append(bounds, p.Bound())
		}

		// Are we really going to worry about MultiLineStrings?
		// That will be a fun (or at least) interesting day when
		// that's a problem...

	default:

		bounds = []orb.Bound{
			orb_geom.Bound(),
		}
	}

	// END OF put me in go-whosonfirst/v4/feature/geometry

	inception := properties.Inception(opts.Body)
	cessation := properties.Cessation(opts.Body)

	inception = ensureValidEDTF(inception)
	cessation = ensureValidEDTF(cessation)

	// START OF put me in go-whosonfirst/v4/feature/properties

	rank_rsp := gjson.GetBytes(opts.Body, "properties.wof:population_rank")
	pop_rank := rank_rsp.Int()

	alt_rsp := gjson.GetBytes(opts.Body, "properties.wof:placetype_alt")
	alt_count := len(alt_rsp.Array())

	pt_alt := make([]string, alt_count)

	for i, pt := range alt_rsp.Array() {
		pt_alt[i] = pt.String()
	}

	// END OF put me in go-whosonfirst/v4/feature/properties

	co := properties.Country(opts.Body)
	hiers := properties.Hierarchies(opts.Body)

	tokens := make(map[string]map[string][]string)
	vectors := make([]*VectorEmbeddings, 0)

	for lang_str, names := range properties.Names(opts.Body) {

		lang_tag, err := tags.NewLangTag(lang_str)

		var lang string
		var tag string

		if err != nil {

			parts := strings.Split(lang_str, "_x_")

			if len(parts) != 2 {
				slog.Warn("Failed to parse language tag", "lang", lang_str, "error", err)
				continue
			}

			lang = parts[0]
			tag = parts[1]

		} else {

			lang = lang_tag.Language()
			tag = lang_tag.PrivateUse()
		}

		lang_tokens := make([]string, 0)

		for _, n := range names {

			// To do: Look up ancestors based on language...

			for _, t := range placeholder.Tokenize(n) {

				if !slices.Contains(lang_tokens, t) {
					lang_tokens = append(lang_tokens, t)
				}
			}

		}

		_, ok := tokens[lang]

		if !ok {
			tokens[lang] = make(map[string][]string)
		}

		sort.Strings(lang_tokens)
		tokens[lang][tag] = lang_tokens

		//

		if opts.Embedder != nil {

			for _, m := range opts.EmbedderModels {

				name_embeddings := make([]*Embeddings, len(name))

				for idx, n := range names {

					emb_req := &embeddings.EmbeddingsRequest{
						Id:    n,
						Model: m,
						Body:  []byte(n),
					}

					emb_rsp, err := opts.Embedder.TextEmbeddings(ctx, emb_req)

					if err != nil {
						slog.Warn("Failed to generate embeddings", "model", m, "name", n)
						continue
					}

					e := &Embeddings{
						Language:   lang,
						Tag:        tag,
						Embeddings: emb_rsp.Embeddings(),
					}

					name_embeddings[idx] = e
				}

				v := &VectorEmbeddings{
					Model:      m,
					Embeddings: name_embeddings,
				}

				vectors = append(vectors, v)
			}
		}
	}

	// Add concordances as eng_x_concordance

	concordances := make([]string, 0)

	for k, v := range properties.Concordances(opts.Body) {
		// To do: Support wildcards
		k = strings.ReplaceAll(k, ":", "_")
		concordances = append(concordances, fmt.Sprintf("%s__%v", k, v))
	}

	if len(concordances) > 0 {

		_, ok := tokens["eng"]

		if !ok {
			tokens["eng"] = make(map[string][]string)
		}

		sort.Strings(concordances)
		tokens["eng"]["concordances"] = concordances
	}

	sort.Strings(pt_alt)

	r := &Record{
		Id:               id,
		ParentId:         parent_id,
		Name:             name,
		Placetype:        pt,
		PlacetypeAlt:     pt_alt,
		Hierarchies:      hiers,
		Country:          co,
		Inception:        inception,
		Cessation:        cessation,
		IsCurrent:        is_current.StringFlag(),
		PopulationRank:   pop_rank,
		Centroid:         centroid,
		Bounds:           bounds,
		Tokens:           tokens,
		VectorEmbeddings: vectors,
	}

	return r, nil
}
