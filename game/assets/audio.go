package assets

import (
	"github.com/senior-sigan/alyoep/lib"
)

const (
	AudioNone lib.AudioID = iota

	AudioLoTrash
	AudioHiTrash
	AudioDream
	AudioBass

	AudioHeal
	AudioPowerUp

	AudioNova_1
	AudioNova_2
	AudioNova_3
	AudioNova_4
	AudioNova_5

	AudioHitTower_1
	AudioHitTower_2
	AudioHitTower_3
	AudioHitTower_4

	AudioHitPlayer_1
	AudioHitPlayer_2
	AudioHitPlayer_3
	AudioHitPlayer_4
)

func LoadAudioResources(loader *lib.Loader, progress *float64) {
	audioResources := map[lib.AudioID]string{
		AudioLoTrash:     "music/1_lo-trash.ogg",
		AudioHiTrash:     "music/2_hi-trash.ogg",
		AudioDream:       "music/3_dream.ogg",
		AudioBass:        "music/4_bass.ogg",
		AudioHeal:        "sounds/heal.wav",
		AudioPowerUp:     "sounds/powerup.wav",
		AudioNova_1:      "sounds/nova/1_SID.wav",
		AudioNova_2:      "sounds/nova/2_SID.wav",
		AudioNova_3:      "sounds/nova/3_SID.wav",
		AudioNova_4:      "sounds/nova/4_SID.wav",
		AudioNova_5:      "sounds/nova/5_SID.wav",
		AudioHitTower_1:  "sounds/hits/explosion_1.wav",
		AudioHitTower_2:  "sounds/hits/explosion_2.wav",
		AudioHitTower_3:  "sounds/hits/explosion_3.wav",
		AudioHitTower_4:  "sounds/hits/explosion_4.wav",
		AudioHitPlayer_1: "sounds/hits/hit_1.wav",
		AudioHitPlayer_2: "sounds/hits/hit_2.wav",
		AudioHitPlayer_3: "sounds/hits/hit_3.wav",
		AudioHitPlayer_4: "sounds/hits/hit_4.wav",
	}

	loader.LoadAllAudio(audioResources, progress)
}
