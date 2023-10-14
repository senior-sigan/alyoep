package systems

import (
	"log"
	"math"

	"github.com/senior-sigan/alyoep/game/assets"
	"github.com/senior-sigan/alyoep/lib"
)

const BPM = 94
const beatsPerSlotter = 2
const slotsCount = 8

type BgMusicSystem struct {
	ctx      *lib.Context
	lastSlot int
	slots    []lib.AudioID
}

func NewBgMusicSystem(ctx *lib.Context) *BgMusicSystem {
	s := &BgMusicSystem{
		ctx:      ctx,
		lastSlot: -1,
		slots: []lib.AudioID{
			assets.AudioSh_1,
			assets.AudioSh_2,
			assets.AudioSh_3,
			assets.AudioSh_4,
			assets.AudioSh_5,
			assets.AudioSh_2,
			assets.AudioSh_3,
			assets.AudioSh_4,
		},
	}

	bass := ctx.Loader.Audio[assets.AudioBass]
	bass.Play()

	return s
}

func (s *BgMusicSystem) slotterProgress() float64 {
	bass := s.ctx.Loader.Audio[assets.AudioBass]
	_, progress := math.Modf((bass.Position().Seconds() * BPM / 60) / beatsPerSlotter)
	return progress
}

func (s *BgMusicSystem) Update() {
	progress := s.slotterProgress()
	slot := int(math.Floor(progress * slotsCount))

	if slot != s.lastSlot {
		log.Printf("%d", slot)
		s.lastSlot = slot
		p := s.ctx.Loader.Audio[s.slots[slot]]
		p.Rewind()
		p.Play()
	}
}
