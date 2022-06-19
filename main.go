package main

import (
	"log"
  "math"
  
	types "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_Types"
	render "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_Render3D"
	components "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_Components"

	sort "github.com/illua1/go-helpful/Sort"
	matrix "github.com/illua1/go-helpful/VectorMatrix"

	"github.com/hajimehoshi/ebiten/v2"
)

type Programm struct{
  components.ComponentCollector
}

var (
	time float64

	camera = types.NewCamera(0, 0, 0)
)

var(
  cube0 = render.NewCube(
		render.NewBox(100, 100, 100, Block_wall_face),
		types.NewGeoM(150, 150, 0),
	)
	cube1 = render.NewCube(
		render.NewBox(100, 100, 100, Block_wall_face),
		types.NewGeoM(0, 150, 0),
	)
	cube2 = render.NewCube(
		render.NewBox(100, 100, 100, Block_wall_face),
		types.NewGeoM(-150, 150, 0),
	)
  
	cube3 = render.NewCube(
		render.NewBox(100, 100, 100, Block_wall_face),
		types.NewGeoM(-150, 0, 0),
	)
	cube4 = render.NewCube(
		render.NewBox(100, 100, 100, Block_wall_face),
		types.NewGeoM(150, 0, 0),
	)
	cube5 = render.NewCube(
		render.NewBox(100, 100, 100, Block_wall_face),
		types.NewGeoM(0, 0, 0),
	)
  
	cube6 = render.NewCube(
		render.NewBox(100, 100, 100, Block_wall_face),
		types.NewGeoM(-150, -150, 0),
	)
	cube7 = render.NewCube(
		render.NewBox(100, 100, 100, Block_wall_face),
		types.NewGeoM(0, -150, 0),
	)
	cube8 = render.NewCube(
		render.NewBox(100, 100, 100, Block_wall_face),
		types.NewGeoM(150, -150, 0),
	)

	cube10 = render.NewCube(
		render.NewBox(50, 50, 100, Block_wall_face),
		geom3_3.Concat(types.NewGeoM(0, 50, 100)),
	)
  
  geom3_0 = &cube1.Matrix
  geom3_1 = &cube3.Matrix
  geom3_2 = &cube4.Matrix
  geom3_3 = &cube5.GeoM
  geom3_4 = &cube7.Matrix
  geom3_5 = &cube10.GeoM
)

func (g *Programm) Update() error {

	time += 0.03
/*
	x, y := ebiten.CursorPosition()
	camera.SetAngle(
		float64(y)/100,
		0.0,
		-float64(x)/100,
	)
*/
	camera.SetAngle(-math.Pi/4, 0, -math.Pi/4)
	return nil
}

func (g *Programm) Draw(screen *ebiten.Image) {
  
	x, y := screen.Size()
	var geom = ebiten.GeoM{}
	geom.Scale(float64(sort.MaxF(x, y))/1000, float64(sort.MaxF(x, y))/1000)
	geom.Translate(float64(x/2), float64(y/2))


  *geom3_0 = matrix.Rotate3x3_YXZ[float64]([3]float64{-time, 0, 0})
  *geom3_1 = matrix.Rotate3x3_YXZ[float64]([3]float64{0, time, 0})
  *geom3_2 = matrix.Rotate3x3_YXZ[float64]([3]float64{0, -time, 0})
  geom3_3.Matrix = matrix.Rotate3x3_YXZ[float64]([3]float64{0, 0, time})
  *geom3_4 = matrix.Rotate3x3_YXZ[float64]([3]float64{time, 0, 0})
  *geom3_5 = geom3_3.Concat(types.NewGeoM(0, 50, 100))
  
	g.ComponentCollector.Draw(screen, &camera)
}

func (g *Programm) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func main() {

	ebiten.SetWindowTitle("Azgenol")
	ebiten.SetWindowResizable(true)

	prog := Programm{}

  prog.Add(
		&cube5,
		&cube0,
		&cube1,
		&cube2,
		&cube3,
		&cube4,
		&cube6,
		&cube7,
		&cube8,
		&cube10,
  )

	if err := ebiten.RunGame(&prog); err != nil {
		log.Fatal(err)
	}
}
