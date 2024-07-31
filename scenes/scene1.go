package scenes

import (
	// enemies "gorpg/entities/enemies/blueslime"
	"fmt"
	"gorpg/components"
	"gorpg/enemies"
	"gorpg/player"
	"gorpg/utils"
	"github.com/jakecoffman/cp/v2"
)

type DebugScene Scene

func NewDebugScene(w, h float64, cw, ch int, player *player.Player) *DebugScene {
	s := DebugScene{}
	s.debug = true
	s.space = cp.NewSpace()
	s.space.Iterations = 1
	fmt.Printf("%f %f %d %d", w, h, cw, ch)
	s.player = player
	s.player.AddSpace(s.space)
	s.space.EachShape(func(shape *cp.Shape) {
		fmt.Printf("%v\n", shape.BB().Center())
	})

	// dont think the walls 'bodies' are being added to space, but collision is working
	walls := []*cp.Shape{
		cp.NewSegment(cp.NewStaticBody(), cp.Vector{X: 0, Y: 0}, cp.Vector{X: w, Y: 0}, 1),
        cp.NewSegment(cp.NewStaticBody(), cp.Vector{X: 0, Y: 0}, cp.Vector{X: 0, Y: h}, 1),
	    cp.NewSegment(cp.NewStaticBody(), cp.Vector{X: w, Y: 0}, cp.Vector{X: w, Y: h}, 1),	
	    cp.NewSegment(cp.NewStaticBody(), cp.Vector{X: 0, Y: h}, cp.Vector{X: w, Y: h}, 1),	
	}
	walls = append(walls, cp.NewSegment(cp.NewStaticBody(), cp.Vector{X: 200, Y: 0}, cp.Vector{X: 200, Y: 180}, 1))

	fmt.Println(len(*s.space.ArrayForBodyType(cp.BODY_STATIC)))

	enemy := enemies.NewSlime(cp.Vector{X: 250, Y: 250}, s.space)
	s.enemies = append(s.enemies, enemy)
	enemy.AddToSpace(s.space)
	filter := cp.NewShapeFilter(0, components.ENVIRONMENT_LAYER, components.PLAYER_LAYER | components.ENEMY_LAYER)
	for _, w := range walls {
		w.SetFilter(filter)
        w.SetCollisionType(components.ENVIRONMENT_TYPE)
        w.UserData = utils.Collidable{}
		s.space.AddShape(w)

	}
	fmt.Println(len(*s.space.ArrayForBodyType(cp.BODY_KINEMATIC)))
	SetupCollisionHandlers(s.space)
    

	// enemy := enemies.NewSlime()
	// s.enemies = append(s.enemies, enemy)
	// enemy.AddToSpace(s.space)

	return &s
}
