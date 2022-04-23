package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Stand struct {
  This Box
  Item ImageDrawer
}

type Player Stand

func NewPlayer(x,y,z float64)Player{
  return Player(
    Stand{
      NewBox(
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
