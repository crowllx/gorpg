package enemies

import (
	. "gorpg/components"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
	"golang.org/x/image/colornames"
)

type Enemy struct {
	sprite   *Sprite2D
	collider *Collider
}

func NewEnemy(space *resolv.Space) *Enemy {
	img := ebiten.NewImage(64, 64)
	img.Fill(colornames.Black)
	e := &Enemy{
		sprite: &Sprite2D{
			Sprite: img,
			X:      250,
			Y:      250,
		},
		collider: NewCollider(250, 250, 64, 64, "hit", "solid"),
	}
	space.Add((*resolv.Object)(e.collider))
	return e
}
func (e *Enemy) Collider() *resolv.Object {
	return (*resolv.Object)(e.collider)
}
func (e *Enemy) Sprite() *Sprite2D {
	return e.sprite
}

func (e *Enemy) Update() {
	e.collider.Check(0, 0)
}
