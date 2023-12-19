package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	player *Player
}

const (
	ScreenWidth  = 800
	ScreenHeight = 600
)

func (g *Game) Update() error {
	g.player.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
