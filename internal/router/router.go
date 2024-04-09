package router

import (
	"goTest/internal/infrastructure/component"
	"goTest/internal/modules"
	"net/http"

	"github.com/go-chi/chi"
)

func NewApiRouter(controllers *modules.Controllers, components *component.Components) http.Handler {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Get("/car", controllers.Carer.GetCar)
		r.Delete("/car", controllers.DeleteCar)
		r.Put("/car", controllers.UpdateCar)
		r.Post("/car", controllers.CreateCars)
	})
	return r
}
