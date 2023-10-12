package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/senior-sigan/alyoep/game/assets"
)

type Game struct {
	Loader *assets.Loader
}

func NewGame(loader *assets.Loader) *Game {
	return &Game{Loader: loader}
}

func (g Game) Update() error {
	return nil
}

func (g Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(g.Loader.Image[assets.ImageLogo], op)

	msg := fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f", ebiten.ActualTPS(), ebiten.ActualFPS())
	ebitenutil.DebugPrint(screen, msg)
}

func (g Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1280, 720
}
