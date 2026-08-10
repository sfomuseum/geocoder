package geocoder

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/paulmach/orb"
)

// StringBoundsToOrbBounds converts string bounds in the form of 'minx,miny,maxx,maxy' in to an `orb.Bound` instance.
func StringBoundsToOrbBounds(str_bounds string) (*orb.Bound, error) {

	str_parts := strings.Split(str_bounds, ",")

	if len(str_parts) != 4 {
		return nil, fmt.Errorf("Invalid bounding box. Expected (4) parts but got %d.", len(str_parts))
	}

	parts := make([]float64, 4)

	for i, str_c := range str_parts {

		str_c = strings.TrimSpace(str_c)
		c, err := strconv.ParseFloat(str_c, 64)

		if err != nil {
			return nil, fmt.Errorf("Invalid coordinate '%s'", str_c)
		}

		parts[i] = c
	}

	bounds := orb.Bound{
		Min: orb.Point([2]float64{parts[0], parts[1]}),
		Max: orb.Point([2]float64{parts[2], parts[3]}),
	}

	if bounds.IsEmpty() {
		return nil, fmt.Errorf("Bounds are empty")
	}

	return &bounds, nil
}
