package systems

import (
	"log"
	"math"

	"github.com/senior-sigan/alyoep/game/assets"
	"github.com/senior-sigan/alyoep/lib"
)

const BPM = 94.0

type BgMusicSystem struct {
	ctx             *lib.Context
	lastSlot        int
	slots           []lib.AudioID
	beatsPerSlotter int
	maxSpeed        int
	once            *lib.Once
}

func NewBgMusicSystem(ctx *lib.Context) *BgMusicSystem {
	slots := []lib.AudioID{
		assets.AudioSh_1,
		assets.AudioSh_2,
		assets.AudioSh_3,
		assets.AudioSh_4,
		assets.AudioSh_5,
		assets.AudioSh_2,
		assets.AudioSh_3,
		assets.AudioSh_4,
	}
	s := &BgMusicSystem{
		ctx:             ctx,
		lastSlot:        -1,
		slots:           slots,
		beatsPerSlotter: 2,
		maxSpeed:        int(math.Round(1 + math.Log2(float64(len(slots))))),
		once:            lib.NewOnce(),
	}

	bass := ctx.Loader.Audio[assets.AudioBass]
	bass.Play()

	return s
}

func (s *BgMusicSystem) slotterProgress() float64 {
	bass := s.ctx.Loader.Audio[assets.AudioBass]
	_, progress := math.Modf((bass.Position().Seconds() * BPM / 60.0) / float64(s.beatsPerSlotter))
	return progress
}

func (s *BgMusicSystem) tryShoot(playerId int, speed int) bool {
	// FIXME: тут баг с подвисанием такта музыки  если взять двух игроков со скоростьяю 1 и 3
	slotsCount := len(s.slots)
	slot := int(math.Floor(s.slotterProgress() * float64(slotsCount)))

	sfx := s.ctx.Loader.Audio[s.slots[slot]]

	if slot != s.lastSlot {
		s.lastSlot = slot
		s.once.Reset()
	}

	playerSlot := (slot + playerId) % slotsCount
	if playerSlot%(1<<(s.maxSpeed-speed)) != 0 {
		return false
	}

	if s.once.Invoke() {
		log.Printf("BUM pid=%d slot=%d", playerId, slot)
		sfx.Rewind()
		sfx.Play()
	} else {
		return false
	}

	return true
}

func (s *BgMusicSystem) Update() {
	s.tryShoot(0, 3)
	s.tryShoot(1, 1)
}
