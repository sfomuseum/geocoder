package whosonfirst

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"
	"sync"

	"github.com/paulmach/orb/planar"
	"github.com/whosonfirst/go-openpois"
	"github.com/whosonfirst/go-reader/v2"
	"github.com/whosonfirst/go-whosonfirst/v4/feature/properties"
	wof_reader "github.com/whosonfirst/go-whosonfirst/v4/reader"
	"github.com/whosonfirst/go-whosonfirst/v4/spatial/database"
	"github.com/whosonfirst/go-whosonfirst/v4/spatial/filter"
	"github.com/whosonfirst/go-whosonfirst/v4/spatial/hierarchy"
)

var parent_cache = new(sync.Map)

type AppendWhosOnFirstPropertiesOptions struct {
	Database database.SpatialDatabase
	Resolver *hierarchy.PointInPolygonHierarchyResolver
	Reader   reader.Reader
}

func AppendWhosOnFirstProperties(ctx context.Context, rec *openpois.Record, opts *AppendWhosOnFirstPropertiesOptions) error {

	logger := slog.Default()
	logger = logger.With("id", rec.UnifiedID)
	logger = logger.With("name", rec.PrimaryName())

	geom, err := rec.ToOrbGeometry()

	if err != nil {
		logger.Warn("Failed to derive geometry", "error", err)
		return err
	}

	pt, _ := planar.CentroidArea(geom)

	logger = logger.With("centroid", pt)

	if opts.Database == nil && opts.Resolver == nil {
		return fmt.Errorf("Both database and hierarchy resolver are nil")
	}

	if opts.Resolver == nil {

		resolver_opts := &hierarchy.PointInPolygonHierarchyResolverOptions{
			Database: opts.Database,
		}

		resolver, err := hierarchy.NewPointInPolygonHierarchyResolver(ctx, resolver_opts)

		if err != nil {
			return err
		}

		opts.Resolver = resolver
	}

	inputs := &filter.SPRInputs{
		IsCurrent: []int64{
			1,
		},
	}

	rsp, err := opts.Resolver.PointInPolygonWithPoint(ctx, inputs, &pt, "venue")

	if err != nil {
		logger.Error("Failed to point-in-polygon", "error", err)
		return err
	}

	switch len(rsp) {
	case 0:
		rec.WhosOnFirstParentId = -1
		rec.WhosOnFirstCountry = "XY"
	case 1:

		spr := rsp[0]

		parent_id, err := strconv.ParseInt(spr.Id(), 10, 64)

		if err != nil {
			logger.Error("Failed to parse parent ID", "parent id", spr.Id(), "error", err)
		} else {
			rec.WhosOnFirstParentId = parent_id
		}

		logger = logger.With("parent id", parent_id)

		v, ok := parent_cache.Load(parent_id)

		if ok {
			parent_body := v.([]byte)
			rec.WhosOnFirstHierarchies = properties.Hierarchies(parent_body)
			rec.WhosOnFirstCountry = properties.Country(parent_body)
		} else {

			parent_body, err := wof_reader.LoadBytes(ctx, opts.Reader, parent_id)

			if err != nil {
				logger.Error("Failed to retrieve parent body", "error", err)
			} else {
				parent_cache.Store(parent_id, parent_body)

				rec.WhosOnFirstHierarchies = properties.Hierarchies(parent_body)
				rec.WhosOnFirstCountry = properties.Country(parent_body)
			}
		}

		// slog.Info("Hiers", "h", rec.WhosOnFirstHierarchies)
	default:

		rec.WhosOnFirstParentId = -1
		rec.WhosOnFirstCountry = "XY"

		for i, r := range rsp {
			logger.Info("Multiple hierarchies", "i", i, "r", r.Id(), "n", r.Name(), "p", r.Placetype())
		}
	}

	return nil
}
