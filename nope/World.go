package main

import (
	//"log"
  //"math"
	"github.com/hajimehoshi/ebiten/v2"
  matrix "github.com/illua1/go-helpful/VectorMatrix"
  //volume "github.com/illua1/go-helpful/Volume"
  //UI "github.com/illua1/Game-UI"
)

type Render interface {
  Rendering(camera *Camera)
}

type Camera struct {
  Angle [3]float64
  Matrix matrix.Matrix[float64, [3]float64, [3][3]float64]
  Invert matrix.Matrix[float64, [3]float64, [3][3]float64]
  Location [3]int
  Image *ebiten.Image
}

func NewCamera(location [3]int, angle [3]float64)Camera{
  return Camera{
    Angle : angle,
    Matrix : matrix.Rotate3x3_YXZ[float64](angle),
    Invert : matrix.Rotate3x3_YXZ[float64](angle).Invert(),
    Location : location,
    Image : nil,
  }
}

func(camera *Camera)UpadateImage(image *ebiten.Image){
  camera.Image = image
}

type Layer struct {
  Location [2]int
  Size [2]float64
  Elemenrs []Serface
}

type Serface struct {
  Img *ebiten.Image
  Loacation [2]float64
}

func(serface *Serface)Rendeer(screen *ebiten.Image){
  
}

var(
  t float64 = 0.0
)

func(layer Layer)Render(camera *Camera){
  var geom_matrix ebiten.GeoM
  op := &ebiten.DrawImageOptions{}
  x, y := camera.Image.Size()
  for i := range layer.Elemenrs {
    camera.Matrix.Slise(0,0,2,2).FillTo(geom_matrix.SetElement)
    centre := ImageCentre(Block_plit_face, layer.Size[0], layer.Size[1])
    centre.Concat(geom_matrix)
    op.GeoM = centre
    //op.GeoM.Translate(geom_matrix.Apply(float64(x_)*layer.Size[0],float64(y_)*layer.Size[1]))
    op.GeoM.Translate(float64(x)/2, float64(y)/2)
    //camera.Image.DrawImage(Block_plit_face, op)
    layer.Elemenrs[i].Render(Block_plit_face, op)
  }
}

func ImageCentre(img *ebiten.Image, x_, y_ float64)ebiten.GeoM {
  var geom ebiten.GeoM
  x, y := img.Size()
  geom.Translate(
    -float64(x)/2/x_,
    -float64(y)/2/y_,
  )
  geom.Scale(
    x_/float64(x),
    y_/float64(y),
  )
  return geom
}