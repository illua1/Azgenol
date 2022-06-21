package AKL_TypeComponents

import (
	pipeline "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_RenderPipeLine"
	types "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_Types"
)

type ComponentProcessRender struct {
	pipeline.RenderPipeLine
}

func NewComponentProcessRender() *ComponentProcessRender {
	return &ComponentProcessRender{
		pipeline.NewRenderPipeLine(),
	}
}

type IsRenderObject interface {
	GetRenderObject() pipeline.RenderObject
}

func (cpRender *ComponentProcessRender) Add(in interface{}) {
	if in == nil {
		return
	}
	if instance, ok := in.(IsRenderObject); ok {
		if renderObjectInstance := instance.GetRenderObject(); renderObjectInstance == nil {
			return
		} else {
			cpRender.RenderPipeLine.Add(renderObjectInstance)
		}
	}
}

func (cpRender *ComponentProcessRender) Update(context types.Context) {
	cpRender.RenderPipeLine.Draw(context.Screen, context.Camera)
}
