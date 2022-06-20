package AKL_StaticTypes

import (
	matrix "github.com/illua1/go-helpful/VectorMatrix"

	"github.com/hajimehoshi/ebiten/v2"
)

type GeoM struct {
	Vector3
	Matrix3
}

func NewGeoM(x, y, z float64) GeoM {
	return GeoM{
		NewVector3(x, y, z),
		Matrix3{
			matrix.Matrix3x3[float64](),
		},
	}
}

func (geom GeoM) Project() ebiten.GeoM {
	return geom.Vector3.Project(
		geom.Matrix3.Project(),
	)
}

func (geom GeoM) Concat(in GeoM) (ret GeoM) {
	ret = GeoM{
		geom.Apply(in.Vector3),
		Matrix3{in.Matrix.Mull(geom.Matrix)},
	}
	return
}

func (geom GeoM) Apply(vector Vector3) (ret Vector3) {
	ret = Vector3{geom.MulVector(vector.Vector)}
	ret.Add(geom.Vector)
	return
}
