package main

import (
	"log"
  "math"
  "time"
  
	entities "github.com/illua1/Azgenol/AzgenolExecLib/AEL_Entities"
  
	types "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_Types"
	components "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_Components"

	sort "github.com/illua1/go-helpful/Sort"

	"github.com/hajimehoshi/ebiten/v2"
)

type Programm struct{
  System components.ComponentSystem
}

var (
	time_ float64
	camera = types.NewCamera(0, 0, 0)
)

func (g *Programm) Update() error {

	time_ += 0.03

	camera.SetAngle(-math.Pi/4, 0, -math.Pi/4)
	x, y := ebiten.CursorPosition()
	camera.SetAngle(
		float64(y)/100,
		0.0,
		-float64(x)/100,
	)

	return nil
}

func (g *Programm) Draw(screen *ebiten.Image) {
  
	x, y := screen.Size()
	var geom = ebiten.GeoM{}
	geom.Scale(float64(sort.MaxF(x, y))/1000, float64(sort.MaxF(x, y))/1000)
	geom.Translate(float64(x/2), float64(y/2))
  
  g.System.Update(
    types.Context{
      Screen : screen,
      Camera : camera,
      Time : time.Duration(0),
    },
  )
}

func (g *Programm) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func main() {

	ebiten.SetWindowTitle("Azgenol")
	ebiten.SetWindowResizable(true)

	prog := Programm{
    System : components.NewComponentSystem(
      components.NewComponentProcessRender(),
      components.NewComponentProcessKinematic(),
    ),
  }
  
  var Block1 = entities.NewBlock(
    0,
    0,
    0,
    100,
    100,
    100,
    Block_wall_face,
    false,
  )
  
  var Block2 = entities.NewBlock(
    0,
    0,
    200,
    100,
    100,
    100,
    Block_wall_face,
    true,
  )
  
  prog.System.Add(
    &Block1,
    &Block2,
  )

	if err := ebiten.RunGame(&prog); err != nil {
		log.Fatal(err)
	}
}
