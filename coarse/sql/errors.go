package sql

import (
	"errors"
)

// NotSupportedError is returned by the table implementations for
// any operation that the SQL backend does not support – for
// example, the IndexRecord method is not needed because
// records are added in bulk.
var NotSupportedError = errors.New("Not supported")
