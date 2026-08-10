package coarse

import (
	"github.com/paulmach/orb"
	"github.com/sfomuseum/go-edtf/unix"
	"github.com/whosonfirst/go-whosonfirst/v4/flags"
)

type QueryRequest struct {
	Query      string                `json:"query"`
	Lang       string                `json:"lang,omitempty"`
	Tag        string                `json:"tag,omitempty"`
	Placetype  []string              `json:"placetype,omitempty"`
	Country    []string              `json:country,omitempty"`
	BelongsTo  []int64               `json:"belongsto,omitempty"`
	Bounds     *orb.Bound            `json:"bounds,omitempty"`
	IsCurrent  flags.ExistentialFlag `json:"is_current,omitempty"`
	DateStarts *unix.DateRange       `json:"date_starts,omitempty"`
	DateEnds   *unix.DateRange       `json:"date_ends,omitempty"`
}
