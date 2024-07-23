package scenes

import (
	"fmt"
	"gorpg/components"
	"gorpg/entities/enemies"
	"gorpg/entities/player"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/jakecoffman/cp/v2"
	"golang.org/x/image/colornames"
)

type Drawable interface {
	Draw(*ebiten.Image)
}
type Updateable interface {
	Update()
}
type Scene struct {
	player  *player.Player
	space   *cp.Space
	enemies []enemies.Enemy
	debug   bool
}

func (s *Scene) Draw(screen *ebiten.Image) {
	// for _, o := range s.space.Objects() {
	// 	switch o.Data.(type) {
	// 	case Drawable:
	// 		o.Data.(Drawable).Draw(screen)
	// 	default:
	// 	}
	// }
	s.space.EachBody(func(body *cp.Body) {
		switch body.UserData.(type) {
		case Drawable:
			body.UserData.(Drawable).Draw(screen)
		default:
		}
	})
	if s.debug {
		debug(screen, s)
		s.debugCollisions(screen)
	}
}

func (s *Scene) Update() {
	s.space.Step(1.0)
	s.space.EachBody(func(body *cp.Body) {
		switch body.UserData.(type) {
		case Updateable:
			body.UserData.(Updateable).Update()
		default:
		}
	})
	// for _, o := range s.space.Objects() {
	// 	o.Update()
	// 	switch o.Data.(type) {
	// 	case Updateable:
	// 		o.Data.(Updateable).Update()
	// 	default:
	// 	}
	// }
}

func (s *Scene) debugCollisions(screen *ebiten.Image) {
	s.space.EachShape(func(shape *cp.Shape) {
		switch shape.Class.(type) {
		case *cp.Circle:
			var col color.Color
			switch shape.UserData.(type) {
			case *components.Detection:
				var color colorm.ColorM
				color.Scale(1, 1, 1, .2)
				col = color.Apply(colornames.Green)
			default:
				var color colorm.ColorM
				color.Scale(1, 1, 1, .2)
				col = color.Apply(colornames.Crimson)
				// fmt.Printf("%T\n", shape.UserData)
			}
			// fmt.Printf("shape center %v\n", shape.BB().Center())
			pos := shape.BB().Center()
			rad := shape.BB().R - shape.BB().Center().X
			vector.DrawFilledCircle(screen, float32(pos.X), float32(pos.Y), float32(rad), col, false)
		case *cp.Segment:
			var color colorm.ColorM
			color.Scale(1, 1, 1, .2)
			v1 := shape.Class.(*cp.Segment).A()
			v2 := shape.Class.(*cp.Segment).B()
			vector.StrokeLine(screen, float32(v1.X), float32(v1.Y), float32(v2.X), float32(v2.Y), 4, color.Apply(colornames.Black), false)
		default:
		}
	})
}

func debug(screen *ebiten.Image, s *Scene) {
	hp, _ := s.player.Status.Query("health")
	mp, _ := s.player.Status.Query("mana")
	var ehp, emp int
	if len(s.enemies) > 0 {
		ehp, _ = s.enemies[0].Query("health")
		emp, _ = s.enemies[0].Query("mana")
	}
	ebitenutil.DebugPrint(screen, fmt.Sprintf(`
	  player hp: %v mp: %v
	  enemy  hp: %v mp: %v
	  `, hp, mp, ehp, emp))
}
