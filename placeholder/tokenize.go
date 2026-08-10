package placeholder // or your own package name

import (
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

// ---------------------------------------------------------------------
// Normalisation helpers
// ---------------------------------------------------------------------
func normalizeNFKC(s string) string {
	return norm.NFKC.String(s)
}

// ---------------------------------------------------------------------
// Regular expressions
// ---------------------------------------------------------------------
// Characters that are written from “major to minor” (Arabic, Hebrew,
// Chinese, Cyrillic, etc.).  The regex contains the **real
// Unicode characters**, not escape sequences.
var majorMinor = regexp.MustCompile(`[\x{0591}-\x{07FF}\x{1100}-\x{11FF}\x{3130}-\x{318F}\x{A960}-\x{A97F}\x{AC00}-\x{D7AF}\x{D7B0}-\x{D7FF}\x{0400}-\x{04FF}]`)

// ---------------------------------------------------------------------
// Tokeniser
// ---------------------------------------------------------------------

func Normalize(input string) string {

	// 1. Normalise
	input = normalizeNFKC(input)

	// 2. Trim
	input = strings.TrimSpace(input)

	// 3. Replace punctuation (except apostrophes) with a space
	input = strings.Map(func(r rune) rune {
		if r == '\'' || r == '’' { // keep apostrophes
			return r
		}
		if unicode.IsPunct(r) {
			return ' '
		}
		return r
	}, input)

	return input
}

func Tokenize(input string) []string {

	input = Normalize(input)

	input = strings.Join(strings.Fields(input), " ")

	tokens := strings.Split(input, " ")

	for i, t := range tokens {
		tokens[i] = strings.ToLower(t)
	}

	if len(tokens) > 0 && majorMinor.MatchString(tokens[0]) {
		for i, j := 0, len(tokens)-1; i < j; i, j = i+1, j-1 {
			tokens[i], tokens[j] = tokens[j], tokens[i]
		}
	}

	return tokens
}
