package main

import (
  "image/color"
  volume "github.com/illua1/go-helpful/Volume"
  matrix "github.com/illua1/go-helpful/VectorMatrix"
)

type PhysicsCollectro struct {
  Collection []*Box
}

func NewPhysicsCollectro(in ...*Box)PhysicsCollectro{
  return PhysicsCollectro{
    Collection : in,
  }
}

func(pc *PhysicsCollectro)Append(in ...*Box){
  pc.Collection = append(pc.Collection, in...)
}

func(pc *PhysicsCollectro)Update(){
  for t := 0; t < 100; t++ {
    for i := 0; i < len(pc.Collection); i++{
      pc.Collection[i].DrawColor = color.RGBA{255,255,255,255}
    }
    for i := 0; i < len(pc.Collection); i++{
      
      for e := i+1; e < len(pc.Collection); e++ {
        if offset_box := pc.Collection[e].Core.Add(
          volume.Point[float64]{
            pc.Collection[e].Dynamik.Location.A[0],
            pc.Collection[e].Dynamik.Location.A[1],
            pc.Collection[e].Dynamik.Location.A[2],
          }.Sub(
            volume.Point[float64]{
              pc.Collection[i].Dynamik.Location.A[0],
              pc.Collection[i].Dynamik.Location.A[1],
              pc.Collection[i].Dynamik.Location.A[2],
            },
          ),
        ); pc.Collection[i].Core.Colise(offset_box) {
          if (!pc.Collection[i].Dynamik.Immovable) && (!pc.Collection[e].Dynamik.Immovable) {
            BoxesPushing(pc.Collection[i], pc.Collection[e])
            BoxesPushing(pc.Collection[e], pc.Collection[i])
          }
        }
      }
    }
    for i := 0; i < len(pc.Collection); i++{
      pc.Collection[i].Dynamik.Update(1.0/100.0)
    }
  }
}

func BoxesPushing(this, another *Box){
  if this.Dynamik.Immovable {
    return
  }
  offset := this.Dynamik.Location
  offset.Sub(another.Dynamik.Location)
  
  if !another.Dynamik.Immovable {
    offset.Div(matrix.Vector3[float64](2,2,2))
  }
  
  {
    dx1 := this.Core.Dx()
    dy1 := this.Core.Dy()
    dz1 := this.Core.Dz()
    
    dx2 := another.Core.Dx()
    dy2 := another.Core.Dy()
    dz2 := another.Core.Dz()
    
    x := offset.A[0]*dy2*dz2
    y := offset.A[1]*dx2*dz2
    z := offset.A[2]*dy2*dx2
    
    var m_v float64 = matrix.Abc(x)
    var m_id int = 0
    
    if y_ := matrix.Abc(y); m_v < y_ {
      m_id = 1
      m_v = y_
    }
    
    if z_ := matrix.Abc(z); m_v < z_ {
      m_id = 2
    }
    
    offset.Sub(
      matrix.Vector3[float64](
        (dx1 - dx2),
        (dy1 - dy2),
        (dz1 - dz2),
      ),
    )
    
    this.Dynamik.Location.A[m_id] += offset.A[m_id]/10
    
    this.Dynamik.Velocity.A[m_id] = -this.Dynamik.Velocity.A[m_id]
  }
}