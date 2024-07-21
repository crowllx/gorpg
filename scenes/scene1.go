package scenes

import (
	enemies "gorpg/entities/enemies/blueslime"
	"gorpg/entities/player"

	"github.com/solarlune/resolv"
)

type DebugScene Scene

func NewDebugScene(w, h, cw, ch int, player *player.Player) *DebugScene {
	s := DebugScene{}
	s.debug = true
	space := resolv.NewSpace(w, h, cw, ch)
	s.space = space
	s.player = player
	player.AddToSpace(s.space)

	enemy := enemies.NewSlime()
	s.enemies = append(s.enemies, enemy)
	enemy.AddToSpace(s.space)

	// add walls
	s.space.Add(
		resolv.NewObject(0, 0, 640, 16, "solid"),
		resolv.NewObject(0, 360-16, 640, 16, "solid"),
		resolv.NewObject(0, 16, 16, 360-32, "solid"),
		resolv.NewObject(640-16, 16, 16, 360-32, "solid"),
	)
	for _, o := range s.space.Objects() {
		o.SetShape(resolv.NewRectangle(o.Position.X, o.Position.Y, o.Size.X, o.Size.Y))
	}
	return &s
}
