package ui

import (
	"embed"
)

//go:embed "templates"
var Templates embed.FS

//go:embed "assets/js/node_modules"
var NodeModules embed.FS
