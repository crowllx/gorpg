package scenes

import (
	"fmt"
	"gorpg/entities/enemies"
	"gorpg/entities/player"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/solarlune/resolv"
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
	space   *resolv.Space
	enemies []enemies.Enemy
	debug   bool
}

func (s *Scene) Draw(screen *ebiten.Image) {
	for _, o := range s.space.Objects() {
		switch o.Data.(type) {
		case Drawable:
			o.Data.(Drawable).Draw(screen)
		default:
		}
	}

	if s.debug {
		debug(screen, s)
		s.debugCollisions(screen)
	}
}

func (s *Scene) Update() {
	for _, o := range s.space.Objects() {
		o.Update()
		switch o.Data.(type) {
		case Updateable:
			o.Data.(Updateable).Update()
		default:
		}
	}
}

func (s *Scene) debugCollisions(screen *ebiten.Image) {
	// debug drawing
	for _, o := range s.space.Objects() {
		if o.Shape != nil {
			switch o.Shape.(type) {
			case *resolv.Circle:
				fmt.Println("winning")
			case *resolv.ConvexPolygon:
				pos, size := o.Shape.Bounds()
				vector.DrawFilledCircle(screen, float32(pos.X), float32(pos.Y), float32(size.X), colornames.Black, false)
				// opts.ColorScale.ScaleAlpha(.2)
			}
		}
		if o.HasTags("hurt") {
			pos := o.Position
			size := o.Size
			//debug img
			img := ebiten.NewImage(int(size.X), int(size.Y))
			img.Fill(colornames.Lightcoral)
			opts := ebiten.DrawImageOptions{}
			opts.GeoM.Translate(pos.X, pos.Y)
			opts.ColorScale.ScaleAlpha(.3)
			screen.DrawImage(img, &opts)
		} else if o.HasTags("hit") {
			pos := o.Position
			size := o.Size
			//debug img
			img := ebiten.NewImage(int(size.X), int(size.Y))
			img.Fill(colornames.Green)
			opts := ebiten.DrawImageOptions{}
			opts.GeoM.Translate(pos.X, pos.Y)
			opts.ColorScale.ScaleAlpha(.3)
			screen.DrawImage(img, &opts)
		}
	}
}

func debug(screen *ebiten.Image, s *Scene) {
	hp, _ := s.player.Status.Query("health")
	mp, _ := s.player.Status.Query("mana")
	var ehp, emp int
	if len(s.enemies) > 0 {
		ehp, _ = s.enemies[0].Query("health")
		emp, _ = s.enemies[0].Query("mana")
	}
	o := s.player.Object.Position
	osize := s.player.Object.Size
	imgb := s.player.Sprite().CurrentImg.Draw().Bounds().Size()
	ebitenutil.DebugPrint(screen, fmt.Sprintf(`
		objs: %d
		player obj pos: %v
		player opbj size: %v
		player img bounds: %v
		player health: %d mana: %d
		enemy hp: %d mp: %d`, len(s.space.Objects()), o, osize, imgb, hp, mp, ehp, emp))
}
