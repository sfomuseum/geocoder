package parquet

import (
	"errors"
	"fmt"
	"io"

	"github.com/parquet-go/parquet-go"
)

type Id interface {
	string | int64
}

// FindRecordById scans any parquet schema by an ID field.
// T: Represents the Go struct model to return (e.g., Record)
// K: Represents the search key type (string or int64)
func FindRecordById[T any, K Id](uri string, targetID K) (T, error) {

	var zero T

	file, sz, err := OpenURI(uri)

	if err != nil {
		return zero, err
	}

	defer file.Close()

	pf, err := parquet.OpenFile(file, sz)

	if err != nil {
		return zero, err
	}

	return FindRecordByIdWithParquetFile[T, K](file, pf, targetID)
}

func FindRecordByIdWithParquetFile[T any, K Id](file io.ReaderAt, pf *parquet.File, id K) (T, error) {

	var zero T

	id_idx := -1

	fields := pf.Schema().Fields()

	for i, field := range fields {

		if field.Name() == "id" {
			id_idx = i
			break
		}
	}

	if id_idx == -1 {
		id_idx = 0
	}

	// Prepare a parquet.Value representation of our target ID for Bloom filter evaluation

	var parquet_v parquet.Value

	switch v := any(id).(type) {
	case int64:
		parquet_v = parquet.Int64Value(v)
	case string:
		parquet_v = parquet.ByteArrayValue([]byte(v))
	}

	var row_idx int64 = 0

	for _, rg := range pf.RowGroups() {

		chunks := rg.ColumnChunks()

		if len(chunks) <= id_idx {
			continue
		}

		id_chunk := chunks[id_idx]

		bf := id_chunk.BloomFilter()

		if bf != nil {

			hasValue, err := bf.Check(parquet_v)

			if err == nil && !hasValue {
				row_idx += rg.NumRows()
				continue
			}
		}

		id_pages := id_chunk.Pages()

		for {

			p, err := id_pages.ReadPage()

			if err != nil {

				if errors.Is(err, io.EOF) {
					break
				}

				id_pages.Close()
				return zero, err
			}

			// Branch depending on key type for optimized vector scans
			switch any(id).(type) {
			case int64:

				target := any(id).(int64)

				int64_r, ok := p.Values().(parquet.Int64Reader)

				if ok {

					id_buf := make([]int64, p.NumValues())
					n, _ := int64_r.ReadInt64s(id_buf)

					for i := 0; i < n; i++ {

						if id_buf[i] == target {
							id_pages.Close()
							exactRowIndex := row_idx + int64(i)
							return fetchFullRowAt[T](file, exactRowIndex)
						}
					}

					row_idx += int64(n)
					continue
				}

			case string:

				target := any(id).(string)
				// Strings use the generic underlying Value scanner
				values := make([]parquet.Value, p.NumValues())

				n, _ := p.Values().ReadValues(values)

				for i := 0; i < n; i++ {
					if values[i].String() == target {
						id_pages.Close()
						exact_idx := row_idx + int64(i)
						return fetchFullRowAt[T](file, exact_idx)
					}
				}
				row_idx += int64(n)
				continue
			}

			// Catch-all structural fallback evaluation
			values := make([]parquet.Value, p.NumValues())
			n, _ := p.Values().ReadValues(values)
			for i := 0; i < n; i++ {
				match := false

				int_v, ok := any(id).(int64)

				if ok && values[i].Int64() == int_v {
					match = true
				} else if str_v, ok := any(id).(string); ok && values[i].String() == str_v {
					match = true
				}

				if match {
					id_pages.Close()
					exact_idx := row_idx + int64(i)
					return fetchFullRowAt[T](file, exact_idx)
				}
			}

			row_idx += int64(n)
		}

		id_pages.Close()
	}

	return zero, fmt.Errorf("id not found in parquet archive")
}

// Reconstructs only the requested matching row from all column chunks
func fetchFullRowAt[T any](file io.ReaderAt, row_idx int64) (T, error) {

	var null_row T

	reader := parquet.NewGenericReader[T](file)
	defer reader.Close()

	err := reader.SeekToRow(row_idx)

	if err != nil {
		return null_row, err
	}

	rows := make([]T, 1)
	n, err := reader.Read(rows)

	if n > 0 {
		return rows[0], nil
	}

	if err != nil && !errors.Is(err, io.EOF) {
		return null_row, err
	}

	return null_row, fmt.Errorf("failed to hydrate record at row %d", row_idx)
}
