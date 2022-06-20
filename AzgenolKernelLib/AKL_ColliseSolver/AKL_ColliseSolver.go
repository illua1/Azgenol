package AKL_TypeCollise

import(
  "time"
  
  types "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_Types"
  
  node "github.com/illua1/go-helpful/Node"
)

type ColliseSolver struct {
  list *node.LNode[ColliseObject]
  time.Duration
}

func NewColliseSolver() ColliseSolver {
  return ColliseSolver{}
}

func (cSolver *ColliseSolver) Add (in ColliseObject){
  var object = node.Append(&cSolver.list, in)
  if instance, ok := in.(types.DeleteObject); ok {
    instance.SetDelete(object.Del)
  }
}

type ColliseObject interface {
  //Step(DeltaTime float64)
}

func (cSolver *ColliseSolver) Update (context types.Context) {
  
  //var DeltaTime = (context.Time - cSolver.Duration).Seconds()
  
  node.For(
    &cSolver.list,
    func(kObject ColliseObject){
      
    },
  )
  
  cSolver.Duration = context.Time
}