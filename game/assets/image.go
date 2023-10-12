package assets

import (
	"github.com/senior-sigan/alyoep/lib"
)

const (
	ImageNone lib.ImageID = iota

	ImageLogo

	ImageTitle
	ImageGameOver
)

func LoadImageResources(loader *lib.Loader, progress *float64) {
	resources := map[lib.ImageID]string{
		ImageLogo:     "textures/logo.png",
		ImageTitle:    "textures/title.png",
		ImageGameOver: "textures/game_over.png",
	}

	loader.LoadAllImages(resources, progress)
}
