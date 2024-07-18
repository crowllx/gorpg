package components

import (
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
	Animations [][]*Animation
	Current    int
	Direction  int
	X, Y       float64
}

func (as *AnimatedSprite) AddAnimation(spritesheet *ebiten.Image, anim int, speed float64) {
	as.Animations[anim] = append(as.Animations[anim], &Animation{
		Frames:         spritesheet,
		Index:          0,
		Advance:        0,
		AnimationSpeed: speed,
		w:              64,
		h:              64,
		FrameCount:     spritesheet.Bounds().Dx() / 64,
	})
}

func (as *AnimatedSprite) CurrentAnimation() *Animation {
	a := as.Animations[as.Current][as.Direction]
	return a
}

// helper methods for working with sprites
func LoadSpriteSheet(path string, w, h int) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return img
}
