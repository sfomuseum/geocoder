package placeholder // or your own package name

import (
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

func normalizeNFKC(s string) string {
	return norm.NFKC.String(s)
}

// Characters that are written from “major to minor” (Arabic, Hebrew,
// Chinese, Cyrillic, etc.).  The regex contains the **real
// Unicode characters**, not escape sequences.
var majorMinor = regexp.MustCompile(`[\x{0591}-\x{07FF}\x{1100}-\x{11FF}\x{3130}-\x{318F}\x{A960}-\x{A97F}\x{AC00}-\x{D7AF}\x{D7B0}-\x{D7FF}\x{0400}-\x{04FF}]`)

// Normalize normalises an arbitrary string for tokenisation.  The
// normalisation pipeline includes:
//
//  1. NFKC Unicode normalisation
//  2. Trimming of whitespace
//  3. Replacement of punctuation (except apostrophes) with a space
//
// The function is idempotent – calling it repeatedly yields the
// same result – and is suitable for use with any Unicode text.
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

// Tokenize splits a string into a slice of lower‑case tokens suitable
// for full‑text search.  It first normalises the input, removes
// punctuation, splits on whitespace, and, for strings that start
// with a "major‑minor" character class (e.g. Arabic, Cyrillic),
// reverses the token order to optimise SQLite FTS5 ranking.
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
