package game

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	s       Scene
	queuedS Scene
	gs      *GameState
}

func (g *Game) Init() {
	g.s = &SceneTitle{}
	g.queuedS = nil
	g.gs = MakeGameState()
}

func (g *Game) Update() error {
	if g.queuedS != nil {
		g.s = g.queuedS
		g.queuedS = nil
		g.s.Init(g, g.gs)
	}
	err := g.s.Update()
	if err != nil {
		g.Exit()
	}
	return err
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(BGColor)
	g.s.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Enter() {
	if g.s == nil {
		log.Fatal("no scene provided to gameroot")
	}
	g.s.Init(g, g.gs)
}

func (g *Game) SetScene(ns Scene) {
	if g.s != nil {
		g.s.Exit()
	}
	g.queuedS = ns
}

func (g *Game) Exit() {
	g.s = nil
}

func Main() {
	ebiten.SetWindowSize(WindowWidth, WindowHeight)
	ebiten.SetWindowTitle("Hazard Pay")
	ebiten.SetTPS(TPS)
	ebiten.SetVsyncEnabled(VSyncEnabled)
	g := Game{}
	g.Init()
	g.Enter()
	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}
