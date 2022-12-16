package main

import (
	"embed"
	"image/color"
	_ "image/png"
	"log"

	hero "github.com/LittleB0xes/kokyu/hero"
	monster "github.com/LittleB0xes/kokyu/monster"
	sound "github.com/LittleB0xes/kokyu/soundBox"

	my "github.com/LittleB0xes/kokyu/custom"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//go:embed assets/sounds/*.*

var f embed.FS

type GameState int64

const (
	Intro GameState = iota
	Play
	Win
	End
)

type Game struct {
	level         Level
	hero          hero.Hero
	colliders     []my.Rect
	camera        my.Vec2
	scale         float64
	monsters      []*monster.Ghost
	ghost_text    *ebiten.Image
	monster_timer int
	max_monster   int
	soundBox      *sound.SoundBox
	maxHealth     int
	state         GameState
	fader         uint8
	fadeType      Fading
}

type Fading int

const (
	In Fading = iota
	Out
	None
)

func NewGame() *Game {

	level := NewBaseLevel()
	sound := sound.NewSoundBox(f)

	img, _, err := ebitenutil.NewImageFromFile("./assets/sprites/Hero.png")

	if err != nil {
		log.Fatal("Hero - Error when openning file", err)
	}

	ghost_text, _, err := ebitenutil.NewImageFromFile("./assets/sprites/MonsterOne.png")

	if err != nil {
		log.Fatal("main.go Error when openning file", err)
	}
	monsters := make([]*monster.Ghost, 0)

	health := 20 * 60
	hero := hero.NewHero(10, 10, health, img)
	hero.SetSoundSystem(sound)

	colliders := make([]my.Rect, 3)
	colliders[0] = my.Rect{X: 0, Y: 101, W: 426, H: 11}
	colliders[1] = my.Rect{X: -16, Y: 0, W: 16, H: 112}
	colliders[2] = my.Rect{X: 426, Y: 0, W: 16, H: 112}

	return &Game{
		level:         level,
		hero:          hero,
		colliders:     colliders,
		camera:        my.Vec2{X: 0, Y: -64},
		scale:         3,
		monsters:      monsters,
		ghost_text:    ghost_text,
		monster_timer: 30,
		max_monster:   5,
		soundBox:      sound,
		maxHealth:     health,
		state:         Intro,
		fader:         255,
		fadeType:      In,
	}
}

func (g *Game) Update() error {
	switch g.state {
	case Intro:
		g.UpdateIntroState()
	case Play:
		g.UpdatePlayState()
	case Win:
		g.UpdateWinState()
	case End:
		g.UpdateEndState()

	}

	return nil
}

func (g *Game) Reset() {

	g.hero.Reset(0, 0, g.maxHealth)

	// All monster stuff
	g.monsters = make([]*monster.Ghost, 0)
	g.monster_timer = 30
	g.max_monster = 1

}

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.state {
	case Intro:
		g.RenderIntroState(screen)
	case Play:
		g.RenderPlayState(screen)
	case Win:
		g.RenderWinState(screen)
	case End:
		g.RenderEndState(screen)
	}

	// Fader
	switch g.fadeType {
	case In:
		if g.fader > 0 {
			g.fader -= 3
		}

	case Out:
		if g.fader < 255 {
			g.fader += 3

		}
	}
	color := color.RGBA{R: 0, G: 0, B: 0, A: g.fader}
	ebitenutil.DrawRect(screen, 0, 0, 1280, 720, color)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1278, 720
}

func main() {
	ebiten.SetWindowSize(1278, 720)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
