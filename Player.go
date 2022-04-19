package main

import (
	"github.com/hajimehoshi/ebiten/v2"
  matrix "github.com/illua1/go-helpful/VectorMatrix"
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
          100,
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
  
  var FacePoints1 = a.Core.FaceCentres()
  var FacePoints2 = b.Core.FaceCentres()
  
  p1 := a.Dynamik.Location.A
  v1 := a.Dynamik.Velocity.A
  p2 := b.Dynamik.Location.A
  v2 := b.Dynamik.Velocity.A
  
  x_dist := (FacePoints1[0].Z + p1[2]) - (FacePoints2[0].Z + p2[2])
  
  x_velocity := v1[2] - v2[2]
  
  t := (x_dist / x_velocity)
  
  //log.Print(x_dist, x_velocity, t)
  
  if (t < 1)&&(t > 0) {
    
    var sizes1 = a.Core.FaceArea()
    var sizes2 = b.Core.FaceArea()
    
    if (sizes1[0][0]+sizes2[0][0]) > matrix.Abc(p1[0] - p2[0])&&(sizes1[0][1]+sizes2[0][1]) > matrix.Abc(p1[1] - p2[1]) {
      a.Dynamik.Velocity.A[2] = -a.Dynamik.Velocity.A[2]*0.9
      b.Dynamik.Velocity.A[2] = -b.Dynamik.Velocity.A[2]*0.9
      a.Dynamik.Accelerate.A[2] = -a.Dynamik.Accelerate.A[2]*0.9
      b.Dynamik.Accelerate.A[2] = -b.Dynamik.Accelerate.A[2]*0.9
      
      return true
    }
  }
  return false
}
