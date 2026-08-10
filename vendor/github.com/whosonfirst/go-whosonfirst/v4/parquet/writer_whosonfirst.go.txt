package parquet

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/url"

	"github.com/whosonfirst/go-writer/v3"
)

type ParquetWhosOnFirstWriter struct {
	writer.Writer
	writer     writer.Writer
	parquet_wr ParquetWriter
}

func NewParquetWhosOnFirstWriter(ctx context.Context, uri string) (writer.Writer, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, err
	}

	q := u.Query()

	if !q.Has("writer-uri") {
		return nil, fmt.Errorf("Missing ?writer-uri= parameter")
	}

	wr_uri := q.Get("writer-uri")

	wr, err := writer.NewWriter(ctx, wr_uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to create new writer, %w", err)
	}

	pw, err := NewWriterWithIoWriteCloser(ctx, wr)

	if err != nil {
		return nil, fmt.Errorf("Failed to create new parquet writer, %w", err)
	}

	wof_wr := &ParquetWhosOnFirstWriter{
		writer:     wr,
		parquet_wr: pw,
	}

	return wof_wr, nil
}

func (wof_wr *ParquetWhosOnFirstWriter) Write(ctx context.Context, key string, r io.ReadSeeker) (int64, error) {

	var record *Record

	dec := json.NewDecoder(record)
	err := dec.Decode(&record)

	if err != nil {
		return 0, fmt.Errorf("Failed to decode record, %w", err)
	}

	rows := []*Record{
		r,
	}

	i, err := wof_wr.parquet_wr.Write(rows)

	i64 := int64(i)

	if err != nil {
		return i64, err
	}

	return i64, nil
}

func (wof_wr *ParquetWhosOnFirstWriter) WriteURI(ctx context.Context, key string) string {
	return key
}

func (wof_wr *ParquetWhosOnFirstWriter) Close(ctx context.Context) error {

	err := wof_wr.parquet_wr.Close()

	if err != nil {
		return fmt.Errorf("Failed to close parquet writer, %w", err)
	}

	err = wof_wr.writer.Close(ctx)

	if err != nil {
		return fmt.Errorf("Failed to close writer, %w", err)
	}

	return nil
}
