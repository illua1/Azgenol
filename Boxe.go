package main

import (
  "image/color"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
  matrix "github.com/illua1/go-helpful/VectorMatrix"
  volume "github.com/illua1/go-helpful/Volume"
)

type Box struct{
  Core volume.Box[float64]
  Dynamik Physics
  DrawColor color.Color
}

func NewBox(x, y, z float64, physics Physics)Box{
  x, y, z = x/2, y/2, z/2
  return Box{
    Core : volume.NewBox[float64](-x, -y, -z, x, y, z),
    Dynamik : physics,
    DrawColor : color.RGBA{255,255,255,255},
  }
}

func(Box *Box)Draw(screen *ebiten.Image, screen_geom ebiten.GeoM, worldMatrix *matrix.Matrix[float64, [3]float64, [3][3]float64]){
  points := Box.Core.Points()
  edges := Box.Core.Edges()
  for i := range edges {
    var p1 = worldMatrix.MulVector(matrix.Vector[float64, [3]float64]{
      A:[3]float64{
        float64(points[edges[i][0]].X),
        float64(points[edges[i][0]].Y),
        float64(points[edges[i][0]].Z),
      },
    })
    var p2 = worldMatrix.MulVector(matrix.Vector[float64, [3]float64]{
      A:[3]float64{
        float64(points[edges[i][1]].X),
        float64(points[edges[i][1]].Y),
        float64(points[edges[i][1]].Z),
      },
    })
    x1, y1 := screen_geom.Apply(p1.A[0], p1.A[1])
    x2, y2 := screen_geom.Apply(p2.A[0], p2.A[1])
    ebitenutil.DrawLine(screen, x1, y1, x2, y2, Box.DrawColor)
  }
  Box.Dynamik.Draw(screen, screen_geom, worldMatrix)
}

func (Box *Box)ColiseSolve(do *Box, faceId int){
  if (faceId == volume.BottomFace) || (faceId == volume.TopFace){
    Box.Dynamik.Velocity.A[2] = -Box.Dynamik.Velocity.A[2]*(Box.Dynamik.Resistance * do.Dynamik.Resistance)
    
    //Box.Dynamik.Velocity.A[0] = Box.Dynamik.Velocity.A[2]*(Box.Dynamik.Resistance * do.Dynamik.Resistance)
    //Box.Dynamik.Velocity.A[1] = Box.Dynamik.Velocity.A[2]*(Box.Dynamik.Resistance * do.Dynamik.Resistance)
    return
  }
}