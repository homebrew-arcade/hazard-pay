package game

type GameState struct {
	Lvls    *[]GameLevel
	IdleF   uint // Idle frame ticker for attract mode
	P1Score uint
	LvlInd  uint8
	P1Lives uint8
	P1Bombs uint8
	Credits uint8
	IsIdle  bool
}

func MakeGameState() *GameState {
	gs := GameState{
		LvlInd:  0,
		Lvls:    MakeLevels(),
		P1Score: 0,
		P1Lives: 3,
		P1Bombs: 2,
		Credits: 0,
		IdleF:   0,
		IsIdle:  false,
	}
	return &gs
}
