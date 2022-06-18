package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	node "github.com/illua1/go-helpful/Node"
)

type ToFill struct {
	Contain   *Serface
	SortValue float64
	GeoM      ebiten.GeoM
}

type RenderList []node.BNode[ToFill]

func NewRenderList(size int) (ret RenderList) {
	ret = make([]node.BNode[ToFill], 0, size)
	return
}

func (rl RenderList) Draw(screen *ebiten.Image) {
	node.BNodeForTo(&rl[0], func(_ int, contain ToFill) {
		contain.Contain.Draw(screen, ebiten.GeoM{})
	})
}

type RenderFiller interface {
	IsUpdate() bool
	Len() int
	Fill(func(contain ToFill), *Camera)
}

type RenderListFills struct {
	Body    RenderList
	Fillers []RenderFiller
}

func NewRenderListFills(in ...RenderFiller) RenderListFills {
	return RenderListFills{
		Fillers: in,
	}
}

func (rlf *RenderListFills) Draw(screen *ebiten.Image, camera *Camera) {
	var accumulateLen int = 0
	for i := range rlf.Fillers {
		if rlf.Fillers[i].IsUpdate() {
			for i := range rlf.Fillers {
				accumulateLen += rlf.Fillers[i].Len()
			}
		}
	}
	//if len(rlf.Body) != accumulateLen {
	rlf.Body = NewRenderList(accumulateLen)
	//}
	for i := range rlf.Fillers {
		rlf.Fillers[i].Fill(
			func(contain ToFill) {
				rlf.Body = append(rlf.Body, node.BNode[ToFill]{nil, nil, contain})
			},
			camera,
		)
	}
	for i := 1; i < len(rlf.Body); i++ {
		node.BNodeDescrent(
			&rlf.Body[0],
			&rlf.Body[i],
			func(a, b *node.BNode[ToFill]) bool {
				return a.Contain.SortValue > b.Contain.SortValue
			},
		)
	}
	rlf.Body.Draw(screen)
}
