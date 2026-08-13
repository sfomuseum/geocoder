package coarse

import (
	"context"
	"fmt"
	"net/url"
	"sort"
	"strings"

	"github.com/aaronland/go-pagination"
	"github.com/aaronland/go-roster"
	"github.com/paulmach/orb/geojson"
)

// Geocoder defines the interface used by the command‑line utilities
// and HTTP server to interact with a coarse geocoding database.
type Geocoder interface {
	// PreIndex runs any pre‑indexing logic that must happen before records are added.
	// For example, the SQLite implementation this method to drop existing indices and
	// the full‑text search (FTS) table so that bulk inserts can be performed more quickly.
	PreIndex(context.Context) error
	// PostIndex runs any post‑indexing logic that must happen after all records have been added.
	// For example the SQLite implementation uses this method to recreate the indices, deduplicates
	// rows in the auxiliary tables and rebuilds the FTS table.
	PostIndex(context.Context) error
	// AddRecord queues a record for bulk insertion into the database.
	// The record will be written when the internal buffer reaches its
	// batch size or when Flush is called.
	AddRecord(context.Context, *Record) error
	// RemoveRecord deletes a record and all of its dependent rows
	// from the database.
	RemoveRecord(context.Context, int64) error
	// RecordExists reports whether a record with the supplied ID is present in the database.
	RecordExists(context.Context, int64) (bool, error)
	// HasRecordHashChanged checks whether a record already exists  and, if so, whether its
	// content hash has changed. It returns (exists, changed, error).
	HasRecordHashChanged(context.Context, *Record) (bool, bool, error)
	// Query performs a coarse geocoding query using the supplied
	// QueryRequest and pagination options.  It returns a slice of
	// GeoJSON Features and a pagination.Results object.
	Query(context.Context, *QueryRequest, pagination.Options) ([]*geojson.Feature, pagination.Results, error)
	// Flush forces any buffered records to be written to the database.
	Flush(context.Context) error
	// Close closes the underlying database connection.
	Close() error
}

var geocoder_roster roster.Roster

// GeocoderInitializationFunc is a function defined by individual geocoder package and used to create
// an instance of that geocoder
type GeocoderInitializationFunc func(ctx context.Context, uri string) (Geocoder, error)

// RegisterGeocoder registers 'scheme' as a key pointing to 'init_func' in an internal lookup table
// used to create new `Geocoder` instances by the `NewGeocoder` method.
func RegisterGeocoder(ctx context.Context, scheme string, init_func GeocoderInitializationFunc) error {

	err := ensureGeocoderRoster()

	if err != nil {
		return err
	}

	return geocoder_roster.Register(ctx, scheme, init_func)
}

func MustRegisterGeocoder(ctx context.Context, scheme string, init_func GeocoderInitializationFunc) {

	err := RegisterGeocoder(ctx, scheme, init_func)

	if err != nil {
		panic(err)
	}
}

func ensureGeocoderRoster() error {

	if geocoder_roster == nil {

		r, err := roster.NewDefaultRoster()

		if err != nil {
			return err
		}

		geocoder_roster = r
	}

	return nil
}

// NewGeocoder returns a new `Geocoder` instance configured by 'uri'. The value of 'uri' is parsed
// as a `url.URL` and its scheme is used as the key for a corresponding `GeocoderInitializationFunc`
// function used to instantiate the new `Geocoder`. It is assumed that the scheme (and initialization
// function) have been registered by the `RegisterGeocoder` method.
func NewGeocoder(ctx context.Context, uri string) (Geocoder, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, err
	}

	scheme := u.Scheme

	i, err := geocoder_roster.Driver(ctx, scheme)

	if err != nil {
		return nil, err
	}

	init_func := i.(GeocoderInitializationFunc)
	return init_func(ctx, uri)
}

// GeocoderSchemes returns the list of schemes that have been registered.
func GeocoderSchemes() []string {

	ctx := context.Background()
	schemes := []string{}

	err := ensureGeocoderRoster()

	if err != nil {
		return schemes
	}

	for _, dr := range geocoder_roster.Drivers(ctx) {
		scheme := fmt.Sprintf("%s://", strings.ToLower(dr))
		schemes = append(schemes, scheme)
	}

	sort.Strings(schemes)
	return schemes
}
