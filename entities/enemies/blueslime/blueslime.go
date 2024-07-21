package enemies

import (
	"gorpg/components"
	"gorpg/entities/enemies"
)

const (
	attack = iota
	idle
	walk
	hurt
	death
)

type BlueSlime struct {
	*enemies.BaseEnemy
}

func NewSlime() *BlueSlime {
	e := load()
	return e
}

func load() *BlueSlime {
	prefix := "assets/images/pixelarium-enemy/slime/blue-slime/spr_Blue_slime_"
	e := &BlueSlime{
		enemies.New(),
	}
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
