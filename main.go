package main
import (
	"fmt"
	"gorpg/audio"
	"gorpg/player"
	"gorpg/scenes"
	"gorpg/ui"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	input "github.com/quasilyte/ebitengine-input"
	"github.com/solarlune/ldtkgo"
	"github.com/yohamta/furex/v2"
	"golang.org/x/image/colornames"
)

type World struct {
	game *Game
}
type Game struct {
	player       *player.Player
	inputSystem  input.System
	scene        *scenes.Scene
	project      *ldtkgo.Project
	camera       *Camera
	Height       float64
	Width        float64
	ui           *furex.View
	audioContext *audio.Context
	bgm          *sound.AudioEmitter
}

func NewGame() *Game {
	gh, gw := 360.0, 640.0
	g := &Game{}
	g.inputSystem.Init(input.SystemConfig{
		DevicesEnabled: input.AnyDevice,
	})
	g.project = scenes.LoadProject("gorpg.ldtk")
	g.scene, g.player = scenes.LoadLevel(g.project.Levels[0])
	g.camera = NewCamera(0, 0, g.player)
	fmt.Printf("%T %T\n", g.scene, g.player)
	g.player.AddInputHandler(&g.inputSystem)
	g.Height = gh
	g.Width = gw
	g.ui = ui.LoadUI(int(gw), int(gh))
	// furex.Debug = true

	//sound
	g.bgm = sound.NewEmitter()
    g.bgm.NewPlayer("audio/music/Wav/Ambient 2.wav", true)
    g.bgm.Streams["audio/music/Wav/Ambient 2.wav"].Play()

	for _, l := range g.project.Levels {
		log.Printf("%v", l)
	}
	return g
}
func (g *Game) Update() error {
	g.inputSystem.Update()
	g.camera.Follow(g.Width, g.Height)
	g.ui.Update()
	g.scene.Update()
	// how can i move this into the code for ui/bar?
	hpbar, _ := g.ui.GetByID("hp")
	if hpbar != nil {
		//temporarily max hp is 25
		hp, _ := g.player.Status.Query("health")
		hpbar.Handler.(*ui.Bar).Val = float64(hp) / 25
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(colornames.Darkslateblue)
	g.scene.Draw(screen, g.camera.X, g.camera.Y)
	g.ui.Draw(screen)
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
