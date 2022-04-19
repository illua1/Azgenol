package main

import (
  matrix "github.com/illua1/go-helpful/VectorMatrix"
)

type Physics struct {
  Location, Velocity, Accelerate matrix.Vector[float64, [3]float64]
  Resistance float64
  Gravity float64
}

func NewPhysics(px, py, pz, g float64)Physics{
  return Physics{
    Location : matrix.Vector[float64, [3]float64]{[3]float64{px, py, pz}},
    Resistance : 1,
    Gravity : -g*10,
  }
}

func (physics *Physics)Update(){
  physics.LocationUpdate()
  physics.VelocityUpdate()
  physics.AccelerateUpdate()
  physics.AccelerateAdd(0,0, physics.Gravity)
}

func (physics *Physics)LocationUpdate(){
  for i := range physics.Location.A {
    physics.Location.A[i] += physics.Velocity.A[i]/60
  }
}

func (physics *Physics)VelocityUpdate(){
  for i := range physics.Velocity.A {
    physics.Velocity.A[i] += physics.Accelerate.A[i]/60
  }
  physics.Velocity.Scale(1./physics.Resistance)
}

func (physics *Physics)AccelerateUpdate(){
  physics.Accelerate.Scale(1./60)
}

func (physics *Physics)AccelerateAdd(x, y, z float64){
  physics.Accelerate.A[0] += x
  physics.Accelerate.A[1] += y
  physics.Accelerate.A[2] += z
}