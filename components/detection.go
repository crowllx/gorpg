package components

import (
	"github.com/jakecoffman/cp/v2"
)

type Detection struct {
	Radius  float64
	Enabled bool
	Shape   *cp.Shape
}

func NewDetection(d float64, body *cp.Body, space *cp.Space, collisionFilter uint) *Detection {
	shape := cp.NewCircle(body, d, cp.Vector{0, 0})
	filter := cp.NewShapeFilter(0, DETECTION_LAYER, collisionFilter)
	shape.SetCollisionType(4)
	shape.SetFilter(filter)
	los := &Detection{d, true, shape}
	shape.UserData = los
	space.AddShape(shape)
	return los
}
