package utils

import (
	"fmt"

	"github.com/jakecoffman/cp/v2"
)

type BasicArea struct {
	maxDistance float64
	enabled     bool
	shape       *cp.Shape
}

func NewArea(dist float64, enabled bool, body *cp.Body) {
	if body != nil {
		// shape := cp.NewCircle(body, dist, cp.Vector{0, 0})
	} else {
		// shape := cp.NewCircle(nil, dist, cp.Vector{0, 0})
	}

}
func (b *BasicArea) Update() {
	b.shape.Space().ShapeQuery(b.shape, func(s *cp.Shape, ps *cp.ContactPointSet) {
		fmt.Printf(`
			shape: %T
			ps: %v`, s, ps)
	})
}
