package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	InputL    = ebiten.KeyA
	InputR    = ebiten.KeyD
	InputA    = ebiten.KeyJ
	InputB    = ebiten.KeyK
	InputC    = ebiten.KeyL
	InputBomb = ebiten.KeySpace
	InputCoin = ebiten.KeyControlLeft
	InputP1S  = ebiten.KeyEnter
)

func InputIsLeftPressed() bool {
	return ebiten.IsKeyPressed(InputL)
}

func InputIsLeftJustPressed() bool {
	return inpututil.IsKeyJustPressed(InputL)
}

func InputIsRightPressed() bool {
	return ebiten.IsKeyPressed(InputR)
}

func InputIsRightJustPressed() bool {
	return inpututil.IsKeyJustPressed(InputR)
}

func InputIsBombJustPressed() bool {
	return inpututil.IsKeyJustPressed(InputBomb)
}

func InputIsCoinJustPressed() bool {
	return inpututil.IsKeyJustPressed(InputCoin)
}

func InputIsPlStartJustPressed() bool {
	return inpututil.IsKeyJustPressed(InputP1S)
}

func InputIsHoldPressed(plLen int, wid int) bool {
	if plLen == 1 {
		return ebiten.IsKeyPressed(InputA) || ebiten.IsKeyPressed(InputB) || ebiten.IsKeyPressed(InputC)
	}
	if plLen == 2 {
		if wid == 0 && ebiten.IsKeyPressed(InputA) {
			return true
		}
		if wid == 1 && (ebiten.IsKeyPressed(InputB) || ebiten.IsKeyPressed(InputC)) {
			return true
		}
	} else {
		if wid == 0 && ebiten.IsKeyPressed(InputA) {
			return true
		}
		if wid == 1 && ebiten.IsKeyPressed(InputB) {
			return true
		}
		if wid == 2 && ebiten.IsKeyPressed(InputC) {
			return true
		}
	}
	return false
}

func InputIsAJustPressed() bool {
	return inpututil.IsKeyJustPressed(InputA)
}
