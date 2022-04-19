package main

import (
  //"fmt"
  "math"
  "image/color"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
  matrix "github.com/illua1/go-helpful/VectorMatrix"
  volume "github.com/illua1/go-helpful/Volume"
)

type GameCube struct{
  Body volume.Boxe[int]
  Matrixes [6]matrix.Matrix[float64, [3]float64, [3][3]float64]
  Centres [6]volume.Point[int]
  Src *ebiten.Image
}

type (
  RenderingObject interface {
    Render(screen *ebiten.Image, location [2]int)
  }
  ColisedObject interface {
    BoundingBoxe()volume.Boxe[int]
  }
)

type (
  Camera struct {
    Location [3]int
    Angle [3]float64
    Matrix matrix.Matrix[float64, [3]float64, [3][3]float64]
    Matrix_Invert matrix.Matrix[float64, [3]float64, [3][3]float64]
  }
  
  layer struct{
    Objects []RenderingObject
    Location_z int
  }
  
  WorldRender struct{
    Layers []layer
    Camera Camera
  }
)

func(gm GameCube)Render(screen *ebiten.Image, wm matrix.Matrix[float64, [3]float64, [3][3]float64], camera_location volume.Point[int]){
  op := &ebiten.DrawImageOptions{}
  wm = wm.Mull(matrix.Rotate3x3_YXZ[float64](r))
  
  for i := range gm.Matrixes {
    op.GeoM = ebiten.GeoM{}
    
    mat_invert := wm.Invert().Mull(gm.Matrixes[i])
    mat_invert.Slise(0,0,2,2).FillTo(op.GeoM.SetElement)
    op.GeoM.Translate(op.GeoM.Apply(-250,-250))
    
    location_i := wm.MulVector(
      matrix.Vector[float64, [3]float64]{
        A:[3]float64{
          float64(gm.Centres[i].X),
          float64(gm.Centres[i].Y),
          float64(gm.Centres[i].Z),
        },
      },
    )
    op.GeoM.Translate(location_i.A[0], location_i.A[1])
    op.GeoM.Translate(350,400)
    screen.DrawImage(gm.Src, op)
  }
  a := gm.Body.Points()
  b := gm.Body.Edges()
  for i := range b {
    var p1 = wm.MulVector(matrix.Vector[float64, [3]float64]{A:[3]float64{float64(a[b[i][0]].X), float64(a[b[i][0]].Y), float64(a[b[i][0]].Z)}})
    var p2 = wm.MulVector(matrix.Vector[float64, [3]float64]{A:[3]float64{float64(a[b[i][1]].X), float64(a[b[i][1]].Y), float64(a[b[i][1]].Z)}})
    ebitenutil.DrawLine(screen, p1.A[0]+350, p1.A[1]+400, p2.A[0]+350, p2.A[1]+400, color.RGBA{255,255,255,255})
    
  }
}

var(
  Cube = GameCube{
    Body : volume.NewBoxe[int](-250,-250,-250,250,250,250),
    Matrixes : [6]matrix.Matrix[float64, [3]float64, [3][3]float64]{
      matrix.Matrix3x3[float64](),
      matrix.Matrix3x3[float64](),
      matrix.Rotate3x3_x[float64](math.Pi/2),
      matrix.Rotate3x3_x[float64](-math.Pi/2),
      matrix.Rotate3x3_y[float64](math.Pi/2),
      matrix.Rotate3x3_y[float64](-math.Pi/2),
    },
    Centres : volume.NewBoxe[int](-250,-250,-250,250,250,250).FaceCentres(),
    Src : Block_plit_face,
  }
)

func NewWorld(depth_min, depth_max int, cam Camera)WorldRender{
  cam.Matrix = matrix.Rotate3x3_XYZ[float64](cam.Angle)
  return WorldRender{
    Layers : make([]layer, depth_max-depth_min),
    Camera : cam,
  }
}

func(wr *WorldRender)AppendToLayer(layer_index int, elements ...RenderingObject){
  wr.Layers[layer_index].Objects = append(wr.Layers[layer_index].Objects, elements...)
}

func (wr *WorldRender)Render(screen *ebiten.Image)error{
  var (
    rotate_geom ebiten.GeoM
    op = &ebiten.DrawImageOptions{}
    x_centre, y_centre = screen.Size()
  )
  op.GeoM.Translate(-250, -250)
  wr.Camera.Matrix.Slise(0,0,2,2).FillTo(rotate_geom.SetElement)
  op.GeoM.Concat(rotate_geom)
  for l := range wr.Layers{
    for e := range wr.Layers[l].Objects {
      op.GeoM.Translate(float64(x_centre/2), float64(y_centre/2))
      wr.Layers[l].Objects[e].Render(screen, [2]int{x_centre, y_centre})
    }
  }
  Cube.Render(screen, wr.Camera.Matrix, volume.Point[int]{1000, 1000, 1000})
  
  return nil
}