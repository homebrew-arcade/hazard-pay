package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	X   int16
	Y   int16
	Img *ebiten.Image
}
