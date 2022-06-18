package AKL_Drawers

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type ImageDrawerExec struct {
  ImageDrawing
  ebiten.GeoM
}

func NewImageDrawerExec(iDrawerE ImageDrawing, GeomContext ebiten.GeoM) ImageDrawerExec {
  return ImageDrawerExec{iDrawerE, GeomContext}
}

func (iDrawerE ImageDrawerExec)Draw(screen *ebiten.Image, GeomContext ebiten.GeoM){
  iDrawerE.ImageDrawing.Draw(screen, iDrawerE.GetGeoM(GeomContext))
}

func (iDrawerE ImageDrawerExec) ToImageDrawer(GeomContext ebiten.GeoM) ImageDrawer {
  return iDrawerE.ImageDrawing.ToImageDrawer(iDrawerE.GetGeoM(GeomContext))
}

func (iDrawerE ImageDrawerExec)GetGeoM(GeomContext ebiten.GeoM)ebiten.GeoM{
  iDrawerE.GeoM.Concat(GeomContext)
  return iDrawerE.ImageDrawing.GetGeoM(iDrawerE.GeoM)
}