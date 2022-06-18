package AKL_Drawers

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type ImageSerface struct {
  ImageDrawer
  ebiten.GeoM
}

func NewSerface(iDraver ImageDrawer, geom ebiten.GeoM) ImageSerface {
  return ImageSerface{iDraver, geom}
}

func (iSerface ImageSerface)Draw(screen *ebiten.Image, GeomContext ebiten.GeoM){
  iSerface.ImageDrawer.Draw(screen, iSerface.GetGeoM(GeomContext))
}

func (iSerface ImageSerface) ToImageDrawer(GeomContext ebiten.GeoM) ImageDrawer {
  return iSerface.ImageDrawer.ToImageDrawer(iSerface.GetGeoM(GeomContext))
}

func (iSerface ImageSerface)GetGeoM(GeomContext ebiten.GeoM) ebiten.GeoM {
  iSerface.GeoM.Concat(GeomContext)
  return iSerface.ImageDrawer.GetGeoM(iSerface.GeoM)
}