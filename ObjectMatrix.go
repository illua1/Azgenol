package main

import (
	matrix "github.com/illua1/go-helpful/VectorMatrix"
)

type ObjectMatrix struct {
	Matrix3
	Location3
}

func NewObjectMatrix() ObjectMatrix {
	return ObjectMatrix{Matrix3{matrix.Matrix3x3[Value]()}, Location3{}}
}

func (om ObjectMatrix) Concat(in ObjectMatrix) (ret ObjectMatrix) {
	ret = ObjectMatrix{
		Matrix3{om.Matrix3.Mull(in.Matrix3.Matrix)},
		Location3{om.MulVector(in.Location3.Vector)},
	}
	ret.Location3.Vector.Add(om.Location3.Vector)
	return
}

func (om *ObjectMatrix) Apply(Loaction Location3) (ret Location3) {
	ret = Location3{om.MulVector(Loaction.Vector)}
	ret.Add(om.Location3.Vector)
	return
}
