package scenes

import (
	"fmt"

	"github.com/jakecoffman/cp/v2"
)

const PLAYER_LAYER = 1
const ENVIRONMENT_LAYER = 2
const HIT_LAYER = 4
const ENEMIES_LAYER = 8

func setupCollisionHandlers(space *cp.Space) {

	hitHandler := space.NewCollisionHandler(4, 3)
	hitHandler.PreSolveFunc = func(arb *cp.Arbiter, _ *cp.Space, _ interface{}) bool {
		a, b := arb.Shapes()
		fmt.Printf("a %T\nb %T\n", a.UserData, b.UserData)
		return false
	}

}
