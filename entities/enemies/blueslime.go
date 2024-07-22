package enemies

import (
	"gorpg/components"

	"github.com/jakecoffman/cp/v2"
)

const (
	attack = iota
	idle
	walk
	hurt
	death
)

type BlueSlime struct {
	*BaseEnemy
}

// TODO: statemachine/ai logic & hurtboxes
func NewSlime(pos cp.Vector) *BlueSlime {
	e := load()
	e.body = cp.NewKinematicBody()
	e.body.SetPosition(pos)
	e.shape = cp.NewCircle(e.body, 16, cp.Vector{X: 0, Y: 0})
	filter := cp.NewShapeFilter(0, uint(0b0001000), uint(0b00001111))
	e.shape.SetFilter(filter)
	e.shape.SetCollisionType(3)
	e.shape.UserData = e
	e.Status = components.NewStatus(10, 0, e.Death)
	e.body.UserData = e
	return e
}

func load() *BlueSlime {
	prefix := "assets/images/pixelarium-enemy/slime/blue-slime/spr_Blue_slime_"
	e := &BlueSlime{&BaseEnemy{}}
	var animations [][]*components.Animation
	animations = append(animations, []*components.Animation{})
	animations = append(animations, []*components.Animation{})
	animations = append(animations, []*components.Animation{})
	animations = append(animations, []*components.Animation{})
	animations = append(animations, []*components.Animation{})
	e.Sprite = components.NewAS(animations)

	// attack
	frames := components.LoadSpriteSheet(prefix+"attack.png", 64, 64)
	e.Sprite.AddAnimation(frames, attack, .2, 64, 64, false)
	frames = components.LoadSpriteSheet(prefix+"idle.png", 64, 64)
	e.Sprite.AddAnimation(frames, idle, .2, 64, 64, false)
	frames = components.LoadSpriteSheet(prefix+"walk.png", 64, 64)
	e.Sprite.AddAnimation(frames, walk, .2, 64, 64, false)
	frames = components.LoadSpriteSheet(prefix+"hurt.png", 64, 64)
	e.Sprite.AddAnimation(frames, hurt, .2, 64, 64, false)
	frames = components.LoadSpriteSheet(prefix+"death.png", 64, 64)
	e.Sprite.AddAnimation(frames, death, .2, 64, 64, false)

	e.Sprite.ChangeAnimation(idle, 0)
	return e
}
