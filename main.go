package main

import (
	"gorpg/entities/player"
	"gorpg/scenes"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	input "github.com/quasilyte/ebitengine-input"
	"golang.org/x/image/colornames"
)

type World struct {
	game *Game
}
type Game struct {
	player      *player.Player
	inputSystem input.System
	scene       *scenes.Scene
	Height      float64
	Width       float64
}

func NewGame() *Game {
	gh, gw := 360.0, 640.0
	g := &Game{player: player.New()}
	g.inputSystem.Init(input.SystemConfig{
		DevicesEnabled: input.AnyDevice,
	})
	g.player.AddInputHandler(&g.inputSystem)
	g.Height = gh
	g.Width = gw
	g.scene = (*scenes.Scene)(scenes.NewDebugScene(int(g.Width), int(g.Height), 16, 16, g.player))
	return g
}
func (g *Game) Update() error {
	g.inputSystem.Update()
	g.scene.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(colornames.Darkslateblue)
	g.scene.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHieght int) (screenWdith, screenHeight int) {
	return int(g.Width), int(g.Height)
}

func main() {
	ebiten.SetWindowSize(1600, 900)
	ebiten.SetWindowTitle("gorpg")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
