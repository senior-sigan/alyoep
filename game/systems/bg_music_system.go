package systems

import (
	"github.com/senior-sigan/alyoep/game/assets"
	"github.com/senior-sigan/alyoep/lib"
)

type BgMusicSystem struct {
	ctx *lib.Context
}

func NewBgMusicSystem(ctx *lib.Context) *BgMusicSystem {
	s := &BgMusicSystem{ctx: ctx}
	s.syncMusic()

	// FIXME: remove, it's just an example of playing music
	p := ctx.Loader.Audio[assets.AudioBass]
	p.Play()

	return s
}

func (s *BgMusicSystem) syncMusic() {

}

func (s *BgMusicSystem) Update() {

}
