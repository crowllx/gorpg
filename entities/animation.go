package entities

import (
	"fmt"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Animation struct {
	Frames            []*ebiten.Image
	CurrentFrameIndex int
	Advance           float64
	AnimationSpeed    float64
}

func (a *Animation) Update() {
	a.Advance += a.AnimationSpeed
	fmt.Println(a.Advance)
	fmt.Println(a.CurrentFrameIndex)
	a.CurrentFrameIndex = int(math.Floor(a.Advance))

	if a.CurrentFrameIndex >= len(a.Frames) {
		a.Reset()
	}
}
func (a *Animation) Draw() *ebiten.Image {
	return a.Frames[a.CurrentFrameIndex]
}

func (a *Animation) Reset() {
	a.CurrentFrameIndex = 0
	a.Advance = 0
}
