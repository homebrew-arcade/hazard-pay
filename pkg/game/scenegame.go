package game

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type SceneGame struct {
	gr         GameRoot
	lvl        GameLevel
	gs         *GameState
	obsRow     uint
	wA         *Movable
	wB         *Movable
	wC         *Movable
	mRows      [][]*Movable
	pltSprs    []Sprite
	wSprs      []Sprite
	leftKeyDb  uint8 // debounce counter
	rightKeyDb uint8 // debounce counter
}

type Movable struct {
	X        int16
	Y        int16
	Hold     bool
	IsWorker bool
}

const WHitPad = 5 // 10px of 16
const XVel = 1
const KeyDBLimit = 4

func (s *SceneGame) Init(gr GameRoot, gs *GameState) {
	s.gr = gr
	s.gs = gs
	s.lvl = (*gs.Lvls)[gs.LvlInd]
	s.obsRow = 0
	s.wSprs = make([]Sprite, 3)
}

func (s *SceneGame) Enter() {
	// Set up level
	s.mRows = make([][]*Movable, len(s.lvl.WRows))
	for i, wRow := range s.lvl.WRows {
		// Allocate empty bumper
		mRow := make([]*Movable, 1, 5)
		lbM := Movable{
			X:        int16(wRow.LBnd) * TileSize,
			Y:        ScreenHeight - TileSize - (int16(wRow.ZPos) * TileSize),
			Hold:     true,
			IsWorker: false,
		}
		mRow[0] = &lbM
		s.mRows[i] = mRow

		// Draw Platform
		for i := range wRow.RBnd - 1 - wRow.LBnd {
			s.pltSprs = append(s.pltSprs, Sprite{
				X:   (int16(wRow.LBnd) * TileSize) + (int16(i) * TileSize) + TileSize,
				Y:   ScreenHeight - TileSize - (int16(wRow.ZPos) * TileSize),
				Img: ImgPlatform,
			})
			//fmt.Printf("%+v, %v \n", s.pltSprs[len(s.pltSprs)-1], i)
		}
	}
	for _, wPos := range s.lvl.WPos {
		wMov := Movable{
			X:        int16(wPos.RowPos) * TileSize,
			Y:        ScreenHeight - (TileSize * 3) - (int16(s.lvl.WRows[wPos.RowInd].ZPos) * TileSize),
			Hold:     false,
			IsWorker: true,
		}
		spr := Sprite{
			X:   wMov.X,
			Y:   wMov.Y,
			Img: ImgWorkerStatic,
		}
		switch wPos.WID {
		case 0:
			s.wA = &wMov
			s.wSprs[0] = spr
		case 1:
			s.wB = &wMov
			s.wSprs[1] = spr
		case 2:
			s.wC = &wMov
			s.wSprs[2] = spr
		default:
			log.Fatal("Bad WorkerPos WID")
		}
		s.mRows[wPos.RowInd] = append(s.mRows[wPos.RowInd], &wMov)
	}
	for i, wRow := range s.lvl.WRows {
		// Allocate empty bumper
		rbM := Movable{
			X:        int16(wRow.RBnd) * TileSize,
			Y:        ScreenHeight - TileSize - (int16(wRow.ZPos) * TileSize),
			Hold:     true,
			IsWorker: false,
		}
		s.mRows[i] = append(s.mRows[i], &rbM)
	}
}

func (s *SceneGame) Update() error {
	s.wA.Hold = false
	s.wB.Hold = false
	s.wC.Hold = false
	if ebiten.IsKeyPressed(ebiten.KeyJ) {
		s.wA.Hold = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyK) {
		s.wB.Hold = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyL) {
		s.wC.Hold = true
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		if s.leftKeyDb == 0 || s.leftKeyDb > KeyDBLimit {
			for _, mRow := range s.mRows {
				// start from left, slide unheld
				for i := 0; i < len(mRow); i++ {
					m := mRow[i]
					if m.Hold {
						continue
					}
					mp := mRow[i-1]
					m.X -= XVel
					if m.X < mp.X+TileSize {
						m.X = mp.X + TileSize
					}
				}
			}
		}
		if s.leftKeyDb <= KeyDBLimit {
			s.leftKeyDb++
		}
	} else {
		s.leftKeyDb = 0
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		if s.rightKeyDb == 0 || s.rightKeyDb > KeyDBLimit {
			for _, mRow := range s.mRows {
				// start from right, slide unheld
				for i := len(mRow) - 1; i >= 0; i-- {
					m := mRow[i]
					if m.Hold {
						continue
					}
					mp := mRow[i+1]
					m.X += XVel
					if m.X+TileSize > mp.X {
						m.X = mp.X - TileSize
					}
				}
			}
		}
		if s.rightKeyDb <= KeyDBLimit {
			s.rightKeyDb++
		}
	} else {
		s.rightKeyDb = 0
	}

	s.wSprs[0].X = s.wA.X
	s.wSprs[1].X = s.wB.X
	s.wSprs[2].X = s.wC.X

	return nil
}

func (s *SceneGame) Draw(screen *ebiten.Image) {
	screen.DrawImage(ImgGameBg, ImgGameBgDrawOp)
	screen.DrawImage(ImgForeman, ImgFormanDrawOp)

	for _, sp := range s.pltSprs {
		if sp.Img == nil {
			continue
		}
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(sp.X), float64(sp.Y))
		screen.DrawImage(sp.Img, op)
	}
	for _, sp := range s.wSprs {
		if sp.Img == nil {
			continue
		}
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(sp.X), float64(sp.Y))
		screen.DrawImage(sp.Img, op)
	}
}

func (s *SceneGame) Exit() {}
