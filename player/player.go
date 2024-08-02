package player

import (
	"context"
	"fmt"
	. "gorpg/components"
	"gorpg/utils"
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
	body := cp.NewKinematicBody()
	body.SetPosition(cp.Vector{X: 100, Y: 100})
	body.SetVelocity(0, 0)
	body.UserData = player
	player.Body = body
	player.NewSM()
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
		DETECTION_LAYER

	filter := cp.NewShapeFilter(0, PLAYER_LAYER, mask)
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
		p.attackEnd()
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

	// new velocity based on user input
	dx := float64(dir.X) * p.speed
	dy := float64(dir.Y) * p.speed

	// check collisions given new velocity
	p.shape.Space().ShapeQuery(p.shape, func(s *cp.Shape, cps *cp.ContactPointSet) {
		switch s.UserData.(type) {
		case utils.Collidable:
			fmt.Printf("%T\n", s.UserData)
			normal := cps.Normal
			colX, colY := TerrainCheck(cp.Vector{dx, dy}, normal, p.Body)
			if colX {
				dx = 0
			}
			if colY {
				dy = 0
			}
		default:
		}
	})

	// finally update velocity considering user input and collisions
	p.Body.SetVelocity(dx, dy)

	// what else needs to be done here? can i abstract this out to different module?
	p.sprite.CurrentImg.Update()
	if p.stateMachine.Current() == "attack" && p.sprite.CurrentImg.Done() {
		p.stateMachine.Event(context.Background(), "attack-end")
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	opts := ebiten.DrawImageOptions{}
	pos := p.Body.Position()
	opts.GeoM.Translate(pos.X-32, pos.Y-32)
	screen.DrawImage(p.sprite.CurrentImg.Draw(), &opts)
}
