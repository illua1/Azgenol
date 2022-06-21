package AEL_components

import (
	"fmt"

	components "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_Components"
	pipeline "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_RenderPipeLine"
	types "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_Types"

	noise "github.com/aquilax/go-perlin"
)

type InstancesEmiterComponent struct {
	count *int

	instance_render pipeline.RenderObject
	instance_move_v *types.Vector3
	instance_move_m *types.Matrix3

	GeoMList *[]types.GeoM

	Time float64
}

type InstanceObject interface {
	components.IsRenderObject
	GetObjectToMove() (*types.Vector3, *types.Matrix3) // In the future, it will be moved to interface for system of moving
}

func NewInstancesEmiterComponent(count *int, instance InstanceObject) (*InstancesEmiterComponent, error) {
	inst_move_v, inst_move_m := instance.GetObjectToMove()
	if inst_rend := instance.GetRenderObject(); (inst_rend == nil) || (inst_move_v == nil) || (inst_move_m == nil) {
		return &InstancesEmiterComponent{}, fmt.Errorf("")
	} else {
		var list = make([]types.GeoM, *count)
		for i := range list {
			list[i] = types.NewGeoM(0, 0, 0)
		}
		return &InstancesEmiterComponent{
			count:           count,
			instance_render: inst_rend,
			instance_move_v: inst_move_v,
			instance_move_m: inst_move_m,
			GeoMList:        &list,
		}, nil
	}
}

func (InstancesEmiterC *InstancesEmiterComponent) RenderCustom(call pipeline.RenderCallAppend, camera *types.Camera) {

	if geomLen := len(*InstancesEmiterC.GeoMList); geomLen != *InstancesEmiterC.count {
		if geomLen < *InstancesEmiterC.count {
			(*InstancesEmiterC.GeoMList) = append(*InstancesEmiterC.GeoMList, make([]types.GeoM, *InstancesEmiterC.count-geomLen)...)
			for i := range (*InstancesEmiterC.GeoMList)[geomLen:] {
				(*InstancesEmiterC.GeoMList)[i] = types.NewGeoM(0, 0, 0)
			}
		} else {
			(*InstancesEmiterC.GeoMList) = (*InstancesEmiterC.GeoMList)[:*InstancesEmiterC.count]
		}
	}
	InstancesEmiterC.Time += 0.01

	for i := range *InstancesEmiterC.GeoMList {
		(*InstancesEmiterC.GeoMList)[i].Add(
			types.NewVector3(
				float64(int(noise.NewPerlin(2, 3, 8, int64(i)).Noise2D(InstancesEmiterC.Time, -15.1)*100)/50)*50,
				float64(int(noise.NewPerlin(2, 3, 8, int64(i)).Noise2D(InstancesEmiterC.Time, 42.5)*100)/50)*50,
				/*noise.NewPerlin(2, 3, 8, int64(i)).Noise2D(InstancesEmiterC.Time, -11.9)*20,*/ 0.0,
			).Vector,
		)

		*InstancesEmiterC.instance_move_v = (*InstancesEmiterC.GeoMList)[i].Vector3
		*InstancesEmiterC.instance_move_m = (*InstancesEmiterC.GeoMList)[i].Matrix3
		InstancesEmiterC.instance_render.RenderCustom(
			call,
			camera,
		)
	}
}
