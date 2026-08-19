package coarse

import (
	"embed"
)

//go:embed *.html css/*.css javascript/*.js wasm/*.wasm models/*/*.gguf
var FS embed.FS
