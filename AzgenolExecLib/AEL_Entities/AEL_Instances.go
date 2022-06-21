package AEL_EBlock

import (
	pipeline "github.com/illua1/Azgenol/AzgenolKernelLib/AKL_RenderPipeLine"

	componentExec "github.com/illua1/Azgenol/AzgenolExecLib/AEL_Components"
)

type InstancesEmiter[Object any] struct {
	count    int
	instance Object
}

func NewEmiter[Object any](count int, instance Object) InstancesEmiter[Object] {
	return InstancesEmiter[Object]{
		count:    count,
		instance: instance,
	}
}

func (emiter *InstancesEmiter[Object]) GetRenderObject() pipeline.RenderObject {
	if instance, ok := any(&emiter.instance).(componentExec.InstanceObject); ok {
		if ret, err := componentExec.NewInstancesEmiterComponent(&emiter.count, instance); err == nil {
			return ret
		}
	}
	return nil
}
