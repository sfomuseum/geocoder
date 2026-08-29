package coarse

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"github.com/paulmach/orb"
	"github.com/sfomuseum/go-edtf"
	"github.com/sfomuseum/go-edtf/parser"
)

type Tokens struct {
	Language string
	Tag      string
	Tokens   []string
}

type Embeddings struct {
	Language   string
	Tag        string
	Embeddings []float32
}

type VectorEmbeddings struct {
	Model      string
	Embeddings []*Embeddings
}

// Record represents the data stored in the coarse geocoder database.
// The JSON tags are the same names that appear in Who's On First
// GeoJSON files.  The struct contains a subset of the full
// WOF schema – only the fields that are needed for coarse
// geocoding are retained.
type Record struct {
	// Id is the unique identifier of the place.
	Id string `json:"geocoder:id"`
	// ParentId is the unique identifier of the parent place.
	ParentId string `json:"geocoder:parent_id"`
	// Name is the primary name of the place.
	Name string `json:"geocoder:name"`
	// Country is the ISO 3166‑1 alpha‑2 country code of the place.
	Country string `json:"wof:country"`
	// Placetype is the primary Who's On First placetype of the place.
	Placetype string `json:"wof:placetype"`
	// PlacetypeAlt contains any alternative placetypes stored in
	// the Who's On First `wof:placetype_alt` property.
	PlacetypeAlt []string `json:"wof:placetype_alt"`
	// Hierarchies contains the ancestor hierarchies for the place.
	// Each hierarchy is a map of placetype to ancestor ID.
	Hierarchies []map[string]string `json:"geocoder:hierarchies"`
	// Centroid is the geographic centroid of the place.
	Centroid *orb.Point `json:"geo:centroid"`
	// Bounds is a slice of bounding boxes that enclose the place.
	Bounds []orb.Bound `json:"geo:bounds"`
	// Inception is the EDTF representation of the start date of the place.
	Inception string `json:"edtf:inception,omitempty"`
	// Cessation is the EDTF representation of the end date of the place.
	Cessation string `json:"etdf:cessation,omitempty"`
	// PopulationRank is an integer that indicates relative population size.
	PopulationRank int64 `json:"wof:population_rank,omitempty"`
	// IsCurrent indicates whether the place is current (1), not current (0)
	// or unknown (-1).
	IsCurrent string `json:"mz:is_current,omitempty"`
	// Tokens contains tokenised names and concordances indexed for full‑text search.
	Tokens map[string]map[string][]string `json:"tokens,omitempty"` // please make me something better...
	// Vectors ...
	VectorEmbeddings []*VectorEmbeddings
}

// Hash returns a SHA‑256 digest of the record in JSON form.  The hash
// is used to determine whether a record has changed between indexing
// runs.
func (r *Record) Hash() (string, error) {

	enc, err := json.Marshal(r)

	if err != nil {
		return "", err
	}

	sum := sha256.Sum256(enc)
	return fmt.Sprintf("%x", sum), nil
}

// DateRanges returns the four timestamp boundaries that are stored in a
// "dates" table.  The returned values are:
//
//   - start_outer – the outermost (earliest) start date
//   - start_inner – the inner (most precise) start date
//   - end_inner   – the inner end date
//   - end_outer   – the outermost end date
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
