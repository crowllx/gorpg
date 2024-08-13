package ui

import (
	"bytes"
	"fmt"
	"gorpg/assets/fonts"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/yohamta/furex/v2"
	"golang.org/x/image/colornames"
)

type Bar struct {
	Color     color.Color
	Max       float64
	Val       float64
	OnClick   func()
	mouseover bool
	pressed   bool
	Update    func() (float64, float64)
}

func (p *Bar) Draw(screen *ebiten.Image, frame image.Rectangle, view *furex.View) {
	fgImg := ebiten.NewImage(int(float64(frame.Dx())*p.Val/p.Max), frame.Dy())
	bgImg := ebiten.NewImage(frame.Dx(), frame.Dy())
	bgImg.Fill(colornames.Gray)
    fgImg.Fill(p.Color)
	opts := ebiten.DrawImageOptions{}
	point := frame.Min
	bgImg.DrawImage(fgImg, &opts)
	source, _ := text.NewGoTextFaceSource(bytes.NewReader(fonts.FiraCode_ttf))
	face := &text.GoTextFace{
		Source: source,
		Size:   12,
	}
	textOpts := &text.DrawOptions{}
	text.Draw(bgImg, fmt.Sprintf("%d / %d", int(p.Val), int(p.Max)), face, textOpts)
	opts.GeoM.Translate(float64(point.X), float64(point.Y))
	screen.DrawImage(bgImg, &opts)
}
