package main

import (
	"image"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Level struct {
	background      *ebiten.Image
	ground          *ebiten.Image
	light_text      *ebiten.Image
	healthContainer *ebiten.Image
	healthBar       *ebiten.Image
	title           *ebiten.Image
	particles       []Particle
	lights          []Light
}

func NewBaseLevel() Level {

	background, _, err := ebitenutil.NewImageFromFile("./assets/sprites/Level.png")
	if err != nil {
		log.Fatal("Error when openning file", err)
	}

	ground, _, err := ebitenutil.NewImageFromFile("./assets/sprites/Ground.png")
	if err != nil {
		log.Fatal("Error when openning file", err)
	}
	light_text, _, err := ebitenutil.NewImageFromFile("./assets/sprites/Light.png")
	if err != nil {
		log.Fatal(" Error when openning file", err)
	}
	healthContainer, _, err := ebitenutil.NewImageFromFile("./assets/sprites/Health_deco.png")
	if err != nil {
		log.Fatal("Error when openning file", err)
	}
	healthBar, _, err := ebitenutil.NewImageFromFile("./assets/sprites/Health_bar.png")
	if err != nil {
		log.Fatal("Error when openning file", err)
	}
	title, _, err := ebitenutil.NewImageFromFile("./assets/sprites/Title.png")
	if err != nil {
		log.Fatal("Error when openning file", err)
	}
	part_text, _, err := ebitenutil.NewImageFromFile("./assets/sprites/ParticleOne.png")
	if err != nil {
		log.Fatal("Error when openning file", err)
	}
	particles := make([]Particle, 100)
	for i := 0; i < len(particles); i++ {
		x := rand.Intn(410)
		y := rand.Intn(100)
		particles[i] = NewParticle(float64(x), float64(y), part_text)

	}

	lights := []Light{
		NewLight(118.0, 70.0, 32.0),
		NewLight(117.0, 69.0, 24.0),
		NewLight(119.0, 70.0, 30.0),
		NewLight(329.0, 70.0, 32.0),
		NewLight(328.0, 69.0, 24.0),
		NewLight(330.0, 70.0, 30.0),
	}

	return Level{
		background,
		ground,
		light_text,
		healthContainer,
		healthBar,
		title,
		particles,
		lights,
	}
}

func (l *Level) RenderBackground(screen *ebiten.Image, camX, camY, scale float64) {
	op := &ebiten.DrawImageOptions{}

	// Set the good place of the sprite
	op.GeoM.Translate(-camX, -camY)

	// Apply the global scale
	op.GeoM.Scale(scale, scale)
	screen.DrawImage(l.background, op)

	for i := 0; i < len(l.lights); i++ {
		l.lights[i].Update()
		pos := l.lights[i].GetPosition()
		r := l.lights[i].GetRadius()

		op := &ebiten.DrawImageOptions{}
		op.ColorM.Scale(1, 1, 1, l.lights[i].transparency)
		op.GeoM.Scale(r/32, r/32)
		op.GeoM.Translate(pos.X-camX, pos.Y-camY)
		op.GeoM.Scale(scale, scale)

		screen.DrawImage(l.light_text.SubImage(image.Rect(0, 0, 64, 64)).(*ebiten.Image), op)
	}
}

func (l *Level) RenderGroundMask(screen *ebiten.Image, camX, camY, scale float64) {
	op := &ebiten.DrawImageOptions{}

	// Set the good place of the sprite
	op.GeoM.Translate(-camX, -camY)

	// Apply the global scale
	op.GeoM.Scale(scale, scale)
	screen.DrawImage(l.ground, op)

}

func (l *Level) RenderParticles(screen *ebiten.Image, camX, camY, scale float64) {
	for i := 0; i < len(l.particles); i++ {
		l.particles[i].Draw(screen, camX, camY, scale)
	}
}

func (l *Level) RenderHealth(screen *ebiten.Image, camX, camY, scale, value float64) {

	// Deco
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(81-camX, -48-camY)
	op.GeoM.Scale(scale, scale)
	screen.DrawImage(l.healthContainer, op)

	// Bar
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(89-camX, -36-camY)
	op.GeoM.Scale(scale, scale)
	life := int(240 * value)

	screen.DrawImage(l.healthBar.SubImage(image.Rect(0, 0, life, 8)).(*ebiten.Image), op)

}

func (l *Level) RenderTitle(screen *ebiten.Image, camX, camY, scale float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(149-camX, 8-camY)
	op.GeoM.Scale(scale, scale)
	screen.DrawImage(l.title.SubImage(image.Rect(0, 0, 128, 64)).(*ebiten.Image), op)
}
