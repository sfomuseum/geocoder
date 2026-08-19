package query

import (
	"flag"
	"fmt"
	"os"

	"github.com/sfomuseum/geocoder/x/vec"
	"github.com/sfomuseum/go-flags/flagset"
	"github.com/sfomuseum/go-flags/multi"
)

var geocoder_uri string
var query string
var lang string
var lang_tag string

var placetypes multi.MultiString
var countries multi.MultiString
var belongsto multi.MultiInt64
var str_bounds string
var date_starts string
var date_ends string

var embeddings_search bool
var embeddings_model string
var embedder_uri string

var page int64
var per_page int64

var query_timeout int

var mode string
var verbose bool

func DefaultFlagSet() *flag.FlagSet {

	fs := flagset.NewFlagSet("query")

	fs.StringVar(&geocoder_uri, "geocoder-uri", "null://", "A registered sfomuseum/geocoder/coarse.Geocoder URI.")
	fs.StringVar(&query, "query", "", "The term to query for. Required.")
	fs.Var(&placetypes, "placetype", "Zero or more placetypes to filter results by.")
	fs.Var(&countries, "country", "Zero or more 2-letter country codes to filter results by.")
	fs.Var(&belongsto, "belongs-to", "Zero or more Who's On First ancestor IDs to filter results by.")
	fs.StringVar(&lang, "lang", "", "An optional (3-letter) language code to filter results by,")
	fs.StringVar(&lang_tag, "tag", "", "An option WOF language tag to filter results by.")
	fs.StringVar(&str_bounds, "bounds", "", "Optional bounding box (in the form of 'minx,miny,maxx,mayx') to filter results by.")
	fs.StringVar(&date_starts, "date-starts", "", "Optional ETDF starting date string to filter results by.")
	fs.StringVar(&date_ends, "date-ends", "", "Optional ETDF ending date string to filter results by.")

	fs.BoolVar(&embeddings_search, "embeddings-search", false, "Generate and use vector embeddings for query terms.")
	fs.StringVar(&embedder_uri, "embedder-uri", vec.DEFAULT_EMBEDDER_URI, "A registered sfomuseum/go-embeddings.Embedder URI.")
	fs.StringVar(&embeddings_model, "embeddings-model", vec.DEFAULT_EMBEDDINGS_MODEL, "The URI for the model to use to generate embeddings. For the time being, do NOT change this unless you are using an alternate model with a dimensionality of 384.")

	fs.IntVar(&query_timeout, "query-timeout", 5, "The maximum allowable time in seconds for a query to complete.")

	fs.Int64Var(&page, "page", 1, "The specific page number to query for paginated result sets.")
	fs.Int64Var(&per_page, "per-page", 100, "The number of results to include for paginated result sets.")

	fs.StringVar(&mode, "mode", "tab", "Output mode for results. Valid options are: geojson, tab.")
	fs.BoolVar(&verbose, "verbose", false, "Enable verbose (debug) logging.")

	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "Query a Who's On First (coarse) geocoding database.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t%s [options]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Valid options are:\n")
		fs.PrintDefaults()
	}

	return fs
}
