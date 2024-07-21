package enemies

import (
	. "gorpg/components"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
)

type Enemy interface {
	Draw(*ebiten.Image)
	AddToSpace(*resolv.Space)
	Update()
	Death()
	Query(string) (int, error)
}
type BaseEnemy struct {
	*resolv.Object
	Sprite *AnimatedSprite
	Status *Status
}

func New() *BaseEnemy {
	e := &BaseEnemy{}
	e.Object = resolv.NewObject(250, 250, 64, 64, "solid", "hit")
	e.Status = NewStatus(10, 0, e.Death)
	e.Data = e
	return e
}

func (e *BaseEnemy) AddToSpace(space *resolv.Space) {
	space.Add(e.Object)
}

func (e *BaseEnemy) Draw(screen *ebiten.Image) {
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(e.Position.X, e.Position.Y)
	screen.DrawImage(e.Sprite.CurrentImg.Draw(), &opts)
}
func (e *BaseEnemy) Death() {
	if space := e.Object.Space; space != nil {
		space.Remove(e.Object)
	}
	e = nil
}
func (e *BaseEnemy) Query(q string) (int, error) {
	return e.Status.Query(q)
}

// TODO: what is needed to update an 'enemy'?
func (e *BaseEnemy) Update() {
	e.Check(0, 0)
}
