package AKL_SerfaceRender

import (
	pipeline "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_RenderPipeLine"
	types "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_Types"
)

type Cube struct {
	Box
	types.GeoM
}

func NewCube(box Box, geom types.GeoM) Cube {
	return Cube{box, geom}
}

func (cube *Cube) RenderCustom(call pipeline.RenderCallAppend, camera *types.Camera) {
	CubeRender(
		&cube.Box,
		&cube.Vector3,
		&cube.Matrix3,
		call,
		camera,
	)
}

func CubeRender(
	box *Box,
	vector *types.Vector3,
	matrix *types.Matrix3,
	call pipeline.RenderCallAppend,
	camera *types.Camera,
) {
	var face_matrix = types.Matrix3{camera.MatrixInvert.Mull(matrix.Matrix)}
	var face_location types.Vector3
	for _, i := range Box_driwe_faces(&face_matrix) {

		face_location = Box_Faces_Vectors[i]
		face_location.Scale(box[i].float64)
		face_location.Vector = matrix.MulVector(face_location.Vector)
		face_location.Add(vector.Vector)
		face_location.Sub(camera.Location.Vector)

		face_location.Vector = camera.MatrixInvert.MulVector(
			face_location.Vector,
		)
		face_matrix.Matrix = camera.MatrixInvert.Mull(
			matrix.Mull(
				Box_Faces_Matrixes[i].Matrix,
			),
		)

		call(
			box[i].ImageDrawer.ToImageDrawer(
				face_location.Project(face_matrix.Project()),
			),
			face_location.A[2],
		)
	}
}
