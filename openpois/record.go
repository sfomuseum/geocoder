package openpois

import (
	"fmt"
	"strings"
	"time"

	"github.com/adrg/strutil/metrics"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/encoding/wkb"
)

// BoundingBox represents the nested bounding box struct
type BoundingBox struct {
	XMin float64 `parquet:"xmin,optional"`
	YMin float64 `parquet:"ymin,optional"`
	XMax float64 `parquet:"xmax,optional"`
	YMax float64 `parquet:"ymax,optional"`
}

type Record struct {
	UnifiedID                   string      `parquet:"unified_id,optional"`
	Source                      string      `parquet:"source,optional"`
	OsmID                       float64     `parquet:"osm_id,optional"`
	OsmType                     string      `parquet:"osm_type,optional"`
	OvertureID                  string      `parquet:"overture_id,optional"`
	Name                        string      `parquet:"name,optional"`
	Brand                       string      `parquet:"brand,optional"`
	ConfMean                    float64     `parquet:"conf_mean,optional"`
	ConfLower                   float64     `parquet:"conf_lower,optional"`
	ConfUpper                   float64     `parquet:"conf_upper,optional"`
	MatchScore                  float64     `parquet:"match_score,optional"`
	MatchDistanceM              float64     `parquet:"match_distance_m,optional"`
	AddrHousenumber             string      `parquet:"addr_housenumber,optional"`
	AddrStreet                  string      `parquet:"addr_street,optional"`
	AddrUnit                    string      `parquet:"addr_unit,optional"`
	AddrCity                    string      `parquet:"addr_city,optional"`
	AddrState                   string      `parquet:"addr_state,optional"`
	AddrPostcode                string      `parquet:"addr_postcode,optional"`
	AddrCountry                 string      `parquet:"addr_country,optional"`
	Phone                       string      `parquet:"phone,optional"`
	Website                     string      `parquet:"website,optional"`
	OpeningHours                string      `parquet:"opening_hours,optional"`
	Access                      string      `parquet:"access,optional"`
	OvertureSocials             []string    `parquet:"overture_socials,list"`
	OvertureCategoriesAlternate []string    `parquet:"overture_categories_alternate,list"`
	OsmName                     string      `parquet:"osm_name,optional"`
	OvertureName                string      `parquet:"overture_name,optional"`
	OsmBrand                    string      `parquet:"osm_brand,optional"`
	OvertureBrand               string      `parquet:"overture_brand,optional"`
	OsmAddrHousenumber          string      `parquet:"osm_addr_housenumber,optional"`
	OsmAddrStreet               string      `parquet:"osm_addr_street,optional"`
	OsmAddrUnit                 string      `parquet:"osm_addr_unit,optional"`
	OsmAddrCity                 string      `parquet:"osm_addr_city,optional"`
	OsmAddrState                string      `parquet:"osm_addr_state,optional"`
	OsmAddrPostcode             string      `parquet:"osm_addr_postcode,optional"`
	OsmAddrCountry              string      `parquet:"osm_addr_country,optional"`
	OvertureAddrStreet          string      `parquet:"overture_addr_street,optional"`
	OvertureAddrCity            string      `parquet:"overture_addr_city,optional"`
	OvertureAddrState           string      `parquet:"overture_addr_state,optional"`
	OvertureAddrPostcode        string      `parquet:"overture_addr_postcode,optional"`
	OvertureAddrCountry         string      `parquet:"overture_addr_country,optional"`
	OsmPhone                    string      `parquet:"osm_phone,optional"`
	OverturePhones              []string    `parquet:"overture_phones,list"`
	OsmWebsite                  string      `parquet:"osm_website,optional"`
	OvertureWebsites            []string    `parquet:"overture_websites,list"`
	OsmConfMean                 float64     `parquet:"osm_conf_mean,optional"`
	OvertureConfidence          float64     `parquet:"overture_confidence,optional"`
	Geometry                    []byte      `parquet:"geometry,optional"`
	ShadowMatched               bool        `parquet:"shadow_matched,optional"`
	ShadowGhostID               string      `parquet:"shadow_ghost_id,optional"`
	ShadowEventType             string      `parquet:"shadow_event_type,optional"`
	ShadowEventTimestamp        time.Time   `parquet:"shadow_event_timestamp,optional"`
	ShadowScore                 float64     `parquet:"shadow_score,optional"`
	ShadowDistanceM             float64     `parquet:"shadow_distance_m,optional"`
	OriginalConfMean            float64     `parquet:"original_conf_mean,optional"`
	ConfMeanUncalibrated        float64     `parquet:"conf_mean_uncalibrated,optional"`
	CalibrationFlag             string      `parquet:"calibration_flag,optional"`
	Geohash                     string      `parquet:"geohash,optional"`
	IndexLevel0                 int64       `parquet:"__index_level_0__,optional"`
	BBox                        BoundingBox `parquet:"bbox,optional"`
}

func (r *Record) PrimaryName() string {

	candidates := []string{
		r.Name,
		r.OsmName,
		r.OvertureName,
		r.Brand,
		r.OsmBrand,
		r.OvertureBrand,
	}

	for _, name := range candidates {

		name := strings.TrimSpace(name)

		if name != "" {
			return name
		}
	}

	return ""
}

func (r *Record) ToOrbGeometry() (orb.Geometry, error) {

	if len(r.Geometry) == 0 {
		return nil, fmt.Errorf("geometry byte array is empty")
	}

	geom, err := wkb.Unmarshal(r.Geometry)

	if err != nil {
		return nil, fmt.Errorf("failed to decode WKB geometry: %w", err)
	}

	return geom, nil
}

func (r *Record) Addresses() []Address {

	variants := []Address{
		{
			HouseNumber: r.AddrHousenumber,
			Street:      r.AddrStreet,
			Unit:        r.AddrUnit,
			City:        r.AddrCity,
			State:       r.AddrState,
			Postcode:    r.AddrPostcode,
			Country:     r.AddrCountry,
		},
		{
			HouseNumber: r.OsmAddrHousenumber,
			Street:      r.OsmAddrStreet,
			Unit:        r.OsmAddrUnit,
			City:        r.OsmAddrCity,
			State:       r.OsmAddrState,
			Postcode:    r.OsmAddrPostcode,
			Country:     r.OsmAddrCountry,
		},
		{
			HouseNumber: "",
			Street:      r.OvertureAddrStreet,
			Unit:        "",
			City:        r.OvertureAddrCity,
			State:       r.OvertureAddrState,
			Postcode:    r.OvertureAddrPostcode,
			Country:     r.OvertureAddrCountry,
		},
	}

	return variants
}

// DistinctAddresses compiles and phonetically filters available addresses.
func (r *Record) DistinctAddresses() []string {

	jw := metrics.NewJaroWinkler()

	variants := r.Addresses()
	var final []string

	for _, addr := range variants {

		if strings.TrimSpace(addr.Street) == "" && strings.TrimSpace(addr.City) == "" {
			continue
		}

		addr_string := addr.String()

		if addr_string == "" {
			continue
		}

		// Preprocess this candidate address string
		candidate := cleanAndNormalize(addr_string)

		is_duplicate := false

		for _, existing := range final {

			normalized := cleanAndNormalize(existing)

			// Jaro-Winkler yields 0.0 (totally different) to 1.0 (identical)
			// A threshold of 0.90 catches minor layout or formatting differences perfectly.

			if jw.Compare(candidate, normalized) >= 0.90 {
				is_duplicate = true
				break
			}
		}

		if !is_duplicate {
			final = append(final, addr_string)
		}
	}

	return final
}
