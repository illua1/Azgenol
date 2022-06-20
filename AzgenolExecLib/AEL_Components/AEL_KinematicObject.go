package AEL_components

import(
  types "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_Types"
)

type KinematicComponent struct {
  Location *types.Vector3
  Velocity *types.Vector3
  Move *bool
}

func NewKinematicComponent(Location *types.Vector3, Velocity *types.Vector3, Move *bool)KinematicComponent{
  return KinematicComponent{Location, Velocity, Move}
}

func (kinematicC *KinematicComponent)Step(DeltaTime float64){
  if *kinematicC.Move {
    kinematicC.Velocity.Add(types.NewVector3(0, 0, -10 * DeltaTime).Vector)
    kinematicC.Location.Add(kinematicC.Velocity.Vector)
  }
}