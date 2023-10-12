package lib

import (
	"github.com/hajimehoshi/ebiten/v2/audio"
)

type Context struct {
	Loader *Loader

	FullScreen   bool
	WindowTitle  string
	WindowWidth  float64
	WindowHeight float64
	ScreenWidth  float64
	ScreenHeight float64
}

func NewContext() *Context {
	audioContext := audio.NewContext(44100)
	loader := NewLoader(audioContext)

	ctx := &Context{
		Loader:      loader,
		WindowTitle: "Game",
	}

	return ctx
}
