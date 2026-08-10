package coarse

import (
	"github.com/sfomuseum/go-edtf"
)

// ensureValidEDTF normalises EDTF strings that use the 2012
// Open/Unspecified markers to the 2015 values expected by the
// rest of the code base.
func ensureValidEDTF(input string) string {

	edtf_updates := map[string]string{
		edtf.UNSPECIFIED_2012: edtf.UNSPECIFIED,
		edtf.OPEN_2012:        edtf.OPEN,
	}

	new, exists := edtf_updates[input]

	if exists {
		return new
	}

	return input
}
