package coarse

import (
	"embed"
)

//go:embed *.html css/*.css javascript/*.js
var FS embed.FS
