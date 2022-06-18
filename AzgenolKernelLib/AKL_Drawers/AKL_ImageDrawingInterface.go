package AKL_Drawers

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type ImageDrawing interface {

  GetGeoM(GeomContext ebiten.GeoM) ebiten.GeoM
  
  Draw(img *ebiten.Image, GeomContext ebiten.GeoM)
  
  ToImageDrawer(GeomContext ebiten.GeoM) ImageDrawer
}

type ImageDrawers interface {
  ImageDrawer | ImageSerface | ImageDrawerExec
}