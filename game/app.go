package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/senior-sigan/alyoep/game/assets"
)

func RunApp() error {
	audioContext := audio.NewContext(44100)
	loader := assets.NewLoader(audioContext)
	progress := 0.0
	assets.LoadAudioResources(loader, &progress)
	progress = 0.0
	assets.LoadImageResources(loader, &progress)

	p := loader.Audio[assets.AudioBass]
	p.Play()

	game := NewGame(loader)

	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("ALYOP")
	return ebiten.RunGame(game)
}
