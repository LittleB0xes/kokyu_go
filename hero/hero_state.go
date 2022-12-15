package hero

import (
	sound "github.com/LittleB0xes/kokyu/soundBox"
)

func (e *Hero) StateManager() {
	current_state := e.state

	if e.hited {
		e.state = Hit
	}

	if e.health <= 0 && e.onTheFloor && e.state != Dead {
		e.state = Dying
	}

	switch e.state {
	case Idle:
		if e.dir != 0 {
			e.state = Walk
		}

		if !e.onTheFloor {
			e.state = Jump
		}

		switch e.attack {
		case AttackDouble:
			e.state = Double
		case AttackHeavy:
			e.state = Heavy
		default:
			e.attack = NoAttack
		}

	case Walk:
		if e.dir == 0 {
			e.state = Idle
		}
		if !e.onTheFloor {
			e.state = Jump
		}
		switch e.attack {
		case AttackDash:
			e.state = Dash
			e.dirDash = e.dir
			e.dashTimer = 10
		case AttackHeavy:
			e.state = Heavy
		default:
			e.attack = NoAttack
		}
	case Jump:
		if e.onTheFloor {
			e.state = Idle
		}
		if e.attack == AttackDouble {
			e.state = Double
		} else if e.attack == AttackDash {
			e.state = AirDash
			e.dirDash = e.dir
			e.dashTimer = 10
		} else {
			e.attack = NoAttack
		}
	case Double:
		if e.sprite.GetCurrentFrame() == 3 {
			e.sound.SBPlay(sound.Sword1)
		}
		if e.sprite.GetCurrentFrame() == 10 {
			e.sound.SBPlay(sound.Sword2)
		}
		if e.sprite.IsAnimationEnded() {
			e.state = Idle
			e.attack = NoAttack
		}
	case Heavy:
		if e.sprite.GetCurrentFrame() == 0 {
			e.sound.SBPlay(sound.Heavy)
		}
		e.velocity.X = 0
		e.dir = 0
		if e.sprite.IsAnimationEnded() {
			e.state = Idle
			e.attack = NoAttack
		}
	case Dash:
		e.dir = e.dirDash
		e.dashTimer -= 1
		e.velocity.X = 8 * e.dir
		if e.dashTimer < 0 {
			e.velocity.X *= 0.5
			e.state = Idle
			e.attack = NoAttack
		}

	case AirDash:
		e.dir = e.dirDash
		e.dashTimer -= 1
		e.velocity.X = 8 * e.dir
		if e.dashTimer < 0 {

			e.state = Idle
			e.attack = NoAttack
		}
	case Hit:
		if e.sprite.GetCurrentFrame() == 0 {
			e.sound.SBPlay(sound.Huh1)
		}
		if e.sprite.GetCurrentFrame() == 0 {
			e.velocity.X = 2 * e.dirDash
		} else if e.sprite.IsAnimationEnded() {
			e.hited = false
			e.state = Idle

		}
	case Dying:
		if e.sprite.GetCurrentFrame() == 0 {
			e.sound.SBPlay(sound.Death)
		} else if e.sprite.IsAnimationEnded() {
			e.state = Dead
		}

	case Dead:

	}

	if e.dir > 0 {
		e.sprite.SetFlipX(false)
	} else if e.dir < 0 {
		e.sprite.SetFlipX(true)
	}

	if e.state != current_state {
		e.sprite.SetAnimation(e.animations[e.state])
		e.sprite.Play()
	}
}
