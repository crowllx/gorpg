package components

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/yohamta/ganim8/v2"
)

type Drawable interface {
	Draw(screen *ebiten.Image) *ebiten.Image
}

// static 2d sprite
type Sprite2D struct {
	Sprite *ebiten.Image
	X, Y   float64
}

func (s *Sprite2D) Draw(screen *ebiten.Image) *ebiten.Image {
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(s.X, s.Y)
	screen.DrawImage(s.Sprite, &opts)
	return s.Sprite
}

// animated 2d sprite
type AnimatedSprite struct {
	animations  map[string]*ganim8.Animation
	CurrentAnim *ganim8.Animation
	Current     string
	// Cardinal   int
	// X, Y       float64
}

func NewAS(key string, anim *ganim8.Animation) *AnimatedSprite {
	animMap := make(map[string]*ganim8.Animation)
	animMap[key] = anim
	return &AnimatedSprite{
		animations:  animMap,
		CurrentAnim: animMap[key],
		Current:     key,
	}
}

func (as *AnimatedSprite) AddAnimation(key string, anim *ganim8.Animation) {
	// probably needs handling of more cases. ie check if animation exists already
	as.animations[key] = anim
}

func (as *AnimatedSprite) ChangeAnimation(key string) {
	as.CurrentAnim.PauseAtStart()
	as.Current = key
	as.CurrentAnim = as.animations[key]
	as.CurrentAnim.Resume()
}

// func (as *AnimatedSprite) AddAnimation(spritesheet *ebiten.Image, anim int, speed float64, w, h int, loop bool) {
// 	as.animations[anim] = append(as.animations[anim], &Animation{
// 		Frames:         spritesheet,
// 		Index:          0,
// 		Advance:        0,
// 		AnimationSpeed: speed,
// 		w:              w,
// 		h:              h,
// 		FrameCount:     spritesheet.Bounds().Dx() / w,
// 		loop:           loop,
// 		finished:       false,
// 	})
// }

// func (as *AnimatedSprite) CurrentAnimation(cardinal int) *Animation {
// 	a := as.animations[as.Current][cardinal]
// 	return a
// }

// helper methods for working with sprites
func LoadSpriteSheet(path string, w, h int) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return img
}
