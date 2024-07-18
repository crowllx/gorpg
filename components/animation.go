package components

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Animation struct {
	Frames         *ebiten.Image
	Index          int
	Advance        float64
	AnimationSpeed float64
	FrameCount     int
	w              int
	h              int
}

func (a *Animation) Update() {
	a.Advance += a.AnimationSpeed
	a.Index = int(math.Floor(a.Advance))

	if a.Index >= a.FrameCount {
		a.Reset()
	}
}
func (a *Animation) Draw() *ebiten.Image {
	offset := a.Index * a.w
	dimensions := image.Rect(offset, 0, offset+a.w, a.h)
	return a.Frames.SubImage(dimensions).(*ebiten.Image)
}

func (a *Animation) Reset() {
	a.Index = 0
	a.Advance = 0
}
