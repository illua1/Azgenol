package main

import (
	node "github.com/illua1/go-helpful/Node"
)

type RenderCallContain struct {
	deph Value
	ImageDrawer
}

type RenderCall node.BNode[RenderCallContain]

func NewRenderCall(imgd ImageDrawer, deph Value) RenderCall {
	return RenderCall(node.BNode[RenderCallContain]{nil, nil, RenderCallContain{deph, imgd}})
}

type RenderCallAppend func(ImageDrawer, Value)

type RenderObject interface {
	RenderCustom(RenderCallAppend, ObjectMatrix, *Camera)
}
