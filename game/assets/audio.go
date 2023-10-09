package assets

import (
	"fmt"
	"runtime"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
)

func RegisterMusicResource(progress *float64) {
	audioResources := map[int]string{
		MusicBass: "music/4_bass.ogg",
	}

	singleThread := runtime.GOMAXPROCS(-1) == 1
	progressPerItem := 1.0 / float64(len(audioResources))

	sampleRate := 44100
	audioContext := audio.NewContext(sampleRate)

	for id, res := range audioResources {
		LoadOgg(audioContext, res)
		if progress != nil {
			*progress += progressPerItem
		}
		if singleThread {
			runtime.Gosched()
		}
	}
}

func LoadOgg(ctx *audio.Context, path string) {
	r := OpenAsset(path)
	oggStream, err := vorbis.DecodeWithoutResampling(r)
	if err != nil {
		panic(fmt.Sprintf("decode %q ogg: %v", path, err))
	}
	loopedStream := audio.NewInfiniteLoop(oggStream, oggStream.Length())
	player, err := ctx.NewPlayer(loopedStream)
	if err != nil {
		panic(err.Error())
	}
	return player
}

const (
	AudioNone int = iota

	MusicBass
)
