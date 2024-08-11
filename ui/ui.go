package ui

import (
	_ "embed"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/furex/v2"
	"golang.org/x/image/colornames"
)


type Panel struct {
    Color color.Color
    OnClick func()
    mouseover bool
    pressed bool
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
func LoadUI(w,h int) *furex.View{
    ui := furex.Parse(mainHTML, &furex.ParseOptions{
        Width: w,
        Height: h,
        Components: furex.ComponentsMap{
            "player-status": &Panel{},
            "gauge": &Bar{
                Val: .7,  
            },
        },
    })
    return ui

}
