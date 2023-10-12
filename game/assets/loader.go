package assets

// Almost all code got from https://github.com/quasilyte/ebitengine-resource/blob/master/loader.go
// And simplified in my flavour

import (
	"image"
	"io"
	"log"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

type AudioID int
type ImageID int

type Loader struct {
	audioContext *audio.Context
	Audio        map[AudioID]*audio.Player
	Image        map[ImageID]*ebiten.Image
}

func NewLoader(audioContext *audio.Context) *Loader {
	return &Loader{
		audioContext: audioContext,
		Audio:        make(map[AudioID]*audio.Player),
		Image:        make(map[ImageID]*ebiten.Image),
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
		log.Panicf("decode %q ogg: %v", path, err)
	}
	loopedStream := audio.NewInfiniteLoop(stream, stream.Length())
	player, err := l.audioContext.NewPlayer(loopedStream)
	if err != nil {
		log.Panicf("player create %q ogg: %v", path, err)
	}

	l.Audio[id] = player
	log.Printf("+ loaded ogg audio %q", path)
	return player
}

func (l *Loader) LoadWAV(id AudioID, path string) *audio.Player {
	r := OpenAsset(path)
	defer func() {
		if err := r.Close(); err != nil {
			log.Panicf("closing %q wav reader: %v", path, err)
		}
	}()

	stream, err := wav.DecodeWithoutResampling(r)
	if err != nil {
		log.Panicf("decode %q wav: %v", path, err)
	}

	wavData := make([]byte, stream.Length())
	if _, err := io.ReadFull(stream, wavData); err != nil {
		log.Panicf("read %q wav: %v", path, err)
	}

	player := l.audioContext.NewPlayerFromBytes(wavData)

	l.Audio[id] = player
	log.Printf("+ loaded wav audio %q", path)
	return player
}

func (l *Loader) LoadImage(id ImageID, path string) *ebiten.Image {
	r := OpenAsset(path)
	defer func() {
		if err := r.Close(); err != nil {
			log.Panicf("closing %q image reader: %v", path, err)
		}
	}()
	rawImage, _, err := image.Decode(r)
	if err != nil {
		log.Panicf("decode %q image: %v", path, err)
	}
	image := ebiten.NewImageFromImage(rawImage)

	l.Image[id] = image
	log.Printf("+ loaded image %q", path)
	return image
}
