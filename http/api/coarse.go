package api

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/aaronland/go-http/v4/sanitize"
	"github.com/aaronland/go-http/v4/slog"
	"github.com/aaronland/go-pagination"
	"github.com/aaronland/go-pagination/countable"
	"github.com/paulmach/orb/geojson"
	"github.com/sfomuseum/geocoder"
	"github.com/sfomuseum/geocoder/coarse"
	"github.com/sfomuseum/go-edtf/unix"
)

// APIResponse represents the JSON structure returned by the
// HTTP /api/query endpoint.  It contains a pagination object
// and a GeoJSON FeatureCollection.
type APIResponse struct {
	// Pagination contains pagination metadata such as total results, current page and number of pages.
	Pagination pagination.Results `json:"pagination"`
	// Results is the GeoJSON FeatureCollection returned by the query.
	Results *geojson.FeatureCollection `json:"results"`
}

// CoarseGeocoderHandlerOptions contains configuration options
// for the HTTP handler that implements the /api/query endpoint.
type CoarseGeocoderHandlerOptions struct {
	// Geocoder is the Geocoder instance that will be used  to execute the query.
	Geocoder coarse.Geocoder
	// PaginationPerPage is the maximum number of results that should be returned per API request.
	PaginationPerPage int64
	// The maximum allowable time in seconds for a query to complete.
	QueryTimeout int
	// AllowQueryEmbeddings in API requests.
	AllowQueryEmbeddings bool
}

// CoarseGeocoderHandler creates an HTTP handler that exposes the geocoder as a REST API.
// The handler validates request parameters, performs a query and returns the results in
// JSON format.
func CoarseGeocoderHandler(opts *CoarseGeocoderHandlerOptions) (http.Handler, error) {

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		logger := slog.LoggerWithRequest(req, nil)

		if req.Method != http.MethodPost {
			http.Error(rsp, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		query, err := sanitize.PostString(req, "query")

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

		countries, err := sanitize.PostStringMulti(req, "country")

		if err != nil {
			logger.Error("Failed to derive country parameter(s)", "error", err)
			http.Error(rsp, "Invalid country", http.StatusBadRequest)
			return
		}

		if len(countries) > 0 {
			query_req.Country = countries
		}

		// Belongs to

		belongsto, err := sanitize.PostInt64Multi(req, "belongs-to")

		if err != nil {
			logger.Error("Failed to derive belongs-to parameter(s)", "error", err)
			http.Error(rsp, "Invalid belongs-to", http.StatusBadRequest)
			return
		}

		if len(belongsto) > 0 {
			query_req.BelongsTo = belongsto
		}

		// Placetypes

		placetypes, err := sanitize.PostStringMulti(req, "placetype")

		if err != nil {
			logger.Error("Failed to derive placetype parameter(s)", "error", err)
			http.Error(rsp, "Invalid placetype", http.StatusBadRequest)
			return
		}

		if len(placetypes) > 0 {
			query_req.Placetype = placetypes
		}

		// Language and tag

		lang, err := sanitize.PostString(req, "lang")

		if err != nil {
			logger.Error("Failed to derive lang parameter", "error", err)
			http.Error(rsp, "Invalid lang", http.StatusBadRequest)
			return
		}

		if lang != "" {
			query_req.Lang = lang
		}

		tag, err := sanitize.PostString(req, "tag")

		if err != nil {
			logger.Error("Failed to derive tag parameter", "error", err)
			http.Error(rsp, "Invalid tag", http.StatusBadRequest)
			return
		}

		if tag != "" {
			query_req.Tag = tag
		}

		// Bounds

		str_bounds, err := sanitize.PostString(req, "bounds")

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

		date_starts, err := sanitize.PostString(req, "date-starts")

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

		date_ends, err := sanitize.PostString(req, "date-ends")

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

		// Embeddings

		str_embeddings, err := sanitize.PostString(req, "query-embeddings")

		if err != nil {
			logger.Error("Failed to derive ?query-embeddings= parameter", "error", err)
			http.Error(rsp, "Invalid ?query-embeddings= parameter", http.StatusBadRequest)
			return
		}

		if str_embeddings != "" {

			if !opts.AllowQueryEmbeddings {
				http.Error(rsp, "Query embeddings are not supported", http.StatusBadRequest)
				return
			}

			var query_embeddings []float32

			err = json.Unmarshal([]byte(str_embeddings), &query_embeddings)

			if err != nil {
				logger.Error("Failed to unmarshal ?query-embeddings= parameter", "error", err)
				http.Error(rsp, "Invalid ?query-embeddings= parameter", http.StatusBadRequest)
				return
			}

			query_req.QueryEmbeddings = query_embeddings
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

		page, err := sanitize.PostInt64(req, "page")

		if err != nil {
			logger.Error("Failed to derive page number", "error", err)
			http.Error(rsp, "Invalid page number", http.StatusBadRequest)
			return
		}

		if page > 0 {
			pg_opts.Pointer(page)
		}

		timeout := time.Duration(opts.QueryTimeout) * time.Second

		ctx, cancel := context.WithTimeout(req.Context(), timeout)
		defer cancel()

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
