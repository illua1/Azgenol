package AKL_TypeComponents

import (
  types "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_Types"
  collise "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_ColliseSolver"
)

type ComponentProceseCollise struct{
  collise.ColliseSolver
}

func NewComponentProcessCollise()*ComponentProceseCollise{
  return &ComponentProceseCollise{
    collise.NewColliseSolver(),
  }
}

type IsColliseObject interface {
  GetColliseObject()collise.ColliseObject
}

func (cpCollise *ComponentProceseCollise) Add (in interface{}){
  if instance, ok := in.(IsColliseObject); ok {
    cpCollise.ColliseSolver.Add(instance.GetColliseObject())
  }
}

func (cpCollise *ComponentProceseCollise) Update (context types.Context){
  cpCollise.ColliseSolver.Update(context)
}