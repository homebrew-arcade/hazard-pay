package game

import (
	randv2 "math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
)

var rand = func() *randv2.Rand {
	s := randv2.NewPCG(0, 0)
	return randv2.New(s)
}()

func MakeRand(a uint64, b uint64) *randv2.Rand {
	s := randv2.NewPCG(a, b)
	return randv2.New(s)
}

const (
	Debug = true

	ScreenWidth  = 480 // 30 tiles
	ScreenHeight = 256 // 16 tiles

	WindowTitle  = "Hazard Pay"
	WindowWidth  = 1920
	WindowHeight = 1080

	TPS          = 100
	VSyncEnabled = true

	TileSize     = 16
	HalfTileSize = 8

	Esc     = ebiten.KeyEscape
	P1Coin  = ebiten.Key5
	P1Start = ebiten.Key1
	P1Right = ebiten.KeyArrowRight
	P1Left  = ebiten.KeyArrowLeft
	P1Up    = ebiten.KeyArrowUp
	P1Down  = ebiten.KeyArrowDown
	P1B1    = ebiten.KeyControlLeft
	P1B2    = ebiten.KeyAltLeft
	P1B3    = ebiten.KeySpace
	P1B4    = ebiten.KeyShiftLeft
	P1B5    = ebiten.KeyZ
	P1B6    = ebiten.KeyX
	P1B7    = ebiten.KeyC
	P1B8    = ebiten.KeyV

	P2Coin  = ebiten.Key6
	P2Start = ebiten.Key2
	P2Right = ebiten.KeyG
	P2Left  = ebiten.KeyD
	P2Up    = ebiten.KeyR
	P2Down  = ebiten.KeyF
	P2B1    = ebiten.KeyA
	P2B2    = ebiten.KeyS
	P2B3    = ebiten.KeyQ
	P2B4    = ebiten.KeyW
	P2B5    = ebiten.KeyI
	P2B6    = ebiten.KeyK
	P2B7    = ebiten.KeyJ
	P2B8    = ebiten.KeyL
)

/*
INPUT 		NORMAL CODES 			CODES WITH SHIFT
COIN 1 			5
START 1 		1
1 RIGHT 		R arrow				Tab
1 LEFT 			L arrow				Enter
1 UP 			U arrow				~ (tilde)
1 DOWN 			D arrow				P (pause)
1 SW 1 			L-ctrl				5 (Coin A)
1 SW 2 			L-alt
1 SW 3 			space
1 SW 4 			L-shift
1 SW 5 			Z
1 SW 6 			X
1 SW 7 			C
1 SW 8 			V
1 A 			P
1 B 			ENTER

PLAYER 2		NORMAL CODES 		CODES WITH SHIFT
COIN 2 			6
START 2 		2				ESC
2 RIGHT 		G
2 LEFT 			D
2 UP 			R
2 DOWN 			F
2 SW 1 			A
2 SW 2 			S
2 SW 3 			Q
2 SW 4 			W
2 SW 5 			I
2 SW 6 			K
2 SW 7 			J
2 SW 8 			L
2 A 			TAB
2 B 			ESC
*/
