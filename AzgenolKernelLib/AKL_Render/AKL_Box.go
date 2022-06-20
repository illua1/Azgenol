package AKL_SerfaceRender

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"

	draw "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_Drawers"
	types "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_Types"

	matrix "github.com/illua1/go-helpful/VectorMatrix"
	volume "github.com/illua1/go-helpful/Volume"
)

var (
	Box_Faces_Matrixes = volume.BoxContainerFaces[types.Matrix3]{
		types.Matrix3{matrix.Rotate3x3_x[float64](math.Pi * 0.0)},
		types.Matrix3{matrix.Rotate3x3_x[float64](math.Pi * 1.0)},
		types.Matrix3{matrix.Rotate3x3_x[float64](math.Pi * 0.5)},
		types.Matrix3{matrix.Rotate3x3_x[float64](-math.Pi * 0.5)},
		types.Matrix3{matrix.Rotate3x3_y[float64](math.Pi * 0.5)},
		types.Matrix3{matrix.Rotate3x3_y[float64](-math.Pi * 0.5)},
	}

	Box_Faces_Vectors = volume.BoxContainerFaces[types.Vector3]{
		types.NewVector3(0, 0, -0.5),
		types.NewVector3(0, 0, 0.5),
		types.NewVector3(0, -0.5, 0),
		types.NewVector3(0, 0.5, 0),
		types.NewVector3(-0.5, 0, 0),
		types.NewVector3(0.5, 0, 0),
	}
)

func Box_driwe_face_check(i int, matrix *types.Matrix3) bool {
	switch i {
	case 0:
		{
			return matrix.A[2][2] < 0
		}
	case 1:
		{
			return matrix.A[2][2] >= 0
		}
	case 2:
		{
			return matrix.A[2][1] < 0
		}
	case 3:
		{
			return matrix.A[2][1] >= 0
		}
	case 4:
		{
			return matrix.A[2][0] < 0
		}
	case 5:
		{
			return matrix.A[2][0] >= 0
		}
	}
	return false
}

func NewBox(sx, sy, sz float64, img *ebiten.Image) Box {
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
			draw.NewImageDrawer(sx, sz, img).Flip(false, true),
			sy,
		},
		BoxFace{
			draw.NewImageDrawer(sx, sz, img).Flip(true, false),
			sy,
		},
		BoxFace{
			draw.NewImageDrawerR(sz, sy, 1, img).Rot(-1.0).Flip(true, true),
			sx,
		},
		BoxFace{
			draw.NewImageDrawerR(sz, sy, 1, img).Rot(1.0).Flip(true, true),
			sx,
		},
	}
}

type BoxFace struct {
	draw.ImageDrawer
	float64
}

type Box volume.BoxContainerFaces[BoxFace]
