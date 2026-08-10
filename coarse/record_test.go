package coarse

import (
	"context"
	"encoding/json"
	"io"
	"os"
	"testing"
)

func TestNewRecord(t *testing.T) {

	ctx := context.Background()

	body, err := os.ReadFile("../fixtures/sf.geojson")

	if err != nil {
		t.Fatalf("Failed to read data, %v", err)
	}

	r, err := NewWhosOnFirstRecord(ctx, body)

	if err != nil {
		t.Fatalf("Failed to create new record, %v", err)
	}

	enc := json.NewEncoder(io.Discard)
	err = enc.Encode(r)

	if err != nil {
		t.Fatalf("Failed to encode record, %v", err)
	}
}

func TestRecordHash(t *testing.T) {

	ctx := context.Background()

	body, err := os.ReadFile("../fixtures/sf.geojson")

	if err != nil {
		t.Fatalf("Failed to read data, %v", err)
	}

	r, err := NewWhosOnFirstRecord(ctx, body)

	if err != nil {
		t.Fatalf("Failed to create new record, %v", err)
	}

	h1, err := r.Hash()

	if err != nil {
		t.Fatalf("Failed to hash record, %v", err)
	}

	h2, err := r.Hash()

	if err != nil {
		t.Fatalf("Failed to hash record (2), %v", err)
	}

	if h1 != h2 {
		t.Fatalf("Record hashes not match")
	}
}
