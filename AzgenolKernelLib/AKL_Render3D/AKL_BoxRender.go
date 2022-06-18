package AKL_SerfaceRender

import(
  "math"
  
	"github.com/hajimehoshi/ebiten/v2"
  
  types "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_Types"
  draw "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_Drawers"
  
  volume "github.com/illua1/go-helpful/Volume"
  matrix "github.com/illua1/go-helpful/VectorMatrix"
)

var (
  Box_Faces_Matrixes = volume.BoxContainerFaces[types.Matrix3]{
    types.Matrix3{matrix.Rotate3x3_x[float64](math.Pi*0.0)},
    types.Matrix3{matrix.Rotate3x3_x[float64](math.Pi*1.0)},
    types.Matrix3{matrix.Rotate3x3_x[float64](math.Pi*0.5)},
    types.Matrix3{matrix.Rotate3x3_x[float64](-math.Pi*0.5)},
    types.Matrix3{matrix.Rotate3x3_y[float64](math.Pi*0.5)},
    types.Matrix3{matrix.Rotate3x3_y[float64](-math.Pi*0.5)},
  }
  
  Box_Faces_Vectors = volume.BoxContainerFaces[types.Vector3]{
    types.NewVector3(0,0,-0.5),
    types.NewVector3(0,0,0.5),
    types.NewVector3(0,-0.5,0),
    types.NewVector3(0,0.5,0),
    types.NewVector3(-0.5,0,0),
    types.NewVector3(0.5,0,0),
  }
)

func box_driwe_face_check(i int, matrix *types.Matrix3) bool {
  switch i {
    case 0 : {
      return matrix.A[2][2] < 0
    }
    case 1 : {
      return matrix.A[2][2] >= 0
    }
    case 2 : {
      return matrix.A[2][1] < 0
    }
    case 3 : {
      return matrix.A[2][1] >= 0
    }
    case 4 : {
      return matrix.A[2][0] < 0
    }
    case 5 : {
      return matrix.A[2][0] >= 0
    }
  }
  return false
}

type BoxFace struct {
  draw.ImageDrawer
  float64
}

type Box [6]BoxFace

func NewBox(sx, sy, sz float64, img *ebiten.Image)Box{
  return Box{
    BoxFace{
      draw.NewImageDrawer(sx, sy, img),
      sz,
    },
    BoxFace{
      draw.NewImageDrawer(sx, sy, img),
      sz,
    },
    BoxFace{
      draw.NewImageDrawer(sx, sz, img),
      sy,
    },
    BoxFace{
      draw.NewImageDrawer(sx, sz, img),
      sy,
    },
    BoxFace{
      draw.NewImageDrawer(sz, sy, img),
      sx,
    },
    BoxFace{
      draw.NewImageDrawer(sz, sy, img),
      sx,
    },
  }
}

func(box *Box)Render(camera *types.Camera, location types.Vector3) (ret [3]draw.ImageDrawer) {
  var offset int = 0
  for i := range *box {
    if box_driwe_face_check(i, &camera.Matrix) {
      var location = location
      location.Sub(camera.Location.Vector)
      var face_location = Box_Faces_Vectors[i]
      face_location.Scale(-box[i].float64)
      location.Add(face_location.Vector)
      location.Vector = camera.MatrixInvert.MulVector(location.Vector)
      var face_matrix = types.Matrix3{camera.MatrixInvert.Mull(Box_Faces_Matrixes[i].Matrix)}
      ret[offset] = box[i].ImageDrawer.ToImageDrawer(
        location.Project(face_matrix.Project()),
      )
      offset++
    }
  }
  return
}