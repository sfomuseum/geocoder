package index

import (
	"flag"
	"fmt"
	"os"

	"github.com/sfomuseum/go-flags/flagset"
)

var iterator_uri string
var offset int64
var geocoder_uri string
var fresh bool
var prune bool
var index_juggling bool

var exclude_deprecated bool
var exclude_superseded bool
var exclude_funky bool
var exclude_nullisland bool

var embeddings_index bool
var embeddings_model string
var embedder_uri string

var verbose bool

func DefaultFlagSet() *flag.FlagSet {

	fs := flagset.NewFlagSet("index")

	fs.StringVar(&iterator_uri, "iterator-uri", "repo://", "A registered whosonfirst/go-whosonfirst/v4/iterate.Iterate URI.")
	fs.StringVar(&geocoder_uri, "geocoder-uri", "sql://sqlite?dsn=:memory:", "A registered sfomuseum/geocoder/coarse.Geocoder URI.")
	fs.Int64Var(&offset, "offset", 0, "Optional document offset to start indexing from.")
	fs.BoolVar(&fresh, "fresh", false, "This flags signals that a fresh database is being indexed disabling checks for existing or updated records.")
	fs.BoolVar(&prune, "prune", false, "Prune existing records before (re)adding them to the database.")
	fs.BoolVar(&index_juggling, "index-juggling", true, "Perform indexing speed optiomizations. This will include dropping existing indices and the FTS table prior to indexing and (re)adding them at the end.")

	fs.BoolVar(&exclude_deprecated, "exclude-deprecated", true, "Do not index records which have been deprecated.")
	fs.BoolVar(&exclude_superseded, "exclude-superseded", true, "Do not index records which have been superseded.")
	fs.BoolVar(&exclude_funky, "exclude-funky", true, "Do not index records which have been flagged as \"funky\".")
	fs.BoolVar(&exclude_nullisland, "exclude-nullisland", true, "Do not index records that are \"visiting\" Null Island (have 0,0 coordinate data).")

	fs.BoolVar(&embeddings_index, "embeddings-index", false, "...")
	fs.StringVar(&embedder_uri, "embedder-uri", "ollama://", "...")
	fs.StringVar(&embeddings_model, "embeddings-model", "embeddinggemma", "...")

	fs.BoolVar(&verbose, "verbose", false, "Enable verbose (debug) logging.")

	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "Index one or more Who's On First data sources in a (coarse) geocoding database.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t%s [options] uri(N) uri(N) uri(N)\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Valid options are:\n")
		fs.PrintDefaults()
	}

	return fs
}
