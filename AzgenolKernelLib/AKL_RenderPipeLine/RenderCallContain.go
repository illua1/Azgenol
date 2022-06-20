package AKL_RenderPipeLine

import (
	"fmt"

	draw "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_Drawers"
	types "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_Types"

	node "github.com/illua1/go-helpful/Node"
)

type RenderCallContain struct {
	deph float64
	draw.ImageDrawer
}

type RenderCall node.BNode[RenderCallContain]

func (rCallContain RenderCallContain) String() string {
	return fmt.Sprint(rCallContain.deph)
}

func NewRenderCall(IDrawer draw.ImageDrawer, deph float64) RenderCall {
	return RenderCall(node.BNode[RenderCallContain]{nil, nil, RenderCallContain{deph, IDrawer}})
}

type RenderCallAppend func(draw.ImageDrawer, float64)

type RenderObject interface {
	RenderCustom(RenderCallAppend, *types.Camera)
}
