package main

import (
	t "gospace/basic"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	player           *Player
	meteorSpawnTimer *t.Timer
	meteors          []*Meteor
}

const (
	ScreenWidth  = 800
	ScreenHeight = 600
)

func (g *Game) Update() error {
	g.player.Update()

	g.meteorSpawnTimer.Update()
	if g.meteorSpawnTimer.IsReady() {
		g.meteorSpawnTimer.Reset()
		m := NewMeteor()
		g.meteors = append(g.meteors, m)
	}

	for _, m := range g.meteors {
		m.Update()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)

	for _, m := range g.meteors {
		m.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
