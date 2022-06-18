package AKL_Drawers

import (
  "testing"
  
	"github.com/hajimehoshi/ebiten/v2"
)

func TestInterface(t *testing.T) {
  var img = ebiten.NewImage(10, 10)
  NewImageDrawerExec(
    NewImageDrawerExec(
      NewSerface(
        NewImageDrawer(
          10,
          10,
          img,
        ),
        ebiten.GeoM{},
      ),
      ebiten.GeoM{},
    ),
    ebiten.GeoM{},
  )
  
  NewImageDrawerExec(
    NewImageDrawer(
      10,
      10,
      img,
    ),
    ebiten.GeoM{},
  )
}