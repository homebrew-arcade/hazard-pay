package game

import (
	"bytes"
	"log"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
)

// bgm

func MakeAudioPlayerEmbedPath(path string, sampleRate int) *audio.Player {
	sf, err := embedFS.ReadFile(path)
	if err != nil {
		log.Fatal("Unable to read bgm wav")
	}
	s, err := vorbis.DecodeF32(bytes.NewReader(sf))
	if err != nil {
		log.Fatal("Unable to decode bgm wav")
	}
	var ctx = audio.NewContext(sampleRate)
	p, err := ctx.NewPlayerF32(s)
	if err != nil {
		log.Fatal("Unable to make player for bgm wav")
	}
	return p
}

var BgmPlayer = MakeAudioPlayerEmbedPath("assets/Jake Schofield - HardHatBop.ogg", 44100)
