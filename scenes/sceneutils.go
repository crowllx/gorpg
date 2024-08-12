package scenes

import (
	"embed"
	"fmt"
	"gorpg/components"
	"gorpg/enemies"
	"gorpg/player"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/jakecoffman/cp/v2"
	"github.com/solarlune/ldtkgo"
	"golang.org/x/image/colornames"
)

type Drawable interface {
	Draw(*ebiten.Image, float64, float64)
}
type Updateable interface {
	Update()
}
type Scene struct {
	player  *player.Player
	space   *cp.Space
	tileSet *ebiten.Image
	tiles   []*ldtkgo.Tile
	enemies []enemies.Enemy
	debug   bool
}

func (s *Scene) Draw(screen *ebiten.Image, camX, camY float64) {
	for _, t := range s.tiles {
		x, y := t.Src[0], t.Src[1]
		sub := s.tileSet.SubImage(image.Rect(x, y, x+16, y+16))
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64(t.Position[0]), float64(t.Position[1]))
		opts.GeoM.Translate(camX, camY)
		screen.DrawImage(sub.(*ebiten.Image), opts)
	}
	s.space.EachBody(func(body *cp.Body) {
		switch body.UserData.(type) {
		case enemies.Enemy:
			ePos := body.Position()
			distance := ePos.Distance(s.player.Body.Position())
			if distance <= 40 {
				e := body.UserData.(enemies.Enemy)
				bar, x, y := e.DrawHealthBar()
				opts := &ebiten.DrawImageOptions{}
				opts.GeoM.Translate(x+camX, y+camY)
				screen.DrawImage(bar, opts)
			}
			body.UserData.(enemies.Enemy).Draw(screen, camX, camY)
		case Drawable:
			body.UserData.(Drawable).Draw(screen, camX, camY)
		default:
		}
	})
	if s.debug {
		debug(screen, s)
		s.debugCollisions(screen, camX, camY)
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
}

var assets embed.FS

func (s *Scene) debugCollisions(screen *ebiten.Image, camX, camY float64) {
	s.space.EachShape(func(shape *cp.Shape) {
		switch shape.Class.(type) {
		case *cp.Circle:
			var col color.Color
			switch shape.UserData.(type) {
			case *components.Detection:
				var color colorm.ColorM
				color.Scale(1, 1, 1, .2)
				col = color.Apply(colornames.Green)
			case *components.HurtBox:
				var color colorm.ColorM
				color.Scale(1, 1, 1, .2)
				col = color.Apply(colornames.Blueviolet)
			default:
				var color colorm.ColorM
				color.Scale(1, 1, 1, .2)
				col = color.Apply(colornames.Crimson)
				// fmt.Printf("%T\n", shape.UserData)
			}
			// fmt.Printf("shape center %v\n", shape.BB().Center())
			pos := shape.BB().Center()
			pos.X += camX
			pos.Y += camY
			rad := shape.BB().R - shape.BB().Center().X
			vector.DrawFilledCircle(screen, float32(pos.X), float32(pos.Y), float32(rad), col, false)
		case *cp.Segment:
			var color colorm.ColorM
			color.Scale(1, 1, 1, .2)
			v1 := shape.Class.(*cp.Segment).A()
			v2 := shape.Class.(*cp.Segment).B()
			v1.Add(cp.Vector{camX, camY})
			v2.Add(cp.Vector{camX, camY})
			vector.StrokeLine(screen, float32(v1.X), float32(v1.Y), float32(v2.X), float32(v2.Y), 4, color.Apply(colornames.Black), false)
		default:
		}
	})
}

func debug(screen *ebiten.Image, s *Scene) {
	hp, _, _ := s.player.Status.Query("health")
	mp, _, _ := s.player.Status.Query("mana")
	var ehp, emp int
	if len(s.enemies) > 0 {
		ehp, _, _ = s.enemies[0].Query("health")
		emp, _, _ = s.enemies[0].Query("mana")
	}
	ebitenutil.DebugPrint(screen, fmt.Sprintf(`
	  player
		hp: %d
		mp: %d
	  enemy
		hp: %d
		mp: %d
	  `, hp, mp, ehp, emp))
}
