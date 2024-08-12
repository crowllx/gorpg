package player

import (
	"fmt"
	sound "gorpg/audio"
	. "gorpg/components"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/ganim8/v2"
)

// sfx paths
const walkDirt = "audio/sfx/WAV Files/SFX/Footsteps/Dirt/Dirt Walk 1.wav"
const basicAttack = "audio/sfx/WAV Files/SFX/Attacks/Sword Attacks Hits and Blocks/Sword Attack 1.wav"
const basicAttackHit = "audio/sfx/WAV Files/SFX/Attacks/Sword Attacks Hits and Blocks/Sword Impact Hit 1.wav"

func loadImg(img *ebiten.Image, fw, fh int) *ganim8.Animation {
	imgW := img.Bounds().Dx()
	imgH := img.Bounds().Dy()
	grid := ganim8.NewGrid(fw, fh, imgW, imgH)
	cols := imgW / fw
	framesRange := fmt.Sprintf("%d-%d", 1, cols)
	frames := grid.Frames(framesRange, 1)
	anim := ganim8.New(img, frames, time.Millisecond*100)
	return anim
}
func load() *Player {
	prefix := "assets/images/pixelarium-character/"
	frameW := 64
	frameH := 64
	// back
	img := LoadSpriteSheet(prefix+"back-animations/spr_player_back_attack.png", 64, 64)
	anim := loadImg(img, frameW, frameH)
	anim.SetOnLoop(ganim8.PauseAtStart)
	sprite := NewAS("back-attack", anim)

	img = LoadSpriteSheet(prefix+"back-animations/spr_player_back_hit.png", 64, 64)
	anim = loadImg(img, frameW, frameH)
	sprite.AddAnimation("back-hit", anim)

	img = LoadSpriteSheet(prefix+"back-animations/spr_player_back_idle.png", 64, 64)
	anim = loadImg(img, frameW, frameH)
	sprite.AddAnimation("back-idle", anim)

	img = LoadSpriteSheet(prefix+"back-animations/spr_player_back_walk.png", 64, 64)
	anim = loadImg(img, frameW, frameH)
	sprite.AddAnimation("back-walk", anim)
	//right
	img = LoadSpriteSheet(prefix+"side-animations/spr_player_right_attack.png", 64, 64)
	anim = loadImg(img, frameW, frameH)
	anim.SetOnLoop(ganim8.PauseAtStart)
	sprite.AddAnimation("right-attack", anim)
	img = LoadSpriteSheet(prefix+"side-animations/spr_player_right_hit.png", 64, 64)
	anim = loadImg(img, frameW, frameH)
	sprite.AddAnimation("right-hit", anim)
	img = LoadSpriteSheet(prefix+"side-animations/spr_player_right_idle.png", 64, 64)
	anim = loadImg(img, frameW, frameH)
	sprite.AddAnimation("right-idle", anim)
	img = LoadSpriteSheet(prefix+"side-animations/spr_player_right_walk.png", 64, 64)
	anim = loadImg(img, frameW, frameH)
	sprite.AddAnimation("right-walk", anim)
	img = LoadSpriteSheet(prefix+"front-animations/spr_player_front_attack.png", 64, 64)
	anim = loadImg(img, frameW, frameH)
	anim.SetOnLoop(ganim8.PauseAtStart)
	sprite.AddAnimation("front-attack", anim)
	img = LoadSpriteSheet(prefix+"front-animations/spr_player_front_hit.png", 64, 64)
	anim = loadImg(img, frameW, frameH)
	sprite.AddAnimation("front-hit", anim)
	img = LoadSpriteSheet(prefix+"front-animations/spr_player_front_idle.png", 64, 64)
	anim = loadImg(img, frameW, frameH)
	sprite.AddAnimation("front-idle", anim)
	img = LoadSpriteSheet(prefix+"front-animations/spr_player_front_walk.png", 64, 64)
	anim = loadImg(img, frameW, frameH)
	sprite.AddAnimation("front-walk", anim)
	img = LoadSpriteSheet(prefix+"side-animations/spr_player_left_attack.png", 64, 64)
	anim = loadImg(img, frameW, frameH)
	anim.SetOnLoop(ganim8.PauseAtStart)
	sprite.AddAnimation("left-attack", anim)
	img = LoadSpriteSheet(prefix+"side-animations/spr_player_left_hit.png", 64, 64)
	anim = loadImg(img, frameW, frameH)
	sprite.AddAnimation("left-hit", anim)
	img = LoadSpriteSheet(prefix+"side-animations/spr_player_left_idle.png", 64, 64)
	anim = loadImg(img, frameW, frameH)
	sprite.AddAnimation("left-idle", anim)
	img = LoadSpriteSheet(prefix+"side-animations/spr_player_left_walk.png", 64, 64)
	anim = loadImg(img, frameW, frameH)
	sprite.AddAnimation("left-walk", anim)
	player := Player{sprite: sprite}
	player.sprite.ChangeAnimation("left-idle")

	// SFX
	emitter := sound.NewEmitter()
	emitter.NewPlayer(walkDirt, true)
    emitter.NewPlayer(basicAttack, false)
	player.sfxEmitter = emitter
	return &player
}
