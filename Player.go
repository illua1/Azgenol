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
          //10,
          0,
          false,
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

func(player *Player)Draw(screen *ebiten.Image, objectm ObjectMatrix, camera *Camera){
  //player.This.Draw(screen, screen_geom, &camera.Matrix)
}

func(player *Player) RenderCustom(append_draw RenderCallAppend, objectm ObjectMatrix, camera *Camera){
  append_draw(
    player.Item,
    Value(player.This.Dynamik.Location.A[0]),
  )
}