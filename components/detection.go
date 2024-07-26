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
	target      *cp.Shape
	state       int
	speed       float64
}

func NewDetection(d float64, body *cp.Body, space *cp.Space, collisionFilter uint) *Detection {
	shape := cp.NewCircle(body, d, cp.Vector{0, 0})
	filter := cp.NewShapeFilter(0, DETECTION_LAYER, collisionFilter)
	shape.SetCollisionType(4)
	fmt.Printf("%v\n", shape)
	shape.SetFilter(filter)
	los := &Detection{d, true, shape, body, nil, 0, 1}
	shape.UserData = los
	space.AddShape(shape)
	return los
}

func (l *Detection) onEnter(shape *cp.Shape, _ *cp.ContactPointSet) {
}
func (l *Detection) Update() {
	if l.enabled && l.state == 0 {
		info := l.shape.Space().PointQueryNearest(l.body.Position(), l.maxDistance, l.shape.Filter)
        velocity := info.Point.Sub(l.shape.BB().Center()).Normalize().Mult(l.speed)
		if info.Shape != nil {
			l.shape.Body().SetVelocityVector(velocity)
		} else {
            l.shape.Body().SetVelocity(0,0)
        }
	}
}
