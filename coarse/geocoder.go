package coarse

import (
	"context"

	"github.com/aaronland/go-pagination"
	"github.com/paulmach/orb/geojson"
)

// To do: Replace QueryRequest with some sort of generate QueryOptions interface
// and move this in to the root geocoder package.

type Geocoder interface {
	PreIndex(context.Context) error
	PostIndex(context.Context) error
	AddRecord(context.Context, *Record) error
	RemoveRecord(context.Context, int64) error
	RecordExists(context.Context, int64) (bool, error)
	HasRecordHashChanged(context.Context, *Record) (bool, bool, error)
	Query(context.Context, *QueryRequest, pagination.Options) ([]*geojson.Feature, pagination.Results, error)
	Flush(context.Context) error
	Close() error
}
