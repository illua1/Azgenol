package main

import (
	"log"
  //"math"
	"github.com/hajimehoshi/ebiten/v2"
  sort "github.com/illua1/go-helpful/Sort"
  matrix "github.com/illua1/go-helpful/VectorMatrix"
)

type Programm struct {}

var(
  r [3]float64
  
  
  camera = NewCamera(0,0,0)
  player = NewPlayer(0, 0, 100)
  cube = NewCube(200, 200, 10, 0, 0, 0, Block_plit_face)
)

func (g *Programm) Update() error {
	
  
  
  
  if ebiten.IsKeyPressed(ebiten.KeyW) {
    player.This.Dynamik.VelocityAdd(0, -5, 0)
  }
  if ebiten.IsKeyPressed(ebiten.KeyS) {
    player.This.Dynamik.VelocityAdd(0, 5, 0)
  }
  if ebiten.IsKeyPressed(ebiten.KeyD) {
    player.This.Dynamik.VelocityAdd(-5, 0, 0)
  }
  if ebiten.IsKeyPressed(ebiten.KeyA) {
    player.This.Dynamik.VelocityAdd(5, 0, 0)
  }
  if ebiten.IsKeyPressed(ebiten.KeySpace) {
    player.This.Dynamik.AccelerateAdd(0, 0, 2)
  }
  player.This.Dynamik.AccelerateAdd(0,0, player.This.Dynamik.Gravity)
  
  BoxesColiseTest(&player.This, &cube.This)
  
  player.This.Dynamik.Update()
  
  {
    //r[0] = math.Pi/4
    //r[0] = math.Pi/2
    //r[2] = math.Pi/4
    /*
    */
    x, y := ebiten.CursorPosition()
    r[2] = float64(x)/100
    r[0] = float64(y)/100
    camera.SetMatrix(matrix.Rotate3x3_YXZ[float64](r))
  }
  
  for i := range camera.Location.A {
    camera.Location.A[i] = matrix.Lerp(camera.Location.A[i], player.This.Dynamik.Location.A[i], 0.06)
  }
  
  return nil
}

func (g *Programm) Draw(screen *ebiten.Image) {
  x,y := screen.Size()
  //op := &ebiten.DrawImageOptions{}
  ScreenGeom := ebiten.GeoM{}
  ScreenGeom.Scale(float64(sort.MaxF(x, y))/1000, float64(sort.MaxF(x, y))/1000)
  ScreenGeom.Translate(float64(x/2), float64(y/2))
  GeomMatrix := ebiten.GeoM{}
  camera.Matrix.Slise(0,0,2,2).FillTo(GeomMatrix.SetElement)
  GeomMatrix.Translate(GeomMatrix.Apply(-camera.Location.A[0], -camera.Location.A[1]))
  
  cube.Draw(screen, ScreenGeom, &camera)
  player.Draw(screen, ScreenGeom, &camera)
}

func (g *Programm) Layout(outsideWidth, outsideHeight int) (int, int) {
  return outsideWidth, outsideHeight
}

func main() {
  
	ebiten.SetWindowTitle("Azgenol")
  ebiten.SetWindowResizable(true)
  
  prog := Programm{}
  
	if err := ebiten.RunGame(&prog); err != nil {
		log.Fatal(err)
	}
}