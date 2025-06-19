package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type SceneToken int

type Scene interface {
	Init(GameRoot)
	Enter()
	Update() error
	Draw(screen *ebiten.Image)
	Exit()
	Root() GameRoot
}

func SceneProvider(token SceneToken) Scene {
	return nil
}
