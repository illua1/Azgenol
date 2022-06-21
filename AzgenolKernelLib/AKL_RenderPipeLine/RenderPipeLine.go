package AKL_RenderPipeLine

import (
	"fmt"

	draw "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_Drawers"
	types "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_Types"

	node "github.com/illua1/go-helpful/Node"
	sort "github.com/illua1/go-helpful/Sort"

	"github.com/hajimehoshi/ebiten/v2"
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

type RenderPipeLine struct {
	Objects *node.LNode[RenderObject]
	First   *RenderCall
	Debug   bool
}

type RenderObjectDebug interface {
	Draw(*ebiten.Image, *types.Camera)
}

func NewRenderPipeLine(in ...RenderObject) (ret RenderPipeLine) {
	for i := range in {
		ret.Add(in[i])
	}
	return
}

func (rp *RenderPipeLine) Add(in RenderObject) {
	var object = node.Append(&rp.Objects, in)
	if instance, ok := in.(types.DeleteObject); ok {
		instance.SetDelete(object.Del)
	}
}

func (rp *RenderPipeLine) Draw(screen *ebiten.Image, camera *types.Camera) {
	node.For(
		&rp.Objects,
		func(in RenderObject) {
			in.RenderCustom(
				RenderCallAppend(func(IDrawer draw.ImageDrawer, deph float64) {
					var call = NewRenderCall(IDrawer, deph)
					if rp.First != nil {
						node.BNodeDescrent(
							(*node.BNode[RenderCallContain])(rp.First),
							(*node.BNode[RenderCallContain])(&call),
							func(a, b *node.BNode[RenderCallContain]) bool {
								return a.Contain.deph < b.Contain.deph
							},
						)
					} else {
						rp.First = &call
					}
				}),
				&*camera, // Protect original camera info
			)
		},
	)

	x, y := screen.Size()
	var geom = ebiten.GeoM{}
	geom.Scale(float64(sort.MaxF(x, y))/1000, float64(sort.MaxF(x, y))/1000)
	geom.Translate(float64(x/2), float64(y/2))

	if rp.First != nil {
		node.BNodeForTo((*node.BNode[RenderCallContain])(rp.First), func(_ int, contain RenderCallContain) {
			contain.Draw(screen, geom)
		})
	}

	if rp.Debug {
		node.For(
			&rp.Objects,
			func(in RenderObject) {
				if obj, ok := in.(RenderObjectDebug); ok {
					obj.Draw(screen, camera)
				}
			},
		)
	}
	rp.First = nil
}
