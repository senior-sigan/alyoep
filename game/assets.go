package game

import (
	"bytes"

	_ "embed"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
)

//go:embed assets/music/4_bass.ogg
var Bass_ogg []byte

const sampleRate = 44100

func NewAudioSystem() error {
	s, err := vorbis.DecodeWithoutResampling(bytes.NewReader(Bass_ogg))
	if err != nil {
		return nil
	}

	audioContext := audio.NewContext(sampleRate)
	p, err := audioContext.NewPlayer(s)
	if err != nil {
		return err
	}
	p.Play()
	return nil
}
