package parquet

import (
	"encoding/json"
	"fmt"
	"io"

	_ "github.com/whosonfirst/go-whosonfirst/v4/geojson"

	"github.com/paulmach/orb/encoding/wkb"
	"github.com/paulmach/orb/geojson"
)

type Record struct {
	Id         int64  `parquet:"id,type=int64"`
	ParentId   int64  `parquet:"parent_id,type="int64"`
	Placetype  string `parquet:"placetype,type=dict,zstd"`
	Country    string `parquet:"country,type=dict,zstd"`
	Geometry   []byte `parquet:"geometry,geometry"`
	Properties []byte `parquet:"properties,json"`
}

func (r *Record) AsGeoJSON() (*geojson.Feature, error) {

	geom, err := wkb.Unmarshal(r.Geometry)

	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal geometry, %w", err)
	}

	var props map[string]any

	err = json.Unmarshal(r.Properties, &props)

	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal properties, %w", err)
	}

	f := geojson.NewFeature(geom)
	f.Properties = props

	return f, nil
}

func (r *Record) AsGeoJSONBytes() ([]byte, error) {

	f, err := r.AsGeoJSON()

	if err != nil {
		return nil, err
	}

	return json.Marshal(f)
}

func RecordFromGeoJSONReader(r io.Reader) (*Record, error) {

	body, err := io.ReadAll(r)

	if err != nil {
		return nil, err
	}

	return RecordFromGeoJSONBytes(body)
}

func RecordFromGeoJSONBytes(body []byte) (*Record, error) {

	f, err := geojson.UnmarshalFeature(body)

	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal feature, %w", err)
	}

	return RecordFromGeoJSONFeature(f)
}

func RecordFromGeoJSONFeature(f *geojson.Feature) (*Record, error) {

	geom, err := wkb.Marshal(f.Geometry, wkb.DefaultByteOrder)

	if err != nil {
		return nil, fmt.Errorf("Failed to marshal geometry, %w", err)
	}

	id_fl64 := f.Properties.MustFloat64("wof:id", -1)

	if id_fl64 == -1 {
		return nil, fmt.Errorf("Failed to determine wof:id")
	}

	pid_fl64 := f.Properties.MustFloat64("wof:parent_id", -1)

	id := int64(id_fl64)
	pid := int64(pid_fl64)

	pt := f.Properties.MustString("wof:placetype", "custom")
	co := f.Properties.MustString("wof:country", "XX")

	props, err := json.Marshal(f.Properties)

	record := &Record{
		Id:         id,
		ParentId:   pid,
		Placetype:  pt,
		Country:    co,
		Geometry:   geom,
		Properties: props,
	}

	return record, nil
}

/*

memory D SELECT id, parent_id, placetype, properties."wof:name" AS name, ST_Centroid(geometry) FROM read_parquet('test.parquet') WHERE properties->>'$.sfomuseum:placetype'='airport' AND country = 'CA' LIMIT 5;
┌───────────┬────────────┬───────────┬────────────────────────────────────┬─────────────────────────────────────┐
│    id     │ parent_id  │ placetype │                name                │        st_centroid(geometry)        │
│   int64   │   int64    │  varchar  │                json                │        geometry('ogc:crs84')        │
├───────────┼────────────┼───────────┼────────────────────────────────────┼─────────────────────────────────────┤
│ 102554351 │  101736475 │ campus    │ "Montreal-Pierre Elliott Trudeau I │ POINT (-73.7436343824761 45.4673683 │
│           │            │           │ nternational Airport"              │ 15411896)                           │
├───────────┼────────────┼───────────┼────────────────────────────────────┼─────────────────────────────────────┤
│ 102554473 │         -1 │ campus    │ "Edmonton International Airport"   │ POINT (-113.58489936246232 53.30860 │
│           │            │           │                                    │ 794988534)                          │
├───────────┼────────────┼───────────┼────────────────────────────────────┼─────────────────────────────────────┤
│ 102554519 │   85682123 │ campus    │ "Goose Bay Airport"                │ POINT (-60.425833 53.319167)        │
├───────────┼────────────┼───────────┼────────────────────────────────────┼─────────────────────────────────────┤
│ 102554947 │ 1108964229 │ campus    │ "Calgary International Airport"    │ POINT (-114.00691541254936 51.12049 │
│           │            │           │                                    │ 9437034056)                         │
├───────────┼────────────┼───────────┼────────────────────────────────────┼─────────────────────────────────────┤
│ 102555091 │ 1108905413 │ campus    │ "Toronto Lester B Pearson Internat │ POINT (-79.63120282896976 43.676962 │
│           │            │           │ ional Airport"                     │ 210845375)                          │
└───────────┴────────────┴───────────┴────────────────────────────────────┴─────────────────────────────────────┘

*/
