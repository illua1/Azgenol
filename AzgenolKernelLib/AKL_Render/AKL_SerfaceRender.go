package AKL_SerfaceRender

import (
	draw "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_Drawers"
	types "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_Types"
)

func RenderDrawer(img draw.ImageDrawer, camera *types.Camera, location types.Vector3) draw.ImageDrawer {
	location.Sub(camera.Location.Vector)
	location.Vector = camera.MatrixInvert.MulVector(location.Vector)
	return img.ToImageDrawer(
		location.Project(camera.MatrixInvert.Project()),
	)
}

func RenderSerface(img draw.ImageSerface, camera *types.Camera, location types.Vector3) draw.ImageDrawer {
	location.Sub(camera.Location.Vector)
	location.Vector = camera.MatrixInvert.MulVector(location.Vector)
	return img.ToImageDrawer(
		location.Project(camera.MatrixInvert.Project()),
	)
}
