package AKL_TypeComponents

import (
	kinematic "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_KinematicSolver"
	types "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_Types"
)

type ComponentProcessKinematic struct {
	kinematic.KinematicSolver
}

func NewComponentProcessKinematic() *ComponentProcessKinematic {
	return &ComponentProcessKinematic{
		kinematic.NewKinematicSolver(),
	}
}

type IsKinematicObject interface {
	GetKinematicObject() kinematic.KinematicObject
}

func (cpKinematic *ComponentProcessKinematic) Add(in interface{}) {
	if in == nil {
		return
	}
	if instance, ok := in.(IsKinematicObject); ok {
		cpKinematic.KinematicSolver.Add(instance.GetKinematicObject())
	}
}

func (cpKinematic *ComponentProcessKinematic) Update(context types.Context) {
	cpKinematic.KinematicSolver.Update(context)
}
