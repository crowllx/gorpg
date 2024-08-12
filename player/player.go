package player

import (
	"context"
	"fmt"
	sound "gorpg/audio"
	"gorpg/components"
	. "gorpg/components"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jakecoffman/cp/v2"
	"github.com/looplab/fsm"
	input "github.com/quasilyte/ebitengine-input"
	"github.com/yohamta/ganim8/v2"
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
	direction    int
	area         *BasicArea
	sfxEmitter   *sound.AudioEmitter
}

func New(x, y float64) *Player {
	player := load()
	body := cp.NewKinematicBody()
	body.SetPosition(cp.Vector{x, y})
	body.SetVelocity(0, 0)
	body.UserData = player
	player.Body = body
	player.NewSM()
	player.speed = 1.5
	player.Body.SetAngle(UP)
	player.Status = NewStatus(25, 10, func() {
		fmt.Println("I'm immune to death")
	}, player._onHealthChanged)
	return player
}
func (p *Player) _onHealthChanged() {
	fmt.Println("what happened?")
}
func (p *Player) Sprite() *AnimatedSprite {
	return p.sprite
}
func (p *Player) AddInputHandler(s *input.System) {
	p.input = s.NewHandler(0, GenerateKeyMap())
}

func (p *Player) AddSpace(space *cp.Space) {
	hb := NewHurtBox(16, space, p.Body, &cp.Vector{X: 0, Y: -16},
		cp.NewShapeFilter(0, HIT_LAYER, ENEMY_LAYER))
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
func (p *Player) updateDirection(input cp.Vector) {
	switch input {
	// player is facing north
	case cp.Vector{0, -1}, cp.Vector{1, -1}, cp.Vector{-1, -1}:
		p.direction = UP
	// player is facing south
	case cp.Vector{0, 1}, cp.Vector{1, 1}, cp.Vector{-1, 1}:
		p.direction = DOWN
	default:
		if input.X == 1 {
			p.direction = RIGHT
		} else if input.X == -1 {
			p.direction = LEFT
		}
	}
}
func (p *Player) switchAnim(key string) {
	var fullKey string
	var dir string
	switch p.direction {
	case UP:
		dir = "back"
	case DOWN:
		dir = "front"
	case LEFT:
		dir = "left"
	case RIGHT:
		dir = "right"
	}
	fullKey = fmt.Sprintf("%s-%s", dir, key)
	p.sprite.ChangeAnimation(fullKey)
}
func (p *Player) enterState(e *fsm.Event) {
	fmt.Println(e.Event)
	select {
	case smPipe <- 1:
	default:
	}
	switch e.Event {
	case "idle":
		p.switchAnim("idle")
	case "walk":
		p.enterWalk()
	case "attack":
		p.enterAttack()
	case "attack-end":
		p.attackEnd()
	}
}

func (p *Player) Update() {
	dir := p.InputVec()
	p.updateDirection(dir)
	angle := math.Atan2(float64(dir.X), float64(dir.Y*-1))
	zero := cp.Vector{X: 0, Y: 0}
	if dir != zero && (p.stateMachine.Can("walk") || p.stateMachine.Current() == "walk") {
		p.Body.SetAngle(angle)
		p.stateMachine.Event(context.Background(), "walk")
		p.switchAnim("walk")
	} else {
		p.stateMachine.Event(context.Background(), "idle")
	}
	if p.input.ActionIsJustPressed(ActionAttack) {
		p.stateMachine.Event(context.Background(), "attack")
	}

	dx := float64(dir.X) * p.speed
	dy := float64(dir.Y) * p.speed

	// check collisions given new velocity
	dx, dy = components.Move(p.shape, dx, dy)
	p.Body.SetVelocity(dx, dy)

	// what else needs to be done here? can i abstract this out to different module?
	p.sprite.CurrentAnim.Update()
	if p.stateMachine.Current() == "attack" &&
		p.sprite.CurrentAnim.Status() == ganim8.Paused {
		p.stateMachine.Event(context.Background(), "attack-end")
	}
	switch p.stateMachine.Current() {
	case "walk":
	default:
	}
}

func (p *Player) Draw(screen *ebiten.Image, camX, camY float64) {
	pos := p.Body.Position()
	opts := ganim8.DrawOpts(pos.X-32, pos.Y-32)
	opts.X += camX
	opts.Y += camY
	// opts.OriginX = camX
	// opts.OriginY = camY
	p.sprite.CurrentAnim.Draw(screen, opts)
}
