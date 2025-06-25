package game

import (
	"fmt"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type SGState = uint8

const (
	SGStatePlaying SGState = iota
	SGStateDying
	SGStateDeload
)

type SceneGame struct {
	gr         GameRoot
	lvl        GameLevel
	obsPreview LObsRow // Preview obstacles
	gs         *GameState
	cm         *CharacterMap
	msgText    *FontText
	pls        []*Player
	pRows      []PlayerRow
	pltSprs    []Sprite
	obsFalling []FallingObstacle // Falling obstacles May have many
	obsDequeue []uint16          // falling indicies to remove before next loop
	delayRmnd  int16
	dyingF     int16
	sgState    SGState
	leftKeyDb  uint8 // debounce counter
	rightKeyDb uint8 // debounce counter
	obsRowInd  uint8
}

type Player struct {
	X     int16
	Y     int16
	DazeF uint16 // Daze frames to Hold on bucket
	AnimF uint8
	Hold  bool
	DirR  bool // used for flip scale
}

type PlayerRow struct {
	Pls []*Player
	LB  int16
	RB  int16
}

type FallingObstacle struct {
	X    int16
	Y    int16
	Type uint8
}

const (
	ObsHitPad   = 3 // 10px of 16
	XVel        = 1
	FallVel     = 1
	KeyDBLimit  = 4
	MaxDaze     = TPS * 3
	CashPts     = 10000
	SandwichPts = 5000
)

func (s *SceneGame) Init(gr GameRoot, gs *GameState) {
	s.gr = gr
	s.gs = gs
	s.lvl = (*gs.Lvls)[gs.LvlInd]
	s.obsPreview = s.lvl.Obs[s.obsRowInd]
	s.delayRmnd = s.obsPreview.Delay
	s.sgState = SGStatePlaying
	s.dyingF = TPS * 4
}

func (s *SceneGame) Enter() {
	// Set up level
	s.makeWorkerPlatforms()
	s.cm = MakeCharacterMap(ImgFont)
	s.msgText = MakeFontText(s.cm, []string{})
	s.msgText.X = 16
	s.msgText.Y = 32
	s.msgText.LineSpace = 4
	s.msgText.SetText(Messages[s.obsPreview.MsgInd])
}

func (s *SceneGame) Update() error {
	if s.sgState == SGStateDying {
		if s.dyingF <= 0 {
			s.dyingF = 0
			s.sgState = SGStateDeload
		}
		s.updateObstacles()
		s.dyingF--
		return nil
	}
	if s.sgState == SGStateDeload {
		if s.gs.LvlInd == 0 {
			// Tutorial reset
			s.gs.P1Score = 0
			s.gs.P1Lives = 3
			s.gs.P1Bombs = 2
		}
		if s.gs.P1Lives == 0 {
			log.Fatal("Game over")
		}
		s.gr.SetScene(&SceneGame{})
		return nil
	}

	s.gs.P1Score++
	s.updateWorkerMovement()
	s.updateObstacles()
	s.updateCollisions()

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

	for tx, obsType := range s.obsPreview.Obs {
		if obsType == 0 {
			continue
		}
		if int(obsType) >= len(ImgsObstacles) {
			log.Fatal("Obstacle preview type not mappable to sprite sheet")
		}
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(tx*TileSize+TileSize), 0)
		screen.DrawImage(ImgsObstacles[obsType], op)
	}

	for _, pl := range s.pls {
		op := &ebiten.DrawImageOptions{}
		if !pl.DirR {
			op.GeoM.Translate(float64(TileSize*-1), 0)
			op.GeoM.Scale(-1.0, 1.0)
		}
		op.GeoM.Translate(float64(pl.X), float64(pl.Y))
		// TODO UPDATE FRAME
		screen.DrawImage(ImgWorkerStatic, op)
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
	s.msgText.Draw(screen)
}

func (s *SceneGame) Exit() {
	// nil out pointers and empty slices
}

func (s *SceneGame) makeWorkerPlatforms() {
	for _, wRow := range s.lvl.WRows {
		s.pRows = append(s.pRows, PlayerRow{
			LB:  int16(wRow.LBnd) * TileSize,
			RB:  int16(wRow.RBnd) * TileSize,
			Pls: make([]*Player, 0),
		})

		// Draw Platform
		for i := range wRow.RBnd - 1 - wRow.LBnd {
			s.pltSprs = append(s.pltSprs, Sprite{
				X:   (int16(wRow.LBnd) * TileSize) + (int16(i) * TileSize) + TileSize,
				Y:   ScreenHeight - TileSize - (int16(wRow.ZPos) * TileSize),
				Img: ImgPlatform,
			})
		}
	}
	s.pls = make([]*Player, len(s.lvl.WPos))
	for _, wPos := range s.lvl.WPos {
		pl := &Player{
			X:    int16(wPos.RowPos) * TileSize,
			Y:    ScreenHeight - (TileSize * 3) - (int16(s.lvl.WRows[wPos.RowInd].ZPos) * TileSize),
			Hold: false,
		}

		switch wPos.WID {
		case 0:
			s.pls[0] = pl
		case 1:
			s.pls[1] = pl
		case 2:
			s.pls[2] = pl
		default:
			log.Fatal("Bad WorkerPos WID")
		}

		s.pRows[wPos.RowInd].Pls = append(s.pRows[wPos.RowInd].Pls, pl)
	}
}

func (s *SceneGame) updateWorkerMovement() {
	plLen := len(s.pls)

	// Handle Hold state with input and Daze
	for i := range plLen {
		s.pls[i].Hold = InputIsHoldPressed(plLen, i)
		if s.pls[i].DazeF > 0 {
			s.pls[i].DazeF--
			s.pls[i].Hold = true
		}
	}

	if InputIsLeftPressed() {
		if s.leftKeyDb == 0 || s.leftKeyDb > KeyDBLimit {
			for _, pRow := range s.pRows {
				// start from left, slide unheld
				prevBnd := pRow.LB + TileSize
				for i := 0; i < len(pRow.Pls); i++ {
					pl := pRow.Pls[i]
					if pl.Hold {
						prevBnd = pl.X + TileSize
						continue
					}
					pl.DirR = false
					pl.X -= XVel
					if pl.X < prevBnd {
						pl.X = prevBnd
					}
					prevBnd = pl.X + TileSize
				}
			}
		}
		if s.leftKeyDb <= KeyDBLimit {
			s.leftKeyDb++
		}
	} else {
		s.leftKeyDb = 0
	}
	if InputIsRightPressed() {
		if s.rightKeyDb == 0 || s.rightKeyDb > KeyDBLimit {
			for _, pRow := range s.pRows {
				// start from right, slide unheld
				prevBnd := pRow.RB
				for i := len(pRow.Pls) - 1; i >= 0; i-- {
					pl := pRow.Pls[i]
					if pl.Hold {
						prevBnd = pl.X
						continue
					}
					pl.DirR = true
					pl.X += XVel
					if pl.X+TileSize > prevBnd {
						pl.X = prevBnd - TileSize
					}
					prevBnd = pl.X
				}
			}
		}
		if s.rightKeyDb <= KeyDBLimit {
			s.rightKeyDb++
		}
	} else {
		s.rightKeyDb = 0
	}

	if InputIsBombJustPressed() {
		if s.gs.LvlInd == 0 {
			// Tutorial skip
			s.gs.LvlInd++
			s.sgState = SGStateDeload
			return
		}

		if s.gs.P1Bombs > 0 {
			fmt.Println("Bomb activated")
			s.gs.P1Bombs--

			// set all s.obsPreview.Obs to smoke tile
			for i, currType := range s.obsPreview.Obs {
				if currType == 0 {
					continue
				}
				s.obsPreview.Obs[i] = ObstacleCloud
			}
			for i, _ := range s.obsFalling {
				s.obsFalling[i].Type = ObstacleCloud
			}
		}
	}
}

func (s *SceneGame) updateObstacles() {
	s.delayRmnd--
	if s.delayRmnd <= 0 && s.sgState == SGStatePlaying {
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
			s.gs.LvlInd++
			if int(s.gs.LvlInd) >= len(*s.gs.Lvls) {
				// TODO: Game Over
				log.Fatal("No more levels")
			}
			s.sgState = SGStateDeload
			return
		}
		s.obsPreview = s.lvl.Obs[s.obsRowInd]
		if s.obsPreview.MsgInd > 0 && int(s.obsPreview.MsgInd) < len(Messages) {
			s.msgText.SetText(Messages[s.obsPreview.MsgInd])
		} else {
			s.msgText.SetText([]string{})
		}
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
	pRects := make([]image.Rectangle, 3)
	for i, pl := range s.pls {
		ix := int(pl.X)
		iy := int(pl.Y)
		pRects[i] = image.Rect(
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

		for pI, pR := range pRects {
			if pR.Overlaps(obsR) {
				s.obsDequeue = append(s.obsDequeue, uint16(obsI))
				lostLife := s.handleCollision(obs.Type, pI)
				if lostLife {
					s.sgState = SGStateDying
					// TODO use Message index based on lives
					switch s.gs.P1Lives {
					case 2:
						s.msgText.SetText(Messages[3])
					case 1:
						s.msgText.SetText(Messages[4])
					case 0:
						s.msgText.SetText(Messages[5])
					default:
						s.msgText.SetText(Messages[2])
					}
					return
				}
			}
		}
	}
}

func (s *SceneGame) handleCollision(obsType uint8, pInd int) bool {
	switch obsType {
	case ObstacleBucket:
		if s.pls[pInd].DazeF > 0 && s.pls[pInd].DazeF < MaxDaze-32 {
			s.gs.P1Lives--
			return true
		}
		s.pls[pInd].DazeF = MaxDaze
	case ObstacleBeam:
		if s.gs.P1Lives > 0 {
			s.gs.P1Lives--
			return true
		}
	case ObstacleSandwich:
		s.gs.P1Score += SandwichPts
	case ObstacleCash:
		s.gs.P1Score += CashPts
	}
	return false
}
