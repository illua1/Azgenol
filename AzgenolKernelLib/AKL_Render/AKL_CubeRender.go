package AKL_SerfaceRender

import(
  types "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_Types"
  pipeline "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_RenderPipeLine"
)

type Cube struct {
  Box
  types.GeoM
}

func NewCube(box Box, geom types.GeoM) Cube {
  return Cube{box, geom}
}

func(cube *Cube) RenderCustom(call pipeline.RenderCallAppend, camera *types.Camera) {
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
  for i := range box {
    var matrix_l = types.Matrix3{camera.MatrixInvert.Mull(matrix.Matrix)}
    if Box_driwe_face_check(i, &matrix_l) {
      var location = *vector
      location.Sub(camera.Location.Vector)
      var face_location = Box_Faces_Vectors[i]
      face_location.Scale(box[i].float64)
      face_location.Vector = matrix.MulVector(face_location.Vector)
      location.Add(face_location.Vector)
      location.Vector = camera.MatrixInvert.MulVector(location.Vector)
      var face_matrix = types.Matrix3{camera.MatrixInvert.Mull(matrix.Mull(Box_Faces_Matrixes[i].Matrix))}
      call(
        box[i].ImageDrawer.ToImageDrawer(
          location.Project(face_matrix.Project()),
        ),
        location.A[2],
      )
    }
  }
}