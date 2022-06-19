package AKL_TypeComponents

import (
  pipline "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_RenderPipeLine"
)

type ComponentProcessRender struct{
  pipline.RenderPipeLine
}

func NewComponentProcessRender()*ComponentProcessRender{
  return &ComponentProcessRender{
    pipline.NewRenderPipeLine(),
  }
}

func (cpRender *ComponentProcessRender) Add (in interface{}){
  if instance, ok := in.(pipline.RenderObject); ok {
    cpRender.RenderPipeLine.Add(instance)
  }
}

func (cpRender *ComponentProcessRender) Update (context Context){
  cpRender.RenderPipeLine.Draw(context.Screen, &context.Camera)
}