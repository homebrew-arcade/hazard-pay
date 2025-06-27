package game

import (
	"bytes"
	"log"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
)

var sndCtx = audio.NewContext(44100)

func MakeAudioPlayerEmbedPath(path string) *audio.Player {
	sf, err := embedFS.ReadFile(path)
	if err != nil {
		log.Fatal("Unable to read audio file", err)
	}
	s, err := vorbis.DecodeF32(bytes.NewReader(sf))
	if err != nil {
		log.Fatal("Unable to decode audio file", err)
	}
	p, err := sndCtx.NewPlayerF32(s)
	if err != nil {
		log.Fatal("Unable to make player for audio file", err)
	}
	return p
}

var BgmPlayer = MakeAudioPlayerEmbedPath("assets/Jake Schofield - HardHatBop.ogg")
var SndBeamBonk = MakeAudioPlayerEmbedPath("assets/beam_bonk.ogg")
var SndBucketBonk = MakeAudioPlayerEmbedPath("assets/bucket_bonk.ogg")
var SndWalking = MakeAudioPlayerEmbedPath("assets/walking.ogg")
var SndForeman = MakeAudioPlayerEmbedPath("assets/foreman_talk.ogg")
var SndBomb = MakeAudioPlayerEmbedPath("assets/bomb.ogg")
var SndCash = MakeAudioPlayerEmbedPath("assets/cash.ogg")
var SndPowerup = MakeAudioPlayerEmbedPath("assets/sandwich.ogg")
