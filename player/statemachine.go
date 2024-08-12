package player

import (
	"context"
	"fmt"

	"github.com/looplab/fsm"
)

var smPipe = make(chan int)

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
	p.sprite.ChangeAnimation("idle")
}

func (p *Player) enterWalk() {
	p.switchAnim("walk")
    go func() {
        p.sfxEmitter.Streams[walkDirt].Play()
        <- smPipe
        p.sfxEmitter.Streams[walkDirt].Pause()
        p.sfxEmitter.Streams[walkDirt].Rewind()
    } ()
    fmt.Printf("walking\n")
}

// TODO
// start animation
// create hurtbox and add to space
func (p *Player) enterAttack() {
	p.switchAnim("attack")
	p.hurtboxes[0].Enabled = true
    go func() {
        p.sfxEmitter.Streams[basicAttack].Play()
        <- smPipe
        p.sfxEmitter.Streams[basicAttack].Rewind()
    } ()
}

func (p *Player) attackEnd() {
	p.switchAnim("idle")
	p.hurtboxes[0].Enabled = false
	p.hurtboxes[0].Reset()
}

func enterDodge() {

}
