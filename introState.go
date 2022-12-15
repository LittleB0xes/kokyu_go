package main

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) RenderIntroState(screen *ebiten.Image) {
	g.level.RenderBackground(screen, g.camera.X, g.camera.Y, g.scale)

	//g.level.RenderParticles(screen, g.camera.X, g.camera.Y, g.scale)

	g.level.RenderGroundMask(screen, g.camera.X, g.camera.Y, g.scale)

	g.level.RenderTitle(screen, g.camera.X, g.camera.Y, g.scale)

	// Press Space
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(149-g.camera.X, 80-g.camera.Y)
	op.GeoM.Scale(g.scale, g.scale)
	screen.DrawImage(g.level.title.SubImage(image.Rect(0, 64, 128, 80)).(*ebiten.Image), op)

	// Letterbox effect
	ebitenutil.DrawRect(screen, 0, 0, 1280, 192, color.Black)
	ebitenutil.DrawRect(screen, 0, 528, 1280, 192, color.Black)

}

func (g *Game) UpdateIntroState() {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.state = Play
	}

}
