package assets

import (
	"embed"
)

//go:embed audio
var AudioAssets embed.FS

//go:embed images
var ImageAssets embed.FS
