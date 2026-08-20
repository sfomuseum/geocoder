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
	"github.com/sfomuseum/geocoder/http/api"
	www_coarse "github.com/sfomuseum/geocoder/http/www/coarse"
)

func Run(ctx context.Context) error {

	fs := DefaultFlagSet()
	return RunWithFlagSet(ctx, fs)
}

func RunWithFlagSet(ctx context.Context, fs *flag.FlagSet) error {

	opts, err := OptionsFromFlagSet(ctx, fs)

	if err != nil {
		return fmt.Errorf("Failed to derive options, %w", err)
	}

	return RunWithOptions(ctx, opts)
}

func RunWithOptions(ctx context.Context, opts *Options) error {

	if opts.Verbose {
		slog.SetLogLoggerLevel(slog.LevelDebug)
		slog.Debug("Verbose logging enabled")
	}

	if opts.Geocoder == nil {
		return fmt.Errorf("Missing geocoder")
	}

	defer opts.Geocoder.Close()

	mux := http.NewServeMux()

	if demo {

		map_opts := &maps.AssignMapConfigHandlerOptions{
			MapProvider: "leaflet",
			MapTileURI:  maps.LEAFLET_OSM_TILE_URL,
			InitialView: "-180,-90,180,90",
		}

		map_uri := "/map.json"

		if opts.Prefix != "" {

			u, err := url.JoinPath(opts.Prefix, map_uri)

			if err != nil {
				return fmt.Errorf("Failed to add prefix to map config URL, %w", err)
			}

			map_uri = u
		}

		maps.AssignMapConfigHandler(map_opts, mux, map_uri)

		www_handler := http.FileServerFS(www_coarse.FS)

		if opts.Prefix == "" {
			mux.Handle("/", www_handler)
		} else {

			root_uri, err := url.JoinPath(opts.Prefix, "/")

			if err != nil {
				return fmt.Errorf("Failed to apply prefix to root (demo) URL, %w", err)
			}

			mux.Handle(root_uri, http.StripPrefix(opts.Prefix, www_handler))
		}
	}

	api_opts := &api.CoarseGeocoderHandlerOptions{
		Geocoder:             opts.Geocoder,
		PaginationPerPage:    opts.PaginationPerPage,
		QueryTimeout:         opts.QueryTimeout,
		AllowQueryEmbeddings: opts.AllowQueryEmbeddings,
	}

	api_handler, err := api.CoarseGeocoderHandler(api_opts)

	if err != nil {
		return fmt.Errorf("Failed to create coarse geocoder handler, %w", err)
	}

	if opts.Prefix == "" {
		mux.Handle("POST /api/query/", api_handler)
	} else {

		api_uri, err := url.JoinPath(opts.Prefix, "/api/query/")

		if err != nil {
			return fmt.Errorf("Failed to apply prefix to API (query) URL, %w", err)
		}

		api_uri = fmt.Sprintf("POST %s", api_uri)
		mux.Handle(api_uri, api_handler)
	}

	s, err := server.NewServer(ctx, opts.ServerURI)

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
