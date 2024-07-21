package player

import . "gorpg/components"

func load() *Player {
	prefix := "assets/images/pixelarium-character/"
	var animations [][]*Animation
	animations = append(animations, []*Animation{})
	animations = append(animations, []*Animation{})
	animations = append(animations, []*Animation{})
	animations = append(animations, []*Animation{})

	player := Player{sprite: NewAS(animations)}
	// back
	frames := LoadSpriteSheet(prefix+"back-animations/spr_player_back_attack.png", 64, 64)
	player.sprite.AddAnimation(frames, ATTACK, .2, 64, 64, false)
	frames = LoadSpriteSheet(prefix+"back-animations/spr_player_back_hit.png", 64, 64)
	player.sprite.AddAnimation(frames, HIT, .2, 64, 64, true)
	frames = LoadSpriteSheet(prefix+"back-animations/spr_player_back_idle.png", 64, 64)
	player.sprite.AddAnimation(frames, IDLE, .2, 64, 64, true)
	frames = LoadSpriteSheet(prefix+"back-animations/spr_player_back_walk.png", 64, 64)
	player.sprite.AddAnimation(frames, WALK, .2, 64, 64, true)
	// front
	frames = LoadSpriteSheet(prefix+"front-animations/spr_player_front_attack.png", 64, 64)
	player.sprite.AddAnimation(frames, ATTACK, .2, 64, 64, false)
	frames = LoadSpriteSheet(prefix+"front-animations/spr_player_front_hit.png", 64, 64)
	player.sprite.AddAnimation(frames, HIT, .2, 64, 64, true)
	frames = LoadSpriteSheet(prefix+"front-animations/spr_player_front_idle.png", 64, 64)
	player.sprite.AddAnimation(frames, IDLE, .2, 64, 64, true)
	frames = LoadSpriteSheet(prefix+"front-animations/spr_player_front_walk.png", 64, 64)
	player.sprite.AddAnimation(frames, WALK, .2, 64, 64, true)

	// side
	frames = LoadSpriteSheet(prefix+"side-animations/spr_player_left_attack.png", 64, 64)
	player.sprite.AddAnimation(frames, ATTACK, .2, 64, 64, false)
	frames = LoadSpriteSheet(prefix+"side-animations/spr_player_left_hit.png", 64, 64)
	player.sprite.AddAnimation(frames, HIT, .2, 64, 64, true)
	frames = LoadSpriteSheet(prefix+"side-animations/spr_player_left_idle.png", 64, 64)
	player.sprite.AddAnimation(frames, IDLE, .2, 64, 64, true)
	frames = LoadSpriteSheet(prefix+"side-animations/spr_player_left_walk.png", 64, 64)
	player.sprite.AddAnimation(frames, WALK, .2, 64, 64, true)

	//right
	frames = LoadSpriteSheet(prefix+"side-animations/spr_player_right_attack.png", 64, 64)
	player.sprite.AddAnimation(frames, ATTACK, .2, 64, 64, false)
	frames = LoadSpriteSheet(prefix+"side-animations/spr_player_right_hit.png", 64, 64)
	player.sprite.AddAnimation(frames, HIT, .2, 64, 64, true)
	frames = LoadSpriteSheet(prefix+"side-animations/spr_player_right_idle.png", 64, 64)
	player.sprite.AddAnimation(frames, IDLE, .2, 64, 64, true)
	frames = LoadSpriteSheet(prefix+"side-animations/spr_player_right_walk.png", 64, 64)
	player.sprite.AddAnimation(frames, WALK, .2, 64, 64, true)

	player.sprite.ChangeAnimation(IDLE, RIGHT)
	return &player
}
