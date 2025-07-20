package ui

import (
	"embed"
)

//go:embed "templates"
var Templates embed.FS

//go:embed "assets/node_modules"
var NodeModules embed.FS

//go:embed "assets/css"
var CSS embed.FS
