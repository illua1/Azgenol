package AKL_TypeComponents

import(
  "time"
  
  types "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_Types"
  
  "github.com/hajimehoshi/ebiten/v2"
)

type ComponentProcess interface {
  Add(interface{})
  Update(Context)
}

type Context struct {
  Screen *ebiten.Image
  Camera types.Camera
  Time time.Duration
}