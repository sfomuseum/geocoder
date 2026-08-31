package coarse

import (
	"context"

	"github.com/aaronland/go-pagination"
	"github.com/aaronland/go-pagination/countable"
	"github.com/paulmach/orb/geojson"
)

func init() {
	ctx := context.Background()
	MustRegisterGeocoder(ctx, "null", NewNullGeocoder)
}

type NullGeocoder struct {
	Geocoder
}

func NewNullGeocoder(ctx context.Context, uri string) (Geocoder, error) {

	gc := &NullGeocoder{}
	return gc, nil
}

func (gc *NullGeocoder) PreIndex(ctx context.Context) error {
	return nil
}

func (gc *NullGeocoder) PostIndex(ctx context.Context) error {
	return nil
}

func (gc *NullGeocoder) AddRecord(ctx context.Context, rec *Record) error {
	return nil
}

func (gc *NullGeocoder) RemoveRecord(ctx context.Context, id string) error {
	return nil
}

func (gc *NullGeocoder) RecordExists(ctx context.Context, id string) (bool, error) {
	return false, nil
}

func (gc *NullGeocoder) HasRecordHashChanged(ctx context.Context, rec *Record) (bool, bool, error) {
	return false, false, nil
}

func (qc *NullGeocoder) Query(ctx context.Context, req *QueryRequest, pg_opts pagination.Options) ([]*geojson.Feature, pagination.Results, error) {

	pg_rsp, err := countable.NewResultsFromCountWithOptions(pg_opts, 0)

	if err != nil {
		return nil, nil, err
	}

	rsp := make([]*geojson.Feature, 0)
	return rsp, pg_rsp, nil
}

func (gc *NullGeocoder) Flush(ctx context.Context) error {
	return nil
}

func (gc *NullGeocoder) Close() error {
	return nil
}
