package player

import (
	"github.com/jakecoffman/cp/v2"
	input "github.com/quasilyte/ebitengine-input"
)

const (
	ActionUnkown input.Action = iota
	ActionMoveLeft
	ActionMoveRight
	ActionMoveUp
	ActionMoveDown
	ActionAttack
	ActionDodge
)

func GenerateKeyMap() input.Keymap {
	keymap := input.Keymap{
		ActionMoveLeft:  {input.KeyLeft, input.KeyA},
		ActionMoveRight: {input.KeyRight, input.KeyD},
		ActionMoveUp:    {input.KeyUp, input.KeyW},
		ActionMoveDown:  {input.KeyDown, input.KeyS},
		ActionAttack:    {input.Key2},
	}
	return keymap
}

func (p *Player) InputVec() cp.Vector {
	dir := cp.Vector{X: 0, Y: 0}
	if p.input.ActionIsPressed(ActionMoveLeft) {
		dir.X = -1
	}
	if p.input.ActionIsPressed(ActionMoveRight) {
		dir.X = 1
	}
	if p.input.ActionIsPressed(ActionMoveUp) {
		dir.Y = -1
	}
	if p.input.ActionIsPressed(ActionMoveDown) {
		dir.Y = 1
	}
	return dir
}
