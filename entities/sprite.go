package entities

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Sprite interface {
	Draw() *ebiten.Image
}

// static 2d sprite
type Sprite2D struct {
	Sprite *ebiten.Image
	X, Y   float64
}

func (s *Sprite2D) Draw() *ebiten.Image {
	return s.Sprite
}

// TODO: refactor velocity into separate component
func (s *Sprite2D) UpdateVelocity(x, y float64) {
	s.X += x
	s.Y += y
}
func (s *AnimatedSprite) UpdateVelocity(x, y float64) {
	s.X += x
	s.Y += y
}

// animated 2d sprite
type AnimatedSprite struct {
	animations map[string]*Animation
	current    string
	X, Y       float64
}

func (as *AnimatedSprite) AddAnimation(spritesheet []*ebiten.Image, key string, speed float64) {
	as.animations[key] = &Animation{
		Frames:            spritesheet,
		CurrentFrameIndex: 0,
		Advance:           0,
		AnimationSpeed:    speed,
	}
}

func (as *AnimatedSprite) CurrentFrame() *ebiten.Image {
	a := as.animations[as.current]
	return a.Draw()
}

func (as *AnimatedSprite) CurrentAnimation() *Animation {
	anim := as.animations[as.current]
	return anim
}

// helper methods for working with sprites
func LoadSpriteSheet(path string, w, h int) []*ebiten.Image {
	frames := []*ebiten.Image{}
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		log.Fatal(err)
	}
	n := img.Bounds().Dx() / w
	for i := range n {
		dimensions := image.Rect(i*w, 0, (i+1)*w, h)
		frame := img.SubImage(dimensions).(*ebiten.Image)
		frames = append(frames, frame)
	}
	return frames
}
