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
	p.sprite.ChangeAnimation(IDLE, p.Cardinal)
}

func (p *Player) enterWalk() {
	p.sprite.ChangeAnimation(WALK, p.Cardinal)
}

// TODO
// start animation
// create hurtbox and add to space
func (p *Player) enterAttack() {
	p.sprite.ChangeAnimation(ATTACK, p.Cardinal)
	pos := p.Object.Center()
	switch p.Cardinal {
	case UP:
		p.hurtboxes[0].SetCenter(pos.X, pos.Y-16)
	case DOWN:
		p.hurtboxes[0].SetCenter(pos.X, pos.Y+20)
	case LEFT:
		p.hurtboxes[0].SetCenter(pos.X-16, pos.Y+4)
	case RIGHT:
		p.hurtboxes[0].SetCenter(pos.X+16, pos.Y+4)
	}
	p.hurtboxes[0].Enable()
	fmt.Println(p.hurtboxes[0].Tags())
}

func (p *Player) attackEnd() {
	p.sprite.ChangeAnimation(IDLE, p.Cardinal)
	p.hurtboxes[0].Disable()
	fmt.Println(p.hurtboxes[0].Tags())
}

func enterDodge() {

}
