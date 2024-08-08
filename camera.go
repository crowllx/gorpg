package main

import "gorpg/player"

type Vector2 struct {
	X float64
	Y float64
}

type Camera struct {
	X, Y   float64
	target *player.Player
}

func NewCamera(x, y float64, target *player.Player) *Camera {
	return &Camera{
		x, y, target,
	}
}

func (cam *Camera) SetTarget(t *player.Player) {
	cam.target = t
}

func (cam *Camera) Follow(screenW, screenH float64) {
	if cam.target != nil {
		targetPos := cam.target.Body.Position()
		cam.X = -targetPos.X + screenW/2.0
		cam.Y = -targetPos.Y + screenH/2.0
	}
}
