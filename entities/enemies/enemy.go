package enemies

import (
	. "gorpg/components"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jakecoffman/cp/v2"
)

type Enemy interface {
	Draw(*ebiten.Image)
	AddToSpace(*cp.Space)
	Update()
	Death()
	Query(string) (int, error)
	Modify(string, int)
}
type BaseEnemy struct {
	body      *cp.Body
	shape     *cp.Shape
	hurtboxes *[]HurtBox
	Sprite    *AnimatedSprite
	Status    *Status
}

func (e *BaseEnemy) AddToSpace(space *cp.Space) {
	space.AddBody(e.body)
	space.AddShape(e.shape)
}

func (e *BaseEnemy) Draw(screen *ebiten.Image) {
	opts := ebiten.DrawImageOptions{}
	pos := e.body.Position()
	opts.GeoM.Translate(pos.X-32, pos.Y-32)
	screen.DrawImage(e.Sprite.CurrentImg.Draw(), &opts)
}
func (e *BaseEnemy) Death() {
	if space := e.shape.Space(); space != nil {
        e.body.EachShape(func(s *cp.Shape) {
            space.RemoveShape(s)
        })
		space.RemoveBody(e.body)
	}
	e = nil
}
func (e *BaseEnemy) Query(q string) (int, error) {
	return e.Status.Query(q)
}

func (e *BaseEnemy) Modify(q string, v int) {
	e.Status.Modify(q, v)
}

// TODO: what is needed to update an 'enemy'?
func (e *BaseEnemy) Update() {

}
