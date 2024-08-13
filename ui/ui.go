package ui

import (
	_ "embed"
	"gorpg/player"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/furex/v2"
	"golang.org/x/image/colornames"
)

type Panel struct {
	Color     color.Color
	OnClick   func()
	mouseover bool
	pressed   bool
}

func (p *Panel) Draw(screen *ebiten.Image, frame image.Rectangle, view *furex.View) {
	img := ebiten.NewImage(frame.Dx(), frame.Dy())
	img.Fill(colornames.Darkcyan)
	opts := ebiten.DrawImageOptions{}
	point := frame.Min
	opts.GeoM.Translate(float64(point.X), float64(point.Y))
	screen.DrawImage(img, &opts)
}

//go:embed main.html
var mainHTML string

func SetGameUIValues(ui *furex.View, player *player.Player) {
	hpBar, _ := ui.GetByID("hp")
	manaBar, _ := ui.GetByID("mana")
	hp, maxHp, _ := player.Status.Query("health")
	mp, maxMp, _ := player.Status.Query("mana")
    hpBar.Handler.(*Bar).Val = float64(hp)
    hpBar.Handler.(*Bar).Max = float64(maxHp)
    manaBar.Handler.(*Bar).Val = float64(mp)
    manaBar.Handler.(*Bar).Max = float64(maxMp)
}
func LoadUI(w, h int, player *player.Player) *furex.View {
	ui := furex.Parse(mainHTML, &furex.ParseOptions{
		Width:  w,
		Height: h,
		Components: furex.ComponentsMap{
			"player-status": &Panel{},
			"gauge": func() furex.Handler {
				return &Bar{
					Val: 1,
					Max: 1,
				}
			},
		},
	})

	// initalize gauge values
	hp, _ := ui.GetByID("hp")
	mana, _ := ui.GetByID("mana")
	hp.Handler.(*Bar).Color = colornames.Darkred
	mana.Handler.(*Bar).Color = colornames.Darkblue
    SetGameUIValues(ui, player)

	return ui

}
