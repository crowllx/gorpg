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
	switch event.Event {
	case "idle":
		e.Sprite.ChangeAnimation("idle")
	case "chase":
		e.Sprite.ChangeAnimation("idle")
	case "attack":
		e.enterAttack()
	case "attack-end":
		e.attackEnd()
	default:
	}
}

func (p *BaseEnemy) enterAttack() {
	p.Sprite.ChangeAnimation("attack")
	p.hurtboxes[0].Enabled = true
}

func (p *BaseEnemy) attackEnd() {
	p.Sprite.ChangeAnimation("walk")
	p.hurtboxes[0].Enabled = false
	p.hurtboxes[0].Reset()
}
