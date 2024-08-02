package enemies

import (
	"fmt"
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
func NewSlime(pos cp.Vector, space *cp.Space) *BlueSlime {
	e := load()
	e.body = cp.NewKinematicBody()
	e.body.SetPosition(pos)
	e.shape = cp.NewCircle(e.body, 16, cp.Vector{X: 0, Y: 0})
	mask := components.ENEMY_LAYER |
		components.ENVIRONMENT_LAYER | components.HIT_LAYER
	fmt.Println(mask)
	filter := cp.NewShapeFilter(0, components.ENEMY_LAYER, mask)
	e.shape.SetFilter(filter)
	e.shape.SetCollisionType(components.ENEMY_TYPE)
	e.shape.UserData = e
	e.Status = components.NewStatus(10, 0, e.Death)
	e.body.UserData = e
	e.body.AddShape(e.shape)
	e.aggroRadius = components.NewDetection(100, e.body, space, components.PLAYER_LAYER)
	e.stateMachine = e.EnemyStateMachine()
	return e
}

func load() *BlueSlime {
	prefix := "assets/images/pixelarium-enemy/slime/blue-slime/spr_Blue_slime_"
	e := &BlueSlime{BaseEnemy: &BaseEnemy{}}
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
	e.Sprite.AddAnimation(frames, idle, .2, 64, 64, true)
	frames = components.LoadSpriteSheet(prefix+"walk.png", 64, 64)
	e.Sprite.AddAnimation(frames, walk, .2, 64, 64, true)
	frames = components.LoadSpriteSheet(prefix+"hurt.png", 64, 64)
	e.Sprite.AddAnimation(frames, hurt, .2, 64, 64, true)
	frames = components.LoadSpriteSheet(prefix+"death.png", 64, 64)
	e.Sprite.AddAnimation(frames, death, .2, 64, 64, true)
	e.Sprite.ChangeAnimation(idle, 0)
	return e
}
