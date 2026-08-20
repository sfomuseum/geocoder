package server

import (
	"flag"
	"fmt"
	"os"

	"github.com/sfomuseum/go-flags/flagset"
)

var geocoder_uri string
var server_uri string
var prefix string

var allow_query_embeddings bool

var query_timeout int
var per_page int64

var demo bool
var verbose bool

func DefaultFlagSet() *flag.FlagSet {

	fs := flagset.NewFlagSet("query")

	fs.StringVar(&geocoder_uri, "geocoder-uri", "null://", "A registered sfomuseum/geocoder/coarse.Geocoder URI.")
	fs.StringVar(&server_uri, "server-uri", "http://localhost:8080", "A registered aaronland/go-http/v4/server.Server URI.")
	fs.StringVar(&prefix, "prefix", "", "An optional URL prefix to listen for requests on.")

	fs.IntVar(&query_timeout, "query-timeout", 5, "The maximum allowable time in seconds for a query to complete.")
	fs.BoolVar(&demo, "demo", false, "Start a web-based demo on the root URL of the server.")
	fs.Int64Var(&per_page, "pagination-per-page", 50, "The maximum number of results to include per API request.")
	fs.BoolVar(&allow_query_embeddings, "allow-query-embeddings", true, "Enable vector embedding queries in the /api/query endpoint.")

	fs.BoolVar(&verbose, "verbose", false, "Enable verbose (debug) logging.")

	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "HTTP server for handling requests against a Who's On First (coarse) geocoding database.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t%s [options]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Valid options are:\n")
		fs.PrintDefaults()
	}

	return fs
}
