package AKL_StaticTypes

import (
  "github.com/hajimehoshi/ebiten/v2"
  matrix "github.com/illua1/go-helpful/VectorMatrix"
)

func ProjectGeoM(matrix matrix.Matrix[float64, [3]float64, [3][3]float64]) (ret ebiten.GeoM){
  matrix.Slise(0,0,2,2).FillTo(ret.SetElement)
  return
}