package main

import (
	"github.com/hajimehoshi/ebiten/v2"
  sort "github.com/illua1/go-helpful/Sort"
  node "github.com/illua1/go-helpful/Node"
)

type RenderPipeLine struct {
  Objects []RenderObject
  calls []RenderCall
  Debug bool
}

type RenderObjectDebug interface {
  Draw(*ebiten.Image, ObjectMatrix, *Camera)
}

func NewRenderPipeLine(in ...RenderObject) RenderPipeLine {
  return RenderPipeLine{
    Objects : in,
  }
}

func(rp *RenderPipeLine) Add(in ...RenderObject){
  rp.Objects = append(rp.Objects, in...)
}

func(rp *RenderPipeLine) Draw (screen *ebiten.Image, camera *Camera){
  rp.calls = make([]RenderCall, 0)
  
  var object = NewObjectMatrix()
  
  for i := range rp.Objects {
    rp.Objects[i].RenderCustom(
      RenderCallAppend(func(imgd ImageDrawer, deph Value) {
        rp.calls = append(rp.calls, NewRenderCall(imgd, deph))
        if len(rp.calls) == 1 {
          return
        }
        node.BNodeDescrent(
          (*node.BNode[RenderCallContain])(&rp.calls[0]),
          (*node.BNode[RenderCallContain])(&rp.calls[len(rp.calls)-1]),
          func(a, b *node.BNode[RenderCallContain]) bool {
            return a.Contain.deph > b.Contain.deph
          },
        )
      }),
      object,
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
        obj.Draw(screen, object, camera)
      }
    }
  }
  
}