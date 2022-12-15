package hero

import (
	my "github.com/LittleB0xes/kokyu/custom"
)

type Attack uint8

const (
	NoAttack = iota
	AttackHeavy
	AttackHeavyRepeat
	AttackDouble
	AttackDash
	AttackAirDash
)

func GetDamage(attack Attack) int {
	switch attack {
	case AttackAirDash, AttackDouble, AttackDash:
		return 1
	case AttackHeavy, AttackHeavyRepeat:
		return 2
	default:
		return 0
	}
}

func GetHitBox(attack Attack, frame int, flip bool) (my.Rect, bool) {
	hbox := my.Rect{X: 0, Y: 0, W: 0, H: 0}
	isBox := false
	switch attack {
	case AttackDouble:
		switch frame {
		case 6, 7, 8, 9:
			hbox = my.Rect{X: 41.0, Y: 31.0, W: 16.0, H: 16.0}
			isBox = true
		case 13, 14, 15:
			hbox = my.Rect{X: 9.0, Y: 31.0, W: 16.0, H: 16.0}
			isBox = true
		}

	case AttackHeavy:
		switch frame {
		case 12, 13, 14:
			hbox = my.Rect{X: 34.0, Y: 4.0, W: 27.0, H: 44.0}
			isBox = true

		}
	case AttackAirDash, AttackDash:
		hbox = my.Rect{X: 36.0, Y: 32.0, W: 11.0, H: 14.0}
		isBox = true
	}

	if isBox && flip {
		hbox = my.Rect{X: 64.0 - hbox.X - hbox.W, Y: hbox.Y, W: hbox.W, H: hbox.H}
	}

	return hbox, isBox

}
