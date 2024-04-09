package modules

import (
	"goTest/internal/infrastructure/component"
	cService "goTest/internal/modules/car/service"
	"goTest/internal/storages"
)

type Services struct {
	cService.Carer
}

func NewServices(storages *storages.Storages, components *component.Components) *Services {
	return &Services{
		Carer: cService.NewCarService(storages.Carer, components),
	}
}
