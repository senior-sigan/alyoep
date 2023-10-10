package assets

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
)

type AudioID int

type Loader struct {
	audioContext *audio.Context
	Audio        map[AudioID]*audio.Player
}

func NewLoader(audioContext *audio.Context) *Loader {
	return &Loader{
		audioContext: audioContext,
		Audio:        make(map[AudioID]*audio.Player),
	}
}

func (a *Loader) LoadAudioOgg(id AudioID, path string) *audio.Player {
	r := OpenAsset(path)
	stream, err := vorbis.DecodeWithoutResampling(r)
	if err != nil {
		panic(fmt.Sprintf("decode %q ogg: %v", path, err))
	}
	loopedStream := audio.NewInfiniteLoop(stream, stream.Length())
	player, err := a.audioContext.NewPlayer(loopedStream)
	if err != nil {
		panic(err.Error())
	}

	a.Audio[id] = player
	return player
}
