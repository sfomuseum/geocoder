package coarse

import (
	"context"
	"fmt"
	"log/slog"
	"slices"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/dgraph-io/ristretto/v2"
	"github.com/paulmach/orb"
	"github.com/sfomuseum/geocoder/placeholder"
	"github.com/sfomuseum/go-embeddings"
	"github.com/tidwall/gjson"
	"github.com/whosonfirst/go-rfc-5646/tags"
	"github.com/whosonfirst/go-whosonfirst/v4/feature/geometry"
	"github.com/whosonfirst/go-whosonfirst/v4/feature/properties"
)

var langtag_map = new(sync.Map)

type NewWhosOnFirstRecordOptions struct {
	Body           []byte
	Embedder       embeddings.Embedder[float32]
	EmbedderModels []string
	Cache          *ristretto.Cache[string, *VectorEmbeddings]
}

func ParseLangTag(lang_str string) (string, string, error) {

	v, ok := langtag_map.Load(lang_str)

	if ok {

		switch v.(type) {
		case [2]string:
			lt := v.([2]string)
			return lt[0], lt[1], nil
		case error:
			return "", "", v.(error)
		default:
			return "", "", fmt.Errorf("Unexpected cache type for '%s', %v", lang_str, v)
		}
	}

	lang_tag, err := tags.NewLangTag(lang_str)

	var lang string
	var tag string

	if err != nil {

		parts := strings.Split(lang_str, "_x_")

		if len(parts) != 2 {
			err := fmt.Errorf("Failed to parse language tag, %w", err)
			langtag_map.Store(lang_str, err)
			return "", "", err
		}

		lang = parts[0]
		tag = parts[1]

	} else {

		lang = lang_tag.Language()
		tag = lang_tag.PrivateUse()
	}

	langtag_map.Store(lang_str, [2]string{lang, tag})
	return lang, tag, nil
}

// NewWhosOnFirstRecord converts a raw Who's On First GeoJSON document into a Record struct.
// The function parses all required fields, normalises text, tokenises names, collects concordances and
// returns a fully populated Record ready for indexing.
func NewWhosOnFirstRecord(ctx context.Context, opts *NewWhosOnFirstRecordOptions) (*Record, error) {

	logger := slog.Default()

	id, err := properties.Id(opts.Body)

	if err != nil {
		return nil, fmt.Errorf("Failed to derive ID, %w", err)
	}

	logger = logger.With("id", id)

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

		lang, tag, err := ParseLangTag(lang_str)

		if err != nil {
			slog.Warn("Failed to parse language tag", "lang", lang_str, "error", err)
			continue
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

		// Important: Note all the sorting of strings. This is important
		// when generating record hashes.

		sort.Strings(lang_tokens)
		tokens[lang][tag] = lang_tokens
	}

	if opts.Embedder != nil {

		// Set this too high and the cache doesn't really have
		// any effect because everything is happening too fast
		workers := 5

		throttle := make(chan bool, workers)

		for i := 0; i < workers; i++ {
			throttle <- true
		}

		vectors_mu := new(sync.RWMutex)
		wg := new(sync.WaitGroup)

		t1 := time.Now()

		for lang_str, names := range properties.Names(opts.Body) {

			<-throttle

			wg.Go(func() {

				defer func() {
					throttle <- true
				}()

				lang, tag, err := ParseLangTag(lang_str)

				if err != nil {
					slog.Warn("Failed to parse language tag", "lang", lang_str, "error", err)
					return
				}

				// logger.Info("Process name vectors", "lang", lang, "tag", tag)

				names_map := new(sync.Map)

				for _, n := range names {

					_, seen := names_map.LoadOrStore(n, true)

					if seen {
						continue
					}
				}

				names_uniq := make([]string, 0)

				names_map.Range(func(k, v any) bool {
					names_uniq = append(names_uniq, k.(string))
					return true
				})

				if len(names_uniq) == 0 {
					return
				}

				str_names := strings.Join(names_uniq, " ")

				for _, m := range opts.EmbedderModels {

					emb_id := fmt.Sprintf("%d-%s-%s", id, lang, tag)
					emb_key := fmt.Sprintf("%s#%s", m, strings.Replace(str_names, " ", "-", -1))

					var v_emb *VectorEmbeddings

					if opts.Cache != nil {

						v, exists := opts.Cache.Get(emb_key)

						if exists {
							// logger.Info("CACHE HIT", "key", emb_key)
							v_emb = v
						} else {
							// logger.Info("CACHE MISS", "key", emb_key)
						}
					}

					if v_emb == nil {

						emb_req := &embeddings.EmbeddingsRequest{
							Id:    emb_id,
							Model: m,
							Body:  []byte(str_names),
						}

						emb_rsp, err := opts.Embedder.TextEmbeddings(ctx, emb_req)

						if err != nil {
							logger.Warn("Failed to generate embeddings", "model", m, "id", emb_id, "error", err)
							continue
						}

						// logger.Info("Add embeddings", "id", emb_id, "names", str_names)

						e := &Embeddings{
							Language:   lang,
							Tag:        tag,
							Embeddings: emb_rsp.Embeddings(),
						}

						v_emb = &VectorEmbeddings{
							Model:      m,
							Embeddings: []*Embeddings{e},
						}

						if opts.Cache != nil {
							// logger.Info("CACHE SET", "key", emb_key)
							opts.Cache.Set(emb_key, v_emb, 1)
						}
					}

					vectors_mu.Lock()
					vectors = append(vectors, v_emb)
					vectors_mu.Unlock()
				}
			})
		}

		wg.Wait()
		logger.Debug("Processed names", "vectors", len(vectors), "time", time.Since(t1))
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

	count_hiers := len(hiers)

	str_hiers := make([]map[string]string, count_hiers)

	for i, h := range hiers {

		str_h := make(map[string]string)

		for k, id := range h {
			str_h[k] = fmt.Sprintf("wof:id=%d", id)
		}

		str_hiers[i] = str_h
	}

	r := &Record{
		Id:               fmt.Sprintf("wof:id=%d", id),
		ParentId:         fmt.Sprintf("wof:id=%d", parent_id),
		Name:             name,
		Placetype:        pt,
		PlacetypeAlt:     pt_alt,
		Hierarchies:      str_hiers,
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
