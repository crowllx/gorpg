package enemies

import (
	"context"
	"time"

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
	select {
	case e.smPipe <- 1:
	default:
	}
	switch event.Event {
	case "idle":
		e.Sprite.ChangeAnimation("idle")
	case "chase":
		e.enterChase()
	case "attack":
		e.enterAttack()
	case "attack-end":
		e.attackEnd()
	default:
	}
}
func (e *BaseEnemy) enterChase() {
	e.Sprite.ChangeAnimation("walk")
	e.sfx.Streams[walkWater].Play()
	go func() {
		<-e.smPipe
		e.sfx.Streams[walkWater].Pause()
		e.sfx.Streams[walkWater].Rewind()
	}()
}
func (e *BaseEnemy) enterAttack() {
	e.Sprite.ChangeAnimation("attack")
	hitDelay := time.NewTimer(time.Millisecond * 1500)
	go func() {
		<-hitDelay.C
		e.hurtboxes[0].Enabled = true
	}()
}

func (e *BaseEnemy) attackEnd() {
	e.Sprite.ChangeAnimation("walk")
	e.hurtboxes[0].Enabled = false
	e.hurtboxes[0].Reset()
}
