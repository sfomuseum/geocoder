package parquet

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"

	parquet_go "github.com/parquet-go/parquet-go"
)

// nopWriteCloser is an io.WriteCloser that does nothing on Close().
type nopWriteCloser struct{ io.Writer }

func (nopWriteCloser) Close() error { return nil }

// NopWriteCloser returns an io.WriteCloser that wraps w and whose Close() is a no‑op.
func NopWriteCloser(w io.Writer) io.WriteCloser {
	return nopWriteCloser{w}
}

// ParquetWriter is a convenience struct for wrapping the creation of both a Parquet "GenericWriter"
// and the underlying [io.Writer] instance that it writes to.
type ParquetWriter struct {
	writer         io.WriteCloser
	parquet_writer *parquet_go.GenericWriter[*Record]
	batch_size     int
	buffer         []*Record
	mu             *sync.RWMutex
}

// NewWriter returns a new [ParquetWriter] instance configured using 'uri'. If 'uri' is "-"
// then data written (to the writer) will be dispatched to STDOUT. Otherwise 'uri' will be
// treated as the path to a file on the local filesystem.
func NewWriter(ctx context.Context, uri string) (*ParquetWriter, error) {

	var wr io.WriteCloser

	switch uri {
	case "-":
		wr = NopWriteCloser(os.Stdout)
	default:

		abs_uri, err := filepath.Abs(uri)

		if err != nil {
			return nil, err
		}

		w, err := os.OpenFile(abs_uri, os.O_RDWR|os.O_CREATE, 0644)

		if err != nil {
			return nil, fmt.Errorf("Failed to open %s for writing, %w", uri, err)
		}

		wr = w
	}

	return NewWriterWithIoWriteCloser(ctx, wr)
}

func NewWriterWithIoWriteCloser(ctx context.Context, wr io.WriteCloser) (*ParquetWriter, error) {

	p_wr := parquet_go.NewGenericWriter[*Record](wr)

	mu := new(sync.RWMutex)
	buf := make([]*Record, 0)

	pw := &ParquetWriter{
		writer:         wr,
		parquet_writer: p_wr,
		batch_size:     10000,
		buffer:         buf,
		mu:             mu,
	}

	return pw, nil
}

func (pw *ParquetWriter) WriteFromReader(r io.Reader) (int, error) {

	record, err := RecordFromGeoJSONReader(r)

	if err != nil {
		return 0, err
	}

	return pw.Write([]*Record{record})
}

// Write will dispatch 'rows' to the underlying Parquet `GenericWriter` instance.
func (pw *ParquetWriter) Write(rows []*Record) (int, error) {

	pw.mu.Lock()
	defer pw.mu.Unlock()

	pw.buffer = append(pw.buffer, rows...)

	var n int
	var err error

	if len(pw.buffer) >= pw.batch_size {

		n, err = pw.writeBuffer()

		if err != nil {
			return 0, err
		}
	}

	return n, err
}

// Flush will invoke the  underlying Parquet `GenericWriter` instance's `Flush` method.
func (pw *ParquetWriter) Flush() error {
	return pw.parquet_writer.Flush()
}

// Writer returns the underlying [io.WriteCloser] instance.
func (pw *ParquetWriter) Writer() io.WriteCloser {
	return pw.writer
}

// ParquetWriter returns the underlying	Parquet	`GenericWriter`	instance.
func (pw *ParquetWriter) ParquetWriter() *parquet_go.GenericWriter[*Record] {
	return pw.parquet_writer
}

// Close will flush any remaining output and close both the underlying Parquet `GenericWriter`
// and [io.WriteCloser] instances.
func (pw *ParquetWriter) Close() error {

	pw.mu.Lock()
	defer pw.mu.Unlock()

	_, err := pw.writeBuffer()

	if err != nil {
		return err
	}

	pw.parquet_writer.Flush()

	err = pw.parquet_writer.Close()

	if err != nil {
		return fmt.Errorf("Failed to close Parquet writer, %w", err)
	}

	err = pw.writer.Close()

	if err != nil {
		return fmt.Errorf("Failed to close writer, %w", err)
	}

	return nil
}

func (pw *ParquetWriter) writeBuffer() (int, error) {

	var n int
	var err error

	if len(pw.buffer) > 0 {

		n, err = pw.parquet_writer.Write(pw.buffer)

		if err != nil {
			return n, err
		}

		pw.buffer = make([]*Record, 0)
	}

	return n, nil
}
