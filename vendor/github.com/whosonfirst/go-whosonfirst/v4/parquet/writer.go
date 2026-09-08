package parquet

import (
	"context"
	"io"

	sfom_parquet "github.com/sfomuseum/go-parquet"
)

func NewWriter(ctx context.Context, uri string) (*sfom_parquet.ParquetWriter[*Record], error) {
	return sfom_parquet.NewWriter[*Record](ctx, uri)
}

func WriteFromReader(pw *sfom_parquet.ParquetWriter[*Record], r io.Reader) (int, error) {

	record, err := RecordFromGeoJSONReader(r)

	if err != nil {
		return 0, err
	}

	return pw.Write([]*Record{record})
}
