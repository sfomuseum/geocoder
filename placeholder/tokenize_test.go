package placeholder

import (
	"strings"
	"testing"
)

func TestTokenize(t *testing.T) {

	test := []string{
		"St. Louis, MO",
		"北京",
		"New York",
		"São Paulo",
		"Москва",
	}

	for _, s := range test {
		tok := Tokenize(s)
		println("Input:", s)
		println("Tokens:", strings.Join(tok, ", "))
	}
}
