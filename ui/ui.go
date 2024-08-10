package ui

import (
	_ "fmt"
	"image"
	_ "image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
)

type Drawable interface {
	Draw(*ebiten.Image)
}
type UI struct {
	elements []Drawable
}

type ProgressBar struct {
	bgImage      *ebiten.Image
	fullImage    *ebiten.Image
	currentImage *ebiten.Image
	maxW, maxH   int
	val, maximum float64
}

func (pg *ProgressBar) SetVal(v float64) {
	pg.val = v
	width := pg.val / pg.maximum * float64(pg.maxW)
	pg.currentImage = pg.fullImage.SubImage(image.Rect(0, 0, int(width), pg.maxH)).(*ebiten.Image)
}

func NewProgressBar(x, y int, m float64, c color.Color) *ProgressBar {
	bgImage := ebiten.NewImage(x, y)
	bgImage.Fill(colornames.Gray)
	fullImage := ebiten.NewImage(x, y)
	fullImage.Fill(c)

	return &ProgressBar{
		bgImage:      bgImage,
		fullImage:    fullImage,
		currentImage: fullImage,
		maxW:         x,
		maxH:         y,
		val:          m,
		maximum:      m,
	}
}

func (pg *ProgressBar) Draw(screen *ebiten.Image) {
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(screen.Bounds().Dx())/2-100, float64(screen.Bounds().Dy())/2+150)
	screen.DrawImage(pg.bgImage, &opts)
	screen.DrawImage(pg.currentImage, &opts)
}
