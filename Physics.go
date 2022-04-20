package main

import (
  matrix "github.com/illua1/go-helpful/VectorMatrix"
)

var(
  DeltaTime float64 = 1 / 1.5
)

type Physics struct {
  Location, Velocity, Accelerate matrix.Vector[float64, [3]float64]
  Resistance float64
  Gravity float64
  
  Connect_to_Top bool
  Connect_to_Bottom bool
  Connect_to_Left bool
  Connect_to_Right bool
  Connect_to_Fase bool
  Connect_to_Behind bool
}

func NewPhysics(px, py, pz, g float64)Physics{
  return Physics{
    Location : matrix.Vector[float64, [3]float64]{[3]float64{px, py, pz}},
    Resistance : 1,
    Gravity : -g,
  }
}

func (physics *Physics)Update(){
  physics.VelocityUpdate()
  physics.LocationUpdate()
}

func (physics *Physics)LocationUpdate(){
  physics.Velocity.Scale(DeltaTime)
  physics.Location.Add(physics.Velocity)
}

func (physics *Physics)VelocityUpdate(){
  physics.Accelerate.Scale(DeltaTime)
  physics.Velocity.Add(physics.Accelerate)
}

func (physics *Physics)VelocityAdd(x, y, z float64){
  physics.Velocity.A[0] += x
  physics.Velocity.A[1] += y
  
  if z > 0 {
    if !physics.Connect_to_Top {
      physics.Velocity.A[2] += z
    }
    
  } else{
    if !physics.Connect_to_Bottom {
      physics.Velocity.A[2] += z
    }
  }
}

func (physics *Physics)AccelerateAdd(x, y, z float64){
  physics.Accelerate.A[0] += x
  physics.Accelerate.A[1] += y
  if z > 0 {
    if !physics.Connect_to_Top {
      physics.Accelerate.A[2] += z
    }
    
  } else{
    if !physics.Connect_to_Bottom {
      physics.Accelerate.A[2] += z
    }
  }
}