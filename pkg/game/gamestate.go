package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type GameState struct {
	Lvls     *[]GameLevel
	IdleF    uint // Idle frame ticker for attract mode
	P1Score  uint
	LvlInd   uint8
	P1Lives  uint8
	Credits  uint8
	IsIdle   bool
	DebugImg *ebiten.Image
}

func MakeGameState() *GameState {
	gs := GameState{
		LvlInd:  0,
		Lvls:    MakeLevels(),
		P1Score: 0,
		P1Lives: 0,
		Credits: 0,
		IdleF:   0,
		IsIdle:  false,
	}
	if Debug {
		gs.DebugImg = ebiten.NewImage(ScreenWidth, ScreenHeight)
		gs.DebugImg.Fill(color.NRGBA{R: 255, G: 0, B: 0, A: 64})
	}
	return &gs
}
