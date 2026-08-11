package tgn

import (
	"sync"
	"encoding/json"
	"strings"
)

var loadLanguagesMap = sync.OnceValues(func() (map[string]string, error) {

	r, err := FS.Open("languages.json")

	if err != nil {
		return nil, err
	}

	var lang_map map[string]string
	
	dec := json.NewDecoder(r)
	err = dec.Decode(lang_map)

	if err != nil {
		return nil, err
	}

	return lang_map, nil
})

func TgnToWhosOnFirstLanguagesMap() (map[string]string, error) {
	return loadLanguagesMap()
}

func TgnToWhosOnFirstLanguage(tgn_lang string) (string, string) {

	lang_map, err := loadLanguagesMap()

	if err != nil {
		return "und", "preferred"
	}

	lang, ok := lang_map[tgn_lang]

	if !ok { 
		return "und", "preferred"
	}

	parts := strings.Split(lang, "-x-")

	if len(parts) == 1 {
		return parts[0], "preferred"
	}

	return parts[0], parts[1]
}
