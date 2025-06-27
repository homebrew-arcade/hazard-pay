package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Scene interface {
	Init(GameRoot, *GameState)
	Update() error
	Draw(*ebiten.Image)
	Exit()
}
