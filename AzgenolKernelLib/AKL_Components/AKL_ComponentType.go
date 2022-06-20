package AKL_TypeComponents

import(
  types "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_Types"
)

type ComponentProcess interface {
  Add(interface{})
  Update(types.Context)
}

