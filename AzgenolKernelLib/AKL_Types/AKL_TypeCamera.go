package AKL_StaticTypes

import (
  matrix "github.com/illua1/go-helpful/VectorMatrix"
)

type Camera struct {

  Location Vector3
  
  Matrix Matrix3
  
  MatrixInvert Matrix3
}

func NewCamera(x, y, z float64)Camera{
  return Camera{
    Location : NewVector3(x, y, z),
    Matrix : Matrix3{matrix.Matrix3x3[float64]()},
    MatrixInvert : Matrix3{matrix.Matrix3x3[float64]()},
  }
}

func(camera *Camera)SetMatrix(matrix Matrix3){
  camera.Matrix = matrix
  camera.MatrixInvert.Matrix = matrix.Invert()
}

func(camera *Camera)SetAngle(x,y,z float64){
  camera.Matrix.Matrix = matrix.Rotate3x3_ZYX[float64]([3]float64{x,y,z})
  camera.MatrixInvert.Matrix = camera.Matrix.Invert()
}