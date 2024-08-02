package player

import (
	"context"

	"github.com/looplab/fsm"
)

func (p *Player) NewSM() {
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
	p.sprite.ChangeAnimation(IDLE, p.cardinal)
}

func (p *Player) enterWalk() {
	p.sprite.ChangeAnimation(WALK, p.cardinal)
}

// TODO
// start animation
// create hurtbox and add to space
func (p *Player) enterAttack() {
	p.sprite.ChangeAnimation(ATTACK, p.cardinal)
	p.hurtboxes[0].Enabled = true
}

func (p *Player) attackEnd() {
	p.sprite.ChangeAnimation(IDLE, p.cardinal)
	p.hurtboxes[0].Enabled = false
	p.hurtboxes[0].Reset()
}

func enterDodge() {

}
