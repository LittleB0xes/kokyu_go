package sprite

import (
	"image"
	"image/color"
	_ "image/png"

	my "github.com/LittleB0xes/kokyu/custom"
	"github.com/hajimehoshi/ebiten/v2"
)

type AnimationData struct {
	X     int
	Y     int
	W     int
	H     int
	Frame int
	Speed int
}

type AnimatedSprite struct {
	texture        *ebiten.Image
	x              float64
	y              float64
	w              float64
	h              float64
	sx             int
	sy             int
	sw             int
	sh             int
	frame          int
	animationSpeed int
	elapsedTime    int
	currentFrame   int
	camX           float64
	camY           float64
	isPlaying      bool
	isFlipX        bool
	color          color.Color
	transparency   float64
}

func NewAnimatedSprite(xo, yo float64, img *ebiten.Image, animationData AnimationData) AnimatedSprite {

	return AnimatedSprite{
		texture:        img,
		x:              xo,
		y:              yo,
		w:              float64(animationData.W),
		h:              float64(animationData.H),
		sx:             animationData.X,
		sy:             animationData.Y,
		sw:             animationData.W,
		sh:             animationData.H,
		frame:          animationData.Frame,
		animationSpeed: animationData.Speed,
		elapsedTime:    0,
		currentFrame:   0,
		camX:           0,
		camY:           0,
		isPlaying:      true,
		isFlipX:        false,
		color:          color.White,
		transparency:   1.0,
	}

}

func (s *AnimatedSprite) Draw(screen *ebiten.Image, camX, camY float64, scale float64) {
	op := &ebiten.DrawImageOptions{}
	op.ColorM.Scale(1.0, 1.0, 1.0, s.transparency)
	// Apply flipping value before any scalin or other translation
	if s.isFlipX {
		op.GeoM.Translate(-s.w, 0)
		op.GeoM.Scale(-1, 1)
	}

	// Set the good place of the sprite
	op.GeoM.Translate(s.x-camX, s.y-camY)

	// Apply the global scale
	op.GeoM.Scale(scale, scale)

	source_x := s.sx + s.currentFrame*s.sw
	screen.DrawImage(s.texture.SubImage(image.Rect(source_x, s.sy, source_x+s.sw, s.sy+s.sh)).(*ebiten.Image), op)
	if s.isPlaying {
		s.elapsedTime++
	}

	if s.isPlaying && s.elapsedTime >= s.animationSpeed {
		s.currentFrame++
		s.elapsedTime = 0
	}
	if s.currentFrame >= s.frame {
		s.currentFrame = 0
	}
}

func (s *AnimatedSprite) SetAnimation(data AnimationData) {
	s.sx = data.X
	s.sy = data.Y
	s.sw = data.W
	s.sh = data.H
	s.frame = data.Frame
	s.animationSpeed = data.Speed
	s.elapsedTime = 0
	s.currentFrame = 0
}

func (s *AnimatedSprite) GetCurrentFrame() int {
	return s.currentFrame
}

func (s *AnimatedSprite) SetCurrentFrame(f int) {
	s.currentFrame = f
}

func (s *AnimatedSprite) FlipX(value bool) {
	s.isFlipX = value
}

func (s *AnimatedSprite) IsFlipX() bool {
	return s.isFlipX
}

func (s *AnimatedSprite) SetFlipX(value bool) {
	s.isFlipX = value
}

func (s *AnimatedSprite) Pause() {
	s.isPlaying = false
}

func (s *AnimatedSprite) Rewind() {
	s.currentFrame = 0
	s.elapsedTime = 0
}

func (s *AnimatedSprite) Play() {
	s.isPlaying = true
}

func (s *AnimatedSprite) IsAnimationEnded() bool {
	return s.currentFrame == s.frame-1
}

func (s *AnimatedSprite) SetSpeed(speed int) {
	s.animationSpeed = speed
}

func (s *AnimatedSprite) SetPosition(pos my.Vec2) {
	s.x = pos.X
	s.y = pos.Y
}

func (s *AnimatedSprite) SetTransparency(value float64) {
	s.transparency = value
}
