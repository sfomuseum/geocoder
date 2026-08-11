package tgn

import (
	"fmt"
	"math"
)

func TgnToEdtfYear(tgn_year int) (string, error) {

	if tgn_year > 0 {
		// Positive year: format directly to a minimum of 4 digits padded with zeros
		return fmt.Sprintf("%04d", tgn_year), nil
	} else if tgn_year < 0 {
		// Negative year: shift by 1 for astronomical year zero alignment
		shifted := tgn_year + 1

		// Take the absolute value for padding math
		abs_year := int(math.Abs(float64(shifted)))

		// Return with an explicit negative sign prefix
		return fmt.Sprintf("-%04d", abs_year), nil
	} else {
		// TGN data uses -1 for 1 BCE and 1 for 1 CE; a raw 0 is invalid data
		return "", fmt.Errorf("TGN data should not contain a raw 0 integer")
	}
}
