package fs

import (
	"embed"
)

//go:embed *.db
var FS embed.FS

