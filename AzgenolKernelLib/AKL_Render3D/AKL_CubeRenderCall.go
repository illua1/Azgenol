package AKL_SerfaceRender

import(
  //"math"
  
	//"github.com/hajimehoshi/ebiten/v2"
  
  //volume "github.com/illua1/go-helpful/Volume"
  //matrix "github.com/illua1/go-helpful/VectorMatrix"
  types "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_Types"
  //draw "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_Drawers"
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
  for i := range cube.Box {
    var matrix_l = types.Matrix3{camera.MatrixInvert.Mull(cube.GeoM.Matrix)}
    if box_driwe_face_check(i, &matrix_l) {
      var location = cube.Vector3
      location.Sub(camera.Location.Vector)
      var face_location = Box_Faces_Vectors[i]
      face_location.Scale(-cube.Box[i].float64)
      
      face_location.Vector = cube.GeoM.MulVector(face_location.Vector)
      location.Add(face_location.Vector)
      location.Vector = camera.MatrixInvert.MulVector(location.Vector)
      var face_matrix = types.Matrix3{camera.MatrixInvert.Mull(cube.GeoM.Matrix.Mull(Box_Faces_Matrixes[i].Matrix))}
      call(
        cube.Box[i].ImageDrawer.ToImageDrawer(
          location.Project(face_matrix.Project()),
        ),
        location.A[2],
      )
    }
  }
}