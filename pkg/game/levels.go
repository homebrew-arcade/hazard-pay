package game

import "math"

type WorkerRow struct {
	// Placement of workers and platforms
	// All in tile units 0-21
	LBnd uint8 // left bound
	RBnd uint8 // right bound
	ZPos uint8 // T offset from bottom
}

type WorkerPos struct {
	WID    uint8
	RowInd uint8
	RowPos uint8
}

const (
	ObstacleNil uint8 = iota
	ObstacleBeam
	ObstacleBucket
	ObstacleSandwich
	ObstacleCash
)

type ObsRow struct {
	Obs   []uint8 // iotas above
	Delay int16   // (4.25s max)
}

type GameLevel struct {
	WRows []WorkerRow
	WPos  []WorkerPos
	Obs   []ObsRow
}

func DelSec(sec float64) int16 {
	return int16(math.Floor(sec * TPS))
}

func MakeLevels() *[]GameLevel {
	return &[]GameLevel{
		// Level 0
		{
			WRows: []WorkerRow{
				{LBnd: 0, RBnd: 21, ZPos: 1},
				//{LBnd: 0, RBnd: 21, ZPos: 2},
				//{LBnd: 0, RBnd: 21, ZPos: 4},
			},
			WPos: []WorkerPos{
				{WID: 0, RowInd: 0, RowPos: 8},
				{WID: 1, RowInd: 0, RowPos: 10},
				{WID: 2, RowInd: 0, RowPos: 12},
			},
			Obs: []ObsRow{
				{Obs: []uint8{0, 1, 2, 3, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, Delay: DelSec(4)},
				{Obs: []uint8{0, 1, 2, 0, 1, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0}, Delay: DelSec(4)},
				{Obs: []uint8{0, 1, 2, 0, 1, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, Delay: DelSec(4)},
				{Obs: []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Delay: DelSec(99)},
				{Obs: []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Delay: 0},
			},
		},
	}
}
