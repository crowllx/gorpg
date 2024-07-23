package components

import (
	"fmt"

	"github.com/jakecoffman/cp/v2"
)

type Detection struct {
	maxDistance float64
	enabled     bool
	shape       *cp.Shape
	body        *cp.Body
}

func NewDetection(d float64, body *cp.Body, space *cp.Space, collisionFilter uint) *Detection {
	shape := cp.NewCircle(body, d, cp.Vector{0, 0})
	filter := cp.NewShapeFilter(0, DETECTION_LAYER, collisionFilter)
	shape.SetCollisionType(4)
	fmt.Printf("%v\n", shape)
	shape.SetFilter(filter)
	los := &Detection{d, false, shape, body}
	shape.UserData = los
	space.AddShape(shape)
	return los
}

func (l *Detection) onEnter(shape *cp.Shape, _ *cp.ContactPointSet) {
	fmt.Printf("%T\n", shape.UserData)
}
func (l *Detection) Update() {
	if l.enabled {
		res := l.shape.Space().ShapeQuery(l.shape, l.onEnter)
		fmt.Printf("%v\n", res)
	}
}
