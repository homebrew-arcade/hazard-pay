package game

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	s       Scene
	queuedS Scene
}

func (g *Game) Init() {
}

func (g *Game) Update() error {
	if g.queuedS != nil {
		g.s = g.queuedS
		g.queuedS = nil
		g.s.Enter()
	}
	err := g.s.Update()
	if err != nil {
		g.Exit()
	}
	return err
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.s.Draw(screen)

	if Debug {
		ebitenutil.DebugPrint(
			screen,
			fmt.Sprintf("TPS: %.2f, FPS: %.2f", ebiten.ActualTPS(), ebiten.ActualFPS()),
		)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Enter() {
	if g.s == nil {
		log.Fatal("no scene provided to gameroot")
	}
	g.s.Init(g)
	g.s.Enter()
}

func (g *Game) SetScene(t SceneToken) {
	if g.s != nil {
		g.s.Exit()
	}
	g.queuedS = SceneProvider(t)
}

func (g *Game) Exit() {
	g.s = nil
}

func Main() {
	ebiten.SetWindowSize(WindowWidth, WindowHeight)
	ebiten.SetWindowTitle("Hazard Pay")
	g := Game{
		s: nil,
	}
	g.Init()
	g.Enter()
	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}
