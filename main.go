package main

import (
	"fmt"
	. "gorpg/entities/enemies"
	. "gorpg/entities/player"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	input "github.com/quasilyte/ebitengine-input"
	"github.com/solarlune/resolv"
	"golang.org/x/image/colornames"
)

type World struct {
	game *Game
}
type Game struct {
	player          *Player
	inputSystem     input.System
	space           *resolv.Space
	Height          float64
	Width           float64
	enemy           *Enemy
	debugCollisions *ebiten.Image
}

func NewGame() *Game {
	gh, gw := 360.0, 640.0
	space := resolv.NewSpace(int(gw), int(gh), 16, 16)
	g := &Game{player: NewPlayer(space)}
	g.space = space
	g.inputSystem.Init(input.SystemConfig{
		DevicesEnabled: input.AnyDevice,
	})
	g.player.AddInputHandler(&g.inputSystem)
	g.Height = gh
	g.Width = gw

	// collisions
	g.space.Add(
		resolv.NewObject(0, 0, 640, 16, "solid"),
		resolv.NewObject(0, 360-16, 640, 16, "solid"),
		resolv.NewObject(0, 16, 16, 360-32, "solid"),
		resolv.NewObject(640-16, 16, 16, 360-32, "solid"),
	)
	g.enemy = NewEnemy(g.space)

	for _, o := range g.space.Objects() {
		if o.HasTags("hit", "hurt") {
			fmt.Printf("%v %v\n", o.Position, o.Size)
		}
	}
	return g
}

func (g *Game) Update() error {
	g.inputSystem.Update()
	g.player.Update()
	g.enemy.Update()

	for _, o := range g.space.Objects() {
		o.Update()
	}
	x, y, w, h := g.enemy.Collider().BoundsToSpace(0, 0)
	ebitenutil.DebugPrint(ebiten.NewImage(16, 16), fmt.Sprintf("enemy hit box: %d %d %d %d", x, y, w, h))
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(colornames.Darkslateblue)
	opts := ebiten.DrawImageOptions{}
	sprite := g.player.Sprite()
	opts.GeoM.Translate(g.player.X, g.player.Y)
	screen.DrawImage(sprite.CurrentImg.Draw(), &opts)
	opts = ebiten.DrawImageOptions{}
	x := g.enemy.Sprite()
	x.Draw(screen)
	// debug drawing
	for _, o := range g.space.Objects() {
		if o.HasTags("hurt") {
			pos := o.Position
			size := o.Size
			//debug img
			img := ebiten.NewImage(int(size.X), int(size.Y))
			img.Fill(colornames.Lightcoral)
			g.debugCollisions = img
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
			g.debugCollisions = img
			opts := ebiten.DrawImageOptions{}
			opts.GeoM.Translate(pos.X, pos.Y)
			opts.ColorScale.ScaleAlpha(.3)
			screen.DrawImage(img, &opts)

		}
	}

	// enemy

}

func (g *Game) Layout(outsideWidth, outsideHieght int) (screenWdith, screenHeight int) {
	return 640, 360
}

func main() {
	ebiten.SetWindowSize(1600, 900)
	ebiten.SetWindowTitle("gorpg")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
