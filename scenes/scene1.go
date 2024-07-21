package scenes

import (
	// enemies "gorpg/entities/enemies/blueslime"
	"fmt"
	"gorpg/entities/player"

	"github.com/jakecoffman/cp/v2"
)

type DebugScene Scene

func NewDebugScene(w, h float64, cw, ch int, player *player.Player) *DebugScene {
	s := DebugScene{}
	s.debug = true
	s.space = cp.NewSpace()
	s.space.Iterations = 1
	fmt.Printf("%d %d %d %D", w, h, cw, ch)
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
	fmt.Println(len(*s.space.ArrayForBodyType(cp.BODY_STATIC)))

	for _, w := range walls {
		s.space.AddShape(w)
	}
	// s.space.AddShape(shape)
	fmt.Println(len(*s.space.ArrayForBodyType(cp.BODY_KINEMATIC)))

	// enemy := enemies.NewSlime()
	// s.enemies = append(s.enemies, enemy)
	// enemy.AddToSpace(s.space)

	return &s
}
