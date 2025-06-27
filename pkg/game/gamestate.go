package game

type HighScore struct {
	Name  string
	Score uint
}

type GameState struct {
	HighScores []HighScore
	P1Score    uint
	LvlInd     uint8
	P1Lives    uint8
	P1Bombs    uint8
}

func MakeGameState() *GameState {
	gs := &GameState{
		HighScores: []HighScore{
			{Name: "JUSTIN", Score: 50000},
			{Name: "JBONE", Score: 35000},
			{Name: "BOSSR", Score: 30000},
			{Name: "JAKE", Score: 25000},
			{Name: "SCHOF", Score: 20000},
			{Name: "EBITEN", Score: 15000},
			{Name: "GOLANG", Score: 10000},
			{Name: "ARCADE", Score: 7500},
			{Name: "SADTBH", Score: 5000},
			{Name: "FOOBAR", Score: 1000},
		},
	}
	ResetGameState(gs)
	return gs
}

func ResetGameState(gs *GameState) {
	gs.LvlInd = 0
	gs.P1Score = 0
	gs.P1Lives = 3
	gs.P1Bombs = 2
}
