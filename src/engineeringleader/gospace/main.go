package main

import (
	"embed"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed assets/*
var assets embed.FS
var PlayerSprite = mustLoadImage("assets/player.png")

type Vector struct {
	X float64
	Y float64
}

func mustLoadImage(name string) *ebiten.Image {
	println(name)
	f, err := assets.Open(name)
	if err != nil {

		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}

func main() {

	g := &Game{
		player: NewPlayer(),
	}

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
