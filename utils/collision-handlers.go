package utils

import (
	"fmt"
	"gorpg/components"
	"gorpg/entities/enemies"
	"gorpg/entities/player"

	"github.com/jakecoffman/cp/v2"
)

func SetupCollisionHandlers(space *cp.Space) {
	wild := func(_ *cp.Arbiter, _ *cp.Space, _ interface{}) bool {
		return false
	}
	hitBox := func(arb *cp.Arbiter, _ *cp.Space, _ interface{}) bool {
		a, b := arb.Shapes()
		hurt := a.UserData.(*components.HurtBox)
		if !hurt.HitCheck(b) {
			return false
		}
		switch b.UserData.(type) {
		case enemies.Enemy:
			b.UserData.(enemies.Enemy).Modify("health", hurt.Value())
		case *player.Player:
			res, _ := b.UserData.(player.Player).Status.Modify("health", hurt.Value())
			fmt.Println(res)

		default:
		}
		// fmt.Printf("a %T\nb %T\n", a.UserData, b.UserData)
		return false
	}
	space.NewWildcardCollisionHandler(components.HIT_TYPE).PreSolveFunc = wild
	space.NewWildcardCollisionHandler(components.DETECTION_TYPE).PreSolveFunc = wild
	space.NewCollisionHandler(components.HIT_TYPE, components.PLAYER_TYPE).PreSolveFunc = hitBox
	space.NewCollisionHandler(components.HIT_TYPE, components.ENEMY_TYPE).PreSolveFunc = hitBox

}
