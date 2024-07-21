package scenes

import (
	"fmt"
	"gorpg/entities/enemies"
	"gorpg/entities/player"

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
	// debug drawing
	// for _, o := range s.space.Objects() {
	// 	if o.Shape != nil {
	// 		switch o.Shape.(type) {
	// 		case *resolv.Circle:
	// 			fmt.Println("winning")
	// 		case *resolv.ConvexPolygon:
	// 			pos, size := o.Shape.Bounds()
	// 			vector.DrawFilledCircle(screen, float32(pos.X), float32(pos.Y), float32(size.X), colornames.Black, false)
	// 			// opts.ColorScale.ScaleAlpha(.2)
	// 		}
	// 	}
	// 	if o.HasTags("hurt") {
	// 		pos := o.Position
	// 		size := o.Size
	// 		//debug img
	// 		img := ebiten.NewImage(int(size.X), int(size.Y))
	// 		img.Fill(colornames.Lightcoral)
	// 		opts := ebiten.DrawImageOptions{}
	// 		opts.GeoM.Translate(pos.X, pos.Y)
	// 		opts.ColorScale.ScaleAlpha(.3)
	// 		screen.DrawImage(img, &opts)
	// 	} else if o.HasTags("hit") {
	// 		pos := o.Position
	// 		size := o.Size
	// 		//debug img
	// 		img := ebiten.NewImage(int(size.X), int(size.Y))
	// 		img.Fill(colornames.Green)
	// 		opts := ebiten.DrawImageOptions{}
	// 		opts.GeoM.Translate(pos.X, pos.Y)
	// 		opts.ColorScale.ScaleAlpha(.3)
	// 		screen.DrawImage(img, &opts)
	// 	}
	// }
}

func debug(screen *ebiten.Image, s *Scene) {
	// hp, _ := s.player.Status.Query("health")
	// mp, _ := s.player.Status.Query("mana")
	// var ehp, emp int
	// if len(s.enemies) > 0 {
	// 	ehp, _ = s.enemies[0].Query("health")
	// 	emp, _ = s.enemies[0].Query("mana")
	// }
	ebitenutil.DebugPrint(screen, fmt.Sprintf(`
   pos %v
   vel %v
   dir %v`, s.player.Body.Position(), s.player.Body.Velocity(), s.player.Body.Angle()))
	// osize := s.player.Body
	// imgb := s.player.Sprite().CurrentImg.Draw().Bounds().Size()
	// ebitenutil.DebugPrint(screen, fmt.Sprintf(`
	// 	objs: %d
	// 	player obj pos: %vV
	// player opbj size: %v
	// player img bounds: %v
	// player health: %d mana: %d
}
