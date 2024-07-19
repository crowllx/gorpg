package components

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Drawable interface {
	Draw(screen *ebiten.Image)
}

// static 2d sprite
type Sprite2D struct {
	Sprite *ebiten.Image
	X, Y   float64
}

func (s *Sprite2D) Draw(screen *ebiten.Image) *ebiten.Image {
	return s.Sprite
}

// animated 2d sprite
type AnimatedSprite struct {
	animations [][]*Animation
	CurrentImg *Animation
	Current    int
	// Cardinal   int
	// X, Y       float64
}

func NewAS(anims [][]*Animation) *AnimatedSprite {
	return &AnimatedSprite{
		animations: anims,
	}
}
func (as *AnimatedSprite) ChangeAnimation(anim int, cardinal int) {
	as.Current = anim
	as.CurrentImg = as.animations[anim][cardinal]
}

func (as *AnimatedSprite) AddAnimation(spritesheet *ebiten.Image, anim int, speed float64, w, h int, loop bool) {
	as.animations[anim] = append(as.animations[anim], &Animation{
		Frames:         spritesheet,
		Index:          0,
		Advance:        0,
		AnimationSpeed: speed,
		w:              w,
		h:              h,
		FrameCount:     spritesheet.Bounds().Dx() / w,
		loop:           loop,
		finished:       false,
	})
}

func (as *AnimatedSprite) CurrentAnimation(cardinal int) *Animation {
	a := as.animations[as.Current][cardinal]
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
