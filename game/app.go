package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/senior-sigan/alyoep/game/assets"
)

func RunApp() error {
	audioContext := audio.NewContext(44100)
	loader := assets.NewLoader(audioContext)
	p := loader.LoadAudioOgg(assets.AudioBass, "music/4_bass.ogg")
	p.Play()

	game := NewGame()

	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("ALYOP")
	return ebiten.RunGame(game)
}
