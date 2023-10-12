package assets

import "runtime"

const (
	ImageNone ImageID = iota

	ImageLogo

	ImageTitle
	ImageGameOver
)

func LoadImageResources(loader *Loader, progress *float64) {
	resources := map[ImageID]string{
		ImageLogo:     "textures/logo.png",
		ImageTitle:    "textures/title.png",
		ImageGameOver: "textures/game_over.png",
	}

	isSingleThread := runtime.GOMAXPROCS(-1) == 1
	progressPerItem := 1.0 / float64(len(resources))
	for id, path := range resources {
		loader.LoadImage(id, path)
		if progress != nil {
			*progress += progressPerItem
		}
		if isSingleThread {
			runtime.Gosched()
		}
	}
}
