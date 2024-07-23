package components

// collision categories used to determine what shapes can collide with one another
const PLAYER_LAYER uint = 1
const ENVIRONMENT_LAYER uint = 2
const HIT_LAYER uint = 4
const ENEMY_LAYER uint = 8
const DETECTION_LAYER uint = 16

type CollisionType = int

// collision types used for overriding collision handlers
const (
	PLAYER_TYPE = iota
	ENVIRONMENT_TYPE
	HIT_TYPE
	ENEMY_TYPE
	DETECTION_TYPE
)
