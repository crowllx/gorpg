package player

import (
	"context"
	"fmt"

	"github.com/looplab/fsm"
)

func (p *Player) NewStateMachine() {
	sm := fsm.NewFSM("idle",
		fsm.Events{
			{Name: "idle", Src: []string{"walk"}, Dst: "idle"},
			{Name: "walk", Src: []string{"idle"}, Dst: "walk"},
			{Name: "attack", Src: []string{"idle", "walk"}, Dst: "attack"},
			{Name: "attack-end", Src: []string{"attack"}, Dst: "idle"},
		},
		fsm.Callbacks{
			"enter_state": func(_ context.Context, e *fsm.Event) { p.enterState(e) },
		},
	)
	p.stateMachine = sm
	p.stateMachine.SetState("idle")
}

func (p *Player) enterIdle() {
	p.sprite.Current = IDLE
}

func (p *Player) enterWalk() {
	p.sprite.Current = WALK
}

// TODO
// start animation
// create hurtbox and add to space
func (p *Player) enterAttack() {
	p.sprite.ChangeAnimation(ATTACK, p.Cardinal)
	fmt.Println(p.hurtboxes[0].Tags())
	fmt.Println(p.hurtboxes[0].BoundsToSpace(0, 0))
	switch p.Cardinal {
	case UP:
		p.hurtboxes[0].Position.X = p.X + 20
		p.hurtboxes[0].Position.Y = p.Y + 12
	case DOWN:
		p.hurtboxes[0].Position.X = p.X + 20
		p.hurtboxes[0].Position.Y = p.Y + 36
	case LEFT:
		p.hurtboxes[0].Position.X = p.X + 2
		p.hurtboxes[0].Position.Y = p.Y + 24
	case RIGHT:
		p.hurtboxes[0].Position.X = p.X + 38
		p.hurtboxes[0].Position.Y = p.Y + 24
	}
}

func enterDodge() {

}
