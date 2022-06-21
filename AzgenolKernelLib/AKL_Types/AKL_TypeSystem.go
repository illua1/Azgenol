package AKL_StaticTypes

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type DeleteObject interface {
	SetDelete(func())
}

type Context struct {
	Screen    *ebiten.Image
	Camera    *Camera
	Time      time.Duration
	StartTime time.Time
}
