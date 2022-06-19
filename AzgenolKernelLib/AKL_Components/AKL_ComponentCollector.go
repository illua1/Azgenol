package AKL_TypeComponents

import (
  pipline "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_RenderPipeLine"
)

type ComponentCollector struct {
  pipline.RenderPipeLine
}

func(cCollector *ComponentCollector)Add(in ...interface{}){
  for i := range in {
    if instance, ok := in[i].(pipline.RenderObject); ok {
      cCollector.RenderPipeLine.Add(instance)
    }
  }
}

