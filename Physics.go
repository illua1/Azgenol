package main

import (
  
  "image/color"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
  volume "github.com/illua1/go-helpful/Volume"
  matrix "github.com/illua1/go-helpful/VectorMatrix"
)

var(
  DeltaTime float64 = 1/6
)

type Physics struct {
  Location, Velocity matrix.Vector[float64, [3]float64]
  Resistance, Gravity, Mass float64
  
  Connect_to volume.BoxContainerFaces[bool]
}

func NewPhysics(px, py, pz, g float64)Physics{
  return Physics{
    Location : matrix.Vector[float64, [3]float64]{[3]float64{px, py, pz}},
    Resistance : 0.75,
    Gravity : -g,
  }
}

func (physics *Physics)Update(t float64){
  Step := physics.Velocity
  Step.Scale(t)
  physics.Location.Add(Step)
  
  physics.Velocity.Scale(1 - t*physics.Resistance)
  physics.Velocity.Add(matrix.Vector3(0,0,physics.Gravity))
}

func (physics *Physics)VelocityAdd(x, y, z float64){
  physics.Velocity.A[0] += x
  physics.Velocity.A[1] += y
  physics.Velocity.A[2] += z
}
func(physics *Physics)Draw(screen *ebiten.Image, screen_geom ebiten.GeoM, worldMatrix *matrix.Matrix[float64, [3]float64, [3][3]float64]){
  
  x, y := screen_geom.Apply(0,0)
  
  lines_draw := physics.Connect_to
  
  vectors := [6]matrix.Vector[float64, [3]float64]{
    matrix.Vector3[float64](0,0,50),
    matrix.Vector3[float64](0,0,-50),
    matrix.Vector3[float64](0,-50,0),
    matrix.Vector3[float64](0,50,0),
    matrix.Vector3[float64](-50,0,0),
    matrix.Vector3[float64](50,0,0),
  }
  
  for i := range lines_draw {
    if lines_draw[i] {
      var p1 = worldMatrix.MulVector(matrix.Vector[float64, [3]float64]{
        A:[3]float64{
          float64(0),
          float64(0),
          float64(0),
        },
      })
      var p2 = worldMatrix.MulVector(vectors[i])
      x1, y1 := screen_geom.Apply(p1.A[0], p1.A[1])
      x2, y2 := screen_geom.Apply(p2.A[0], p2.A[1])
      ebitenutil.DrawLine(screen, x1, y1, x2, y2, color.RGBA{0,0,255,255})
    }
  }
  const size float64 = 5
  ebitenutil.DrawRect(
    screen,
    x-size/2,
    y-size/2,
    size,
    size,
    color.RGBA{255,255,255,255},
  )
}