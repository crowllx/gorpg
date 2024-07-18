package entities

type Player struct {
	*AnimatedSprite
}

func New() *Player {

	// back
	prefix := "assets/images/pixelarium-character/"
	player := Player{AnimatedSprite: &AnimatedSprite{
		animations: map[string]*Animation{},
		current:    "front-idle",
		X:          100,
		Y:          100,
	}}
	frames := LoadSpriteSheet(prefix+"back-animations/spr_player_back_attack.png", 64, 64)
	player.AddAnimation(frames, "back-attack", .2)
	frames = LoadSpriteSheet(prefix+"back-animations/spr_player_back_hit.png", 64, 64)
	player.AddAnimation(frames, "back-hit", .2)
	frames = LoadSpriteSheet(prefix+"back-animations/spr_player_back_idle.png", 64, 64)
	player.AddAnimation(frames, "back-idle", .2)
	frames = LoadSpriteSheet(prefix+"back-animations/spr_player_back_walk.png", 64, 64)
	player.AddAnimation(frames, "back-walk", .2)
	// front
	frames = LoadSpriteSheet(prefix+"front-animations/spr_player_front_attack.png", 64, 64)
	player.AddAnimation(frames, "front-attack", .2)
	frames = LoadSpriteSheet(prefix+"front-animations/spr_player_front_hit.png", 64, 64)
	player.AddAnimation(frames, "front-hit", .2)
	frames = LoadSpriteSheet(prefix+"front-animations/spr_player_front_idle.png", 64, 64)
	player.AddAnimation(frames, "front-idle", .2)
	frames = LoadSpriteSheet(prefix+"front-animations/spr_player_front_walk.png", 64, 64)
	player.AddAnimation(frames, "front-walk", .2)

	// side
	frames = LoadSpriteSheet(prefix+"side-animations/spr_player_left_attack.png", 64, 64)
	player.AddAnimation(frames, "left-attack", .2)
	frames = LoadSpriteSheet(prefix+"side-animations/spr_player_left_hit.png", 64, 64)
	player.AddAnimation(frames, "left-hit", .2)
	frames = LoadSpriteSheet(prefix+"side-animations/spr_player_left_idle.png", 64, 64)
	player.AddAnimation(frames, "left-idle", .2)
	frames = LoadSpriteSheet(prefix+"side-animations/spr_player_left_walk.png", 64, 64)
	player.AddAnimation(frames, "left-walk", .2)

	return &player
}
