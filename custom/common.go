package custom

import (
	"math"
)

type Vec2 struct {
	X float64
	Y float64
}

func (v *Vec2) Add(vecB Vec2) {
	v.X += vecB.X
	v.Y += vecB.Y
}

type Rect struct {
	X float64
	Y float64
	W float64
	H float64
}

func (box *Rect) Move(v Vec2) {
	box.X += v.X
	box.Y += v.Y
}

func (box1 Rect) IntersectRect(box2 Rect) bool {
	r1_right := box1.X + box1.W
	r1_left := box1.X
	r1_top := box1.Y
	r1_bottom := box1.Y + box1.H

	r2_right := box2.X + box2.W
	r2_left := box2.X
	r2_top := box2.Y
	r2_bottom := box2.Y + box2.H

	collided := r1_left < r2_right && r1_right > r2_left && r1_top < r2_bottom && r1_bottom > r2_top
	return collided
}

// return float64 value of the sign of a number
func SignF(value float64) float64 {
	if !math.Signbit(value) {
		return -1
	} else {
		return 1
	}
}
