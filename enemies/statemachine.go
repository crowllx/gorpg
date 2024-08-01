package enemies

import (
	"context"

	"github.com/looplab/fsm"
)

func EnemyStateMachine() *fsm.FSM {
	sm := fsm.NewFSM("idle",
		fsm.Events{
			{Name: "idle", Src: []string{"chase"}, Dst: "idle"},
			{Name: "chase", Src: []string{"idle", "attack"}, Dst: "chase"},
			{Name: "attack", Src: []string{"idle", "chase"}, Dst: "attack"},
			{Name: "attack-end", Src: []string{"attack"}, Dst: "chase"},
		},
		fsm.Callbacks{
			"enter_state": func(_ context.Context, e *fsm.Event) { enterState(e) },
		},
	)
	return sm
}

func enterState(e *fsm.Event) {

}
