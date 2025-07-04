package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type SceneHighScore struct {
	gr          GameRoot
	gs          *GameState
	scrText     *FontText
	entryText   *FontText
	creditsText *FontText
	scrStrs     []string
	name        string
	idleF       int16
	creditsF    int16
	hsIndex     uint8
	crsX        uint8
	crsY        uint8
	creditsInd  uint8
}

const (
	HSEntryIdleMax  = TPS * 15
	EntryCursorMaxX = 8
	EntryCursorMaxY = 2
	NameMax         = 6
	HSIndNull       = 255
)

var EntryLookup = []string{
	"ABCDEFGHI",
	"JKLMNOPQR",
	"STUVWXYZ>",
}

var CreditsStrs = [][]string{
	{
		"Programming, Art and sound",
		"Justin Horton",
		"github.com/BossRighteous",
	},
	{
		"Music",
		"HardHatBop",
		"By my bestie Jake Schofield",
	},
	{
		"Developed using Go and Ebitengine",
		"Ebitengine Game Jam 2025 Entry",
	},
	{
		"Although it's for a Jam, I intend",
		"to develop it further into a physical",
		"arcade cabinet",
	},
	{
		"Thank you for playing!",
	},
}

func (s *SceneHighScore) Init(gr GameRoot, gs *GameState) {
	s.gr = gr
	s.gs = gs
	s.hsIndex = HSIndNull
	s.name = ""
	s.idleF = 0
	s.crsX = 0
	s.crsY = 0

	s.scrStrs = make([]string, len(s.gs.HighScores))
	if s.InjectScore() {
		s.entryText = MakeFontText(CM, []string{
			EntryLookup[0],
			EntryLookup[1],
			EntryLookup[2],
			"",
			"J: Accept",
			"SPACE: DEL",
		})
		s.entryText.X = 200
		s.entryText.Y = 32
		s.entryText.LineSpace = 4
	}

	s.scrText = MakeFontText(CM, s.scrStrs)
	s.scrText.X = 32
	s.scrText.Y = 32
	s.scrText.LineSpace = 4

	s.creditsText = MakeFontText(CM, CreditsStrs[s.creditsInd])
	s.creditsText.X = 32
	s.creditsText.Y = 192
	s.creditsText.LineSpace = 4
}

func (s *SceneHighScore) Update() error {
	if s.idleF > HSEntryIdleMax {
		s.gr.SetScene(&SceneTitle{})
		return nil
	}

	if s.hsIndex != HSIndNull {
		s.handleInput()
	} else {
		if s.creditsF > TPS*3 {
			s.creditsF = 0
			s.creditsInd++
			if int(s.creditsInd) < len(CreditsStrs) {
				s.creditsText.SetText(CreditsStrs[s.creditsInd])
			}
		}
		s.creditsF++
	}

	s.idleF++
	return nil
}

func (s *SceneHighScore) handleInput() {
	if InputIsLeftJustPressed() {
		s.idleF = 0
		if s.crsX == 0 {
			if s.crsY == 0 {
				s.crsY = EntryCursorMaxY
				s.crsX = EntryCursorMaxX
				return
			}
			s.crsY--
			s.crsX = EntryCursorMaxX
			return
		}
		s.crsX--
	}
	if InputIsRightJustPressed() {
		s.idleF = 0
		if s.crsX == EntryCursorMaxX {
			if s.crsY == EntryCursorMaxY {
				s.crsY = 0
				s.crsX = 0
				return
			}
			s.crsY++
			s.crsX = 0
			return
		}
		s.crsX++
	}
	if InputIsAJustPressed() {
		s.idleF = 0
		if len(s.name) >= NameMax {
			return
		}
		ch := EntryLookup[s.crsY][s.crsX : s.crsX+1]
		if ch == ">" {
			s.hsIndex = HSIndNull
			return
		}
		s.name = s.name + ch
		s.updateName()
	}
	if InputIsBombJustPressed() {
		s.idleF = 0
		if len(s.name) == 0 {
			return
		}
		s.name = s.name[:len(s.name)-1]
		s.updateName()
	}
}

func (s *SceneHighScore) updateName() {
	s.gs.HighScores[s.hsIndex].Name = s.name
	s.scrStrs[s.hsIndex] = s.fmtScoreText(s.gs.HighScores[s.hsIndex])
	s.scrText.SetText(s.scrStrs)
	s.scrText.CharMask = 256 // Skip redraw
}

func (s *SceneHighScore) Draw(screen *ebiten.Image) {
	screen.DrawImage(ImgGameBg, ImgGameBgDrawOp)

	if s.hsIndex != HSIndNull && s.entryText != nil {
		s.entryText.Draw(screen)

		hlop := &ebiten.DrawImageOptions{}
		hlop.GeoM.Translate(float64(s.entryText.X+int16(s.crsX)*8)-1, float64(s.entryText.Y+int16(s.crsY)*(8+int16(s.entryText.LineSpace)))-1)
		screen.DrawImage(ImgNameHightlight, hlop)
	}
	s.scrText.Draw(screen)
	s.creditsText.Draw(screen)

	if !BgmPlayer.IsPlaying() {
		BgmPlayer.Rewind()
		BgmPlayer.Play()
	}

}

func (s *SceneHighScore) Exit() {}

func (s *SceneHighScore) InjectScore() bool {
	sc := s.gs.P1Score
	swap := HighScore{}
	for i, hs := range s.gs.HighScores {
		if s.hsIndex == HSIndNull && sc > hs.Score {
			swap = s.gs.HighScores[i]
			s.hsIndex = uint8(i)
			s.gs.HighScores[i].Name = s.name
			s.gs.HighScores[i].Score = sc
		}
		if s.hsIndex < uint8(i) && i < 10 {
			s2 := s.gs.HighScores[i]
			s.gs.HighScores[i] = swap
			swap = s2
		}

		s.scrStrs[i] = s.fmtScoreText(s.gs.HighScores[i])
	}
	return s.hsIndex != HSIndNull
}

func (s *SceneHighScore) fmtScoreText(hs HighScore) string {
	spaceStr := "                "
	scoreStr := fmt.Sprintf("$%.2f", float32(hs.Score)/100)
	return fmt.Sprintf("%s%s%s", hs.Name, spaceStr[len(hs.Name):len(spaceStr)-len(scoreStr)], scoreStr)
}
