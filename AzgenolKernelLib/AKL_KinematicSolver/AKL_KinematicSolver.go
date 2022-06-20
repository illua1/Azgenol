package AKL_TypeKinematic

import(
 // "time"
  
  types "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_Types"
  
  node "github.com/illua1/go-helpful/Node"
)

type KinematicSolver struct {
  list *node.LNode[KinematicObject]
}

func NewKinematicSolver() KinematicSolver {
  return KinematicSolver{}
}

func (kObject *KinematicSolver) Add (in KinematicObject){
  node.Append(&kObject.list, in)
}

type KinematicObject interface {
  
}

func (kObject *KinematicSolver) Update (context types.Context) {
  
}