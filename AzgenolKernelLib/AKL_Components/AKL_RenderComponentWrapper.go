package AKL_TypeComponents

import (
  types "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_Types"
  pipeline "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_RenderPipeLine"
)

type ComponentProcessRender struct{
  pipeline.RenderPipeLine
}

func NewComponentProcessRender()*ComponentProcessRender{
  return &ComponentProcessRender{
    pipeline.NewRenderPipeLine(),
  }
}

type IsRenderObject interface {
  GetRenderObject()pipeline.RenderObject
}

func (cpRender *ComponentProcessRender) Add (in interface{}){
  if instance, ok := in.(IsRenderObject); ok {
    cpRender.RenderPipeLine.Add(instance.GetRenderObject())
  }
}

func (cpRender *ComponentProcessRender) Update (context types.Context){
  cpRender.RenderPipeLine.Draw(context.Screen, &context.Camera)
}