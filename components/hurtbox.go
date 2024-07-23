package components

import (
	"fmt"
	"slices"

	"github.com/jakecoffman/cp/v2"
)

type HurtBox struct {
	shape   *cp.Shape
	Enabled bool
	hits    []*cp.Shape
	damage  int
}

func NewHurtBox(rad float64, space *cp.Space, body *cp.Body, offset *cp.Vector) *HurtBox {
	obj := HurtBox{
		shape: cp.NewCircle(body, 16, *offset),
	}
	space.AddShape(obj.shape)
	filter := cp.NewShapeFilter(0, HIT_LAYER, ENEMY_LAYER)
	obj.shape.SetFilter(filter)
	obj.shape.SetCollisionType(HIT_TYPE)
	obj.shape.UserData = &obj
	obj.Enabled = false
	obj.damage = 2
	fmt.Printf("bod %v\nbox %v", obj.shape.Body().Position(), obj.shape.BB().Center())
	return &obj
}
func (hb *HurtBox) Value() int {
	return hb.damage
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
