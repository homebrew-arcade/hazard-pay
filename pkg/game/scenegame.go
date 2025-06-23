package game

import (
	"fmt"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type SceneGame struct {
	gr      GameRoot
	lvl     GameLevel
	gs      *GameState
	wA      *Movable     // Player A
	wB      *Movable     // Worker B
	wC      *Movable     // Worker C
	mRows   [][]*Movable // Platform/worker resolvers
	pltSprs []Sprite

	wSprs      []Sprite
	leftKeyDb  uint8 // debounce counter
	rightKeyDb uint8 // debounce counter
	obsRowInd  uint8
	delayRmnd  int16
	obsPreview ObsRow            // Preview obstacles
	obsFalling []FallingObstacle // Falling obstacles May have many
	obsDequeue []uint16          // falling indicies to remove before next loop
	cm         *CharacterMap
	textTest   *FontText
}

type Movable struct {
	X        int16
	Y        int16
	Hold     bool
	IsWorker bool
}

type FallingObstacle struct {
	Type uint8
	X    int16
	Y    int16
}

const ObsHitPad = 3 // 10px of 16
const XVel = 1
const FallVel = 1
const KeyDBLimit = 4

func (s *SceneGame) Init(gr GameRoot, gs *GameState) {
	s.gr = gr
	s.gs = gs
	s.lvl = (*gs.Lvls)[gs.LvlInd]
	s.wSprs = make([]Sprite, 3)
	s.obsPreview = s.lvl.Obs[s.obsRowInd]
	s.delayRmnd = s.obsPreview.Delay
}

func (s *SceneGame) Enter() {
	// Set up level
	s.makeWorkerPlatforms()
	s.cm = MakeCharacterMap(ImgFont)
	s.textTest = MakeFontText(s.cm, []string{
		"Hello World",
		"This is Sample Text",
	})
	s.textTest.X = 8
	s.textTest.Y = 8
	s.textTest.LineSpace = 4
}

func (s *SceneGame) Update() error {
	s.gs.P1Score++
	s.updateWorkerMovement()
	s.updateObstacles()
	s.updateCollisions()

	if s.gs.P1Lives == 0 {
		log.Fatal("Game over")
	}

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
	for tx, obsType := range s.obsPreview.Obs {
		if obsType == 0 {
			continue
		}
		if int(obsType) >= len(ImgsObstacles) {
			log.Fatal("Obstacle Preview type not mappable to sprite sheet")
		}
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(tx*TileSize+TileSize), 0)
		screen.DrawImage(ImgsObstacles[obsType], op)
	}

	for _, obs := range s.obsFalling {
		if int(obs.Type) >= len(ImgsObstacles) {
			log.Fatal("Falling Obstacle type not mappable to sprite sheet")
		}
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(obs.X), float64(obs.Y))
		screen.DrawImage(ImgsObstacles[obs.Type], op)
	}

	if Debug {
		ebitenutil.DebugPrint(
			screen,
			fmt.Sprintf("Lives: %v, Score: %v", s.gs.P1Lives, s.gs.P1Score),
		)
	}

	//screen.DrawImage(ImgFont, &ebiten.DrawImageOptions{})
	s.textTest.Draw(screen)
}

func (s *SceneGame) Exit() {
	// nil out pointers and empty slices
}

func (s *SceneGame) makeWorkerPlatforms() {
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

func (s *SceneGame) updateWorkerMovement() {
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
}

func (s *SceneGame) updateObstacles() {
	s.delayRmnd--
	if s.delayRmnd <= 0 {
		// copy into falling and swap
		for tx, obsType := range s.obsPreview.Obs {
			if obsType == 0 {
				continue
			}
			obs := FallingObstacle{
				X:    int16(tx*TileSize + TileSize),
				Y:    0,
				Type: obsType,
			}
			s.obsFalling = append(s.obsFalling, obs)
		}
		s.obsRowInd++
		if int(s.obsRowInd) >= len(s.lvl.Obs) {
			log.Fatal("end of level rows")
		}
		s.obsPreview = s.lvl.Obs[s.obsRowInd]
		s.delayRmnd = s.obsPreview.Delay
	}

	for i, obs := range s.obsFalling {
		if obs.Type == ObstacleNil || obs.Y > ScreenHeight {
			s.obsDequeue = append(s.obsDequeue, uint16(i))
			continue
		}
		s.obsFalling[i].Y += FallVel
	}
	s.dequeueObstacles()
}

func (s *SceneGame) dequeueObstacles() {
	if len(s.obsDequeue) == 0 {
		return
	}
	// dequeue from falling
	// must iterate in reverse
	for i := len(s.obsDequeue) - 1; i >= 0; i-- {
		di := int(s.obsDequeue[i])
		obsFLen := len(s.obsFalling)
		if di >= obsFLen {
			log.Fatal("Obstable Dequeue index out of range ", di, obsFLen)
		}
		if obsFLen == 1 {
			// straight trim, no swap
			s.obsFalling = s.obsFalling[:0]
			break
		}
		// Swapback routine, copy from end and reslice
		s.obsFalling[di] = s.obsFalling[obsFLen-1]
		s.obsFalling = s.obsFalling[:obsFLen-1]
	}
	// reset dequeue slice without backing mutation
	s.obsDequeue = s.obsDequeue[:0]
}

func (s *SceneGame) updateCollisions() {
	wRects := make([]image.Rectangle, 3)
	for i, mov := range []Movable{*s.wA, *s.wB, *s.wC} {
		ix := int(mov.X)
		iy := int(mov.Y)
		wRects[i] = image.Rect(
			ix+ObsHitPad,
			iy+ObsHitPad,
			ix+TileSize-ObsHitPad,
			iy+TileSize-ObsHitPad,
		)
	}
	for obsI, obs := range s.obsFalling {
		ix := int(obs.X)
		iy := int(obs.Y)
		obsR := image.Rect(
			ix+ObsHitPad,
			iy+ObsHitPad,
			ix+TileSize-ObsHitPad,
			iy+TileSize-ObsHitPad,
		)

		for wI, wR := range wRects {
			if wR.Overlaps(obsR) {
				s.handleCollision(obs.Type, wI)
				s.obsDequeue = append(s.obsDequeue, uint16(obsI))
			}
		}
	}
}

func (s *SceneGame) handleCollision(obsType uint8, wInd int) {
	fmt.Println("COLLISION", wInd, obsType)
	switch obsType {
	case ObstacleBucket:
		//TODO: hold stun
		if s.gs.P1Score <= 1000 {
			s.gs.P1Score = 0
		} else {
			s.gs.P1Score -= 1000
		}
	case ObstacleBeam:
		if s.gs.P1Lives > 0 {
			s.gs.P1Lives--
		}
	case ObstacleSandwich:
		s.gs.P1Score += 500
	case ObstacleCash:
		s.gs.P1Score += 1000
	}

}
