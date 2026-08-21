package coarse

import (
	"github.com/paulmach/orb"
	"github.com/sfomuseum/go-edtf/unix"
	"github.com/whosonfirst/go-whosonfirst/v4/flags"
)

// QueryRequest represents the set of parameters that can be supplied
// by a client to perform a coarse geocoding query.  All fields are
// optional except for Query which is required.  The struct is
// designed to be serialisable to JSON for future API extensions.
type QueryRequest struct {
	// Query is the term to query for.  The field is required.
	Query string `json:"query"`
	// QueryEmbeddings...
	QueryEmbeddings []float32 `json:"query_embeddings"`
	// Lang is an optional 3‑letter language code that limits the
	// search to tokens in that language.
	Lang string `json:"lang,omitempty"`
	// Tag is an optional WOF language tag (e.g. preferred, variant) that limits the search to tokens with that tag.
	Tag string `json:"tag,omitempty"`
	// Placetype is a list of placetypes to restrict the search to.
	Placetype []string `json:"placetype,omitempty"`
	// Country is a list of 2‑letter ISO country codes to restrict the search to.
	Country []string `json:"country,omitempty"`
	// BelongsTo contains a list of ancestor WOF IDs that the results must belong to.
	BelongsTo []int64 `json:"belongsto,omitempty"`
	// Bounds is an optional geographical bounding box that limits the search to results whose bounds intersect this box.
	Bounds *orb.Bound `json:"bounds,omitempty"`
	// IsCurrent is an optional filter that limits results to places that are current, not current or unknown.
	IsCurrent flags.ExistentialFlag `json:"is_current,omitempty"`
	// DateStarts is an optional date range that limits results to those whose start date overlaps the supplied range.
	DateStarts *unix.DateRange `json:"date_starts,omitempty"`
	// DateEnds is an optional date range that limits results to those whose end date overlaps the supplied range.
	DateEnds *unix.DateRange `json:"date_ends,omitempty"`
}
