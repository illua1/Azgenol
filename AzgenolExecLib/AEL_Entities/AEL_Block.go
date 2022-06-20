package AEL_EBlock

import (
  types "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_Types"
	render "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_Render"
  pipeline "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_RenderPipeLine"
  kinematic "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_KinematicSolver"
  
  componentExec "github.com/illua1/Azgenol/AzgenolExecLib/AEL_Components"
  
	volume "github.com/illua1/go-helpful/Volume"
	matrix "github.com/illua1/go-helpful/VectorMatrix"
  
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

func NewBlock(x, y, z float64, xs, ys, zs float64, img *ebiten.Image, isMove bool)Block{
  return Block{
    Location : types.NewVector3(x,y,z),
    Matrix : types.Matrix3{matrix.Matrix3x3[float64]()},
    BoundBox : volume.NewBox[float64](-xs, -ys, -zs, xs, ys, zs),
    Box : render.NewBox(xs, ys, zs, img),
    PhysicMove : isMove,
  }
}

func(block *Block)GetRenderObject()pipeline.RenderObject{
  var cube = componentExec.NewRenderComponentD(
    &block.Box,
    &block.Location,
    &block.Matrix,
  )
  return &cube
}

func(block *Block)GetKinematicObject()kinematic.KinematicObject{
  var move = componentExec.NewKinematicComponent()
  return &move
}