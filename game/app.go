package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func RunApp() error {
	game := NewGame()

	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("ALYOP")
	return ebiten.RunGame(game)
}
