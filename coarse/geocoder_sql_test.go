package coarse

import (
	"context"
	"encoding/json"
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

		var rec *Record

		err = json.Unmarshal(body, &rec)

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

	db, err := testCreateDatabase(ctx, "../fixtures/sf-coarse.json")

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

	db, err := testCreateDatabase(ctx, "../fixtures/sf-coarse.json")

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

	rec_id := rec.Properties.MustString("geocoder:id", "")

	if rec_id != "wof:id=85922583" {
		t.Fatalf("Invalid ID: '%v'", rec_id)
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
		BelongsTo: []string{
			"wof:id=102087579",
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
	rec6_id := rec6.Properties.MustString("geocoder:id", "")

	if rec6_id != "wof:id=85922583" {
		t.Fatalf("Invalid ID: '%s' (6)", rec6_id)
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
		"../fixtures/sfo-coarse.json",
		"../fixtures/sf-coarse.json",
		"../fixtures/ca-coarse.json",
		"../fixtures/nyc-coarse.json",
		"../fixtures/ny-coarse.json",
		"../fixtures/us-coarse.json",
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
	sfo_rec_id := sfo_rec.Properties.MustString("geocoder:id", "")

	if sfo_rec_id != "wof:id=102527513" {
		t.Fatalf("Invalid ID: '%s'", sfo_rec_id)
	}

	sfo_label := sfo_rec.Properties.MustString("geocoder:label", "")

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

	nyc_id := nyc_rec.Properties.MustString("geocoder:id", "")
	nyc_label := nyc_rec.Properties.MustString("geocoder:label", "")

	if nyc_id != "wof:id=85977539" {
		t.Fatalf("Invalid ID: '%s'", nyc_id)
	}

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
		"../fixtures/t3-coarse.json",
		"../fixtures/sf-coarse.json",
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

	rsp1_id := rsp1[0].Properties.MustString("geocoder:id", "")

	if rsp1_id != "wof:id=1477855699" {
		t.Fatalf("Unexpected ID '%s'", rsp1_id)
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
