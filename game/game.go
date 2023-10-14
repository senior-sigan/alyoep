package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/senior-sigan/alyoep/game/assets"
	"github.com/senior-sigan/alyoep/game/systems"
	"github.com/senior-sigan/alyoep/lib"
)

type Game struct {
	Context *lib.Context

	bgMusicSystem *systems.BgMusicSystem
}

func NewGame(ctx *lib.Context) *Game {
	game := &Game{Context: ctx}

	game.bgMusicSystem = systems.NewBgMusicSystem(ctx)

	return game
}

func (g Game) Update() error {
	g.bgMusicSystem.Update()
	return nil
}

func (g Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(g.Context.Loader.Image[assets.ImageLogo], op)

	msg := fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f", ebiten.ActualTPS(), ebiten.ActualFPS())
	ebitenutil.DebugPrint(screen, msg)
}

func (g Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1280, 720
}
