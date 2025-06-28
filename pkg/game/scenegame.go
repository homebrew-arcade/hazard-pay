package game

import (
	"fmt"
	"image"
	"log"
	randv2 "math/rand/v2"

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
	msgText    *FontText
	scoreText  *FontTextBasic
	obsRand    *randv2.Rand
	pls        []*Player
	pRows      []PlayerRow
	pltSprs    []Sprite
	obsFalling []FallingObstacle // Falling obstacles May have many
	obsDequeue []uint16          // falling indicies to remove before next loop
	delayRmnd  int16
	dyingF     int16
	blinkF     int16
	mouthF     int16
	sgState    SGState
	leftKeyDb  uint8 // debounce counter
	rightKeyDb uint8 // debounce counter
	obsRowInd  uint8
	mouthInd   uint8
	debugPause bool
}

type Player struct {
	X         int16
	Y         int16
	DazeF     uint16 // Daze frames to Hold on bucket
	AnimF     uint8  // Sheet index of current walk frame
	AnimFHold uint8  // remaining TPS until next frame
	IsMoving  bool
	Hold      bool
	DirR      bool // used for flip scale
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
	MaxDaze     = TPS
	CashPts     = 10000
	SandwichPts = 5000
	CloudPts    = 1000
)

func (s *SceneGame) Init(gr GameRoot, gs *GameState) {
	if Debug {
		gs.LvlInd = 5
		s.obsRowInd = 13
	}
	s.gr = gr
	s.gs = gs
	s.lvl = (*Levels)[gs.LvlInd]
	s.obsRand = MakeRand(s.lvl.SeedA, s.lvl.SeedB)
	s.obsPreview = s.lvl.Obs[s.obsRowInd]
	s.delayRmnd = s.obsPreview.Delay
	s.sgState = SGStatePlaying
	s.dyingF = TPS * 4

	// Set up level
	s.makeWorkerPlatforms()
	s.msgText = MakeFontText(CM, []string{})
	s.msgText.X = 16
	s.msgText.Y = 32
	s.msgText.LineSpace = 4
	s.msgText.CharDelay = 1
	s.msgText.SetText(Messages[s.obsPreview.MsgInd])

	s.scoreText = MakeFontTextBasic(CM, "$0.00")
	s.scoreText.RightAlign = true
	s.scoreText.X = 464
	s.scoreText.Y = 48

	BgmPlayer.SetVolume(0.07)
}

func (s *SceneGame) Update() error {

	if Debug {
		if InputIsPJustPressed() {
			s.debugPause = !s.debugPause
		}
		if s.debugPause {
			return nil
		}
	}

	if s.sgState == SGStateDying {
		if s.dyingF <= 0 {
			s.dyingF = 0
			s.sgState = SGStateDeload
		}
		if SndWalking.IsPlaying() {
			SndWalking.Pause()
		}
		s.updateObstacles()
		s.dyingF--
		return nil
	}
	if s.sgState == SGStateDeload {
		if s.gs.P1Lives == 0 {
			s.gr.SetScene(&SceneHighScore{})
			return nil
		}
		s.gr.SetScene(&SceneGame{})
		return nil
	}
	s.blinkF--
	s.gs.P1Score++
	s.scoreText.SetText(fmt.Sprintf("$%.2f", float32(s.gs.P1Score)/100))
	s.updateWorkerMovement()
	s.updateObstacles()
	s.updateCollisions()
	return nil
}

