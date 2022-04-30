package main

import (
  "log"
  //"image/color"
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

var(
  ignore_A int = -1
  ignore_B int = -1
)

func(pc *PhysicsCollectro)Update(){
  var min_t float64 = 1
  var idA, idB int = -1, -1
  var colise bool = false
  var AfaceId int = 0
  /*
    Use all DtStep vectors (what contain offset per next frame)
    And check, if Boxs will colised bu this data, and this first action,
    We apply all DtStep`s for all Boxz, and after this, just compute 
    Actual colise-DtStep date for curret Boxs
    And at new date, just repest this things, 
    Before we not get 1.0 time at satart of allgoritm
  */
  for i := 0; i < len(pc.Collection); i++{
    for e := i+1; e < len(pc.Collection); e++ {
      if moment, polyFace, ok := BoxsColiseTest(pc.Collection[i], pc.Collection[e], min_t); ok {
        if !(((ignore_A == i)||(ignore_B == i)) && ((ignore_A == e)||(ignore_B == e))) {
          min_t = moment
          idA, idB = i, e
          colise = true
          AfaceId = polyFace
        }
      }
    }
  }
  for i := 0; i < len(pc.Collection); i++{
    pc.Collection[i].Dynamik.Update(min_t*0.9)
  }
  if colise {
    pc.Collection[idA].ColiseSolve(pc.Collection[idB], AfaceId)
    if !pc.Collection[idA].Dynamik.Immovable && !pc.Collection[idB].Dynamik.Immovable {
      log.Print("Do colise solve :", min_t)
    }
  }
  ignore_A = idA
  ignore_B = idB
  
}

func BoxsColiseTest(a, b *Box, min_t float64)(t float64, f_index int, ok bool){
  
  var FacePoints1 = a.Core.FaceCentres()
  var FacePoints2 = b.Core.FaceCentres()
  
  p1 := a.Dynamik.Location.A
  v1 := a.Dynamik.Velocity.A
  p2 := b.Dynamik.Location.A
  v2 := b.Dynamik.Velocity.A
  
  t = min_t
  f_index = -1
  
  if tx := Future(p1[2], p2[2], p1[2]+v1[2], p2[2]+v2[2], FacePoints1[volume.BottomFace].Z, FacePoints2[volume.TopFace].Z); (tx > 0)&&(tx <= t) {
    t = tx
    f_index = volume.TopFace
    ok = true
  }
  if tx := Future(p1[2], p2[2], p1[2]+v1[2], p2[2]+v2[2], FacePoints1[volume.TopFace].Z, FacePoints2[volume.BottomFace].Z); (tx > 0)&&(tx <= t) {
    t = tx
    f_index = volume.TopFace
    ok = true
  }
  return
}

func Future(A_x, B_x, A_x_t, B_x_t, A_Max, B_min float64)float64{
  return matrix.UnLerp[float64, float64](
    ((A_x + A_Max) - (B_x + B_min)),
    ((A_x_t + A_Max) - (B_x_t + B_min)),
    0,
  )
}