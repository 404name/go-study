package public

import (
	"embed"
)

//go:embed dist
var StaticContent embed.FS
