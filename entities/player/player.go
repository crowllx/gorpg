package player

import (
	"context"
	"fmt"
	. "gorpg/components"
	"gorpg/entities/enemies"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/looplab/fsm"
	input "github.com/quasilyte/ebitengine-input"
	"github.com/solarlune/resolv"
)

type direction int

// direction the player is facing
const (
	UP = iota
	DOWN
	LEFT
	RIGHT
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
	*resolv.Object
	*Movement
	sprite       *AnimatedSprite
	input        *input.Handler
	stateMachine *fsm.FSM
	speed        int
	hurtboxes    []*HurtBox
	Status       *Status
}

func New() *Player {
	player := load()
	player.NewStateMachine()
	player.speed = 1
	player.Movement = &Movement{
		X:        100,
		Y:        100,
		Dir:      Direction{X: 0, Y: 0},
		Cardinal: RIGHT,
	}
	// creation of player obj
	player.Object = resolv.NewObject(100, 100, 16, 16, "solid", "hit")
	player.Object.SetCenter(100, 100)
	player.Object.Update()
	player.Object.Data = player

	hb := NewHurtBox(player.X, player.Y, 24, 24)
	hb.Disable()
	player.hurtboxes = append(player.hurtboxes, hb)
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
func (p *Player) AddToSpace(space *resolv.Space) {
	space.Add(p.Object)
	shape := resolv.NewCircle(100, 100, 64)
	p.Object.SetShape(shape)
	p.Object.Update()
	for _, o := range p.hurtboxes {
		space.Add(o.Object)
		p.AddToIgnoreList(o.Object)
		o.AddToIgnoreList(p.Object)
	}
}

func (p *Player) enterState(e *fsm.Event) {
	fmt.Println(e.Event)
	for i := range 4 {
		p.sprite.CurrentAnimation(i).Reset()
	}
	switch e.Event {
	case "idle":
		p.sprite.ChangeAnimation(IDLE, p.Cardinal)
	case "walk":
		p.sprite.ChangeAnimation(WALK, p.Cardinal)
	case "attack":
		p.enterAttack()
	case "attack-end":
		p.sprite.ChangeAnimation(IDLE, p.Cardinal)
	}
}

func (p *Player) Update() {
	dir := p.InputVec()
	p.Dir = dir
	switch {
	case dir.Y == 1:
		p.Cardinal = DOWN
		if p.stateMachine.Can("walk") || p.stateMachine.Current() == "walk" {
			p.stateMachine.Event(context.Background(), "walk")
			p.sprite.ChangeAnimation(WALK, p.Cardinal)
		}
	case dir.Y == -1:
		p.Cardinal = UP
		if p.stateMachine.Can("walk") || p.stateMachine.Current() == "walk" {
			p.stateMachine.Event(context.Background(), "walk")
			p.sprite.ChangeAnimation(WALK, p.Cardinal)
		}
	case dir.X == -1:
		p.Cardinal = LEFT
		if p.stateMachine.Can("walk") || p.stateMachine.Current() == "walk" {
			p.stateMachine.Event(context.Background(), "walk")
			p.sprite.ChangeAnimation(WALK, p.Cardinal)
		}
	case dir.X == 1:
		p.Cardinal = RIGHT
		if p.stateMachine.Can("walk") || p.stateMachine.Current() == "walk" {
			p.stateMachine.Event(context.Background(), "walk")
			p.sprite.ChangeAnimation(WALK, p.Cardinal)
		}
	default:
		p.stateMachine.Event(context.Background(), "idle")
	}
	if p.input.ActionIsJustPressed(ActionAttack) {
		p.stateMachine.Event(context.Background(), "attack")
	}
	if p.stateMachine.Is("attack") && p.sprite.Current == ATTACK {
		dir.X, dir.Y = 0, 0
	}
	dx := float64(dir.X * p.speed)
	dy := float64(dir.Y * p.speed)
	if collision := p.Object.Check(dx, 0); collision != nil {
		dx = 0
	}
	if collision := p.Object.Check(0, dy); collision != nil {
		dy = 0
	}
	p.Object.Position.X += dx
	p.Object.Position.Y += dy
	p.hurtboxes[0].Position.X += dx
	p.hurtboxes[0].Position.Y += dy

	if col := p.hurtboxes[0].Check(dx, dy, "hit"); col != nil && p.hurtboxes[0].HasTags("hurt") {
		fmt.Println("i hit something")
		fmt.Printf("%T\n", col.Objects[0].Data)
		switch col.Objects[0].Data.(type) {
		case *enemies.BaseEnemy:
			hb := p.hurtboxes[0]
			col.Objects[0].Data.(*enemies.BaseEnemy).Status.Modify("health", -2)
			hb.AddHits(col.Objects...)
		default:
			fmt.Println("not a valid target")
		}
	}

	p.sprite.CurrentImg.Update()
	if p.stateMachine.Current() == "attack" && p.sprite.CurrentImg.Done() {
		p.stateMachine.Event(context.Background(), "attack-end")
		p.attackEnd()
	}
}
func (p *Player) Draw(screen *ebiten.Image) {
	opts := ebiten.DrawImageOptions{}
	pos := p.Object.Center()
	opts.GeoM.Translate(pos.X-32, pos.Y-32)
	screen.DrawImage(p.sprite.CurrentImg.Draw(), &opts)
}
