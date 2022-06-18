package main

import (
	"log"
  "math"
  //"time"
  "image/color"
	"github.com/hajimehoshi/ebiten/v2"
  sort "github.com/illua1/go-helpful/Sort"
  matrix "github.com/illua1/go-helpful/VectorMatrix"
  
  types "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_Types"
  //draw "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_Drawers"
  render "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_Render3D"
  pipeline "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_RenderPipeLine"
)

type Programm struct {}

var(
  r [3]float64
  
  time float64
  
  JumpFlag bool = false
  
  camera = types.NewCamera(0,0,0)
  
  //player0 = NewPlayer(0, 0, 1400)
  player = NewPlayer(0, 0, 100)
  
  cube = NewCube(200, 200, 10, 0, 0, 0, Block_plit_face)
  cube1 = NewCube(200, 200, 10, 50, 0, -150, Block_plit_face)
  cube2 = NewCube(200, 200, 10, 100, 0, -300, Block_plit_face)
  
  phys = NewPhysicsCollectro(
    //&player0.This,
    &player.This,
    &cube.This,
    &cube1.This,
    &cube2.This,
  )
  /*
  render = NewRenderPipeLine(
    &player,
    //&player0,
    &cube,
    &cube1,
    &cube2,
  )*/
)

func (g *Programm) Update() error {
  //time.Sleep(time.Second/14)
  
  time += 0.1
  
  player.This.DrawColor = color.RGBA{255,0,0,255}
  
  if ebiten.IsKeyPressed(ebiten.KeyW) {
    player.This.Dynamik.VelocityAdd(0, -5, 0)
  }
  if ebiten.IsKeyPressed(ebiten.KeyS) {
    player.This.Dynamik.VelocityAdd(0, 5, 0)
  }
  if ebiten.IsKeyPressed(ebiten.KeyD) {
    player.This.Dynamik.VelocityAdd(5, 0, 0)
  }
  if ebiten.IsKeyPressed(ebiten.KeyA) {
    player.This.Dynamik.VelocityAdd(-5, 0, 0)
  }
  if ebiten.IsKeyPressed(ebiten.KeySpace) {
    if !JumpFlag {
      player.This.Dynamik.VelocityAdd(0, 0, 400)
      JumpFlag = true
    }
  } else {
    JumpFlag = false
  }
  
  phys.Update()
  
  {
    //r[0] = math.Pi/4
    r[0] = math.Pi/4
    r[2] = math.Pi/4
    x, y := ebiten.CursorPosition()
    r[2] = float64(x)/100
    r[0] = float64(y)/100
    /*
    */
    camera.SetMatrix(types.Matrix3{matrix.Rotate3x3_ZYX[float64](r)})
  }
  /*
  for i := range camera.Location.A {
    camera.Location.A[i] = matrix.Lerp(camera.Location.A[i], player.This.Dynamik.Location.A[i], 0.06)
  }
  */
  return nil
}

func (g *Programm) Draw(screen *ebiten.Image) {
  //var Frame_Camear = camera
  
  x,y := screen.Size()
  var geom = ebiten.GeoM{}
  geom.Scale(float64(sort.MaxF(x, y))/1000, float64(sort.MaxF(x, y))/1000)
  geom.Translate(float64(x/2), float64(y/2))
  
  var cube0 = render.NewCube(
    render.NewBox(100,100,100, Block_plit_face),
    types.NewGeoM(150,150,0),
  )
  var geom3_0 = types.NewGeoM(0,150,0)
  geom3_0.Matrix = matrix.Rotate3x3_YXZ[float64]([3]float64{-time,0,0})
  var cube1 = render.NewCube(
    render.NewBox(100,100,100, Block_plit_face),
    geom3_0,
  )
  var cube2 = render.NewCube(
    render.NewBox(100,100,100, Block_plit_face),
    types.NewGeoM(-150,150,0),
  )
  
  var geom3_1 = types.NewGeoM(-150,0,0)
  geom3_1.Matrix = matrix.Rotate3x3_YXZ[float64]([3]float64{0,-time,0})
  var geom3_2 = types.NewGeoM(150,0,0)
  geom3_2.Matrix = matrix.Rotate3x3_YXZ[float64]([3]float64{0,time,0})
  var geom3_3 = types.NewGeoM(0,0,0)
  geom3_3.Matrix = matrix.Rotate3x3_YXZ[float64]([3]float64{0,0,time})
  var cube3 = render.NewCube(
    render.NewBox(100,100,100, Block_plit_face),
    geom3_1,
  )
  var cube4 = render.NewCube(
    render.NewBox(100,100,100, Block_plit_face),
    geom3_3,
  )
  var cube5 = render.NewCube(
    render.NewBox(100,100,100, Block_plit_face),
    geom3_2,
  )
  var geom3 = types.NewGeoM(0,-150,0)
  geom3.Matrix = matrix.Rotate3x3_YXZ[float64]([3]float64{time,0,0})
  var cube6 = render.NewCube(
    render.NewBox(100,100,100, Block_plit_face),
    types.NewGeoM(-150,-150,0),
  )
  var cube7 = render.NewCube(
    render.NewBox(100,100,100, Block_plit_face),
    geom3,
  )
  var cube8 = render.NewCube(
    render.NewBox(100,100,100, Block_plit_face),
    types.NewGeoM(150,-150,0),
  )
  
  var cube10 = render.NewCube(
    render.NewBox(50,50,100, Block_plit_face),
    geom3_3.Concat(types.NewGeoM(0,50,100)),
  )
  
  var pip = pipeline.NewRenderPipeLine(
    &cube0,
    &cube1,
    &cube2,
    &cube3,
    &cube4,
    &cube5,
    &cube6,
    &cube7,
    &cube8,
    &cube10,
  )
  pip.Draw(screen, &camera)
  
  
  /*
  var box = render.NewBox(100,100,100, Block_plit_face)
  for _, img := range box.Render(&camera, types.NewVector3(0,0,0)) {
    img.Draw(screen, geom)
  for _, img := range box.Render(&camera, types.NewVector3(0,150,0)) {
    img.Draw(screen, geom)
  }
  for _, img := range box.Render(&camera, types.NewVector3(0,-150,0)) {
    img.Draw(screen, geom)
  }
  for _, img := range box.Render(&camera, types.NewVector3(150,0,0)) {
    img.Draw(screen, geom)
  }
  for _, img := range box.Render(&camera, types.NewVector3(-150,0,0)) {
    img.Draw(screen, geom)
  }
  }*/
  
  //render.RenderDrawer(draw.NewImageDrawer(50, 50, Block_plit_face), &camera, types.NewVector3(0,0,100)).Draw(screen, geom)
  //render.RenderDrawer(draw.NewImageDrawer(50, 50, Block_plit_face), &camera, types.NewVector3(50,0,0)).Draw(screen, geom)
  //render.RenderDrawer(draw.NewImageDrawer(50, 50, Block_plit_face), &camera, types.NewVector3(0,50,0)).Draw(screen, geom)
  //render.RenderDrawer(draw.NewImageDrawer(50, 50, Block_plit_face), &camera, types.NewVector3(50,50,0)).Draw(screen, geom)
  
  
  
  //render.Draw(screen, &Frame_Camear)
  
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