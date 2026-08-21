package demo

import (
	"embed"
)

//go:embed *.html css/*.css javascript/*.js wasm/* models/*/*
var FS embed.FS
