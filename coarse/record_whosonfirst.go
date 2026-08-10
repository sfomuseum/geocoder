package coarse

import (
	"context"
	"fmt"
	"log/slog"
	"slices"
	"sort"
	"strings"

	"github.com/paulmach/orb"
	"github.com/tidwall/gjson"
	"github.com/sfomuseum/geocoder/placeholder"
	"github.com/whosonfirst/go-rfc-5646/tags"
	"github.com/whosonfirst/go-whosonfirst/v4/feature/geometry"
	"github.com/whosonfirst/go-whosonfirst/v4/feature/properties"
)

func NewWhosOnFirstRecord(ctx context.Context, body []byte) (*Record, error) {

	// Important: Note all the sorting of strings. This is important
	// when generating record hashes.

	id, err := properties.Id(body)

	if err != nil {
		return nil, fmt.Errorf("Failed to derive ID, %w", err)
	}

	parent_id, err := properties.ParentId(body)

	if err != nil {
		return nil, fmt.Errorf("Failed to derive parent ID, %w", err)
	}

	name, err := properties.Name(body)

	if err != nil {
		return nil, fmt.Errorf("Failed to derive name, %w", err)
	}

	pt, err := properties.Placetype(body)

	if err != nil {
		return nil, fmt.Errorf("Failed to derive placetype, %w", err)
	}

	is_current, err := properties.IsCurrent(body)

	if err != nil {
		return nil, fmt.Errorf("Failed to derive is current, %w", err)
	}

	centroid, _, err := properties.Centroid(body)

	if err != nil {
		return nil, fmt.Errorf("Failed to derive centroid, %w", err)
	}

	// START OF put me in go-whosonfirst/v4/feature/geometry

	bounds := make([]orb.Bound, 0)

	geom, err := geometry.Geometry(body)

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

	inception := properties.Inception(body)
	cessation := properties.Cessation(body)

	inception = ensureValidEDTF(inception)
	cessation = ensureValidEDTF(cessation)

	// START OF put me in go-whosonfirst/v4/feature/properties

	rank_rsp := gjson.GetBytes(body, "properties.wof:population_rank")
	pop_rank := rank_rsp.Int()

	alt_rsp := gjson.GetBytes(body, "properties.wof:placetype_alt")
	alt_count := len(alt_rsp.Array())

	pt_alt := make([]string, alt_count)

	for i, pt := range alt_rsp.Array() {
		pt_alt[i] = pt.String()
	}

	// END OF put me in go-whosonfirst/v4/feature/properties

	co := properties.Country(body)
	hiers := properties.Hierarchies(body)

	tokens := make(map[string]map[string][]string)

	for lang_str, names := range properties.Names(body) {

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
	}

	// Add concordances as eng_x_concordance

	concordances := make([]string, 0)

	for k, v := range properties.Concordances(body) {
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
		Id:             id,
		ParentId:       parent_id,
		Name:           name,
		Placetype:      pt,
		PlacetypeAlt:   pt_alt,
		Hierarchies:    hiers,
		Country:        co,
		Inception:      inception,
		Cessation:      cessation,
		IsCurrent:      is_current.StringFlag(),
		PopulationRank: pop_rank,
		Centroid:       centroid,
		Bounds:         bounds,
		Tokens:         tokens,
	}

	return r, nil
}
