package AKL_StaticTypes

import (
	matrix "github.com/illua1/go-helpful/VectorMatrix"

	"github.com/hajimehoshi/ebiten/v2"
)

type Vector3 struct {
	matrix.Vector[float64, [3]float64]
}

func NewVector3(x, y, z float64) Vector3 {
	return Vector3{matrix.Vector[float64, [3]float64]{[3]float64{x, y, z}}}
}

func (vector3 Vector3) Project(geom ebiten.GeoM) ebiten.GeoM {
	geom.Translate(vector3.Vector.A[0], vector3.Vector.A[1])
	return geom
}

type Matrix3 struct {
	matrix.Matrix[float64, [3]float64, [3][3]float64]
}

func (matrix3 Matrix3) Project() (ret ebiten.GeoM) {
	matrix3.Slise(0, 0, 2, 2).FillTo(ret.SetElement)
	return
}
