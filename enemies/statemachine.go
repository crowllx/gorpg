package enemies

import (
	"context"

	"github.com/looplab/fsm"
)

func (enemy *BaseEnemy) EnemyStateMachine() *fsm.FSM {
	sm := fsm.NewFSM("idle",
		fsm.Events{
			{Name: "idle", Src: []string{"chase"}, Dst: "idle"},
			{Name: "chase", Src: []string{"idle", "attack"}, Dst: "chase"},
			{Name: "attack", Src: []string{"idle", "chase"}, Dst: "attack"},
			{Name: "attack-end", Src: []string{"attack"}, Dst: "chase"},
		},
		fsm.Callbacks{
			"enter_state": func(_ context.Context, e *fsm.Event) { enemy.enterState(e) },
		},
	)
	return sm
}

func (e *BaseEnemy) enterState(event *fsm.Event) {
	e.Sprite.CurrentAnimation(0).Reset()
	switch event.Event {
	case "idle":
		e.Sprite.ChangeAnimation(idle, 0)
	case "chase":
		e.Sprite.ChangeAnimation(walk, 0)
	case "attack":
		e.enterAttack()
	case "attack-end":
		e.attackEnd()
	default:
	}
}

func (p *BaseEnemy) enterAttack() {
	p.Sprite.ChangeAnimation(attack, 0)
	// p.hurtboxes[0].Enabled = true
}

func (p *BaseEnemy) attackEnd() {
	p.Sprite.ChangeAnimation(walk, 0)
	// p.hurtboxes[0].Enabled = false
	// p.hurtboxes[0].Reset()
}
