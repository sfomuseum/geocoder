package coarse

import (
	"context"
	"os"
	"testing"

	"github.com/aaronland/go-pagination/countable"
	"github.com/sfomuseum/go-edtf/unix"
	"github.com/whosonfirst/go-whosonfirst/v4/flags/existential"
)

func testCreateDatabase(ctx context.Context, to_index ...string) (Geocoder, error) {

	db, err := NewSQLGeocoder(ctx, "sql://sqlite?dsn=:memory:")

	if err != nil {
		return nil, err
	}

	err = db.PreIndex(ctx)

	if err != nil {
		return nil, err
	}

	for _, path := range to_index {

		body, err := os.ReadFile(path)

		if err != nil {
			return nil, err
		}

		rec, err := NewWhosOnFirstRecord(ctx, body)

		if err != nil {
			return nil, err
		}

		err = db.AddRecord(ctx, rec)

		if err != nil {
			return nil, err
		}
	}

	err = db.PostIndex(ctx)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func TestSQLGeocoder(t *testing.T) {

	ctx := context.Background()

	db, err := testCreateDatabase(ctx)

	if err != nil {
		t.Fatalf("Failed to create SQL geocoder, %v", err)
	}

	err = db.Close()

	if err != nil {
		t.Fatalf("Failed to close database, %v", err)
	}
}

func TestSQLGeocoderIndex(t *testing.T) {

	ctx := context.Background()

	db, err := testCreateDatabase(ctx, "../fixtures/sf.geojson")

	if err != nil {
		t.Fatalf("Failed to create SQL geocoder, %v", err)
	}

	err = db.Close()

	if err != nil {
		t.Fatalf("Failed to close database, %v", err)
	}
}

func TestSQLGeocoderQuery(t *testing.T) {

	ctx := context.Background()

	db, err := testCreateDatabase(ctx, "../fixtures/sf.geojson")

	if err != nil {
		t.Fatalf("Failed to create SQL geocoder, %v", err)
	}

	pg_opts, err := countable.NewCountableOptions()

	if err != nil {
		t.Fatalf("Failed to create pagination options, %v", err)
	}

	req := &QueryRequest{
		Query: "San Francisco",
	}

	rsp, _, err := db.Query(ctx, req, pg_opts)

	if err != nil {
		t.Fatalf("Failed to query database (1), %v", err)
	}

	if len(rsp) != 1 {
		t.Fatalf("Expected count of 1 but got %d", len(rsp))
	}

	rec := rsp[0]

	if rec.ID != int64(85922583) {
		t.Fatalf("Invalid ID: %d", rec.ID)
	}

	req2 := &QueryRequest{
		Query: "San Francisco",
		Placetype: []string{
			"region",
		},
	}

	rsp2, _, err := db.Query(ctx, req2, pg_opts)

	if err != nil {
		t.Fatalf("Failed to query database (2), %v", err)
	}

	if len(rsp2) != 0 {
		t.Fatalf("Expected count of 0 but got %d (2)", len(rsp))
	}

	req3 := &QueryRequest{
		Query: "SF",
		BelongsTo: []int64{
			102087579,
		},
	}

	rsp3, _, err := db.Query(ctx, req3, pg_opts)

	if err != nil {
		t.Fatalf("Failed to query database (3), %v", err)
	}

	if len(rsp3) != 1 {
		t.Fatalf("Expected count of 1 but got %d (3)", len(rsp3))
	}

	//

	is_current, err := existential.NewKnownUnknownFlag(1)

	if err != nil {
		t.Fatalf("Failed to create existential flag, %v", err)
	}

	req4 := &QueryRequest{
		Query:     "Frisco",
		IsCurrent: is_current,
	}

	rsp4, _, err := db.Query(ctx, req4, pg_opts)

	if err != nil {
		t.Fatalf("Failed to query database (4), %v", err)
	}

	if len(rsp4) != 1 {
		t.Fatalf("Expected count of 1 but got %d (4)", len(rsp4))
	}

	//

	not_current, err := existential.NewKnownUnknownFlag(0)

	if err != nil {
		t.Fatalf("Failed to create existential flag, %v", err)
	}

	req5 := &QueryRequest{
		Query:     "Frisco",
		IsCurrent: not_current,
	}

	rsp5, _, err := db.Query(ctx, req5, pg_opts)

	if err != nil {
		t.Fatalf("Failed to query database (5), %v", err)
	}

	if len(rsp5) != 0 {
		t.Fatalf("Expected count of 0 but got %d (5)", len(rsp5))
	}

	//

	req6 := &QueryRequest{
		Query: "サンフランシスコ",
	}

	rsp6, _, err := db.Query(ctx, req6, pg_opts)

	if err != nil {
		t.Fatalf("Failed to query database (6), %v", err)
	}

	if len(rsp6) != 1 {
		t.Fatalf("Expected count of 1 but got %d (6)", len(rsp))
	}

	rec6 := rsp6[0]

	if rec6.ID != int64(85922583) {
		t.Fatalf("Invalid ID: %d (6)", rec6.ID)
	}

	//

	err = db.Close()

	if err != nil {
		t.Fatalf("Failed to close database, %v", err)
	}
}

func TestSQLGeocoderLabels(t *testing.T) {

	ctx := context.Background()

	to_index := []string{
		"../fixtures/sfo.geojson",
		"../fixtures/sf.geojson",
		"../fixtures/ca.geojson",
		"../fixtures/nyc.geojson",
		"../fixtures/ny.geojson",
		"../fixtures/us.geojson",
	}

	db, err := testCreateDatabase(ctx, to_index...)

	if err != nil {
		t.Fatalf("Failed to create SQL geocoder, %v", err)
	}

	pg_opts, err := countable.NewCountableOptions()

	if err != nil {
		t.Fatalf("Failed to create pagination options, %v", err)
	}

	// SFO label

	sfo_req := &QueryRequest{
		Query: "SFO",
	}

	sfo_rsp, _, err := db.Query(ctx, sfo_req, pg_opts)

	if err != nil {
		t.Fatalf("SFO query failed, %v", err)
	}

	if len(sfo_rsp) != 1 {
		t.Fatalf("Expected count of 1 but got %d", len(sfo_rsp))
	}

	sfo_rec := sfo_rsp[0]

	if sfo_rec.ID != int64(102527513) {
		t.Fatalf("Invalid ID: %d", sfo_rec.ID)
	}

	sfo_label := sfo_rec.Properties.MustString("wof:label", "")

	if sfo_label != "San Francisco International Airport, San Francisco, California, US" {
		t.Fatalf("Invalid label for SFO: '%s'", sfo_label)
	}

	// NYC labels

	nyc_req := &QueryRequest{
		Query: "NYC",
	}

	nyc_rsp, _, err := db.Query(ctx, nyc_req, pg_opts)

	if err != nil {
		t.Fatalf("NYC query failed, %v", err)
	}

	if len(nyc_rsp) != 1 {
		t.Fatalf("Expected count of 1 but got %d", len(nyc_rsp))
	}

	nyc_rec := nyc_rsp[0]

	if nyc_rec.ID != int64(85977539) {
		t.Fatalf("Invalid ID: %d", nyc_rec.ID)
	}

	nyc_label := nyc_rec.Properties.MustString("wof:label", "")

	if nyc_label != "New York, New York, US" {
		t.Fatalf("Invalid label for NYC: '%s'", nyc_label)
	}

	// Clean up

	err = db.Close()

	if err != nil {
		t.Fatalf("Failed to close database, %v", err)
	}

}

func TestSQLGeocoderDates(t *testing.T) {

	ctx := context.Background()

	to_index := []string{
		"../fixtures/t3.geojson",
		"../fixtures/sf.geojson",
	}

	db, err := testCreateDatabase(ctx, to_index...)

	if err != nil {
		t.Fatalf("Failed to create SQL geocoder, %v", err)
	}

	//

	ok, date_starts, err := unix.DeriveRanges("2017/2018")

	if err != nil {
		t.Fatalf("Failed to parse start range, %v", err)
	}

	if !ok {
		t.Fatalf("Failed to derive start range")
	}

	ok, date_ends, err := unix.DeriveRanges("2020")

	if err != nil {
		t.Fatalf("Failed to parse end range, %v", err)
	}

	if !ok {
		t.Fatalf("Failed to derive end range")
	}

	req1 := &QueryRequest{
		Query:      "T3",
		DateStarts: date_starts,
		DateEnds:   date_ends,
	}

	pg_opts, err := countable.NewCountableOptions()

	if err != nil {
		t.Fatalf("Failed to create pagination options, %v", err)
	}

	rsp1, _, err := db.Query(ctx, req1, pg_opts)

	if err != nil {
		t.Fatalf("Failed to query database, %v (1)", err)
	}

	if len(rsp1) != 1 {
		t.Fatalf("Expected 1 result but got %d (1)", len(rsp1))
	}

	if rsp1[0].ID != int64(1477855699) {
		t.Fatalf("Unexpected ID %d", rsp1[0].ID)
	}

	//

	ok, date_starts, err = unix.DeriveRanges("2017/2018")

	if err != nil {
		t.Fatalf("Failed to parse start range, %v", err)
	}

	if !ok {
		t.Fatalf("Failed to derive start range")
	}

	ok, date_ends, err = unix.DeriveRanges("2019")

	if err != nil {
		t.Fatalf("Failed to parse end range, %v", err)
	}

	if !ok {
		t.Fatalf("Failed to derive end range")
	}

	req2 := &QueryRequest{
		Query:      "T3",
		DateStarts: date_starts,
		DateEnds:   date_ends,
	}

	rsp2, _, err := db.Query(ctx, req2, pg_opts)

	if err != nil {
		t.Fatalf("Failed to query database, %v (2)", err)
	}

	if len(rsp2) != 0 {
		t.Fatalf("Expected 0 results but got %d (2)", len(rsp2))
	}

	//

	err = db.Close()

	if err != nil {
		t.Fatalf("Failed to close database, %v", err)
	}

}
