package main

import (
	"fmt"
	. "gorpg/entities/player"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	input "github.com/quasilyte/ebitengine-input"
	"github.com/solarlune/resolv"
	"golang.org/x/image/colornames"
)

type World struct {
	game *Game
}
type Game struct {
	player      *Player
	inputSystem input.System
	space       *resolv.Space
	Height      float64
	Width       float64
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
	for _, obj := range g.space.Objects() {
		fmt.Printf("%v %v\n", obj.Position, obj.Size)
	}
	return g
}

func (g *Game) Update() error {
	g.inputSystem.Update()
	g.player.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(colornames.Darkslateblue)
	opts := ebiten.DrawImageOptions{}
	sprite := g.player.Sprite()
	opts.GeoM.Translate(g.player.X, g.player.Y)
	screen.DrawImage(sprite.CurrentImg.Draw(), &opts)
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
