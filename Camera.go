package main

import (
  matrix "github.com/illua1/go-helpful/VectorMatrix"
)

type Camera struct {
  Location matrix.Vector[float64, [3]float64]
  Matrix matrix.Matrix[float64, [3]float64, [3][3]float64]
  MatrixInvert matrix.Matrix[float64, [3]float64, [3][3]float64]
}

func NewCamera(x, y, z float64)Camera{
  return Camera{
    Location : matrix.Vector[float64, [3]float64]{
      [3]float64{x, y, z},
    },
  }
}

func(camera *Camera)SetMatrix(matrix matrix.Matrix[float64, [3]float64, [3][3]float64]){
  camera.Matrix = matrix
  camera.MatrixInvert = matrix.Invert()
}