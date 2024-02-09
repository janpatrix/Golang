package main

import (
	b "gospace/basic"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

var PlayerSprite = b.MustLoadImage("assets/ships/player.png")

type Player struct {
	rotation float64
	positon  Vector
	sprite   *ebiten.Image
}

func NewPlayer() *Player {
	sprite := PlayerSprite

	bounds := sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	pos := Vector{
		X: ScreenWidth/2 - halfW,
		Y: ScreenHeight/2 - halfH,
	}
	return &Player{
		positon: pos,
		sprite:  sprite,
	}
}

func (p *Player) Update() {
	speed := math.Pi / float64(ebiten.TPS())

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.rotation -= speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.rotation += speed
	}

}

func (p *Player) Draw(screen *ebiten.Image) {
	bounds := p.sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Rotate(p.rotation)

	op.GeoM.Translate(halfW, halfH)

	op.GeoM.Scale(0.25, 0.25)

	op.GeoM.Translate(p.positon.X, p.positon.Y)
	screen.DrawImage(p.sprite, op)
}
