package assets

// Almost all code got from https://github.com/quasilyte/ebitengine-resource/blob/master/loader.go
// And simplified in my flavour

import (
	"fmt"
	"io"
	"strings"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
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

func (l *Loader) LoadAudio(id AudioID, path string) *audio.Player {
	if strings.HasSuffix(path, ".ogg") {
		return l.LoadOGG(id, path)
	}
	return l.LoadWAV(id, path)
}

func (l *Loader) LoadOGG(id AudioID, path string) *audio.Player {
	r := OpenAsset(path)
	// Note: we're not closing the "r" above because the ogg stream
	// needs it to be kept open.
	stream, err := vorbis.DecodeWithoutResampling(r)
	if err != nil {
		panic(fmt.Sprintf("decode %q ogg: %v", path, err))
	}
	loopedStream := audio.NewInfiniteLoop(stream, stream.Length())
	player, err := l.audioContext.NewPlayer(loopedStream)
	if err != nil {
		panic(err.Error())
	}

	l.Audio[id] = player
	return player
}

func (l *Loader) LoadWAV(id AudioID, path string) *audio.Player {
	r := OpenAsset(path)
	defer func() {
		if err := r.Close(); err != nil {
			panic(fmt.Sprintf("closing %q wav reader: %v", path, err))
		}
	}()

	stream, err := wav.DecodeWithoutResampling(r)
	if err != nil {
		panic(fmt.Sprintf("decode %q wav: %v", path, err))
	}

	wavData := make([]byte, stream.Length())
	if _, err := io.ReadFull(stream, wavData); err != nil {
		panic(fmt.Sprintf("read %q wav: %v", path, err))
	}

	player := l.audioContext.NewPlayerFromBytes(wavData)

	l.Audio[id] = player
	return player
}
