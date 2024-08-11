package ui

import (
	"fmt"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/furex/v2"
	"golang.org/x/image/colornames"
)

type Bar struct {
	Color     color.Color
	Val       float64
	OnClick   func()
	mouseover bool
	pressed   bool
}

func (p *Bar) Draw(screen *ebiten.Image, frame image.Rectangle, view *furex.View) {
	fgImg := ebiten.NewImage(int(float64(frame.Dx())*p.Val), frame.Dy())
	bgImg := ebiten.NewImage(frame.Dx(), frame.Dy())
	bgImg.Fill(colornames.Gray)
	switch view.Attrs["color"] {
	case "red":
		fgImg.Fill(colornames.Darkred)
	case "blue":
		fgImg.Fill(colornames.Darkblue)
	default:
		fgImg.Fill(colornames.Gray)
	}
	opts := ebiten.DrawImageOptions{}
	fmt.Printf("%v\n", frame.Max)
	fmt.Printf("%v\n", frame.Min)
	point := frame.Min
	bgImg.DrawImage(fgImg, &opts)
	opts.GeoM.Translate(float64(point.X), float64(point.Y))
	screen.DrawImage(bgImg, &opts)
}
