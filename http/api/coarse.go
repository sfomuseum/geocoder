package api

import (
	"encoding/json"
	"net/http"

	"github.com/aaronland/go-http/v4/sanitize"
	"github.com/aaronland/go-http/v4/slog"
	"github.com/aaronland/go-pagination"
	"github.com/aaronland/go-pagination/countable"
	"github.com/paulmach/orb/geojson"
	"github.com/sfomuseum/go-edtf/unix"
	"github.com/sfomuseum/geocoder"
	"github.com/sfomuseum/geocoder/coarse"
)

type APIResponse struct {
	Pagination pagination.Results         `json:"pagination"`
	Results    *geojson.FeatureCollection `json:"results"`
}

type CoarseGeocoderHandlerOptions struct {
	Geocoder          coarse.Geocoder
	PaginationPerPage int64
}

func CoarseGeocoderHandler(opts *CoarseGeocoderHandlerOptions) (http.Handler, error) {

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		logger := slog.LoggerWithRequest(req, nil)

		query, err := sanitize.GetString(req, "query")

		if err != nil {
			logger.Error("Failed to derive query parameter", "error", err)
			http.Error(rsp, "Invalid query", http.StatusBadRequest)
			return
		}

		if query == "" {
			logger.Debug("Query string is empty")
			http.Error(rsp, "Missing query", http.StatusBadRequest)
			return
		}

		query_req := &coarse.QueryRequest{
			Query: query,
		}

		// Countries

		countries, err := sanitize.GetStringMulti(req, "country")

		if err != nil {
			logger.Error("Failed to derive country parameter(s)", "error", err)
			http.Error(rsp, "Invalid country", http.StatusBadRequest)
			return
		}

		if len(countries) > 0 {
			query_req.Country = countries
		}

		// Belongs to

		belongsto, err := sanitize.GetInt64Multi(req, "belongs-to")

		if err != nil {
			logger.Error("Failed to derive belongs-to parameter(s)", "error", err)
			http.Error(rsp, "Invalid belongs-to", http.StatusBadRequest)
			return
		}

		if len(belongsto) > 0 {
			query_req.BelongsTo = belongsto
		}

		// Placetypes

		placetypes, err := sanitize.GetStringMulti(req, "placetype")

		if err != nil {
			logger.Error("Failed to derive placetype parameter(s)", "error", err)
			http.Error(rsp, "Invalid placetype", http.StatusBadRequest)
			return
		}

		if len(placetypes) > 0 {
			query_req.Placetype = placetypes
		}

		// Language and tag

		lang, err := sanitize.GetString(req, "lang")

		if err != nil {
			logger.Error("Failed to derive lang parameter", "error", err)
			http.Error(rsp, "Invalid lang", http.StatusBadRequest)
			return
		}

		if lang != "" {
			query_req.Lang = lang
		}

		tag, err := sanitize.GetString(req, "tag")

		if err != nil {
			logger.Error("Failed to derive tag parameter", "error", err)
			http.Error(rsp, "Invalid tag", http.StatusBadRequest)
			return
		}

		if tag != "" {
			query_req.Tag = tag
		}

		// Bounds

		str_bounds, err := sanitize.GetString(req, "bounds")

		if err != nil {
			logger.Error("Failed to derive bounds parameter", "error", err)
			http.Error(rsp, "Invalid bounds", http.StatusBadRequest)
			return
		}

		if str_bounds != "" {

			bounds, err := geocoder.StringBoundsToOrbBounds(str_bounds)

			if err != nil {
				logger.Error("Failed to parse bounds", "error", err)
				http.Error(rsp, "Invalid bounds", http.StatusBadRequest)
			}

			query_req.Bounds = bounds
		}

		// Dates

		date_starts, err := sanitize.GetString(req, "date-starts")

		if err != nil {
			logger.Error("Failed to derive ?date-starts= parameter", "error", err)
			http.Error(rsp, "Invalid ?date-start= parameter", http.StatusBadRequest)
			return
		}

		if date_starts != "" {

			ok, ranges, err := unix.DeriveRanges(date_starts)

			if err != nil {
				logger.Error("Failed to parse date start", "error", err)
				http.Error(rsp, "Internal server error", http.StatusInternalServerError)
				return
			}

			if !ok {
				logger.Error("Failed to derive date range (start)")
				http.Error(rsp, "Invalid date range (start)", http.StatusBadRequest)
				return
			}

			query_req.DateStarts = ranges
		}

		date_ends, err := sanitize.GetString(req, "date-ends")

		if err != nil {
			logger.Error("Failed to derive ?date-ends= parameter", "error", err)
			http.Error(rsp, "Invalid ?date-end= parameter", http.StatusBadRequest)
			return
		}

		if date_ends != "" {

			ok, ranges, err := unix.DeriveRanges(date_ends)

			if err != nil {
				logger.Error("Failed to parse date end", "error", err)
				http.Error(rsp, "Internal server error", http.StatusInternalServerError)
				return
			}

			if !ok {
				logger.Error("Failed to derive date range (end)")
				http.Error(rsp, "Invalid date range (end)", http.StatusBadRequest)
				return
			}

			query_req.DateEnds = ranges
		}

		//

		// To do: derive options from URI constructor...
		pg_opts, err := countable.NewCountableOptions()

		if err != nil {
			logger.Error("Failed to create new pagination options", "error", err)
			http.Error(rsp, "Internal server error", http.StatusInternalServerError)
			return
		}

		pg_opts.PerPage(opts.PaginationPerPage)
		pg_opts.Spill(0)

		//

		page, err := sanitize.GetInt64(req, "page")

		if err != nil {
			logger.Error("Failed to derive page number", "error", err)
			http.Error(rsp, "Invalid page number", http.StatusBadRequest)
			return
		}

		if page > 0 {
			pg_opts.Pointer(page)
		}

		ctx := req.Context()

		query_rsp, pg_rsp, err := opts.Geocoder.Query(ctx, query_req, pg_opts)

		if err != nil {
			logger.Error("Failed to query geocoder", "query", query, "error", err)
			http.Error(rsp, "Internal server error", http.StatusInternalServerError)
			return
		}

		fc := geojson.NewFeatureCollection()
		fc.Features = query_rsp

		api_rsp := APIResponse{
			Pagination: pg_rsp,
			Results:    fc,
		}

		rsp.Header().Set("Content-type", "application/json")
		enc := json.NewEncoder(rsp)

		err = enc.Encode(api_rsp)

		if err != nil {
			logger.Error("Failed to encode query response", "error", err)
			http.Error(rsp, "Internal server error", http.StatusInternalServerError)
			return
		}

		return
	}

	return http.HandlerFunc(fn), nil
}