func (s *SceneGame) Draw(screen *ebiten.Image) {
	screen.DrawImage(ImgGameBg, ImgGameBgDrawOp)
	screen.DrawImage(ImgForeman, ImgFormanDrawOp)

	for _, sp := range s.pltSprs {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(sp.X), float64(sp.Y))
		screen.DrawImage(ImgPlatform, op)
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

		if pl.DazeF > 0 {
			screen.DrawImage(ImgsWorker[1], op)
		} else if pl.IsMoving {
			if pl.AnimFHold == 0 {
				pl.AnimFHold = 4
				pl.AnimF++
				if pl.AnimF >= 8 {
					pl.AnimF = 0
				}
			}
			pl.AnimFHold--
			screen.DrawImage(ImgsWorker[pl.AnimF+2], op)
		} else {
			screen.DrawImage(ImgsWorker[0], op)
		}
	}

	for _, obs := range s.obsFalling {
		if int(obs.Type) >= len(ImgsObstacles) {
			log.Fatal("Falling Obstacle type not mappable to sprite sheet")
		}
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(obs.X), float64(obs.Y))
		screen.DrawImage(ImgsObstacles[obs.Type], op)
	}

	// Draw lives icons
	for i := range s.gs.P1Lives {
		op := &ebiten.DrawImageOptions{}
		xPos := 464 - (int(s.gs.P1Lives) * TileSize)
		op.GeoM.Translate(float64(xPos+int(i)*TileSize), 80.0)
		screen.DrawImage(ImgsObstacles[IconLife], op)
	}

	// Draw bomb icons
	for i := range s.gs.P1Bombs {
		op := &ebiten.DrawImageOptions{}
		xPos := 464 - (int(s.gs.P1Bombs) * TileSize)
		op.GeoM.Translate(float64(xPos+int(i)*TileSize), 112)
		screen.DrawImage(ImgsObstacles[ObstacleBomb], op)
	}

	// Blink
	if s.blinkF <= 0 {
		screen.DrawImage(ImgsFaceSheet[7], ImgBlinkOp)
		if s.blinkF < -10 {
			s.blinkF = 400 + int16(s.gs.P1Score%100)
		}
	}

	// Move mouth if message on screen
	if s.msgText.CharMask+1 < s.msgText.TotalChars {
		if !SndForeman.IsPlaying() {
			SndForeman.Rewind()
			SndForeman.Play()
		}
		if s.mouthF <= 0 {
			s.mouthF = 2 + int16(s.gs.P1Score%10)
			s.mouthInd = uint8(rand.IntN(6))
		}
		s.mouthF--
		screen.DrawImage(ImgsFaceSheet[s.mouthInd], ImgMouthOp)
	} else {
		if SndForeman.IsPlaying() {
			SndForeman.Pause()
		}
		screen.DrawImage(ImgsFaceSheet[0], ImgMouthOp)
	}

	// Maybe rework as an infinite stream, couldn't figure it out
	if !BgmPlayer.IsPlaying() {
		BgmPlayer.Rewind()
		BgmPlayer.Play()
	}
	s.scoreText.Draw(screen)
	s.msgText.Draw(screen)

	if Debug {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("%v:%v:%v:%v", s.gs.LvlInd, s.obsRowInd, s.obsPreview.Delay, s.delayRmnd))
	}
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
				X: (int16(wRow.LBnd) * TileSize) + (int16(i) * TileSize) + TileSize,
				Y: ScreenHeight - TileSize - (int16(wRow.ZPos) * TileSize),
				//Img: ImgPlatform, Implied
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
			pl.AnimF = 0
			s.pls[0] = pl
		case 1:
			pl.AnimF = 3
			s.pls[1] = pl
		case 2:
			pl.AnimF = 6
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
		s.pls[i].IsMoving = false
		s.pls[i].Hold = InputIsHoldPressed(plLen, i)
		if s.pls[i].DazeF > 0 {
			s.pls[i].DazeF--
			s.pls[i].Hold = true
		}
	}
	anyMoving := false
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
					pl.IsMoving = true
					pl.DirR = false
					pl.X -= XVel
					if pl.X < prevBnd {
						pl.X = prevBnd
						pl.IsMoving = false
					}
					prevBnd = pl.X + TileSize
					if pl.IsMoving {
						anyMoving = true
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
					pl.IsMoving = true
					pl.DirR = true
					pl.X += XVel
					if pl.X+TileSize > prevBnd {
						pl.X = prevBnd - TileSize
						pl.IsMoving = false
					}
					prevBnd = pl.X
					if pl.IsMoving {
						anyMoving = true
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

	if !anyMoving && SndWalking.IsPlaying() {
		SndWalking.Pause()
	} else if anyMoving && !SndWalking.IsPlaying() {
		SndWalking.Rewind()
		SndWalking.Play()
	}

	if InputIsBombJustPressed() {
		if s.gs.LvlInd == 0 {
			ResetGameState(s.gs)
			s.gs.LvlInd++
			s.sgState = SGStateDeload
			return
		}

		if s.gs.P1Bombs > 0 {
			s.gs.P1Bombs--

			if !SndBomb.IsPlaying() {
				SndBomb.Rewind()
				SndBomb.Play()
			}

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

		if s.obsPreview.Repeat == 0 {
			s.obsRowInd++
			if int(s.obsRowInd) >= len(s.lvl.Obs) {
				if s.gs.LvlInd == 0 {
					ResetGameState(s.gs)
				}
				s.gs.LvlInd++
				if int(s.gs.LvlInd) >= len(*Levels) {
					s.gr.SetScene(&SceneHighScore{})
				}
				s.sgState = SGStateDeload
				return
			}
		}

		lrpt := s.obsPreview.Repeat
		s.obsPreview = s.lvl.Obs[s.obsRowInd]
		if lrpt > 0 {
			s.obsPreview.Repeat = lrpt - 1
		}
		if s.obsPreview.Obs[0] >= ObstaclePatternRandom {
			// Randomize
			// Use copy with new backing
			orgObs := s.obsPreview.Obs
			obsLen := len(orgObs)
			s.obsPreview.Obs = make([]uint8, obsLen)
			for i := range obsLen {
				s.obsPreview.Obs[i] = orgObs[i]
			}
			s.obsPreview.Obs[0] -= ObstaclePatternRandom
			/*for i := len(s.obsPreview.Obs) - 1; i > 0; i-- { // Fisher–Yates shuffle
				j := rand.IntN(i + 1)
				s.obsPreview.Obs[i], s.obsPreview.Obs[j] = s.obsPreview.Obs[j], s.obsPreview.Obs[i]
			}*/
			s.obsRand.Shuffle(len(s.obsPreview.Obs), func(i, j int) {
				s.obsPreview.Obs[i], s.obsPreview.Obs[j] = s.obsPreview.Obs[j], s.obsPreview.Obs[i]
			})

		}
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
			fmt.Println("Obstable Dequeue index out of range ", di, obsFLen)
			continue
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
					for i := range s.pls {
						s.pls[i].IsMoving = false
						s.pls[i].DazeF = 1000
					}
					return
				}
			}
		}
	}
}

