package main

import (
	"gorpg/entities"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	player *entities.Player
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.player.UpdateVelocity(2, 0)
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.player.UpdateVelocity(-2, 0)
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.player.UpdateVelocity(0, -2)
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.player.UpdateVelocity(0, 2)
	}
	g.player.AnimatedSprite.CurrentAnimation().Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{120, 100, 255, 255})
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(g.player.X, g.player.Y)
	screen.DrawImage(g.player.CurrentFrame(), &opts)
}

func (g *Game) Layout(outsideWidth, outsideHieght int) (screenWdith, screenHeight int) {
	return 640, 360
}

func main() {
	ebiten.SetWindowSize(1600, 900)
	ebiten.SetWindowTitle("gorpg")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	player := entities.New()
	if err := ebiten.RunGame(&Game{player: player}); err != nil {
		log.Fatal(err)
	}
}
