package main

import (
	"log"

	render "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_Render3D"
	pipeline "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_RenderPipeLine"
	types "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_Types"

	sort "github.com/illua1/go-helpful/Sort"
	matrix "github.com/illua1/go-helpful/VectorMatrix"

	"github.com/hajimehoshi/ebiten/v2"
)

type Programm struct{}

var (
	time float64

	camera = types.NewCamera(0, 0, 0)
)

func (g *Programm) Update() error {

	time += 0.1

	x, y := ebiten.CursorPosition()
	camera.SetAngle(
		float64(y)/100,
		0.0,
		float64(x)/100,
	)

	return nil
}

func (g *Programm) Draw(screen *ebiten.Image) {

	x, y := screen.Size()
	var geom = ebiten.GeoM{}
	geom.Scale(float64(sort.MaxF(x, y))/1000, float64(sort.MaxF(x, y))/1000)
	geom.Translate(float64(x/2), float64(y/2))

	var cube0 = render.NewCube(
		render.NewBox(100, 100, 100, Block_plit_face),
		types.NewGeoM(150, 150, 0),
	)
	var geom3_0 = types.NewGeoM(0, 150, 0)
	geom3_0.Matrix = matrix.Rotate3x3_YXZ[float64]([3]float64{-time, 0, 0})
	var cube1 = render.NewCube(
		render.NewBox(100, 100, 100, Block_plit_face),
		geom3_0,
	)
	var cube2 = render.NewCube(
		render.NewBox(100, 100, 100, Block_plit_face),
		types.NewGeoM(-150, 150, 0),
	)

	var geom3_1 = types.NewGeoM(-150, 0, 0)
	geom3_1.Matrix = matrix.Rotate3x3_YXZ[float64]([3]float64{0, -time, 0})
	var geom3_2 = types.NewGeoM(150, 0, 0)
	geom3_2.Matrix = matrix.Rotate3x3_YXZ[float64]([3]float64{0, time, 0})
	var geom3_3 = types.NewGeoM(0, 0, 0)
	geom3_3.Matrix = matrix.Rotate3x3_YXZ[float64]([3]float64{0, 0, time})
	var cube3 = render.NewCube(
		render.NewBox(100, 100, 100, Block_plit_face),
		geom3_1,
	)
	var cube4 = render.NewCube(
		render.NewBox(100, 100, 100, Block_plit_face),
		geom3_3,
	)
	var cube5 = render.NewCube(
		render.NewBox(100, 100, 100, Block_plit_face),
		geom3_2,
	)
	var geom3 = types.NewGeoM(0, -150, 0)
	geom3.Matrix = matrix.Rotate3x3_YXZ[float64]([3]float64{time, 0, 0})
	var cube6 = render.NewCube(
		render.NewBox(100, 100, 100, Block_plit_face),
		types.NewGeoM(-150, -150, 0),
	)
	var cube7 = render.NewCube(
		render.NewBox(100, 100, 100, Block_plit_face),
		geom3,
	)
	var cube8 = render.NewCube(
		render.NewBox(100, 100, 100, Block_plit_face),
		types.NewGeoM(150, -150, 0),
	)

	var cube10 = render.NewCube(
		render.NewBox(50, 50, 100, Block_plit_face),
		geom3_3.Concat(types.NewGeoM(0, 50, 100)),
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
