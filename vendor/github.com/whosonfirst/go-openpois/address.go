package openpois

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	// Matches common standalone street suffix words to standardize abbreviations
	suffixRegex = regexp.MustCompile(`\b(drive|street|avenue|road|lane|court|boulevard|highway|way|place)\b`)

	suffixReplacements = map[string]string{
		"drive":     "dr",
		"street":    "st",
		"avenue":    "ave",
		"road":      "rd",
		"lane":      "ln",
		"court":     "ct",
		"boulevard": "blvd",
		"highway":   "hwy",
	}
)

// cleanAndNormalize simplifies string formatting for accurate fuzzy comparison
func cleanAndNormalize(s string) string {

	s = strings.ToLower(strings.TrimSpace(s))

	s = strings.TrimSuffix(s, ", us")
	s = strings.TrimSuffix(s, ", usa")

	s = suffixRegex.ReplaceAllStringFunc(s, func(match string) string {
		if sub, exists := suffixReplacements[match]; exists {
			return sub
		}
		return match
	})

	fields := strings.Fields(s)

	if len(fields) >= 2 {

		// If the first two fields are identical numbers, drop the first one
		if fields[0] == fields[1] {
			if _, err := strconv.Atoi(fields[0]); err == nil {
				fields = fields[1:]
			}
		}
	}

	// 4. Uniform whitespace cleanup
	return strings.Join(fields, " ")
}

// Address represents a clean, parsed components set to avoid structural duplicates
type Address struct {
	HouseNumber string
	Street      string
	Unit        string
	City        string
	State       string
	Postcode    string
	Country     string
}

// String normalizes the address components into a single readable line
func (a Address) String() string {

	var parts []string

	streetLine := a.Street

	if a.HouseNumber != "" {
		streetLine = a.HouseNumber + " " + streetLine
	}

	if a.Unit != "" {
		streetLine = streetLine + " " + a.Unit
	}

	streetLine = strings.TrimSpace(streetLine)

	if streetLine != "" {
		parts = append(parts, streetLine)
	}

	if a.City != "" {
		parts = append(parts, a.City)
	}

	stateZip := strings.TrimSpace(a.State + " " + a.Postcode)

	if stateZip != "" {
		parts = append(parts, stateZip)
	}

	if a.Country != "" {
		parts = append(parts, a.Country)
	}

	return strings.Join(parts, ", ")
}
