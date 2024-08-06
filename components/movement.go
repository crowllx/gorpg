package components

import (
	"fmt"
	"gorpg/utils"

	"github.com/jakecoffman/cp/v2"
)

func TerrainCheck(vel cp.Vector, normal cp.Vector, body *cp.Body) (bool, bool) {
	collisionX := false
	collisionY := false
	if (int(normal.X) > 0 && vel.X > 0) || (int(normal.X) < 0 && vel.X < 0) {
		collisionX = true
	}
	if (int(normal.Y) > 0 && vel.Y > 0) || (int(normal.Y) < 0 && vel.Y < 0) {
		collisionY = true

	}
	return collisionX, collisionY
}

// shape to be moved, x and y values to move the object this update (velocity)
// returns the x and y values that the object is able to move
func Move(shape *cp.Shape, dx, dy float64) (float64, float64) {
	shape.Space().ShapeQuery(shape, func(s *cp.Shape, cps *cp.ContactPointSet) {
		switch s.UserData.(type) {
		case utils.Collidable:
			fmt.Printf("%T\n", s.UserData)
			normal := cps.Normal
			colX, colY := TerrainCheck(cp.Vector{dx, dy}, normal, shape.Body())
			if colX {
				dx = 0
			}
			if colY {
				dy = 0
			}
		default:
		}
	})
	return dx, dy
}
