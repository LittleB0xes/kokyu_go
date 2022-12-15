package main

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) RenderEndState(screen *ebiten.Image) {
	g.level.RenderBackground(screen, g.camera.X, g.camera.Y, g.scale)

	g.level.RenderParticles(screen, g.camera.X, g.camera.Y, g.scale)

	g.hero.Draw(screen, g.camera.X, g.camera.Y, g.scale)
	g.level.RenderGroundMask(screen, g.camera.X, g.camera.Y, g.scale)

	// Press Space
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(149-g.camera.X, 40-g.camera.Y)
	op.GeoM.Scale(g.scale, g.scale)
	screen.DrawImage(g.level.title.SubImage(image.Rect(0, 80, 128, 112)).(*ebiten.Image), op)

	// Letterbox effect
	ebitenutil.DrawRect(screen, 0, 0, 1280, 192, color.Black)
	ebitenutil.DrawRect(screen, 0, 528, 1280, 192, color.Black)

}

func (g *Game) UpdateEndState() {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.state = Play
		g.Reset()
	}

}
