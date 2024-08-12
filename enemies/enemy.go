package enemies

import (
	"context"
	"fmt"
	sound "gorpg/audio"
	"gorpg/components"
	. "gorpg/components"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jakecoffman/cp/v2"
	"github.com/looplab/fsm"
	"github.com/yohamta/ganim8/v2"
	"golang.org/x/image/colornames"
)

type Enemy interface {
	Draw(*ebiten.Image, float64, float64)
    DrawHealthBar() (*ebiten.Image,float64,float64)
	AddToSpace(*cp.Space)
	Update()
	Death()
	Query(string) (int, int, error)
	GetStatus() *components.Status
}
type BaseEnemy struct {
	body         *cp.Body
	shape        *cp.Shape
	hurtboxes    []*HurtBox
	Sprite       *AnimatedSprite
	Status       *Status
	aggroRadius  *Detection
	stateMachine *fsm.FSM
	smPipe       chan int
	sfx          *sound.AudioEmitter
}

func _onHealthChanged() {
	fmt.Println("enemy: what happened??")
}
func (e *BaseEnemy) AddToSpace(space *cp.Space) {
	space.AddBody(e.body)
	space.AddShape(e.shape)
}
func (e *BaseEnemy) DrawHealthBar() (*ebiten.Image, float64,float64) {
	w := e.Sprite.CurrentAnim.W()
	h := 5
	hp, maxHp, _ := e.Status.Query("health")

	x := e.body.Position().X - float64(w/2)
	y := e.body.Position().Y - float64(w/2)
	bar := ebiten.NewImage(w, h)
	bar.Fill(colornames.Darkgray)
    percentHp := float64(hp)/float64(maxHp)
	fg := ebiten.NewImage(w,h)
    fg.Fill(colornames.Darkred)
    opts := &ebiten.DrawImageOptions{}
    opts.GeoM.Scale(percentHp, 1)
	bar.DrawImage(fg, opts)
	return bar, x, y
}

func (e *BaseEnemy) Draw(screen *ebiten.Image, camX, camY float64) {
	pos := e.body.Position()
	pos.X += camX
	pos.Y += camY
	opts := ganim8.DrawOpts(pos.X-32, pos.Y-32)
	e.Sprite.CurrentAnim.Draw(screen, opts)
}
func (e *BaseEnemy) Death() {
	if space := e.shape.Space(); space != nil {
        for _, hb := range e.hurtboxes {
            space.RemoveShape(hb.Shape())
        }
        space.RemoveShape(e.aggroRadius.Shape)
        space.RemoveShape(e.shape)
		space.RemoveBody(e.body)
	}
	e = nil

}
func (e *BaseEnemy) GetStatus() *components.Status {
	return e.Status
}

// TODO get rid of these
func (e *BaseEnemy) Query(q string) (int, int, error) {
	return e.Status.Query(q)
}

func (e *BaseEnemy) aggro(pqi *cp.PointQueryInfo) {
	var velocity cp.Vector
	var state string
	if pqi.Shape != nil && e.stateMachine.Current() != "attack" {
		if pqi.Distance < 6 {
			velocity = cp.Vector{0, 0}
			state = "attack"
		} else {
			//multiply by speed
			velocity = pqi.Point.Sub(e.body.Position()).Normalize().Mult(1)
			state = "chase"
		}
	} else {
		velocity = cp.Vector{0, 0}
		state = "idle"
	}
	if e.stateMachine.Current() != state {
		e.stateMachine.Event(context.Background(), state)
	}

	dx, dy := components.Move(e.shape, velocity.X, velocity.Y)
	e.body.SetVelocity(dx, dy)
}

// TODO: what is needed to update an 'enemy'?
func (e *BaseEnemy) Update() {
	e.aggroRadius.Update()
	e.Sprite.CurrentAnim.Update()
	if e.stateMachine.Current() == "attack" && e.Sprite.CurrentAnim.Status() == ganim8.Paused {
		e.stateMachine.Event(context.Background(), "attack-end")

	}
	// status queries
	hp, _, _ := e.Status.Query("health")
	if hp <= 0 {
		e.Death()
	}
}
