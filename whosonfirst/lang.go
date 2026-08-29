package whosonfirst

import (
	"fmt"
	"strings"
	"sync"

	"github.com/whosonfirst/go-rfc-5646/tags"
)

var langtag_map = new(sync.Map)

func ParseLangTag(lang_str string) (string, string, error) {

	v, ok := langtag_map.Load(lang_str)

	if ok {

		switch v.(type) {
		case [2]string:
			lt := v.([2]string)
			return lt[0], lt[1], nil
		case error:
			return "", "", v.(error)
		default:
			return "", "", fmt.Errorf("Unexpected cache type for '%s', %v", lang_str, v)
		}
	}

	lang_tag, err := tags.NewLangTag(lang_str)

	var lang string
	var tag string

	if err != nil {

		parts := strings.Split(lang_str, "_x_")

		if len(parts) != 2 {
			err := fmt.Errorf("Failed to parse language tag, %w", err)
			langtag_map.Store(lang_str, err)
			return "", "", err
		}

		lang = parts[0]
		tag = parts[1]

	} else {

		lang = lang_tag.Language()
		tag = lang_tag.PrivateUse()
	}

	langtag_map.Store(lang_str, [2]string{lang, tag})
	return lang, tag, nil
}
