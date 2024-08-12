package components

import (
	sound "gorpg/audio"
	"slices"

	"github.com/jakecoffman/cp/v2"
)
type hurtable interface {
    Modify(string, int) (int, error)
}
type HurtBox struct {
	shape   *cp.Shape
	Enabled bool
	hits    []*cp.Shape
	damage  int  
    sfx *sound.AudioEmitter
}

func NewHurtBox(rad float64, space *cp.Space, body *cp.Body, offset *cp.Vector, filter cp.ShapeFilter) *HurtBox {
	obj := HurtBox{
		shape: cp.NewCircle(body, rad, *offset),
	}
	space.AddShape(obj.shape)
	obj.shape.SetFilter(filter)
	obj.shape.SetCollisionType(HIT_TYPE)
	obj.shape.UserData = &obj
	obj.Enabled = false
	obj.damage = 2
    obj.sfx = sound.NewEmitter()
	return &obj
}
func (hb *HurtBox) OnHit(stat hurtable) {
    for _, v := range hb.sfx.Streams {
        v.Play()
        go func() {
            for v.IsPlaying(){}
            v.Rewind()
        }()
    }
    stat.Modify("health", hb.damage)
}
func (hb *HurtBox) AddSFX(id string, loop bool) {
    hb.sfx.NewPlayer(id, loop)
}
func (hb *HurtBox) Shape() *cp.Shape {
	return hb.shape
}
func (hb *HurtBox) Reset() {
	hb.hits = nil
}
func (hb *HurtBox) HitCheck(s *cp.Shape) bool {
	if !hb.Enabled || slices.Contains(hb.hits, s) {
		return false
	}
	hb.hits = append(hb.hits, s)
	return true
}
