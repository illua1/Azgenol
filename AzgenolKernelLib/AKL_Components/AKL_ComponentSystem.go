package AKL_TypeComponents

import(
  node "github.com/illua1/go-helpful/Node"
)

type ComponentSystem struct {
  list *node.LNode[ComponentProcess]
}

func NewComponentSystem(in ...ComponentProcess)(ret ComponentSystem){
  for i := range in {
    node.Append(&ret.list, in[i])
  }
  return
}

func(cSystem *ComponentSystem)Add(in ...interface{}){
  node.For(
    &cSystem.list,
    func(cProcess ComponentProcess){
      for e := range in {
        cProcess.Add(in[e])
      }
    },
  )
}

func(cSystem *ComponentSystem)Update(context Context){
  node.For(
    &cSystem.list,
    func(cProcess ComponentProcess){
      cProcess.Update(context)
    },
  )
}
