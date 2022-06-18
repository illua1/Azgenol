package AKL_Drawers

import (
	"github.com/hajimehoshi/ebiten/v2"
  values "github.com/illua1/go-helpful"
)

type ImageDrawer struct {
  *ebiten.Image
  ebiten.GeoM
}

func NewImageDrawer[Value values.Values](x_size, y_size Value, img *ebiten.Image)(ret ImageDrawer){
  ret.Image = img
  
  x, y := img.Size()
  ret.GeoM.Translate(
    (-float64(x)/2),
    (-float64(y)/2),
  )
  ret.GeoM.Scale(
    float64(x_size)/float64(x),
    float64(y_size)/float64(y),
  )
  
  return
}

func (iDrawe ImageDrawer)Draw(screen *ebiten.Image, GeomContext ebiten.GeoM){
  screen.DrawImage(iDrawe.Image, &ebiten.DrawImageOptions{GeoM : iDrawe.GetGeoM(GeomContext)})
}

func (iDrawe ImageDrawer)ToImageDrawer(GeomContext ebiten.GeoM)ImageDrawer{
  iDrawe.GeoM = iDrawe.GetGeoM(GeomContext)
  return iDrawe
}

func (iDrawe ImageDrawer) GetGeoM(GeomContext ebiten.GeoM) ebiten.GeoM {
  iDrawe.GeoM.Concat(GeomContext)
  return iDrawe.GeoM
}