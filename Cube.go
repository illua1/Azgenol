package main

import (
  "math"
	"github.com/hajimehoshi/ebiten/v2"
  volume "github.com/illua1/go-helpful/Volume"
  matrix "github.com/illua1/go-helpful/VectorMatrix"
)

var (
  Cube_Matrixes_Constant = volume.BoxContainerFaces[matrix.Matrix[float64, [3]float64, [3][3]float64]]{
    matrix.Matrix3x3[float64](),
    matrix.Matrix3x3[float64](),
    matrix.Rotate3x3_x[float64](math.Pi/2),
    matrix.Rotate3x3_x[float64](-math.Pi/2),
    matrix.Rotate3x3_y[float64](math.Pi/2),
    matrix.Rotate3x3_y[float64](-math.Pi/2),
  }
)

type Cube struct {
  This Box
  Faces [6]Serface
}

func NewCube(sx, sy, sz, px, py, pz float64, img *ebiten.Image)Cube{
  
  var this = NewBox(
    sx,
    sy,
    sz,
    NewPhysics(
      px,
      py,
      pz,
      0,
      true,
    ),
  )
  FaceSizes := this.Core.FaceArea()
  return Cube{
    This : this,
    Faces : [6]Serface{
      NewSerface(
        NewImageDrawer[float64](
          FaceSizes[0][0]*2,
          FaceSizes[0][1]*2,
          img,
        ),
      ),
      NewSerface(
        NewImageDrawer[float64](
          FaceSizes[1][0]*2,
          FaceSizes[1][1]*2,
          img,
        ),
      ),
      NewSerface(
        NewImageDrawer[float64](
          FaceSizes[2][0]*2,
          FaceSizes[2][1]*2,
          img,
        ),
      ),
      NewSerface(
        NewImageDrawer[float64](
          FaceSizes[3][0]*2,
          FaceSizes[3][1]*2,
          img,
        ),
      ),
      NewSerface(
        NewImageDrawer[float64](
          FaceSizes[4][0]*2,
          FaceSizes[4][1]*2,
          img,
        ),
      ),
      NewSerface(
        NewImageDrawer[float64](
          FaceSizes[5][0]*2,
          FaceSizes[5][1]*2,
          img,
        ),
      ),
    },
  }
}

func(cube *Cube) RenderCustom(append_draw RenderCallAppend, objectm ObjectMatrix, camera *Camera){
  global_location := cube.This.Dynamik.Location
  global_location.Sub(camera.Location)
  
  is_driw := [6]bool{
    camera.Matrix.A[2][2] < 0,
    camera.Matrix.A[2][2] >= 0,
    camera.Matrix.A[2][1] < 0,
    camera.Matrix.A[2][1] >= 0,
    camera.Matrix.A[2][0] < 0,
    camera.Matrix.A[2][0] >= 0,
  }
  
  for i, location := range cube.This.Core.FaceCentres() {
    if is_driw[i] {
      var location_global = matrix.Vector3[float64](float64(location.X), float64(location.Y), float64(location.Z))
      location_global.Sub(global_location)
      var location_ = camera.Matrix.MulVector(location_global)
      cube.Faces[i].MatrixUpdate(camera.MatrixInvert.Mull(Cube_Matrixes_Constant[i]))
      var img_draw = cube.Faces[i].ToImageDrawer()
      img_draw.GeoM.Translate(-location_.A[0], -location_.A[1])
      append_draw(
        img_draw,
        Value(cube.This.Dynamik.Location.A[0]),
      )
    }
  }
}