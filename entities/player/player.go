package player

import (
	"context"
	"fmt"
	. "gorpg/components"

	"github.com/looplab/fsm"
	input "github.com/quasilyte/ebitengine-input"
	"github.com/solarlune/resolv"
)

type direction int

const (
	UP = iota
	DOWN
	LEFT
	RIGHT
)

type animationState int

const (
	ATTACK = iota
	HIT
	IDLE
	WALK
)

type Player struct {
	*Movement
	sprite       *AnimatedSprite
	input        *input.Handler
	stateMachine *fsm.FSM
	speed        int
	collider     *Collider
	hurtboxes    []*HurtBox
}

func NewPlayer(space *resolv.Space) *Player {
	prefix := "assets/images/pixelarium-character/"
	var animations [][]*Animation
	animations = append(animations, []*Animation{})
	animations = append(animations, []*Animation{})
	animations = append(animations, []*Animation{})
	animations = append(animations, []*Animation{})

	player := Player{sprite: NewAS(animations)}
	player.NewStateMachine()
	player.speed = 1
	player.Movement = &Movement{
		X:        100,
		Y:        100,
		Dir:      Direction{X: 0, Y: 0},
		Cardinal: RIGHT,
	}
	player.collider = (NewCollider(100, 100, 16, 16, "solid", "hit"))
	(*resolv.Object)(player.collider).SetCenter(132, 132)
	space.Add((*resolv.Object)(player.collider))
	// back
	frames := LoadSpriteSheet(prefix+"back-animations/spr_player_back_attack.png", 64, 64)
	player.sprite.AddAnimation(frames, ATTACK, .2, 64, 64, false)
	frames = LoadSpriteSheet(prefix+"back-animations/spr_player_back_hit.png", 64, 64)
	player.sprite.AddAnimation(frames, HIT, .2, 64, 64, true)
	frames = LoadSpriteSheet(prefix+"back-animations/spr_player_back_idle.png", 64, 64)
	player.sprite.AddAnimation(frames, IDLE, .2, 64, 64, true)
	frames = LoadSpriteSheet(prefix+"back-animations/spr_player_back_walk.png", 64, 64)
	player.sprite.AddAnimation(frames, WALK, .2, 64, 64, true)
	// front
	frames = LoadSpriteSheet(prefix+"front-animations/spr_player_front_attack.png", 64, 64)
	player.sprite.AddAnimation(frames, ATTACK, .2, 64, 64, false)
	frames = LoadSpriteSheet(prefix+"front-animations/spr_player_front_hit.png", 64, 64)
	player.sprite.AddAnimation(frames, HIT, .2, 64, 64, true)
	frames = LoadSpriteSheet(prefix+"front-animations/spr_player_front_idle.png", 64, 64)
	player.sprite.AddAnimation(frames, IDLE, .2, 64, 64, true)
	frames = LoadSpriteSheet(prefix+"front-animations/spr_player_front_walk.png", 64, 64)
	player.sprite.AddAnimation(frames, WALK, .2, 64, 64, true)

	// side
	frames = LoadSpriteSheet(prefix+"side-animations/spr_player_left_attack.png", 64, 64)
	player.sprite.AddAnimation(frames, ATTACK, .2, 64, 64, false)
	frames = LoadSpriteSheet(prefix+"side-animations/spr_player_left_hit.png", 64, 64)
	player.sprite.AddAnimation(frames, HIT, .2, 64, 64, true)
	frames = LoadSpriteSheet(prefix+"side-animations/spr_player_left_idle.png", 64, 64)
	player.sprite.AddAnimation(frames, IDLE, .2, 64, 64, true)
	frames = LoadSpriteSheet(prefix+"side-animations/spr_player_left_walk.png", 64, 64)
	player.sprite.AddAnimation(frames, WALK, .2, 64, 64, true)

	//right
	frames = LoadSpriteSheet(prefix+"side-animations/spr_player_right_attack.png", 64, 64)
	player.sprite.AddAnimation(frames, ATTACK, .2, 64, 64, false)
	frames = LoadSpriteSheet(prefix+"side-animations/spr_player_right_hit.png", 64, 64)
	player.sprite.AddAnimation(frames, HIT, .2, 64, 64, true)
	frames = LoadSpriteSheet(prefix+"side-animations/spr_player_right_idle.png", 64, 64)
	player.sprite.AddAnimation(frames, IDLE, .2, 64, 64, true)
	frames = LoadSpriteSheet(prefix+"side-animations/spr_player_right_walk.png", 64, 64)
	player.sprite.AddAnimation(frames, WALK, .2, 64, 64, true)

	player.sprite.ChangeAnimation(IDLE, RIGHT)
	hb := NewHurtBox(player.X, player.Y, 24, 24)
	hb.Position.X = player.X + 32
	hb.Position.Y = player.Y + 24
	hb.AddToIgnoreList((*resolv.Object)(player.collider))
	(*resolv.Object)(player.collider).AddToIgnoreList(hb.Object)
	fmt.Printf("%+v", hb.Center())
	space.Add(hb.Object)
	player.hurtboxes = append(player.hurtboxes, hb)
	return &player
}
func (p *Player) AddInputHandler(s *input.System) {
	p.input = s.NewHandler(0, GenerateKeyMap())
}

func (p *Player) Pos() (x, y float64) {
	return p.X, p.Y
}
func (p *Player) Sprite() *AnimatedSprite {
	return p.sprite
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
	switch {
	case dir.Y == 1:
		p.Cardinal = DOWN
		p.stateMachine.Event(context.Background(), "walk")
	case dir.Y == -1:
		p.Cardinal = UP
		p.stateMachine.Event(context.Background(), "walk")
	case dir.X == -1:
		p.Cardinal = LEFT
		p.stateMachine.Event(context.Background(), "walk")
	case dir.X == 1:
		p.Cardinal = RIGHT
		p.stateMachine.Event(context.Background(), "walk")
	default:
		p.stateMachine.Event(context.Background(), "idle")
	}
	if p.input.ActionIsJustPressed(ActionAttack) {
		p.stateMachine.Event(context.Background(), "attack")
	}
	if p.stateMachine.Is("attack") {
		dir.X, dir.Y = 0, 0
	}
	dx := float64(dir.X * p.speed)
	dy := float64(dir.Y * p.speed)
	dx, dy, _ = p.collider.Check(dx, dy)
	p.X += dx
	p.collider.Position.X += dx
	p.Y += dy
	p.collider.Position.Y += dy
	p.hurtboxes[0].Position.X += dx
	p.hurtboxes[0].Position.Y += dy
	if col := p.hurtboxes[0].Check(dx, dy, "hurt"); col != nil {
		fmt.Print("i hit something")
	}

	p.sprite.CurrentImg.Update()
	if p.sprite.CurrentImg.Done() {
		p.stateMachine.Event(context.Background(), "attack-end")
	}
}
