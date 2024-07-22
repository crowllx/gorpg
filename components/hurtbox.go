package components

import (
	"fmt"

	"github.com/jakecoffman/cp/v2"
)

type HurtBox struct {
	Shape *cp.Shape
}

const HURTBOX = 4

func NewHurtBox(rad float64, space *cp.Space, body *cp.Body, offset *cp.Vector) *HurtBox {
	obj := HurtBox{
		Shape: cp.NewCircle(body, 16, *offset),
	}
	space.AddShape(obj.Shape)
	filter := cp.NewShapeFilter(0, HURTBOX, 0b00001000)
	obj.Shape.SetFilter(filter)
	obj.Shape.SetCollisionType(HURTBOX)
	obj.Shape.UserData = obj
	fmt.Printf("bod %v\nbox %v", obj.Shape.Body().Position(), obj.Shape.BB().Center())
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
