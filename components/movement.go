package components

import "github.com/jakecoffman/cp/v2"

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