func (s *SceneGame) handleCollision(obsType uint8, pInd int) bool {
	//TODO: BUG WITH BEAM WHILE MOVING
	switch obsType {
	case ObstacleBucket:
		if !SndBucketBonk.IsPlaying() {
			SndBucketBonk.Rewind()
			SndBucketBonk.Play()
		}
		if s.pls[pInd].DazeF > 0 && s.pls[pInd].DazeF < MaxDaze-32 {
			s.gs.P1Lives--
			return true
		}
		s.pls[pInd].DazeF = MaxDaze
		s.pls[pInd].IsMoving = false
	case ObstacleBeam:
		if !SndBeamBonk.IsPlaying() {
			SndBeamBonk.Rewind()
			SndBeamBonk.Play()
		}
		if s.gs.P1Lives > 0 {
			s.gs.P1Lives--
			return true
		}
	case ObstacleSandwich:
		if !SndCash.IsPlaying() {
			SndCash.Rewind()
			SndCash.Play()
		}
		s.gs.P1Score += SandwichPts
	case ObstacleCash:
		if !SndCash.IsPlaying() {
			SndCash.Rewind()
			SndCash.Play()
		}
		s.gs.P1Score += CashPts
	case ObstacleBomb:
		if s.gs.P1Bombs < 6 {
			if !SndPowerup.IsPlaying() {
				SndPowerup.Rewind()
				SndPowerup.Play()
			}
			s.gs.P1Bombs++
		}
	case ObstacleCloud:
		s.gs.P1Score += CloudPts
	case IconLife:
		if s.gs.P1Lives < 6 {
			if !SndPowerup.IsPlaying() {
				SndPowerup.Rewind()
				SndPowerup.Play()
			}
			s.gs.P1Lives++
		}
	}
	return false
}
