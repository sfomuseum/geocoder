package server

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"

	"github.com/aaronland/go-http-maps/v2"
	"github.com/aaronland/go-http/v4/server"
	"github.com/sfomuseum/geocoder/coarse"
	"github.com/sfomuseum/geocoder/http/api"
	www_coarse "github.com/sfomuseum/geocoder/http/www/coarse"
	"github.com/sfomuseum/go-flags/flagset"
)

func Run(ctx context.Context) error {

	fs := DefaultFlagSet()
	return RunWithFlagSet(ctx, fs)
}

func RunWithFlagSet(ctx context.Context, fs *flag.FlagSet) error {

	flagset.Parse(fs)

	err := flagset.SetFlagsFromEnvVars(fs, "GEOCODER")

	if err != nil {
		return fmt.Errorf("Failed to set flags from environment variables, %w", err)
	}

	if verbose {
		slog.SetLogLoggerLevel(slog.LevelDebug)
		slog.Debug("Verbose logging enabled")
	}

	gc, err := coarse.NewSQLGeocoder(ctx, geocoder_uri)

	if err != nil {
		return fmt.Errorf("Failed to create geocoder, %w", err)
	}

	defer gc.Close()

	mux := http.NewServeMux()

	if demo {

		map_opts := &maps.AssignMapConfigHandlerOptions{
			MapProvider: "leaflet",
			MapTileURI:  maps.LEAFLET_OSM_TILE_URL,
			InitialView: "-180,-90,180,90",
		}

		map_uri := "/map.json"

		if prefix != "" {

			u, err := url.JoinPath(prefix, map_uri)

			if err != nil {
				return fmt.Errorf("Failed to add prefix to map config URL, %w", err)
			}

			map_uri = u
		}

		maps.AssignMapConfigHandler(map_opts, mux, map_uri)

		www_handler := http.FileServerFS(www_coarse.FS)

		if prefix == "" {
			mux.Handle("/", www_handler)
		} else {

			root_uri, err := url.JoinPath(prefix, "/")

			if err != nil {
				return fmt.Errorf("Failed to apply prefix to root (demo) URL, %w", err)
			}

			mux.Handle(root_uri, http.StripPrefix(prefix, www_handler))
		}
	}

	api_opts := &api.CoarseGeocoderHandlerOptions{
		Geocoder:          gc,
		PaginationPerPage: per_page,
		QueryTimeout:      query_timeout,
	}

	api_handler, err := api.CoarseGeocoderHandler(api_opts)

	if err != nil {
		return fmt.Errorf("Failed to create coarse geocoder handler, %w", err)
	}

	if prefix == "" {
		mux.Handle("/api/query/", api_handler)
	} else {

		api_uri, err := url.JoinPath(prefix, "/api/query/")

		if err != nil {
			return fmt.Errorf("Failed to apply prefix to API (query) URL, %w", err)
		}

		mux.Handle(api_uri, api_handler)
	}

	s, err := server.NewServer(ctx, server_uri)

	if err != nil {
		return fmt.Errorf("Failed to create new server, %w", err)
	}

	slog.Info("Listening for requests", "address", s.Address())

	err = s.ListenAndServe(ctx, mux)

	if err != nil {
		return fmt.Errorf("Failed to serve requests, %w", err)
	}

	return nil
}
