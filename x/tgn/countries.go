package tgn

import (
	"encoding/json"
	"log/slog"
	"sync"
)

var loadCountriesMap = sync.OnceValues(func() (map[string]string, error) {

	r, err := FS.Open("countries.json")

	if err != nil {
		return nil, err
	}

	var co_map map[string]string

	dec := json.NewDecoder(r)
	err = dec.Decode(&co_map)

	if err != nil {
		return nil, err
	}

	return co_map, nil
})

func TgnToWhosOnFirstCountryMap() (map[string]string, error) {
	return loadCountriesMap()
}

func TgnToWhosOnFirstCountry(tgn_co string) string {

	co_map, err := loadCountriesMap()

	if err != nil {
		slog.Error("Unable to load countries map", "error", err)
		return "XZ"
	}

	pt, ok := co_map[tgn_co]

	if !ok {
		// slog.Warn("TGN country not found", "pt", tgn_pt)
		return "XZ"
	}

	return pt
}
