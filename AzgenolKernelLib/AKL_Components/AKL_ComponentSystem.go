package AKL_TypeComponents

import (
	node "github.com/illua1/go-helpful/Node"
)

type ComponentProcess[ComponentContext any] interface {
	Add(interface{})
	Update(ComponentContext)
}

type ComponentSystem[ComponentContext any] struct {
	list *node.LNode[ComponentProcess[ComponentContext]]
}

func NewComponentSystem[ComponentContext any](in ...ComponentProcess[ComponentContext]) (ret ComponentSystem[ComponentContext]) {
	if len(in) > 0 {
		ret.list = *node.NewLNode(in[0])
		for i := range in[1:] {
			node.Append(&ret.list, in[i])
		}
	}
	return
}

func (cSystem *ComponentSystem[ComponentContext]) Add(in ...interface{}) {
	node.For(
		&cSystem.list,
		func(cProcess ComponentProcess[ComponentContext]) {
			for e := range in {
				cProcess.Add(in[e])
			}
		},
	)
}

func (cSystem *ComponentSystem[ComponentContext]) Update(context ComponentContext) {
	node.For(
		&cSystem.list,
		func(cProcess ComponentProcess[ComponentContext]) {
			cProcess.Update(context)
		},
	)
}
