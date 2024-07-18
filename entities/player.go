package entities

import (
	. "gorpg/components"
)

type Direction int

const (
	UP = iota
	DOWN
	SIDE
)

type AnimationState int

const (
	ATTACK = iota
	HIT
	IDLE
	WALK
)

type Player struct {
	*AnimatedSprite
}

func New() *Player {

	// back
	prefix := "assets/images/pixelarium-character/"
	var animations [][]*Animation
	animations = append(animations, []*Animation{})
	animations = append(animations, []*Animation{})
	animations = append(animations, []*Animation{})
	animations = append(animations, []*Animation{})

	player := Player{AnimatedSprite: &AnimatedSprite{
		Animations: animations,
		Current:    IDLE,
		Direction:  SIDE,
		X:          100,
		Y:          100,
	}}
	frames := LoadSpriteSheet(prefix+"back-animations/spr_player_back_attack.png", 64, 64)
	player.AddAnimation(frames, ATTACK, .2)
	frames = LoadSpriteSheet(prefix+"back-animations/spr_player_back_hit.png", 64, 64)
	player.AddAnimation(frames, HIT, .2)
	frames = LoadSpriteSheet(prefix+"back-animations/spr_player_back_idle.png", 64, 64)
	player.AddAnimation(frames, IDLE, .2)
	frames = LoadSpriteSheet(prefix+"back-animations/spr_player_back_walk.png", 64, 64)
	player.AddAnimation(frames, WALK, .2)
	// front
	frames = LoadSpriteSheet(prefix+"front-animations/spr_player_front_attack.png", 64, 64)
	player.AddAnimation(frames, ATTACK, .2)
	frames = LoadSpriteSheet(prefix+"front-animations/spr_player_front_hit.png", 64, 64)
	player.AddAnimation(frames, HIT, .2)
	frames = LoadSpriteSheet(prefix+"front-animations/spr_player_front_idle.png", 64, 64)
	player.AddAnimation(frames, IDLE, .2)
	frames = LoadSpriteSheet(prefix+"front-animations/spr_player_front_walk.png", 64, 64)
	player.AddAnimation(frames, WALK, .2)

	// side
	frames = LoadSpriteSheet(prefix+"side-animations/spr_player_left_attack.png", 64, 64)
	player.AddAnimation(frames, ATTACK, .2)
	frames = LoadSpriteSheet(prefix+"side-animations/spr_player_left_hit.png", 64, 64)
	player.AddAnimation(frames, HIT, .2)
	frames = LoadSpriteSheet(prefix+"side-animations/spr_player_left_idle.png", 64, 64)
	player.AddAnimation(frames, IDLE, .2)
	frames = LoadSpriteSheet(prefix+"side-animations/spr_player_left_walk.png", 64, 64)
	player.AddAnimation(frames, WALK, .2)

	return &player
}
