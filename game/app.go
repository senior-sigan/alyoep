package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/senior-sigan/alyoep/game/assets"
	"github.com/senior-sigan/alyoep/lib"
)

func RunApp() error {
	ctx := lib.NewContext()
	ctx.Loader.OpenAsset = assets.OpenAsset

	progress := 0.0
	assets.LoadAudioResources(ctx.Loader, &progress)
	progress = 0.0
	assets.LoadImageResources(ctx.Loader, &progress)

	game := NewGame(ctx)

	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("ALYOP")
	return ebiten.RunGame(game)
}
