package main

import (
  "log"
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
  var t float64 = 0
  var min_t float64 = 1
  var idA, idB int = 0, 1
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
  for count := 0; count < 1000; count++ {
    for i := 0; i < len(pc.Collection); i++{
      for e := i+1; e < len(pc.Collection); e++ {
        if moment, polyFace := BoxsColiseTest(pc.Collection[i], pc.Collection[e]); ((moment > 0)&&((moment+t)<=1)&&(moment < min_t)) {
          min_t = moment
          idA, idB = i, e
          colise = true
          AfaceId = polyFace
        }
      }
    }
    for i := 0; i < len(pc.Collection); i++{
      pc.Collection[i].Dynamik.Update(min_t)
    }
    if colise {
      log.Print("Colise ", t)
      
      pc.Collection[idA].ColiseSolve(pc.Collection[idB], AfaceId)
      pc.Collection[idB].ColiseSolve(pc.Collection[idA], volume.FaceOpposite(AfaceId))
    }
    t += min_t
    if t >= 1 {
      return
    }
  }
}

func BoxsColiseTest(a, b *Box)(float64, int){
  
  var FacePoints1 = a.Core.FaceCentres()
  var FacePoints2 = b.Core.FaceCentres()
  
  p1 := a.Dynamik.Location.A
  v1 := a.Dynamik.Velocity.A
  p2 := b.Dynamik.Location.A
  v2 := b.Dynamik.Velocity.A
  
  if a.Dynamik.Connect_to[volume.BottomFace] {
    v2[2] = -100
  }
  
  x_dist := -(FacePoints1[volume.TopFace].Z + p1[2]) + (FacePoints2[volume.BottomFace].Z + p2[2])
  
  x_velocity := v1[2] - v2[2]
  
  t := (x_dist / x_velocity)
  
  a.DrawColor = color.RGBA{255,255,255,255}
  if (t < 1)&&(t > 0) {
    var sizes1 = a.Core.FaceArea()
    var sizes2 = b.Core.FaceArea()
    if (sizes1[0][0]+sizes2[0][0]) > matrix.Abc(p1[0] - p2[0])&&(sizes1[0][1]+sizes2[0][1]) > matrix.Abc(p1[1] - p2[1]) {
      a.DrawColor = color.RGBA{0,0,255,255}
      return t, volume.BottomFace
    }
  }
  return 100, -1
}