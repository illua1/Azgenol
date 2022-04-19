package main

import (
	"github.com/hajimehoshi/ebiten/v2"
  matrix "github.com/illua1/go-helpful/VectorMatrix"
)

type Serface struct {
  Src ImageDrawer
  GeoM ebiten.GeoM
}

func NewSerface(imageDraver ImageDrawer)Serface {
  return Serface{
    Src : imageDraver,
  }
}

func (serface Serface)Draw(screen *ebiten.Image, geom ebiten.GeoM){
  serface.GeoM.Concat(geom)
  serface.Src.Draw(screen, serface.GeoM)
}

func (serface *Serface)MatrixUpdate(matrix matrix.Matrix[float64, [3]float64, [3][3]float64]){
  matrix.Slise(0,0,2,2).FillTo(serface.GeoM.SetElement)
}