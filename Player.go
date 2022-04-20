package main

import (
  "image/color"
	"github.com/hajimehoshi/ebiten/v2"
  matrix "github.com/illua1/go-helpful/VectorMatrix"
  //ternary "github.com/illua1/go-helpful/If"
)

type Stand struct {
  This Boxe
  Item ImageDrawer
}

type Player Stand

func NewPlayer(x,y,z float64)Player{
  return Player(
    Stand{
      NewBoxe(
        25,
        25,
        25,
        NewPhysics(
          x,
          y,
          z,
          10,
        ),
      ),
      NewImageDrawer[float64](
        100,
        100,
        Player_First_body_1_45,
      ),
    },
  )
}

func(player *Player)Draw(screen *ebiten.Image, screen_geom ebiten.GeoM, camera *Camera){
  location := player.This.Dynamik.Location
  location.Sub(camera.Location)
  global_location := camera.Matrix.MulVector(location)
  {
    sx := screen_geom.Element(0,0)*global_location.A[0] + screen_geom.Element(1,0)*global_location.A[1]
    sy := screen_geom.Element(0,1)*global_location.A[0] + screen_geom.Element(1,1)*global_location.A[1]
  
    screen_geom.Translate(sx, sy)
  }
  player.Item.Draw(screen, screen_geom)
  player.This.Draw(screen, screen_geom, &camera.Matrix)
}



func BoxesColiseTest(a, b *Boxe)bool{
  
  if !a.Dynamik.Connect_to_Bottom{
    
    var FacePoints1 = a.Core.FaceCentres()
    var FacePoints2 = b.Core.FaceCentres()
    
    p1 := a.Dynamik.Location.A
    v1 := a.Dynamik.Velocity.A
    p2 := b.Dynamik.Location.A
    v2 := b.Dynamik.Velocity.A
    
    x_dist := (FacePoints1[0].Z + p1[2]) - (FacePoints2[1].Z + p2[2])
    
    x_velocity := v1[2] - v2[2]
    
    t := (x_dist / x_velocity)
    
    if (t < 1)&&(t > 0) {
        a.DrawColor = color.RGBA{0,0,255,255}
      var sizes1 = a.Core.FaceArea()
      var sizes2 = b.Core.FaceArea()
      
      if (sizes1[0][0]+sizes2[0][0]) > matrix.Abc(p1[0] - p2[0])&&(sizes1[0][1]+sizes2[0][1]) > matrix.Abc(p1[1] - p2[1]) {
        a.Dynamik.Connect_to_Bottom = true
        a.Dynamik.Velocity.A[2] = 0
        a.Dynamik.Accelerate.A[2] = 0
        return true
      }
    }
  }else{
    
    var FacePoints1 = a.Core.FaceCentres()
    var FacePoints2 = b.Core.FaceCentres()
    
    p1 := a.Dynamik.Location.A
    v1 := a.Dynamik.Velocity.A
    p2 := b.Dynamik.Location.A
    v2 := b.Dynamik.Velocity.A
    
    v1[2] = -100
    
    x_dist := (FacePoints1[0].Z + p1[2]) - (FacePoints2[1].Z + p2[2])
    
    x_velocity := v1[2] - v2[2]
    
    t := (x_dist / x_velocity)
    
    if (t < 1)&&(t > 0) {
        a.DrawColor = color.RGBA{0,0,255,255}
      var sizes1 = a.Core.FaceArea()
      var sizes2 = b.Core.FaceArea()
      
      if (sizes1[0][0]+sizes2[0][0]) > matrix.Abc(p1[0] - p2[0])&&(sizes1[0][1]+sizes2[0][1]) > matrix.Abc(p1[1] - p2[1]) {
        return true
      }
    }
  }
  a.Dynamik.Connect_to_Bottom = false
  a.DrawColor = color.RGBA{255,255,255,255}
  return false
}
