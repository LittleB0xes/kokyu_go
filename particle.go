package main

import (
	"math/rand"

	sprite "github.com/LittleB0xes/kokyu/sprite"
	"github.com/hajimehoshi/ebiten/v2"
)

type Particle struct {
	sprite       sprite.AnimatedSprite
	transparency []float64
}

func NewParticle(x, y float64, texture *ebiten.Image) Particle {

	part_data := sprite.AnimationData{X: 0, Y: 0, W: 16, H: 16, Frame: 60, Speed: 10}
	sprite := sprite.NewAnimatedSprite(x, y, texture, part_data)

	// Add a "dephasage"
	sprite.SetCurrentFrame(rand.Intn(part_data.Frame))

	transparency := make([]float64, part_data.Frame)
	for i := 0; i < part_data.Frame; i++ {
		transparency[i] = float64(30+rand.Intn(25)) / 255
	}

	return Particle{
		sprite,
		transparency,
	}
}

func (e *Particle) Draw(screen *ebiten.Image, camX, camY, scale float64) {
	e.sprite.SetTransparency(e.transparency[e.sprite.GetCurrentFrame()])
	e.sprite.Draw(screen, camX, camY, scale)

}
