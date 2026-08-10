package coarse

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"github.com/paulmach/orb"
	"github.com/sfomuseum/go-edtf"
	"github.com/sfomuseum/go-edtf/parser"
)

type Record struct {
	Id             int64                          `json:"wof:id"`
	ParentId       int64                          `json:"wof:parent_id"`
	Name           string                         `json:"wof:name"`
	Country        string                         `json:"wof:country"`
	Placetype      string                         `json:"wof:placetype"`
	PlacetypeAlt   []string                       `json:"wof:placetype_alt"`
	Hierarchies    []map[string]int64             `json:"wof:hierarchies"`
	Centroid       *orb.Point                     `json:"wof:centroid"`
	Bounds         []orb.Bound                    `json:"wof:bounds"`
	Inception      string                         `json:"edtf:inception,omitempty"`
	Cessation      string                         `json:"etdf:cessation,omitempty"`
	PopulationRank int64                          `json:"wof:population_rank,omitempty"`
	IsCurrent      string                         `json:"mz:is_current,omitempty"`
	Tokens         map[string]map[string][]string `json:"tokens,omitempty"` // please make me something better...
}

func (r *Record) Hash() (string, error) {

	enc, err := json.Marshal(r)

	if err != nil {
		return "", err
	}

	sum := sha256.Sum256(enc)
	return fmt.Sprintf("%x", sum), nil
}

func (r *Record) DateRanges() (*edtf.Timestamp, *edtf.Timestamp, *edtf.Timestamp, *edtf.Timestamp) {

	// Note that there are also date range functions in sfomuseum/go-edtf/unix
	// but those are designed specfically to ensure a range which is not necessarily
	// what we want here.

	start_outer, _ := r.DateStartOuter()
	start_inner, _ := r.DateStartInner()
	end_inner, _ := r.DateEndInner()
	end_outer, _ := r.DateEndOuter()

	return start_outer, start_inner, end_inner, end_outer
}

func (r *Record) DateStartOuter() (*edtf.Timestamp, error) {

	d, err := parser.ParseString(r.Inception)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse inception string, %w", err)
	}

	if d.Start == nil {
		return nil, fmt.Errorf("Missing start (range)")
	}

	if d.Start.Lower == nil {
		return nil, fmt.Errorf("Missing start (lower)")
	}

	if d.Start.Lower.Timestamp == nil {
		return nil, fmt.Errorf("Missing start (lower timestamp)")
	}

	return d.Start.Lower.Timestamp, nil
}

func (r *Record) DateStartInner() (*edtf.Timestamp, error) {

	d, err := parser.ParseString(r.Inception)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse inception string, %w", err)
	}

	if d.Start == nil {
		return nil, fmt.Errorf("Missing start (range)")
	}

	if d.Start.Upper == nil {
		return nil, fmt.Errorf("Missing start (lower)")
	}

	if d.Start.Upper.Timestamp == nil {
		return nil, fmt.Errorf("Missing start (lower timestamp)")
	}

	return d.Start.Upper.Timestamp, nil
}

func (r *Record) DateEndInner() (*edtf.Timestamp, error) {

	d, err := parser.ParseString(r.Cessation)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse cessation string, %w", err)
	}

	if d.End == nil {
		return nil, fmt.Errorf("Missing end (range)")
	}

	if d.End.Lower == nil {
		return nil, fmt.Errorf("Missing end (lower)")
	}

	if d.End.Lower.Timestamp == nil {
		return nil, fmt.Errorf("Missing end (lower timestamp)")
	}

	return d.End.Lower.Timestamp, nil
}

func (r *Record) DateEndOuter() (*edtf.Timestamp, error) {

	d, err := parser.ParseString(r.Cessation)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse cessation string, %w", err)
	}

	if d.End == nil {
		return nil, fmt.Errorf("Missing end (range)")
	}

	if d.End.Upper == nil {
		return nil, fmt.Errorf("Missing end (upper)")
	}

	if d.End.Upper.Timestamp == nil {
		return nil, fmt.Errorf("Missing end (upper timestamp)")
	}

	return d.End.Upper.Timestamp, nil
}
