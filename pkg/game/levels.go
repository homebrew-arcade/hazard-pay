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

const ObstacleTileCount = 9
const (
	ObstacleNil uint8 = iota
	ObstacleBeam
	ObstacleBucket
	ObstacleSandwich
	ObstacleCash
	ObstacleBomb
	ObstacleCloud

	IconLife
	TilePlatform
)

// If Obs[0] >100, randomize positions of following indexes and Obs[0] -= 100
const ObstaclePatternRandom = 100

type LObsRow struct {
	Obs    []uint8 // iotas above
	Delay  int16
	MsgInd uint8
	Repeat uint8
}

type GameLevel struct {
	WRows []LWorkerRow
	WPos  []LWorkerPos
	Obs   []LObsRow
	SeedA uint64
	SeedB uint64
}

func DelSec(sec float64) int16 {
	return int16(math.Floor(sec * TPS))
}

var Levels = &[]GameLevel{
	// Level 0
	{
		SeedA: 0,
		SeedB: 0,
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
			{Obs: []uint8{6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6}, Delay: DelSec(7), MsgInd: 17},
			{Obs: []uint8{5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5}, Delay: DelSec(4), MsgInd: 18},
			// standard tail
			{Obs: []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Delay: DelSec(4)},
			{Obs: []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Delay: DelSec(4), MsgInd: 1},
			{Obs: []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Delay: 0},
		},
	},
	// Level 1
	{
		SeedA: 9388503,
		SeedB: 20349869920041,
		WRows: []LWorkerRow{
			{LBnd: 0, RBnd: 21, ZPos: 1},
			//{LBnd: 0, RBnd: 21, ZPos: 2},
			//{LBnd: 0, RBnd: 21, ZPos: 3},
		},
		WPos: []LWorkerPos{
			{WID: 0, RowInd: 0, RowPos: 10},
			//{WID: 1, RowInd: 0, RowPos: 11},
			//{WID: 2, RowInd: 0, RowPos: 11},
		},
		Obs: []LObsRow{
			{Obs: []uint8{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 0, 2, 2, 2, 2, 2, 2, 2, 2, 2}, Delay: DelSec(5), MsgInd: 21},
			{Obs: []uint8{101, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 3, 0, 0, 0, 0, 0, 0}, Delay: DelSec(1), Repeat: 10},
			{Obs: []uint8{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 0}, Delay: DelSec(3.5)},
			{Obs: []uint8{0, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, Delay: DelSec(3.5)},
			{Obs: []uint8{1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1}, Delay: DelSec(2)},
			{Obs: []uint8{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}, Delay: DelSec(1)},
			{Obs: []uint8{4, 2, 2, 3, 2, 2, 0, 2, 2, 0, 2, 2, 0, 2, 2, 3, 2, 2, 4, 2}, Delay: DelSec(1.5)},
			{Obs: []uint8{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}, Delay: DelSec(2)},
			{Obs: []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Delay: DelSec(0.5)},
			{Obs: []uint8{101, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 3, 0, 0, 0, 0, 0, 0}, Delay: DelSec(1), Repeat: 10},
			{Obs: []uint8{101, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 1, 1, 1, 4}, Delay: DelSec(2), Repeat: 10},

			{Obs: []uint8{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 0}, Delay: DelSec(3.5)},
			{Obs: []uint8{0, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, Delay: DelSec(3.5)},
			{Obs: []uint8{1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1}, Delay: DelSec(2)},

			{Obs: []uint8{4, 5, 4, 5, 4, 5, 4, 5, 4, 5, 4, 5, 4, 5, 4, 5, 4, 5, 4, 5}, Delay: DelSec(3.5)},
			{Obs: []uint8{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, Delay: DelSec(1)},
			// standard tail
			{Obs: []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Delay: DelSec(4)},
			{Obs: []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Delay: DelSec(4), MsgInd: 1},
			{Obs: []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Delay: 0},
		},
	},
	// Level 2
	{
		SeedA: 94578312135,
		SeedB: 647,
		WRows: []LWorkerRow{
			{LBnd: 0, RBnd: 21, ZPos: 1},
			//{LBnd: 0, RBnd: 21, ZPos: 2},
			//{LBnd: 0, RBnd: 21, ZPos: 3},
		},
		WPos: []LWorkerPos{
			{WID: 0, RowInd: 0, RowPos: 10},
			{WID: 1, RowInd: 0, RowPos: 11},
			//{WID: 2, RowInd: 0, RowPos: 11},
		},
		Obs: []LObsRow{
			{Obs: []uint8{1, 1, 1, 1, 1, 1, 1, 1, 2, 0, 0, 2, 1, 1, 1, 1, 1, 1, 1, 1}, Delay: DelSec(5), MsgInd: 25},
			{Obs: []uint8{1, 1, 1, 1, 1, 1, 1, 1, 0, 2, 2, 0, 1, 1, 1, 1, 1, 1, 1, 1}, Delay: DelSec(1)},
			{Obs: []uint8{1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1}, Delay: DelSec(1)},
			{Obs: []uint8{1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 1, 1, 1, 1}, Delay: DelSec(1)},
			{Obs: []uint8{1, 1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 1, 1}, Delay: DelSec(1)},
			{Obs: []uint8{1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 1}, Delay: DelSec(1)},
			{Obs: []uint8{1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1}, Delay: DelSec(1)},
			{Obs: []uint8{3, 1, 1, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 1, 1, 3}, Delay: DelSec(1.5)},
			{Obs: []uint8{0, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 0}, Delay: DelSec(1.5)},
			{Obs: []uint8{3, 3, 1, 1, 2, 1, 1, 1, 4, 1, 1, 4, 1, 1, 1, 2, 1, 1, 3, 3}, Delay: DelSec(1)},
			{Obs: []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Delay: DelSec(5), MsgInd: 23},
			{Obs: []uint8{2, 0, 2, 0, 2, 0, 2, 0, 2, 2, 2, 2, 0, 2, 0, 2, 0, 2, 0, 2}, Delay: DelSec(1)},
			{Obs: []uint8{0, 2, 0, 2, 0, 2, 0, 2, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 2, 0}, Delay: DelSec(1)},
			{Obs: []uint8{2, 0, 2, 0, 2, 0, 2, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 2, 0, 2}, Delay: DelSec(1)},
			{Obs: []uint8{0, 2, 0, 2, 0, 2, 0, 2, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 2, 0}, Delay: DelSec(1)},
			{Obs: []uint8{2, 0, 2, 0, 2, 0, 2, 0, 2, 0, 2, 0, 0, 2, 0, 2, 0, 2, 0, 2}, Delay: DelSec(1)},
			{Obs: []uint8{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1}, Delay: DelSec(1.5)},
			{Obs: []uint8{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0}, Delay: DelSec(1.5)},
			{Obs: []uint8{101, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 3, 4, 3, 0, 0, 0, 0}, Delay: DelSec(2.5), Repeat: 10},
			{Obs: []uint8{4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4}, Delay: DelSec(5), MsgInd: 24},
			{Obs: []uint8{0, 0, 2, 0, 5, 0, 0, 2, 0, 1, 1, 0, 2, 0, 0, 5, 0, 2, 0, 0}, Delay: DelSec(1)},
			{Obs: []uint8{4, 0, 2, 0, 1, 0, 0, 2, 0, 5, 5, 0, 2, 0, 0, 1, 0, 2, 0, 4}, Delay: DelSec(1)},
			{Obs: []uint8{0, 2, 0, 0, 0, 2, 0, 0, 2, 0, 0, 2, 0, 0, 2, 0, 0, 0, 2, 0}, Delay: DelSec(1)},
			{Obs: []uint8{0, 0, 0, 4, 0, 0, 0, 0, 4, 0, 0, 4, 0, 0, 0, 0, 4, 0, 0, 0}, Delay: DelSec(1)},
			// standard tail
			{Obs: []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Delay: DelSec(4)},
			{Obs: []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Delay: DelSec(4), MsgInd: 1},
			{Obs: []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Delay: 0},
		},
	},
	// Level 3
	{
		SeedA: 5345673,
		SeedB: 2995039946,
		WRows: []LWorkerRow{
			{LBnd: 1, RBnd: 10, ZPos: 1},
			{LBnd: 6, RBnd: 15, ZPos: 2},
			{LBnd: 11, RBnd: 20, ZPos: 1},
		},
		WPos: []LWorkerPos{
			{WID: 0, RowInd: 0, RowPos: 2},
			{WID: 1, RowInd: 1, RowPos: 7},
			{WID: 2, RowInd: 2, RowPos: 12},
		},
		Obs: []LObsRow{
			{Obs: []uint8{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Delay: DelSec(5), MsgInd: 26},
			{Obs: []uint8{101, 1, 1, 1, 2, 2, 2, 3, 3, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Delay: DelSec(4), Repeat: 10},
			{Obs: []uint8{0, 0, 2, 2, 2, 2, 2, 2, 2, 0, 0, 2, 2, 2, 2, 2, 2, 2, 0, 0}, Delay: DelSec(3), MsgInd: 27},
			{Obs: []uint8{0, 0, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 1, 1, 1, 0, 0}, Delay: DelSec(3)},
			{Obs: []uint8{0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 0, 0}, Delay: DelSec(0.8)},
			{Obs: []uint8{0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 0, 0}, Delay: DelSec(0.8)},
			{Obs: []uint8{0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 0, 0}, Delay: DelSec(0.8)},
			{Obs: []uint8{0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 0, 0}, Delay: DelSec(0.8)},
			{Obs: []uint8{0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 0, 0}, Delay: DelSec(0.8)},
			{Obs: []uint8{0, 0, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0}, Delay: DelSec(0.8)},
			{Obs: []uint8{0, 0, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0}, Delay: DelSec(0.8)},
			{Obs: []uint8{0, 0, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0}, Delay: DelSec(0.8)},
			{Obs: []uint8{0, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0}, Delay: DelSec(0.8)},
			{Obs: []uint8{0, 0, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0}, Delay: DelSec(0.8)},
			{Obs: []uint8{0, 0, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0}, Delay: DelSec(0.8)},
			{Obs: []uint8{0, 1, 0, 0, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0}, Delay: DelSec(0.8)},
			{Obs: []uint8{0, 1, 1, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0}, Delay: DelSec(0.8)},
			{Obs: []uint8{0, 1, 1, 1, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0}, Delay: DelSec(0.8)},
			{Obs: []uint8{0, 1, 1, 1, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0}, Delay: DelSec(0.8)},
			{Obs: []uint8{0, 1, 1, 1, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 0, 0, 1, 0}, Delay: DelSec(0.8)},
			{Obs: []uint8{0, 1, 1, 1, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 0, 0, 1, 1, 0}, Delay: DelSec(0.8)},
			{Obs: []uint8{0, 1, 1, 1, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 0, 0, 1, 1, 1, 0}, Delay: DelSec(0.8)},
			{Obs: []uint8{0, 1, 1, 1, 4, 1, 1, 1, 1, 5, 1, 1, 1, 1, 4, 1, 1, 1, 1, 0}, Delay: DelSec(0.8)},
			{Obs: []uint8{101, 3, 3, 3, 4, 4, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Delay: DelSec(1), Repeat: 5},

			// standard tail
			{Obs: []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Delay: DelSec(4)},
			{Obs: []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Delay: DelSec(4), MsgInd: 1},
			{Obs: []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Delay: 0},
			//{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			//{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			//{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
			//{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	},
	// Level 4
	{
		SeedA: 2000498299555,
		SeedB: 55555,
		WRows: []LWorkerRow{
			{LBnd: 0, RBnd: 21, ZPos: 1},
		},
		WPos: []LWorkerPos{
			{WID: 0, RowInd: 0, RowPos: 10},
			{WID: 1, RowInd: 0, RowPos: 11},
			{WID: 2, RowInd: 0, RowPos: 12},
		},
		Obs: []LObsRow{
			{Obs: []uint8{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Delay: DelSec(5), MsgInd: 28},
			{Obs: []uint8{1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1}, Delay: DelSec(2)},
			{Obs: []uint8{1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1}, Delay: DelSec(2)},
			{Obs: []uint8{1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1}, Delay: DelSec(2)},
			{Obs: []uint8{1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 0, 0, 1, 1, 1, 1, 1, 1, 1}, Delay: DelSec(2)},
			{Obs: []uint8{1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 0, 0, 1, 1, 1, 1, 1, 1}, Delay: DelSec(2)},
			{Obs: []uint8{1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 0, 0, 1, 1, 1, 1, 1}, Delay: DelSec(2)},
			{Obs: []uint8{1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 0, 1, 0, 1, 1, 1, 1}, Delay: DelSec(2)},
			{Obs: []uint8{1, 1, 1, 1, 1, 1, 1, 1, 1, 4, 1, 1, 1, 5, 1, 4, 1, 1, 1, 1}, Delay: DelSec(2)},
			{Obs: []uint8{1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 0, 1, 0, 1, 1, 1, 1}, Delay: DelSec(2)},
			{Obs: []uint8{1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 0, 1, 1, 0, 1, 1, 1, 1}, Delay: DelSec(2)},
			{Obs: []uint8{1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 0, 1, 1, 0, 1, 1, 1, 1}, Delay: DelSec(2)},
			{Obs: []uint8{1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 0, 1, 1, 0, 1, 1, 1, 1, 1}, Delay: DelSec(2)},
			{Obs: []uint8{1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 0, 1, 1, 0, 1, 1, 1, 1, 1, 1}, Delay: DelSec(2)},
			{Obs: []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Delay: DelSec(2)},
			{Obs: []uint8{4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4}, Delay: DelSec(2)},
			{Obs: []uint8{3, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3}, Delay: DelSec(5)},
			{Obs: []uint8{0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0}, Delay: DelSec(1.5)},
			{Obs: []uint8{0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1}, Delay: DelSec(1.5)},
			{Obs: []uint8{1, 0, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1}, Delay: DelSec(1.5)},
			{Obs: []uint8{1, 0, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1}, Delay: DelSec(1.5)},
			{Obs: []uint8{1, 1, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1}, Delay: DelSec(1.5)},
			{Obs: []uint8{1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 0, 1}, Delay: DelSec(1.5)},
			{Obs: []uint8{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 0}, Delay: DelSec(1.5)},
			{Obs: []uint8{1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 0, 1}, Delay: DelSec(1.5)},
			{Obs: []uint8{1, 1, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1}, Delay: DelSec(1.5)},
			{Obs: []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Delay: DelSec(2)},
			{Obs: []uint8{4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4}, Delay: DelSec(2)},
			{Obs: []uint8{3, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3}, Delay: DelSec(5)},
			{Obs: []uint8{4, 1, 1, 1, 1, 1, 1, 1, 1, 4, 1, 1, 1, 1, 1, 1, 1, 1, 1, 4}, Delay: DelSec(1.5)},
			{Obs: []uint8{1, 1, 1, 1, 1, 1, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1}, Delay: DelSec(4)},
			{Obs: []uint8{102, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 1, 1, 3, 1, 4}, Delay: DelSec(1.5), Repeat: 25},

			// standard tail
			{Obs: []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Delay: DelSec(4)},
			{Obs: []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Delay: DelSec(5), MsgInd: 30},
			{Obs: []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Delay: DelSec(5), MsgInd: 31},
			{Obs: []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Delay: DelSec(5), MsgInd: 32},
			{Obs: []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Delay: DelSec(5), MsgInd: 33},
			{Obs: []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Delay: DelSec(4), MsgInd: 1},
			{Obs: []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Delay: 0},
			//{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			//{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			//{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
			//{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	},
	// Level 5
	{
		SeedA: 12345,
		SeedB: 67890,
		WRows: []LWorkerRow{
			{LBnd: 0, RBnd: 21, ZPos: 1},
		},
		WPos: []LWorkerPos{
			{WID: 0, RowInd: 0, RowPos: 10},
			{WID: 1, RowInd: 0, RowPos: 11},
		},
		Obs: []LObsRow{
			{Obs: []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Delay: DelSec(5), MsgInd: 29},
			{Obs: []uint8{4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4}, Delay: DelSec(5), MsgInd: 34},
			{Obs: []uint8{101, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 3, 3, 4, 4, 3}, Delay: DelSec(3), Repeat: 10},
			{Obs: []uint8{5, 4, 5, 4, 5, 4, 5, 4, 5, 4, 5, 4, 5, 4, 5, 4, 5, 4, 5, 4}, Delay: DelSec(1)},
			{Obs: []uint8{4, 5, 4, 5, 4, 5, 4, 5, 4, 5, 4, 5, 4, 5, 4, 5, 4, 5, 4, 5}, Delay: DelSec(1)},
			{Obs: []uint8{101, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 3, 3, 4, 4, 3}, Delay: DelSec(2.5), Repeat: 20},
			{Obs: []uint8{5, 4, 5, 4, 5, 4, 5, 4, 5, 4, 5, 4, 5, 4, 5, 4, 5, 4, 5, 4}, Delay: DelSec(1)},
			{Obs: []uint8{4, 5, 4, 5, 4, 5, 4, 5, 4, 5, 4, 5, 4, 5, 4, 5, 4, 5, 4, 5}, Delay: DelSec(1)},
			{Obs: []uint8{101, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 3, 3, 4, 4, 3}, Delay: DelSec(2), Repeat: 30},
			{Obs: []uint8{5, 4, 5, 4, 5, 4, 5, 4, 5, 4, 5, 4, 5, 4, 5, 4, 5, 4, 5, 4}, Delay: DelSec(1)},
			{Obs: []uint8{4, 5, 4, 5, 4, 5, 4, 5, 4, 5, 4, 5, 4, 5, 4, 5, 4, 5, 4, 5}, Delay: DelSec(1)},
			{Obs: []uint8{101, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 3, 3, 4, 4, 3}, Delay: DelSec(1.5), Repeat: 50},
			{Obs: []uint8{5, 4, 5, 4, 5, 4, 5, 4, 5, 4, 5, 4, 5, 4, 5, 4, 5, 4, 5, 4}, Delay: DelSec(1)},
			{Obs: []uint8{4, 5, 4, 5, 4, 5, 4, 5, 4, 5, 4, 5, 4, 5, 4, 5, 4, 5, 4, 5}, Delay: DelSec(1)},
			{Obs: []uint8{101, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 4, 4, 3}, Delay: DelSec(1.5), Repeat: 255},
			{Obs: []uint8{101, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 4, 4, 3}, Delay: DelSec(1.5), Repeat: 255},
			{Obs: []uint8{101, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 4, 4, 3}, Delay: DelSec(1.5), Repeat: 255},

			{Obs: []uint8{4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4}, Delay: DelSec(5), MsgInd: 35},
			{Obs: []uint8{4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4}, Delay: DelSec(0.25), Repeat: 255},

			// standard tail
			{Obs: []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Delay: DelSec(4)},
			{Obs: []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Delay: DelSec(4), MsgInd: 1},
			{Obs: []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, Delay: 0},
			//{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			//{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			//{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
			//{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	},
}
