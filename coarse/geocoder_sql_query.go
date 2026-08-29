package coarse

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"strings"
	"sync"
	"time"
	"unicode"

	"github.com/aaronland/go-pagination"
	"github.com/aaronland/go-pagination/countable"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
	"github.com/sfomuseum/geocoder/placeholder"
)

func (g *SQLGeocoder) Query(ctx context.Context, req *QueryRequest, pg_opts pagination.Options) ([]*geojson.Feature, pagination.Results, error) {

	logger := slog.Default()
	t1 := time.Now()

	defer func() {
		logger.Debug("Time to query", "time", time.Since(t1))
	}()

	args := make([]any, 0)
	sb := strings.Builder{}

	var query_str string

	if len(req.QueryEmbeddings) == 0 {

		if len(req.Query) < g.min_query_length {
			return nil, nil, fmt.Errorf("Query below min query length")
		}

		query_str = g.prepareQuery(req.Query)

		if query_str == "" {
			return nil, nil, fmt.Errorf("empty or invalid search term")
		}

		if len(query_str) < g.min_query_length {
			return nil, nil, fmt.Errorf("Query below min query length")
		}

		logger = logger.With("query", query_str)

		sb.WriteString("SELECT f.rank, COUNT(*) OVER() as total_count, r.id, r.parent_id, r.name, r.placetype, r.country, r.is_current, r.latitude, r.longitude, r.inception, r.cessation, r.hierarchies")
		sb.WriteString(" FROM tokens_fts f JOIN tokens t ON t.row_id = f.rowid JOIN records r ON r.id = t.record_id")

	} else {

		logger = logger.With("query embeddings", true)
		enc_e, err := json.Marshal(req.QueryEmbeddings)

		if err != nil {
			return nil, nil, fmt.Errorf("Failed to encode embeddings, %w", err)
		}

		sb.WriteString("WITH vector_matches AS (SELECT rowid, distance FROM embeddings WHERE embedding MATCH ? AND k = ?)")
		sb.WriteString(" SELECT vm.distance AS distance, COUNT(*) OVER() as total_count, r.id, r.parent_id, r.name, r.placetype, r.country, r.is_current, r.latitude, r.longitude, r.inception, r.cessation, r.hierarchies")
		sb.WriteString(" FROM vector_matches vm JOIN embeddings_records er ON er.id = vm.rowid JOIN records r ON r.id = er.id")

		args = append(args, string(enc_e))
		args = append(args, g.vector_query_k)

	}

	if len(req.Placetype) > 0 {
		sb.WriteString(" LEFT JOIN placetypes_alt p ON r.id = p.record_id")
	}

	if len(req.BelongsTo) > 0 {
		sb.WriteString(" JOIN ancestors a ON r.id = a.record_id")
	}

	// dates

	if req.DateStarts != nil || req.DateEnds != nil {
		sb.WriteString(" JOIN dates d ON r.id = d.record_id")
	}

	// bounds

	if req.Bounds != nil {
		sb.WriteString(" JOIN bounds b ON r.id = b.record_id")
	}

	if req.Source != "" {
		sb.WriteString(" JOIN identifiers i ON i.id = r.id")
	}

	// Query stuff

	if len(req.QueryEmbeddings) == 0 {

		sb.WriteString(" WHERE f.token MATCH ?")

		args = []any{
			query_str,
		}
	}

	// Dates

	if req.DateStarts != nil {

		sb.WriteString(" AND (? <= d.start_outer AND ? <= d.start_inner)")
		args = append(args, req.DateStarts.Outer.Start)
		args = append(args, req.DateStarts.Inner.Start)
	}

	if req.DateEnds != nil {

		sb.WriteString(" AND (d.end_inner <= ? AND d.end_outer <= ?)")
		args = append(args, req.DateEnds.Inner.End)
		args = append(args, req.DateEnds.Outer.End)
	}

	// Bounds

	if req.Bounds != nil {

		coords := []any{
			req.Bounds.Min.X(),
			req.Bounds.Max.X(),
			req.Bounds.Min.Y(),
			req.Bounds.Max.Y(),
		}

		sb.WriteString(" AND (b.maxx >= ? AND b.minx <= ? AND b.maxy >= ? AND b.miny <= ?)")
		args = append(args, coords...)
	}

	// Placetypes

	if len(req.Placetype) > 0 {

		placeholders := make([]string, len(req.Placetype))

		for i, pt := range req.Placetype {
			placeholders[i] = "?"
			args = append(args, pt)
		}

		for i, pt := range req.Placetype {
			placeholders[i] = "?"
			args = append(args, pt)
		}

		str_placeholders := strings.Join(placeholders, ",")

		sb.WriteString(fmt.Sprintf(" AND (r.placetype IN (%s) OR p.placetype IN (%s))", str_placeholders, str_placeholders))
	}

	// Belongs to (ancestors)

	if len(req.BelongsTo) > 0 {

		placeholders := make([]string, len(req.BelongsTo))

		for i, anc_id := range req.BelongsTo {
			placeholders[i] = "?"
			args = append(args, anc_id)
		}

		sb.WriteString(fmt.Sprintf(" AND a.ancestor_id IN (%s)", strings.Join(placeholders, ",")))
	}

	// Language

	if req.Lang != "" {
		sb.WriteString(" AND r.lang = ?")
		args = append(args, req.Tag)
	}

	// Language (x-) tag

	if req.Tag != "" {
		sb.WriteString(" AND r.tag = ?")
		args = append(args, req.Tag)
	}

	// Countries

	if len(req.Country) > 0 {

		placeholders := make([]string, len(req.Country))

		for i, pt := range req.Country {
			placeholders[i] = "?"
			args = append(args, pt)
		}

		sb.WriteString(fmt.Sprintf(" AND r.country IN (%s)", strings.Join(placeholders, ",")))
	}

	// Is current

	if req.IsCurrent != nil {
		sb.WriteString(" AND r.is_current = ?")
		args = append(args, req.IsCurrent.StringFlag())
	}

	// Source

	if req.Source != "" {
		sb.WriteString(" AND i.identifier LIKE ?")
		args = append(args, req.Source+"%")
	}

	sb.WriteString(" GROUP BY r.id")

	sb.WriteString(` ORDER BY (
			CASE r.is_current
				WHEN 1 THEN 0.0
				WHEN -1 THEN 1.0
				ELSE 2.0
                        END
		) ASC`)

	if len(req.QueryEmbeddings) > 0 {
		sb.WriteString(`, vm.distance ASC`)
	} else {

		sb.WriteString(`, f.rank ASC, MIN(CASE t.tag
					WHEN 'concordance' THEN 0.5
					WHEN 'preferred'    THEN 1.0
					WHEN 'offical'    THEN 1.5
					WHEN 'colloquial' THEN 2.0
					WHEN 'variant'    THEN 4.0
					WHEN 'historical'    THEN 5.0
					WHEN 'unknown'   THEN 6.0
					ELSE 10.0
				END) ASC`)

	}

	sb.WriteString(` ,r.population_rank DESC,
			(CASE r.placetype
				WHEN 'microhood' THEN 1.0
				WHEN 'neighbourhood' THEN 1.0
				WHEN 'borough' THEN 1.5
				WHEN 'locality' THEN 2.0
				WHEN 'localadmin' THEN 2.25
				WHEN 'campus' THEN 2.5
				WHEN 'postalcode' THEN 2.9
				WHEN 'county' THEN 3.0
				WHEN 'marinearea' THEN 3.5
				WHEN 'region' THEN 4.0
				WHEN 'country' THEN 5.0	
				ELSE 10.0
                        END) ASC`)

	page := countable.PageFromOptions(pg_opts)
	per_page := pg_opts.PerPage()

	sb.WriteString(fmt.Sprintf(" LIMIT %d", per_page))

	if page > 1 {
		offset := (page - 1) * per_page
		sb.WriteString(fmt.Sprintf(" OFFSET %d", offset))
	}

	q := sb.String()
	// slog.Info(q, "args", args)

	rows, err := g.db.QueryContext(ctx, q, args...)

	if err != nil {
		return nil, nil, err
	}

	defer rows.Close()

	var features []*geojson.Feature
	var total_count int64

	for rows.Next() {

		var rank float64
		var id int64
		var parent_id int64
		var name string
		var country string
		var placetype string
		var is_current int
		var latitude float64
		var longitude float64
		var inception string
		var cessation string
		var enc_hierarchies string

		err := rows.Scan(&rank, &total_count, &id, &parent_id, &name, &placetype, &country, &is_current, &latitude, &longitude, &inception, &cessation, &enc_hierarchies)

		if err != nil {
			return nil, nil, err
		}

		props := map[string]any{
			"geocoder:id":          id,
			"geocoder:parent_id":   parent_id,
			"geocoder:name":        name,
			"wof:country":          country,
			"wof:placetype":        placetype,
			"mz:is_current":        is_current,
			"edtf:inception":       inception,
			"edtf:cessation":       cessation,
			"geocoder:hierarchies": enc_hierarchies,
			"geocoder:rank":        rank,
		}

		pt := orb.Point([2]float64{longitude, latitude})
		f := geojson.NewFeature(pt)

		f.ID = id
		f.Properties = props
		features = append(features, f)
	}

	err = rows.Err()

	if err != nil {
		return nil, nil, err
	}

	logger = logger.With("total", total_count)

	wg := new(sync.WaitGroup)

	for _, f := range features {

		wg.Go(func() {
			g.assignExtra(ctx, f)
		})
	}

	wg.Wait()

	pg_rsp, err := countable.NewResultsFromCountWithOptions(pg_opts, total_count)

	if err != nil {
		return nil, nil, fmt.Errorf("Failed to create pagination response, %w", err)
	}

	return features, pg_rsp, nil
}

func (g *SQLGeocoder) prepareQuery(input string) string {

	if re_machinetag.MatchString(input) {
		// To do: Support wildcards
		input = strings.ReplaceAll(input, ":", "_")
		input = strings.ReplaceAll(input, "=", "__")
		return input
	}

	words := strings.Fields(placeholder.Normalize(input))

	if len(words) == 0 {
		return ""
	}

	var sanitized []string

	for _, word := range words {

		// Strip characters that disrupt FTS5 syntax (like raw quotes or dashes)
		clean := strings.Map(func(r rune) rune {
			// unicode.IsLetter handles all global alphabets (Greek, Cyrillic, CJK, Arabic, etc.)
			// unicode.IsNumber handles global digits (0-9, and non-Arabic numeral systems)
			if unicode.IsLetter(r) || unicode.IsNumber(r) {
				return r
			}

			// Keep Unicode letters/numbers if working with global languages
			return -1
		}, word)

		if clean != "" {
			sanitized = append(sanitized, clean)
		}
	}

	if len(sanitized) == 0 {
		return ""
	}

	lastIdx := len(sanitized) - 1
	sanitized[lastIdx] = sanitized[lastIdx] + "*"

	return strings.Join(sanitized, " AND ")
}
