package main

import (
  "math"
	"github.com/hajimehoshi/ebiten/v2"
  matrix "github.com/illua1/go-helpful/VectorMatrix"
)

var (
  Cube_Matrixes_Constant = [6]matrix.Matrix[float64, [3]float64, [3][3]float64]{
    matrix.Matrix3x3[float64](),
    matrix.Matrix3x3[float64](),
    matrix.Rotate3x3_x[float64](math.Pi/2),
    matrix.Rotate3x3_x[float64](-math.Pi/2),
    matrix.Rotate3x3_y[float64](math.Pi/2),
    matrix.Rotate3x3_y[float64](-math.Pi/2),
  }
)

type Cube struct {
  This Boxe
  Faces [6]Serface
}

func NewCube(sx, sy, sz, px, py, pz float64, img *ebiten.Image)Cube{
  
  var this = NewBoxe(
    sx,
    sy,
    sz,
    NewPhysics(
      px,
      py,
      pz,
      0,
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

func (cube *Cube)Draw(screen *ebiten.Image, GeoM ebiten.GeoM, camera *Camera){
  
  var (
    screen_location_geom = GeoM
  )
  
  location := cube.This.Dynamik.Location
  location.Sub(camera.Location)
  global_location := camera.Matrix.MulVector(location)
  
  
  is_driw := [6]bool{
    camera.Matrix.A[2][2] < 0,
    camera.Matrix.A[2][2] >= 0,
    camera.Matrix.A[2][1] < 0,
    camera.Matrix.A[2][1] >= 0,
    camera.Matrix.A[2][0] < 0,
    camera.Matrix.A[2][0] >= 0,
  }
  {
    sx := GeoM.Element(0,0)*global_location.A[0] + GeoM.Element(1,0)*global_location.A[1]
    sy := GeoM.Element(0,1)*global_location.A[0] + GeoM.Element(1,1)*global_location.A[1]
    screen_location_geom.Translate(sx, sy)
  }
  for i, location := range cube.This.Core.FaceCentres() {
    if is_driw[i]{
      var location_ = camera.Matrix.MulVector(matrix.Vector[float64, [3]float64]{
        A:[3]float64{
          float64(location.X),
          float64(location.Y),
          float64(location.Z),
        },
      })
      {
        GeoM = screen_location_geom
        lx := GeoM.Element(0,0)*location_.A[0] + GeoM.Element(1,0)*location_.A[1]
        ly := GeoM.Element(0,1)*location_.A[0] + GeoM.Element(1,1)*location_.A[1]
        GeoM.Translate(lx, ly)
      }
      cube.Faces[i].MatrixUpdate(camera.MatrixInvert.Mull(Cube_Matrixes_Constant[i]))
      cube.Faces[i].Draw(screen, GeoM)
    }
  }
  cube.This.Draw(screen, screen_location_geom, &camera.Matrix)
}