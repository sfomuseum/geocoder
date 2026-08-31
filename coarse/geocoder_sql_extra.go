package coarse

import (
	"context"
	db_sql "database/sql"
	"encoding/json"
	"fmt"
	"log/slog"
	"strings"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
	"github.com/whosonfirst/go-whosonfirst/v4/hierarchies"
)

func (g *SQLGeocoder) assignExtra(ctx context.Context, f *geojson.Feature) error {

	logger := slog.Default()
	logger = logger.With("id", f.ID)

	id := f.Properties["geocoder:id"]
	parent_id := f.Properties["geocoder:parent_id"]

	str_id, err := g.retrieveStringIdentifier(ctx, id.(int64))

	if err != nil {
		slog.Error("Failed to retrieve string identifier", "id", id, "error", err)
	} else {
		f.Properties["geocoder:id"] = str_id
	}

	str_parent, err := g.retrieveStringIdentifier(ctx, parent_id.(int64))

	if err != nil {
		slog.Error("Failed to retrieve string identifier", "id", parent_id, "error", err)
	} else {
		f.Properties["geocoder:parent_id"] = str_parent
	}

	err = g.assignHierarchiesAndLabel(ctx, f)

	if err != nil {
		logger.Warn("Failed to assign label", "error", err)
	}

	err = g.assignPlacetypeAlt(ctx, f)

	if err != nil {
		logger.Warn("Failed to assign placetype alt", "error", err)
	}

	err = g.assignBBox(ctx, f)

	if err != nil {
		logger.Warn("Failed to assign bbox", "error", err)
	}

	return nil
}

func (g *SQLGeocoder) assignBBox(ctx context.Context, f *geojson.Feature) error {

	bounds_q := fmt.Sprintf("SELECT MIN(b.minx), MIN(b.miny), MAX(b.maxx), MAX(b.maxy) FROM %s b, %s r  WHERE b.record_id = r.id AND r.id = ?", g.tableName("bounds"), g.tableName("records"))

	bounds_row := g.db.QueryRowContext(ctx, bounds_q, f.ID)

	var minx float64
	var miny float64
	var maxx float64
	var maxy float64

	err := bounds_row.Scan(&minx, &miny, &maxx, &maxy)

	if err != nil {
		return fmt.Errorf("Failed to query bounds, %w", err)
	}

	bounds := orb.Bound{
		Min: orb.Point([2]float64{minx, miny}),
		Max: orb.Point([2]float64{maxx, maxy}),
	}

	f.BBox = geojson.NewBBox(bounds)
	return nil
}

func (g *SQLGeocoder) assignHierarchiesAndLabel(ctx context.Context, f *geojson.Feature) error {

	// To do: Update to account for string-based IDs and convert back to WOF int-based IDs
	// where applicable.

	logger := slog.Default()
	logger = logger.With("id", f.ID)

	f_hiers := f.Properties.MustString("geocoder:hierarchies", "")

	if f_hiers == "" {
		f.Properties["geocoder:hierarchies"] = make([]map[string]int64, 0)
		return nil
	}

	var str_hiers []map[string]string

	err := json.Unmarshal([]byte(f_hiers), &str_hiers)

	if err != nil {
		return fmt.Errorf("Failed to unmarshal hierarchies, %w", err)
	}

	f.Properties["geocoder:hierarchies"] = str_hiers

	if len(str_hiers) == 0 {
		return nil
	}

	name := f.Properties.MustString("geocoder:name")

	labels := []string{
		name,
	}

	str_pt := f.Properties.MustString("wof:placetype", "")
	parent_id, ok := f.Properties["geocoder:parent_id"]

	str_parent := ""

	if ok {

		// This is unfortunate but we need to trap instances where parent_id
		// has already been reassigned as a string value (this happens in
		// assignExtra)

		switch parent_id.(type) {
		case int64:

			p, err := g.retrieveStringIdentifier(ctx, parent_id.(int64))

			if err != nil {
				slog.Warn("Failed to rerieve string identifier for parent", "id", parent_id, "error", err)
			} else {
				str_parent = p
			}

		case string:
			str_parent = parent_id.(string)
		default:
			slog.Warn("Unexpected type for parent ID")
		}
	}

	label_opts := &hierarchies.AncestorIdsForLabelOptionsGeneric[string]{
		Hierarchies: str_hiers,
		Placetype:   str_pt,
		ParentId:    str_parent,
	}

	name_ids := hierarchies.AncestorIdsForLabelGeneric[string](label_opts)

	names_q := fmt.Sprintf("SELECT name, placetype, country from %s WHERE id = ?", g.tableName("records"))

	for _, str_id := range name_ids {

		id, err := g.retrieveInt64IdentifierDb(ctx, str_id)

		if err != nil {
			logger.Error("Failed to retrieve int64 id", "id", str_id, "error", err)
			continue
		}

		var id_name string
		var id_placetype string
		var id_country string

		row := g.db.QueryRowContext(ctx, names_q, id)
		err = row.Scan(&id_name, &id_placetype, &id_country)

		switch {
		case err == db_sql.ErrNoRows:
			continue
		case err != nil:
			logger.Warn("Failed to query ID for name", "name id", id)
		default:

			switch id_placetype {
			case "country":
				labels = append(labels, id_country)
			default:
				labels = append(labels, id_name)
			}
		}
	}

	f.Properties["geocoder:label"] = strings.Join(labels, ", ")
	return nil
}

func (g *SQLGeocoder) assignPlacetypeAlt(ctx context.Context, f *geojson.Feature) error {

	pt_q := fmt.Sprintf("SELECT p.placetype from %s p, %s r WHERE r.id = p.record_id AND r.id = ?", g.tableName("placetypes_alt"), g.tableName("records"))

	pt_rows, err := g.db.QueryContext(ctx, pt_q, f.ID)

	if err != nil {
		return err
	}

	defer pt_rows.Close()

	alt_pt := make([]string, 0)

	for pt_rows.Next() {

		var pt string
		err := pt_rows.Scan(&pt)

		if err != nil {
			return fmt.Errorf("Failed to scan placetypes row, %w", err)
		}

		alt_pt = append(alt_pt, pt)
	}

	err = pt_rows.Err()

	if err != nil {
		return fmt.Errorf("Failed to derive alt placetypes, %w", err)
	}

	if len(alt_pt) > 0 {
		f.Properties["wof:placetype_alt"] = alt_pt
	}

	return nil
}
