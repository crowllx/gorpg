package components

import (
	"fmt"

	"github.com/jakecoffman/cp/v2"
)

type IDetection interface {
	_onAreaEnter()
}
type Detection struct {
	Radius       float64
	Enabled      bool
	Shape        *cp.Shape
	_onAreaEnter func(*cp.PointQueryInfo)
}

func NewDetection(d float64, body *cp.Body, space *cp.Space, collisionFilter uint,
	areaEntered func(*cp.PointQueryInfo)) *Detection {

	shape := cp.NewCircle(body, d, cp.Vector{0, 0})
	filter := cp.NewShapeFilter(0, DETECTION_LAYER, collisionFilter)
	shape.SetCollisionType(4)
	shape.SetFilter(filter)
	los := &Detection{d, true, shape, areaEntered}
	shape.UserData = los
	space.AddShape(shape)
	return los
}

func (area *Detection) Update() {
	fmt.Println("detection updating")
	if area.Enabled {
		fmt.Printf("%v\n%d\n", area.Shape.Body().Position(),area.Radius)
		info := area.Shape.Space().PointQueryNearest(
			area.Shape.Body().Position(),
			area.Radius,
			area.Shape.Filter,
		)
		if info != nil {
			area._onAreaEnter(info)
		}
	}
}
