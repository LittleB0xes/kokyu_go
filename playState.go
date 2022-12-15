package main

import (
	"image/color"
	"math/rand"

	monster "github.com/LittleB0xes/kokyu/monster"
	sound "github.com/LittleB0xes/kokyu/soundBox"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func (g *Game) RenderPlayState(screen *ebiten.Image) {
	g.level.RenderBackground(screen, g.camera.X, g.camera.Y, g.scale)

	for i := 0; i < len(g.monsters); i++ {
		g.monsters[i].Draw(screen, g.camera.X, g.camera.Y, g.scale)
	}

	g.hero.Draw(screen, g.camera.X, g.camera.Y, g.scale)

	//g.level.RenderParticles(screen, g.camera.X, g.camera.Y, g.scale)

	g.level.RenderGroundMask(screen, g.camera.X, g.camera.Y, g.scale)

	// Letterbox effect
	ebitenutil.DrawRect(screen, 0, 0, 1280, 192, color.Black)
	ebitenutil.DrawRect(screen, 0, 528, 1280, 192, color.Black)

	// Health bar drawing

	g.level.RenderHealth(screen, g.camera.X, g.camera.Y, g.scale, float64(g.hero.GetHealth())/float64(g.maxHealth))

}

func (g *Game) UpdatePlayState() {
	g.soundBox.SBPlay(sound.Beat)
	g.soundBox.SBPlay(sound.Ambiance)

	g.monster_timer -= 1
	if g.max_monster > 0 && g.monster_timer == 0 {
		g.max_monster -= 1
		// create a monster
		x := 50 + 330*rand.Float64()
		g.monsters = append(g.monsters, monster.NewGhost(x, 52, g.ghost_text))

		// Reset the timer
		g.monster_timer = 50 + rand.Intn(30)
	}

	g.hero.Update(&g.colliders, g.monsters)

	for i := 0; i < len(g.monsters); i++ {
		g.monsters[i].Update(g.hero.GetPosition())
	}

	// Remove dead monsters
	alive := make([]*monster.Ghost, 0)
	for _, m := range g.monsters {
		if m.IsActive() {
			alive = append(alive, m)
		}
	}
	// copy th new array
	g.monsters = alive

	// Check victory

	if len(g.monsters) == 0 && g.max_monster == 0 {
		g.fadeType = Out
		if g.fadeType == Out && g.fader == 255 {
			g.fadeType = In
			g.state = Win
		}
	} else if g.hero.IsDead() {
		g.fadeType = Out
		if g.fadeType == Out && g.fader == 255 {
			g.fadeType = In
			g.state = End
		}
	}

}
