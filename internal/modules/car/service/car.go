package service

import (
	"goTest/internal/infrastructure/component"
	"goTest/internal/infrastructure/godecoder"
	"goTest/internal/models"
	"goTest/internal/modules/car/storage"

	"go.uber.org/zap"
	"golang.org/x/net/context"
)

type Carer interface {
	CreateCars(ctx context.Context, RegNums []string) error
	DeleteCar(ctx context.Context, id int) error
	UpdateCar(ctx context.Context, id int, car models.Car) error
	GetCars(ctx context.Context, filters map[string]string, limit, offset int) ([]models.Car, error)
}

type Car struct {
	storage.Carer
	Decoder godecoder.Decoder
	Logger  *zap.Logger
	api     Api
}

func NewCarService(CarerRep storage.Carer, components *component.Components) Carer {
	return &Car{
		Carer:   CarerRep,
		Logger:  components.Logger,
		Decoder: components.Decoder,
		api:     &externalApi{apiURL: components.Conf.ApiURL, loggger: components.Logger},
	}
}

func (c *Car) CreateCars(ctx context.Context, RegNums []string) error {
	cars := make([]models.Car, 0)

	for _, regNum := range RegNums {
		car, err := c.api.Get(regNum)
		if err != nil {
			return err
		}
		cars = append(cars, car)
	}

	return c.Carer.CreateCars(ctx, cars)
}
func (c *Car) DeleteCar(ctx context.Context, id int) error {
	return c.Carer.DeleteCar(ctx, id)
}
func (c *Car) UpdateCar(ctx context.Context, id int, car models.Car) error {
	return c.Carer.UpdateCar(ctx, id, car)
}

func (c *Car) GetCars(ctx context.Context, filters map[string]string, limit, offset int) ([]models.Car, error) {
	return c.Carer.GetCars(ctx, filters, limit, offset)
}
