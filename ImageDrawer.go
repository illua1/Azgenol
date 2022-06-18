package main

import (
	"github.com/hajimehoshi/ebiten/v2"
  values "github.com/illua1/go-helpful"
)

var(
  op = &ebiten.DrawImageOptions{}
)

type ImageDrawer struct {
  Img *ebiten.Image
  GeoM ebiten.GeoM
}

func NewImageDrawer[Value values.Values](x_size, y_size Value, img *ebiten.Image)ImageDrawer{
  x, y := img.Size()
  var geom ebiten.GeoM
  geom.Translate(
    (-float64(x)/2),
    (-float64(y)/2),
  )
  geom.Scale(
    float64(x_size)/float64(x),
    float64(y_size)/float64(y),
  )
  return ImageDrawer{img, geom}
}

func (imageDrawe ImageDrawer)Draw(screen *ebiten.Image, geom ebiten.GeoM){
  op.GeoM = imageDrawe.GetGeoM(geom)
  screen.DrawImage(imageDrawe.Img, op)
}

func (imageDrawe ImageDrawer) GetGeoM(concat ebiten.GeoM) ebiten.GeoM {
  imageDrawe.GeoM.Concat(concat)
  return imageDrawe.GeoM
}