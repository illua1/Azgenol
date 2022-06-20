package AKL_TypeKinematic

import(
  "time"
  
  types "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_Types"
  
  node "github.com/illua1/go-helpful/Node"
)

type KinematicSolver struct {
  list *node.LNode[KinematicObject]
  time.Duration
}

func NewKinematicSolver() KinematicSolver {
  return KinematicSolver{}
}

func (kSolver *KinematicSolver) Add (in KinematicObject){
  var object = node.Append(&kSolver.list, in)
  if instance, ok := in.(types.DeleteObject); ok {
    instance.SetDelete(object.Del)
  }
}

type KinematicObject interface {
  Step(DeltaTime float64)
}

func (kSolver *KinematicSolver) Update (context types.Context) {
  
  var DeltaTime = (context.Time - kSolver.Duration).Seconds()
  
  node.For(
    &kSolver.list,
    func(kObject KinematicObject){
      kObject.Step(DeltaTime)
    },
  )
  
  kSolver.Duration = context.Time
}