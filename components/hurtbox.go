package components

import (
	"fmt"

	"github.com/jakecoffman/cp/v2"
)

type HurtBox struct {
	shape *cp.Shape
}

func NewHurtBox(rad float64, space *cp.Space, body *cp.Body, offset *cp.Vector) *HurtBox {
	obj := HurtBox{
		shape: cp.NewCircle(body, 16, *offset),
	}
	space.AddShape(obj.shape)
	filter := cp.NewShapeFilter(1, 1, 0)
	obj.shape.SetFilter(filter)

	fmt.Printf("bod %v\nbox %v", obj.shape.Body().Position(), obj.shape.BB().Center())
	return &obj
}

// func (hb *HurtBox) Move(vec *cp.Vector, direction) {

// func (hb *HurtBox) AddHits(objs ...*resolv.Object) {
// 	for _, o := range objs {
// 		hb.hits = append(hb.hits, o)
// 		hb.AddToIgnoreList(o)
// 	}
// }

// returns dx and dy after checking for collisions,
// as well as a boolean indicating if collision occured with
// a hurtbox if c has tag hitbox
