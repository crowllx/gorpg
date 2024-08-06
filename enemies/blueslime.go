package enemies

import (
	"fmt"
	"gorpg/components"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jakecoffman/cp/v2"
	"github.com/yohamta/ganim8/v2"
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

func NewSlime(pos cp.Vector, space *cp.Space) *BlueSlime {
	e := load()

	// body & hitbox
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

	// hurtboxes
	hb := components.NewHurtBox(32, space, e.body, &cp.Vector{0, 0},
		cp.NewShapeFilter(0, components.HIT_LAYER, components.PLAYER_LAYER))
	e.hurtboxes = append(e.hurtboxes, hb)
	// status & detection range
	e.Status = components.NewStatus(10, 0, e.Death, _onHealthChanged)
	e.body.UserData = e
	e.body.AddShape(e.shape)
	e.aggroRadius = components.NewDetection(100, e.body, space, components.PLAYER_LAYER, e.aggro)

	//state machine for enemy logic
	e.stateMachine = e.EnemyStateMachine()
	return e
}

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

func load() *BlueSlime {
	prefix := "assets/images/pixelarium-enemy/slime/blue-slime/spr_Blue_slime_"
	e := &BlueSlime{BaseEnemy: &BaseEnemy{}}
	frameW := 64
	frameH := 64

	img := components.LoadSpriteSheet(prefix+"attack.png", 64, 64)
	anim := loadImg(img, frameW, frameH)
	anim.SetOnLoop(ganim8.PauseAtStart)
	sprite := components.NewAS("attack", anim)
	img = components.LoadSpriteSheet(prefix+"idle.png", 64, 64)
	anim = loadImg(img, frameW, frameH)
	sprite.AddAnimation("idle", anim)
	img = components.LoadSpriteSheet(prefix+"walk.png", 64, 64)
	anim = loadImg(img, frameW, frameH)
	sprite.AddAnimation("walk", anim)
	img = components.LoadSpriteSheet(prefix+"hurt.png", 64, 64)
	anim = loadImg(img, frameW, frameH)
	sprite.AddAnimation("hurt", anim)
	img = components.LoadSpriteSheet(prefix+"death.png", 64, 64)
	anim = loadImg(img, frameW, frameH)
	sprite.AddAnimation("death", anim)
	sprite.ChangeAnimation("idle")
	e.Sprite = sprite
	return e
}
