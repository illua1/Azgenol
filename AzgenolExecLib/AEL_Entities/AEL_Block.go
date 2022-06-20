package AEL_EBlock

import (
	collise "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_ColliseSolver"
	kinematic "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_KinematicSolver"
	render "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_Render"
	pipeline "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_RenderPipeLine"
	types "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_Types"

	componentExec "github.com/illua1/Azgenol/AzgenolExecLib/AEL_Components"

	matrix "github.com/illua1/go-helpful/VectorMatrix"
	volume "github.com/illua1/go-helpful/Volume"

	"github.com/hajimehoshi/ebiten/v2"
)

type Block struct {
	Location types.Vector3

	Velocity types.Vector3

	Matrix types.Matrix3

	BoundBox volume.Box[float64]

	Box render.Box

	PhysicMove bool
}

func NewBlock(x, y, z float64, xs, ys, zs float64, img *ebiten.Image, isMove bool) Block {
	return Block{
		Location:   types.NewVector3(x, y, z),
		Matrix:     types.Matrix3{matrix.Matrix3x3[float64]()},
		BoundBox:   volume.NewBox[float64](-xs, -ys, -zs, xs, ys, zs),
		Box:        render.NewBox(xs, ys, zs, img),
		PhysicMove: isMove,
	}
}

func (block *Block) GetRenderObject() pipeline.RenderObject {
	var cube = componentExec.NewRenderComponent(
		&block.Box,
		&block.Location,
		&block.Matrix,
	)
	return &cube
}

func (block *Block) GetKinematicObject() kinematic.KinematicObject {
	var move = componentExec.NewKinematicComponent(
		&block.Location,
		&block.Velocity,
		&block.PhysicMove,
	)
	return &move
}

func (block *Block) GetColliseObject() collise.ColliseObject {
	var collise = componentExec.NewColliseComponent()
	return &collise
}
