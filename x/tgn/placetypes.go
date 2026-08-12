package tgn

import (
	"encoding/json"
	"log/slog"
	"sync"
)

var loadPlacetypesMap = sync.OnceValues(func() (map[string]string, error) {

	r, err := FS.Open("placetypes.json")

	if err != nil {
		return nil, err
	}

	var pt_map map[string]string

	dec := json.NewDecoder(r)
	err = dec.Decode(&pt_map)

	if err != nil {
		return nil, err
	}

	return pt_map, nil
})

func TgnToWhosOnFirstPlacetypeMap() (map[string]string, error) {
	return loadPlacetypesMap()
}

func TgnToWhosOnFirstPlacetype(tgn_pt string) string {

	pt_map, err := loadPlacetypesMap()

	if err != nil {
		slog.Error("Unable to load placetypes map", "error", err)
		return "custom"
	}

	pt, ok := pt_map[tgn_pt]

	if !ok {
		// slog.Warn("TGN placetype not found", "pt", tgn_pt)
		return "custom"
	}

	return pt
}
