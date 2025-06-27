package game

import (
	"log"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
)

type SceneTitle struct {
	gr         GameRoot
	gs         *GameState
	obsFalling []FallingObstacle
	pls        []*Player
	pltSprs    []Sprite
	msgText    *FontText
}

func (s *SceneTitle) Init(gr GameRoot, gs *GameState) {
	s.gr = gr
	s.gs = gs
	ResetGameState(s.gs)
	if BgmPlayer.IsPlaying() {
		BgmPlayer.Pause()
	}

	fCount := 50
	s.obsFalling = make([]FallingObstacle, fCount)
	for i := range fCount {
		s.obsFalling[i] = FallingObstacle{
			X:    int16(rand.IntN(ScreenWidth/TileSize) * TileSize),
			Y:    int16(rand.IntN(ScreenHeight/TileSize) * TileSize),
			Type: uint8(rand.IntN(6) + 1),
		}
	}
	wT := ScreenWidth / 16
	bC := 3
	s.pltSprs = make([]Sprite, wT*bC)
	pli := 0
	for xi := range wT {
		for yi := range bC {
			s.pltSprs[pli] = Sprite{
				X: int16(xi * TileSize),
				Y: int16(ScreenHeight - (yi * TileSize * 2) - TileSize),
			}
			pli++
		}
	}
	s.pls = []*Player{
		{X: 0, Y: ScreenHeight - (TileSize * 3), DirR: true, AnimF: 0},
		{X: ScreenWidth - TileSize, Y: ScreenHeight - (TileSize * 5), DirR: false, AnimF: 3},
		{X: TileSize * 14, Y: ScreenHeight - (TileSize * 7), DirR: true, AnimF: 6},
	}

	s.msgText = MakeFontText(CM, []string{})
	s.msgText.X = 140
	s.msgText.Y = 72
	s.msgText.LineSpace = 4
	s.msgText.SetText(Messages[22])
}

func (s *SceneTitle) Update() error {
	if InputIsAJustPressed() {
		s.gr.SetScene(&SceneGame{})
		return nil
	}

	s.updateObstacles()
	s.updatePlayers()
	return nil
}

func (s *SceneTitle) Draw(screen *ebiten.Image) {
	for _, sp := range s.pltSprs {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(sp.X), float64(sp.Y))
		screen.DrawImage(ImgPlatform, op)
	}

	for _, obs := range s.obsFalling {
		if int(obs.Type) >= len(ImgsObstacles) {
			log.Fatal("Falling Obstacle type not mappable to sprite sheet")
		}
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(obs.X), float64(obs.Y))
		screen.DrawImage(ImgsObstacles[obs.Type], op)
	}

	for _, pl := range s.pls {
		op := &ebiten.DrawImageOptions{}
		if !pl.DirR {
			op.GeoM.Translate(float64(TileSize*-1), 0)
			op.GeoM.Scale(-1.0, 1.0)
		}
		op.GeoM.Translate(float64(pl.X), float64(pl.Y))

		if pl.AnimFHold == 0 {
			pl.AnimFHold = 4
			pl.AnimF++
			if pl.AnimF >= 8 {
				pl.AnimF = 0
			}
		}
		pl.AnimFHold--
		screen.DrawImage(ImgsWorker[pl.AnimF+2], op)
	}

	screen.DrawImage(ImgTitlecard, ImgTitlecardDrawOp)
	s.msgText.Draw(screen)
}

func (s *SceneTitle) Exit() {}

func (s *SceneTitle) updateObstacles() {
	for i, obs := range s.obsFalling {
		if obs.Y > ScreenHeight {
			s.obsFalling[i] = s.makeRandomObs()
			continue
		}
		s.obsFalling[i].Y += FallVel
	}
}

func (s *SceneTitle) updatePlayers() {
	for _, pl := range s.pls {
		if pl.X >= ScreenWidth-TileSize {
			pl.DirR = false
		} else if pl.X <= 0 {
			pl.DirR = true
		}
		plVel := XVel
		if !pl.DirR {
			plVel *= -1
		}
		pl.X += int16(plVel)
	}
}

func (s *SceneTitle) makeRandomObs() FallingObstacle {
	return FallingObstacle{
		X:    int16(rand.IntN(ScreenWidth/TileSize) * TileSize),
		Y:    int16(0 - rand.IntN(ScreenHeight/TileSize)*TileSize - TileSize),
		Type: uint8(rand.IntN(6) + 1),
	}
}
