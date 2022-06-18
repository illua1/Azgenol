package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	matrix "github.com/illua1/go-helpful/VectorMatrix"
)

type Serface struct {
	ImageDrawer
	GeoM ebiten.GeoM
}

func NewSerface(imageDraver ImageDrawer) Serface {
	return Serface{ImageDrawer: imageDraver}
}

func (serface Serface) Draw(screen *ebiten.Image, geom ebiten.GeoM) {
	serface.GeoM.Concat(geom)
	serface.ImageDrawer.Draw(screen, serface.GeoM)
}

func (serface *Serface) MatrixUpdate(matrix matrix.Matrix[float64, [3]float64, [3][3]float64]) {
	matrix.Slise(0, 0, 2, 2).FillTo(serface.GeoM.SetElement)
}

func (serface Serface) GetGeoM(concat ebiten.GeoM) ebiten.GeoM {
	serface.GeoM.Concat(concat)
	return serface.ImageDrawer.GetGeoM(serface.GeoM)
}

func (serface Serface) ToImageDrawer() ImageDrawer {
	serface.ImageDrawer.GeoM.Concat(serface.GeoM)
	return serface.ImageDrawer
}
