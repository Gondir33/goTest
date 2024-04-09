package modules

import (
	"goTest/internal/infrastructure/component"
	cHandler "goTest/internal/modules/car/controller"
)

type Controllers struct {
	cHandler.Carer
}

func NewControllers(services *Services, components *component.Components) *Controllers {
	return &Controllers{
		Carer: cHandler.NewCarHandler(services.Carer, components.Responder, components.Decoder, components.Logger),
	}
}
