package main

import (
  matrix "github.com/illua1/go-helpful/VectorMatrix"
)

type Value float64

type Number int32

type Location3 struct{
  matrix.Vector[Value, [3]Value]
}

type Matrix3 struct{
  matrix.Matrix[Value, [3]Value, [3][3]Value]
}