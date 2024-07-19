package main

import (
	. "gorpg/entities/player"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	input "github.com/quasilyte/ebitengine-input"
	"golang.org/x/image/colornames"
)

type Game struct {
	player      *Player
	inputSystem input.System
}

func NewGame() *Game {
	g := &Game{player: NewPlayer()}
	g.inputSystem.Init(input.SystemConfig{
		DevicesEnabled: input.AnyDevice,
	})
	g.player.AddInputHandler(&g.inputSystem)
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
