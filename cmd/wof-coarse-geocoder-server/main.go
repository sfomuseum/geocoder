package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"net/url"
	"os"

	"github.com/aaronland/go-http-maps/v2"
	"github.com/aaronland/go-http/v4/server"
	"github.com/sfomuseum/go-flags/flagset"
	"github.com/sfomuseum/geocoder/coarse"
	"github.com/sfomuseum/geocoder/http/api"
	www_coarse "github.com/sfomuseum/geocoder/http/www/coarse"
)

func main() {

	var geocoder_uri string
	var server_uri string
	var prefix string

	var per_page int64

	var demo bool
	var verbose bool

	fs := flagset.NewFlagSet("query")

	fs.StringVar(&geocoder_uri, "geocoder-uri", "", "A registered sfomuseum/geocoder/coarse.Geocoder URI.")
	fs.StringVar(&server_uri, "server-uri", "http://localhost:8080", "A registered aaronland/go-http/v4/server.Server URI.")
	fs.StringVar(&prefix, "prefix", "", "An optional URL prefix to listen for requests on.")

	fs.BoolVar(&demo, "demo", false, "Start a web-based demo on the root URL of the server.")
	fs.Int64Var(&per_page, "pagination-per-page", 50, "The maximum number of results to include per API request.")
	fs.BoolVar(&verbose, "verbose", false, "Enable verbose (debug) logging.")

	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "HTTP server for handling requests against a Who's On First (coarse) geocoding database.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t%s [options]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Valid options are:\n")
		fs.PrintDefaults()
	}

	flagset.Parse(fs)

	if verbose {
		slog.SetLogLoggerLevel(slog.LevelDebug)
		slog.Debug("Verbose logging enabled")
	}

	ctx := context.Background()

	gc, err := coarse.NewSQLGeocoder(ctx, geocoder_uri)

	if err != nil {
		log.Fatalf("Failed to create geocoder, %v", err)
	}

	defer gc.Close()

	mux := http.NewServeMux()

	if demo {

		map_opts := &maps.AssignMapConfigHandlerOptions{
			MapProvider: "leaflet",
			MapTileURI:  maps.LEAFLET_OSM_TILE_URL,
			InitialView: "-180,-90,180,90",
		}

		maps.AssignMapConfigHandler(map_opts, mux, "/map.json")

		www_handler := http.FileServerFS(www_coarse.FS)

		if prefix == "" {
			mux.Handle("/", www_handler)
		} else {

			root_uri, err := url.JoinPath(prefix, "/")

			if err != nil {
				log.Fatalf("Failed to apply prefix to root (demo) URL, %v", err)
			}

			mux.Handle(root_uri, http.StripPrefix(prefix, www_handler))
		}
	}

	api_opts := &api.CoarseGeocoderHandlerOptions{
		Geocoder:          gc,
		PaginationPerPage: per_page,
	}

	api_handler, err := api.CoarseGeocoderHandler(api_opts)

	if err != nil {
		log.Fatalf("Failed to create coarse geocoder handler, %v", err)
	}

	if prefix == "" {
		mux.Handle("/api/query/", api_handler)
	} else {

		api_uri, err := url.JoinPath(prefix, "/api/query/")

		if err != nil {
			log.Fatalf("Failed to apply prefix to API (query) URL, %v", err)
		}

		mux.Handle(api_uri, api_handler)
	}

	s, err := server.NewServer(ctx, server_uri)

	if err != nil {
		log.Fatalf("Failed to create new server, %v", err)
	}

	slog.Info("Listening for requests", "address", s.Address())

	err = s.ListenAndServe(ctx, mux)

	if err != nil {
		log.Fatalf("Failed to serve requests, %v", err)
	}
}
