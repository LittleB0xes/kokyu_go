package main

import (
	"math"
	"math/rand"

	my "github.com/LittleB0xes/kokyu/custom"
)

type Light struct {
	position     my.Vec2
	dt           float64
	transparency float64
	radius       float64
}

func NewLight(x, y, radius float64) Light {
	return Light{
		position:     my.Vec2{X: x, Y: y},
		dt:           1.5 * rand.Float64(),
		transparency: 0,
		radius:       radius,
	}
}
func (e *Light) Update() {
	e.dt += 0.01
	e.transparency = 0.1 + 0.05*math.Sin(0.5*e.dt)
}

func (e *Light) GetRadius() float64 {
	return e.radius + 2*math.Sin(e.dt)
}

func (e *Light) GetPosition() my.Vec2 {
	return my.Vec2{
		X: e.position.X - e.GetRadius() + 2*math.Cos(e.dt),
		Y: e.position.Y - e.GetRadius() + 2*math.Sin(e.dt),
	}
}
