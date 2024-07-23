package components

import (
	"fmt"

	"github.com/jakecoffman/cp/v2"
)

type Area interface {
	Update()
	onEnter(*cp.Shape)
}
type BasicArea struct {
	maxDistance float64
	enabled     bool
	shape       *cp.Shape
}

func NewBasicArea(dist float64, enabled bool, body *cp.Body) *BasicArea {
	var shape *cp.Shape
	if body != nil {
		shape = cp.NewCircle(body, dist, cp.Vector{0, 0})
	} else {
		shape = cp.NewCircle(nil, dist, cp.Vector{0, 0})
	}
	return &BasicArea{dist, enabled, shape}
}

func (b *BasicArea) Update() {
	b.shape.Space().ShapeQuery(b.shape, func(s *cp.Shape, ps *cp.ContactPointSet) {
		fmt.Printf(`
			shape: %T
			ps: %v`, s, ps)
	})
}
