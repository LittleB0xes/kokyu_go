package monster

import (
	"math"

	my "github.com/LittleB0xes/kokyu/custom"
)

// Use Interface for *only* this is a bit overkill
// but it is in the perspective of extending the possible behavior


type Behaviour interface {
	Update() Behaviour
	GetPosition() my.Vec2
}

type StandBy struct {
	position my.Vec2
}

func (b StandBy) Update() Behaviour {
	return StandBy{position: b.position}
}

func (b StandBy) GetPosition() my.Vec2 {
	return b.position
}

type UpDown struct {
	position my.Vec2
	yo       float64
	speed    float64
	dt       float64
}

func (b UpDown) Update() Behaviour {
	y := b.yo + 15.0*math.Sin(b.dt)
	dt := b.dt + b.speed
	speed := b.speed
	yo := b.yo
	position := my.Vec2{X: b.position.X, Y: y}

	return UpDown{
		position,
		yo,
		speed,
		dt,
	}
}

func (b UpDown) GetPosition() my.Vec2 {
	return b.position
}
