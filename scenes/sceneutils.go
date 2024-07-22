package scenes

import (
	"gorpg/entities/enemies"
	"gorpg/entities/player"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
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
	s.space.EachBody(func(body *cp.Body) {
		body.EachShape(func(shape *cp.Shape) {
			switch shape.Class.(type) {
			case *cp.Circle:
				// fmt.Printf("shape center %v\n", shape.BB().Center())
				var color colorm.ColorM
				color.Scale(1, 1, 1, .2)
				pos := shape.BB().Center()
				vector.DrawFilledCircle(screen, float32(pos.X), float32(pos.Y), 16.0, color.Apply(colornames.Crimson), false)
			default:
			}
		})
	})
}

func debug(screen *ebiten.Image, s *Scene) {
	// hp, _ := s.player.Status.Query("health")
	// mp, _ := s.player.Status.Query("mana")
	// var ehp, emp int
	// if len(s.enemies) > 0 {
	// 	ehp, _ = s.enemies[0].Query("health")
	// 	emp, _ = s.enemies[0].Query("mana")
	// }
	// ebitenutil.DebugPrint(screen, fmt.Sprintf(`
	//   player hp: %v mp: %v
	//   enemy  hp: %v mp: %v
	//   `, hp, mp, ehp, emp))
	// osize := s.player.Body
	// imgb := s.player.Sprite().CurrentImg.Draw().Bounds().Size()
	// ebitenutil.DebugPrint(screen, fmt.Sprintf(`
	// 	objs: %d
	// 	player obj pos: %vV
	// player opbj size: %v
	// player img bounds: %v
	// player health: %d mana: %d
}
