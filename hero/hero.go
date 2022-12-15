package hero

import (
	"image/color"

	my "github.com/LittleB0xes/kokyu/custom"
	monster "github.com/LittleB0xes/kokyu/monster"
	sound "github.com/LittleB0xes/kokyu/soundBox"
	sprite "github.com/LittleB0xes/kokyu/sprite"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type HeroState int64

const (
	Walk HeroState = iota
	Jump
	Idle
	Heavy
	Double
	Dash
	AirDash
	RepeatHeavy
	Hit
	Dying
	Dead
)

type Hero struct {
	position  my.Vec2
	velocity  my.Vec2
	speed     float64
	dir       float64
	dirDash   float64
	dashTimer int

	collisionBox my.Rect

	state  HeroState
	attack Attack

	health int

	onTheFloor bool
	hited      bool

	sound *sound.SoundBox

	sprite     sprite.AnimatedSprite
	animations map[HeroState]sprite.AnimationData
}

func NewHero(x, y float64, health int, texture *ebiten.Image) Hero {

	animations := map[HeroState]sprite.AnimationData{
		Walk:        {X: 0, Y: 128, W: 64, H: 64, Frame: 8, Speed: 7},
		Idle:        {X: 0, Y: 192, W: 64, H: 64, Frame: 8, Speed: 7},
		Double:      {X: 0, Y: 384, H: 64, W: 64, Frame: 19, Speed: 2},
		Heavy:       {X: 0, Y: 448, H: 64, W: 64, Frame: 17, Speed: 4},
		RepeatHeavy: {X: 640, Y: 448, H: 64, W: 64, Frame: 7, Speed: 4},
		Jump:        {X: 0, Y: 256, H: 64, W: 64, Frame: 12, Speed: 4},
		AirDash:     {X: 0, Y: 0, W: 64, H: 64, Frame: 7, Speed: 4},
		Dash:        {X: 0, Y: 64, W: 64, H: 64, Frame: 7, Speed: 4},
		Hit:         {X: 128, Y: 512, H: 64, W: 64, Frame: 5, Speed: 3},
		Dying:       {X: 384, Y: 320, H: 64, W: 64, Frame: 13, Speed: 10},
		Dead:        {X: 1152, Y: 320, H: 64, W: 64, Frame: 1, Speed: 2},
	}
	state := Idle
	sprite := sprite.NewAnimatedSprite(x, y, texture, animations[state])

	return Hero{
		position:     my.Vec2{X: x, Y: y},
		velocity:     my.Vec2{X: 0, Y: 0},
		speed:        1.5,
		dir:          0,
		dirDash:      0,
		dashTimer:    10,
		collisionBox: my.Rect{X: 27.0, Y: 28.0, W: 10.0, H: 20.0},
		state:        state,
		attack:       0,
		onTheFloor:   false,
		hited:        false,
		sound:        &sound.SoundBox{},
		sprite:       sprite,
		animations:   animations,
		health:       health,
	}
}

func (e *Hero) SetSoundSystem(s *sound.SoundBox) {
	e.sound = s
}

func (e *Hero) Update(colliders *[]my.Rect, monsters []*monster.Ghost) {

	// Gravity
	e.velocity.Y += 0.5

	e.hited = false
	// This is only when not hited and not dead
	if e.state != Hit && e.state != Dead && e.state != Dying {

		hitBox, isBox := e.GetHitBox()
		for i, m := range monsters {
			if m.IsHitable() && isBox && hitBox.IntersectRect(m.GetCollisionBox(0, 0)) {
				(monsters)[i].Hit(GetDamage(e.attack))

			}

			if !e.hited && m.IsHitable() && m.GetCollisionBox(0, 0).IntersectRect(e.GetCollisionBox(e.velocity.X, e.velocity.Y)) {
				e.hited = true
				e.dirDash = my.SignF(m.GetPosition().X - e.position.X)
			}
		}

		if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
			e.dir = -1
		} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
			e.dir = 1
		} else {
			e.dir = 0
		}

		// Chack action
		if e.onTheFloor && inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			e.velocity.Y = -8
			e.onTheFloor = false
		}

		// Conditionnal attack
		if e.dir == 0 && inpututil.IsKeyJustPressed(ebiten.KeyC) {
			e.attack = AttackDouble
		} else if e.dir != 0 && inpututil.IsKeyJustPressed(ebiten.KeyC) {
			e.attack = AttackDash

		} else if e.onTheFloor && inpututil.IsKeyJustPressed(ebiten.KeyV) {
			e.attack = AttackHeavy
		}

	}

	if e.dir == 0 {
		e.velocity.X *= 0.9
	} else {
		e.velocity.X = e.speed * e.dir
	}

	// Life
	e.health -= 1
	if e.attack != NoAttack {
		e.health -= 1
	}
	if e.health < 0 {
		e.health = 0
	}

	e.StateManager()

	for i := 0; i < len(*colliders); i++ {
		collider := (*colliders)[i]
		if e.GetCollisionBox(e.velocity.X, 0).IntersectRect(collider) {
			e.velocity.X = 0
		}
		if e.GetCollisionBox(0, e.velocity.Y).IntersectRect(collider) {
			e.velocity.Y = 0
			e.onTheFloor = true
		}
	}

	e.position.Add(e.velocity)

	e.sprite.SetPosition(e.position)
}

func (e *Hero) GetHitBox() (my.Rect, bool) {
	hBox, isBox := GetHitBox(e.attack, e.sprite.GetCurrentFrame(), e.sprite.IsFlipX())
	hBox.Move(e.position)
	return hBox, isBox
}

func (e *Hero) DebugCollisionBox(screen *ebiten.Image, camX, camY, scale float64) {
	rect := e.GetCollisionBox(0, 0)
	color := color.RGBA{255, 0, 0, 50}
	if e.hited {
		color.G = 255
	}
	ebitenutil.DrawRect(screen, (rect.X-camX)*scale, (rect.Y-camY)*scale, rect.W*scale, rect.H*scale, color)
}

func (e *Hero) DebugHitBox(screen *ebiten.Image, camX, camY, scale float64) {
	rect, isBox := GetHitBox(e.attack, e.sprite.GetCurrentFrame(), e.sprite.IsFlipX())
	if isBox {

		rect.Move(e.position)
		color := color.RGBA{255, 255, 0, 50}
		ebitenutil.DrawRect(screen, (rect.X-camX)*scale, (rect.Y-camY)*scale, rect.W*scale, rect.H*scale, color)

	}

}

func (e *Hero) SetPosition(position my.Vec2) {
	e.position = position
}

func (e *Hero) GetPosition() my.Vec2 {
	return e.position
}

func (e *Hero) Draw(screen *ebiten.Image, camX, camY, scale float64) {
	e.sprite.Draw(screen, camX, camY, scale)
}

func (e *Hero) GetCollisionBox(dx, dy float64) my.Rect {
	return my.Rect{
		X: e.collisionBox.X + e.position.X + dx,
		Y: e.collisionBox.Y + e.position.Y + dy,
		W: e.collisionBox.W,
		H: e.collisionBox.H,
	}
}

func (e *Hero) GetHealth() int {
	return e.health
}

func (e *Hero) Reset(x, y float64, health int) {
	e.health = health
	e.position = my.Vec2{X: x, Y: y}
	e.state = Idle

	// Sprite Reset
	e.sprite.SetAnimation(e.animations[e.state])
	e.sprite.Play()
}

func (e *Hero) IsDead() bool {
	return e.state == Dead
}
