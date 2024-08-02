package enemies

import (
	"context"
	"fmt"
	. "gorpg/components"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/jakecoffman/cp/v2"
	"github.com/looplab/fsm"
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
	body         *cp.Body
	shape        *cp.Shape
	hurtboxes    *[]HurtBox
	Sprite       *AnimatedSprite
	Status       *Status
	aggroRadius  *Detection
	stateMachine *fsm.FSM
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

	ebitenutil.DebugPrint(screen, fmt.Sprintf(`
		enemy state: %s
		`, e.stateMachine.Current()))
}
func (e *BaseEnemy) Death() {
	if space := e.shape.Space(); space != nil {
		e.body.EachShape(func(s *cp.Shape) {
			fmt.Printf("%T\n", s.UserData)
			space.RemoveShape(s)
		})
		space.RemoveBody(e.body)
		space.EachShape(func(s *cp.Shape) {
			fmt.Printf("space shape %T\n", s.UserData)
		})

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
	var velocity cp.Vector
	// enemy detection & chase
	if e.aggroRadius.Enabled {
		info := e.shape.Space().PointQueryNearest(
			e.body.Position(),
			e.aggroRadius.Radius,
			e.aggroRadius.Shape.Filter,
		)

		//TODO write this ina way that doesn't require nesting
		var state string
		if info.Shape != nil && e.stateMachine.Current() != "attack" {
			if info.Distance < 8 {
				velocity = cp.Vector{0, 0}
				state = "attack"
			} else {
				velocity = info.Point.Sub(e.body.Position()).Normalize().Mult(1)
				state = "chase"
			}
		} else {
			velocity = cp.Vector{0, 0}
			state = "idle"
		}
		if e.stateMachine.Current() != state {
			e.stateMachine.Event(context.Background(), state)
		}
		e.body.SetVelocityVector(velocity)
	}
	e.Sprite.CurrentImg.Update()
	if e.stateMachine.Current() == "attack" && e.Sprite.CurrentImg.Done() {
		e.stateMachine.Event(context.Background(), "attack-end")

	}
	// status queries
	hp, _ := e.Status.Query("health")
	if hp <= 0 {
		e.Death()
	}
}
