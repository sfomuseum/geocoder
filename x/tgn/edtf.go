package tgn

import (
	"errors"
	"fmt"
	"math"
)

func TgnToEdtfYear(tgnInt int) (string, error) {

	if tgnInt > 0 {
		// Positive year: format directly to a minimum of 4 digits padded with zeros
		return fmt.Sprintf("%04d", tgnInt), nil
	} else if tgnInt < 0 {
		// Negative year: shift by 1 for astronomical year zero alignment
		shiftedYear := tgnInt + 1

		// Take the absolute value for padding math
		absYear := int(math.Abs(float64(shiftedYear)))

		// Return with an explicit negative sign prefix
		return fmt.Sprintf("-%04d", absYear), nil
	} else {
		// TGN data uses -1 for 1 BCE and 1 for 1 CE; a raw 0 is invalid data
		return "", errors.New("TGN data should not contain a raw 0 integer")
	}
}
