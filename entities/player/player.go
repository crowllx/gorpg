package player

import (
	"context"
	"fmt"
	. "gorpg/components"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jakecoffman/cp/v2"
	"github.com/looplab/fsm"
	input "github.com/quasilyte/ebitengine-input"
)

type direction int

// direction the player is facing
const (
	UP = iota
	RIGHT
	DOWN
	LEFT
)

type animationState int

// action the player is taking
const (
	ATTACK = iota
	HIT
	IDLE
	WALK
)

type Player struct {
	Body         *cp.Body
	shape        *cp.Shape
	sprite       *AnimatedSprite
	input        *input.Handler
	stateMachine *fsm.FSM
	speed        float64
	hurtboxes    []*HurtBox
	Status       *Status
	cardinal     int
	area         *BasicArea
}

func New() *Player {
	player := load()
	body := cp.NewBody(3000, cp.INFINITY)
	body.SetPosition(cp.Vector{X: 100, Y: 100})
	body.SetVelocity(0, 0)
	body.UserData = player
	player.Body = body
	player.NewStateMachine()
	player.speed = 1.5
	player.Body.SetAngle(UP)
	player.Status = NewStatus(25, 10, func() {
		fmt.Println("I'm immune to death")
	})
	return player
}
func (p *Player) Sprite() *AnimatedSprite {
	return p.sprite
}
func (p *Player) AddInputHandler(s *input.System) {
	p.input = s.NewHandler(0, GenerateKeyMap())
}

func (p *Player) AddSpace(space *cp.Space) {
	hb := NewHurtBox(16, space, p.Body, &cp.Vector{X: 0, Y: -16})
	p.hurtboxes = append(p.hurtboxes, hb)
	space.AddBody(p.Body)
	shape := space.AddShape(cp.NewCircle(p.Body, 16, cp.Vector{X: 0, Y: 0}))
	mask := PLAYER_LAYER | ENVIRONMENT_LAYER | HIT_LAYER |
		ENEMY_LAYER | DETECTION_LAYER

	filter := cp.NewShapeFilter(0, 1, mask)
	shape.SetFilter(filter)
	shape.SetCollisionType(PLAYER_TYPE)
	p.shape = shape
	p.shape.UserData = p
	space.EachShape(func(s *cp.Shape) {
		fmt.Printf("%v %v\n", s, s.BB())
	})
}
func (p *Player) enterState(e *fsm.Event) {
	fmt.Println(e.Event)
	for i := range 4 {
		p.sprite.CurrentAnimation(i).Reset()
	}
	switch e.Event {
	case "idle":
		p.sprite.ChangeAnimation(IDLE, p.cardinal)
	case "walk":
		p.sprite.ChangeAnimation(WALK, p.cardinal)
	case "attack":
		p.enterAttack()
	case "attack-end":
		p.sprite.ChangeAnimation(IDLE, p.cardinal)
	}
}

func (p *Player) Update() {
	dir := p.InputVec()
	angle := math.Atan2(float64(dir.X), float64(dir.Y*-1))
	zero := cp.Vector{X: 0, Y: 0}
	if dir != zero && (p.stateMachine.Can("walk") || p.stateMachine.Current() == "walk") {
		p.Body.SetAngle(angle)
		p.stateMachine.Event(context.Background(), "walk")
		if dir.Y < 0 {
			p.cardinal = UP
		} else if dir.Y > 0 {
			p.cardinal = DOWN
		} else {
			if dir.X < 0 {
				p.cardinal = LEFT
			} else {
				p.cardinal = RIGHT
			}
		}
		p.sprite.ChangeAnimation(WALK, p.cardinal)
	} else {
		p.stateMachine.Event(context.Background(), "idle")
	}
	if p.input.ActionIsJustPressed(ActionAttack) {
		p.stateMachine.Event(context.Background(), "attack")
	}
	if p.stateMachine.Is("attack") && p.sprite.Current == ATTACK {
		dir.X, dir.Y = 0, 0
	}
	dx := float64(dir.X) * p.speed
	dy := float64(dir.Y) * p.speed
    p.Body.EachArbiter(func( arb *cp.Arbiter) {
        _,b := arb.Shapes()
        if b.Body().GetType() == cp.BODY_STATIC {
            fmt.Println("found static body")
        }
    })
	p.Body.SetVelocity(dx, dy)

	p.sprite.CurrentImg.Update()
	if p.stateMachine.Current() == "attack" && p.sprite.CurrentImg.Done() {
		p.stateMachine.Event(context.Background(), "attack-end")
		p.attackEnd()
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	opts := ebiten.DrawImageOptions{}
	pos := p.Body.Position()
	opts.GeoM.Translate(pos.X-32, pos.Y-32)
	screen.DrawImage(p.sprite.CurrentImg.Draw(), &opts)
}
