package game

import "math"

type LWorkerRow struct {
	// Placement of workers and platforms
	// All in tile units 0-21
	LBnd uint8 // left bound
	RBnd uint8 // right bound
	ZPos uint8 // T offset from bottom
}

type LWorkerPos struct {
	WID    uint8
	RowInd uint8
	RowPos uint8
}

const ObstacleTileCount = 7
const (
	ObstacleNil uint8 = iota
	ObstacleBeam
	ObstacleBucket
	ObstacleSandwich
	ObstacleCash
	ObstacleBomb
	ObstacleCloud
)

type LObsRow struct {
	Obs    []uint8 // iotas above
	Delay  int16   // (4.25s max)
	MsgInd uint8
}

type GameLevel struct {
	WRows []LWorkerRow
	WPos  []LWorkerPos
	Obs   []LObsRow
}

func DelSec(sec float64) int16 {
	return int16(math.Floor(sec * TPS))
}

func MakeLevels() *[]GameLevel {
	return &[]GameLevel{
		// Level 0
		{
			WRows: []LWorkerRow{
				{LBnd: 5, RBnd: 16, ZPos: 1},
				//{LBnd: 0, RBnd: 21, ZPos: 2},
				//{LBnd: 0, RBnd: 21, ZPos: 3},
			},
			WPos: []LWorkerPos{
				{WID: 0, RowInd: 0, RowPos: 9},
				{WID: 1, RowInd: 0, RowPos: 10},
				{WID: 2, RowInd: 0, RowPos: 11},
			},
			Obs: []LObsRow{
				{Obs: []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Delay: DelSec(4), MsgInd: 6},
				{Obs: []uint8{1, 2, 3, 4, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 4, 3, 2, 1}, Delay: DelSec(4), MsgInd: 7},
				{Obs: []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Delay: DelSec(5), MsgInd: 8},
				{Obs: []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Delay: DelSec(4), MsgInd: 9},
				{Obs: []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Delay: DelSec(4), MsgInd: 10},
				{Obs: []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Delay: DelSec(4), MsgInd: 11},
				{Obs: []uint8{3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3}, Delay: DelSec(5), MsgInd: 12},
				{Obs: []uint8{4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4}, Delay: DelSec(7), MsgInd: 13},
				{Obs: []uint8{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}, Delay: DelSec(7), MsgInd: 14},
				{Obs: []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Delay: DelSec(5), MsgInd: 15},
				{Obs: []uint8{1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1}, Delay: DelSec(5), MsgInd: 16},
				{Obs: []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Delay: DelSec(7), MsgInd: 17},
				{Obs: []uint8{5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5}, Delay: DelSec(4), MsgInd: 18},
				{Obs: []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Delay: DelSec(4), MsgInd: 19},
				{Obs: []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Delay: DelSec(4), MsgInd: 1},
				{Obs: []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Delay: 0},
			},
		},
		{
			WRows: []LWorkerRow{
				{LBnd: 0, RBnd: 21, ZPos: 1},
				//{LBnd: 0, RBnd: 21, ZPos: 2},
				//{LBnd: 0, RBnd: 21, ZPos: 3},
			},
			WPos: []LWorkerPos{
				{WID: 0, RowInd: 0, RowPos: 10},
				//{WID: 1, RowInd: 0, RowPos: 10},
				//{WID: 2, RowInd: 0, RowPos: 11},
			},
			Obs: []LObsRow{
				{Obs: []uint8{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, Delay: DelSec(5), MsgInd: 21},
				{Obs: []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Delay: DelSec(3)},
				{Obs: []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Delay: DelSec(4), MsgInd: 1},
				{Obs: []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Delay: 0},
			},
		},
	}
}
