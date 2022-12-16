package monster

import (
	"image/color"

	my "github.com/LittleB0xes/kokyu/custom"
	sprite "github.com/LittleB0xes/kokyu/sprite"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type State uint8

const (
	Idle State = iota
	Hit
	Dead
	Birth
)

type Ghost struct {
	position     my.Vec2
	sprite       sprite.AnimatedSprite
	animations   map[State]sprite.AnimationData
	collisionBox my.Rect
	state        State
	dir          float64
	behaviour    Behaviour
	hited        bool
	hitable      bool
	active       bool
	health       int
}

func NewGhost(x, y float64, texture *ebiten.Image) *Ghost {
	anim := map[State]sprite.AnimationData{
		Idle:  {X: 0, Y: 0, H: 64, W: 64, Frame: 5, Speed: 8},
		Hit:   {X: 0, Y: 64, H: 64, W: 64, Frame: 10, Speed: 4},
		Dead:  {X: 0, Y: 128, H: 64, W: 64, Frame: 10, Speed: 8},
		Birth: {X: 0, Y: 192, H: 64, W: 64, Frame: 13, Speed: 8},
	}

	state := Birth

	return &Ghost{
		position:     my.Vec2{X: x, Y: y},
		collisionBox: my.Rect{X: 25.0, Y: 19.0, W: 15.0, H: 22.0},
		sprite:       sprite.NewAnimatedSprite(x, y, texture, anim[state]),
		animations:   anim,
		state:        state,
		dir:          0,
		behaviour:    StandBy{my.Vec2{X: x, Y: y}},
		hited:        false,
		hitable:      true,
		active:       true,
		health:       3,
	}

}

func (e *Ghost) Update(heroPosition my.Vec2) {
	if e.position.X-heroPosition.X > 0 {
		e.dir = -1
	} else if e.position.X-heroPosition.X < 0 {
		e.dir = 1
	}

	currentState := e.state
	switch e.state {
	case Birth:
		if e.sprite.IsAnimationEnded() {
			e.state = Idle
			e.behaviour = UpDown{position: e.position, yo: e.position.Y, speed: 0.01, dt: 0}
		}
	case Idle:
		e.state = Idle
	case Hit:
		e.hitable = false
		if e.sprite.IsAnimationEnded() {
			e.state = Idle
			e.hited = false
			e.hitable = true
			if e.health <= 0 {
				e.state = Dead
				e.hitable = false
				e.hited = false
			}
		}
	case Dead:
		if e.sprite.IsAnimationEnded() {
			e.active = false
			e.hitable = false
		}
	}

	if e.hited {
		e.state = Hit
	}

	if e.dir > 0 {
		e.sprite.FlipX(false)
	} else if e.dir < 0 {
		e.sprite.FlipX(true)
	}

	if e.state != currentState {
		e.sprite.SetAnimation(e.animations[e.state])
		e.sprite.Play()
	}

	e.behaviour = e.behaviour.Update()
	e.position = e.behaviour.GetPosition()
	e.sprite.SetPosition(e.position)

}
func (e *Ghost) GetCollisionBox(dx, dy float64) my.Rect {
	return my.Rect{
		X: e.collisionBox.X + e.position.X + dx,
		Y: e.collisionBox.Y + e.position.Y + dy,
		W: e.collisionBox.W,
		H: e.collisionBox.H,
	}
}

func (e *Ghost) Draw(screen *ebiten.Image, camX, camY, scale float64) {
	e.sprite.Draw(screen, camX, camY, scale)

}
func (e *Ghost) DebugCollisionBox(screen *ebiten.Image, camX, camY, scale float64) {
	rect := e.GetCollisionBox(0, 0)
	color := color.RGBA{255, 0, 0, 50}
	ebitenutil.DrawRect(screen, (rect.X-camX)*scale, (rect.Y-camY)*scale, rect.W*scale, rect.H*scale, color)
}
func (e *Ghost) Hit(value int) {
	e.hited = true
	e.hitable = false
	e.health -= value
}

func (e *Ghost) IsHitable() bool {
	return e.hitable
}

func (e *Ghost) IsActive() bool {
	return e.active
}
func (e *Ghost) IsDead() bool {
	return e.state == Dead
}

func (e *Ghost) GetPosition() my.Vec2 {
	return e.position
}
