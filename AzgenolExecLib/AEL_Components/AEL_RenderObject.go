package AEL_components

import (
	render "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_Render"
	pipeline "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_RenderPipeLine"
	types "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_Types"
)

type RenderComponent struct {
	*render.Box
	*types.Vector3
	*types.Matrix3
}

func NewRenderComponent(box *render.Box, v *types.Vector3, m *types.Matrix3) RenderComponent {
	return RenderComponent{box, v, m}
}

func (renderC RenderComponent) RenderCustom(call pipeline.RenderCallAppend, camera *types.Camera) {
	render.CubeRender(
		renderC.Box,
		renderC.Vector3,
		renderC.Matrix3,
		call,
		camera,
	)
}
