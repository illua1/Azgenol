package AKL_RenderPipeLine

import (
  draw "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_Drawers"
  types "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_Types"
  
  sort "github.com/illua1/go-helpful/Sort"
  node "github.com/illua1/go-helpful/Node"
  
	"github.com/hajimehoshi/ebiten/v2"
)

type RenderPipeLine struct {
  Objects []RenderObject
  calls []RenderCall
  Debug bool
}

type RenderObjectDebug interface {
  Draw(*ebiten.Image, *types.Camera)
}

func NewRenderPipeLine(in ...RenderObject) RenderPipeLine {
  return RenderPipeLine{
    Objects : in,
  }
}

func(rp *RenderPipeLine) Add(in ...RenderObject){
  rp.Objects = append(rp.Objects, in...)
}

func(rp *RenderPipeLine) Draw (screen *ebiten.Image, camera *types.Camera){
  rp.calls = make([]RenderCall, 0)
  
  for i := range rp.Objects {
    rp.Objects[i].RenderCustom(
      RenderCallAppend(func(IDrawer draw.ImageDrawer, deph float64) {
        rp.calls = append(rp.calls, NewRenderCall(IDrawer, deph))
        if len(rp.calls) == 1 {
          return
        }
        node.BNodeDescrent(
          (*node.BNode[RenderCallContain])(&rp.calls[0]),
          (*node.BNode[RenderCallContain])(&rp.calls[len(rp.calls)-1]),
          func(a, b *node.BNode[RenderCallContain]) bool {
            return a.Contain.deph < b.Contain.deph
          },
        )
      }),
      camera,
    )
  }
  
  x,y := screen.Size()
  var geom = ebiten.GeoM{}
  geom.Scale(float64(sort.MaxF(x, y))/1000, float64(sort.MaxF(x, y))/1000)
  geom.Translate(float64(x/2), float64(y/2))
  
  node.BNodeForTo((*node.BNode[RenderCallContain])(&rp.calls[0]), func(_ int, contain RenderCallContain){
    contain.Draw(screen, geom)
  })
  
  if rp.Debug {
    for i := range rp.Objects {
      if obj, ok := rp.Objects[i].(RenderObjectDebug); ok {
        obj.Draw(screen, camera)
      }
    }
  }
  
}