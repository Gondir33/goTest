package controller

import (
	"goTest/internal/infrastructure/responder"
	"goTest/internal/models"
	"goTest/internal/modules/car/service"
	"net/http"
	"strconv"

	"goTest/internal/infrastructure/godecoder"

	"go.uber.org/zap"
)

type Carer interface {
	GetCar(w http.ResponseWriter, r *http.Request)
	DeleteCar(w http.ResponseWriter, r *http.Request)
	UpdateCar(w http.ResponseWriter, r *http.Request)
	CreateCars(w http.ResponseWriter, r *http.Request)
}

type Car struct {
	service.Carer
	Responder responder.Responder
	Decoder   godecoder.Decoder
	Logger    *zap.Logger
}

func NewCarHandler(Сar service.Carer, respond responder.Responder, Decoder godecoder.Decoder, Logger *zap.Logger) Carer {
	return &Car{
		Carer:     Сar,
		Responder: respond,
		Decoder:   Decoder,
		Logger:    Logger,
	}
}

// @Summary	GetCars with filter and pagination
// @Tags		car
// @Accept		json
// @Produce	json
// @Param		regNum		query		string	false	"string"
// @Param		mark		query		string	false	"string"
// @Param		model		query		string	false	"string"
// @Param		year		query		int		false	"2002"
// @Param		name		query		string	false	"string"
// @Param		surname		query		string	false	"string"
// @Param		patronymic	query		string	false	"string"
// @Success	200			{object}	[]models.Car
// @Failure      400
// @Failure      500
// @Router		/car [get]
func (c Car) GetCar(w http.ResponseWriter, r *http.Request) {
	filters := getFilters(r.URL)
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		c.Logger.Error("get car bad request", zap.Error(err))
		c.Responder.ErrorBadRequest(w, err)
		return
	}
	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil {
		c.Logger.Error("get car bad request", zap.Error(err))
		c.Responder.ErrorBadRequest(w, err)
		return
	}
	cars, err := c.Carer.GetCars(r.Context(), filters, limit, offset)
	if err != nil {
		c.Logger.Error("get car internal error", zap.Error(err))
		c.Responder.ErrorInternal(w, err)
		return
	}
	c.Responder.OutputJSON(w, cars)
}

// @Summary	DeleteCars By ID
// @Tags		car
// @Accept		json
// @Produce	json
// @Param		id	query	string	true "1"
// @Success	200
// @Failure      400
// @Failure      500
// @Router		/car [delete]
func (c Car) DeleteCar(w http.ResponseWriter, r *http.Request) {

	idQuery := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idQuery)
	if err != nil {
		c.Logger.Error("delete car bad request", zap.Error(err))
		c.Responder.ErrorBadRequest(w, err)
		return
	}
	if err := c.Carer.DeleteCar(r.Context(), id); err != nil {
		c.Logger.Error("delete car internal error", zap.Error(err))
		c.Responder.ErrorInternal(w, err)
		return
	}
	c.Responder.OutputJSON(w, http.StatusOK)
}

// @Summary	UpdateCars By ID
// @Tags		car
// @Accept		json
// @Produce	json
// @Param		id		query	string		true "1"
// @Param		request	body	models.Car	true "for update"
// @Success	200
// @Failure      400
// @Failure      500
// @Router		/car [put]
func (c Car) UpdateCar(w http.ResponseWriter, r *http.Request) {
	var reqBody models.Car

	idQuery := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idQuery)
	if err != nil {
		c.Logger.Error("delete car bad request", zap.Error(err))
		c.Responder.ErrorBadRequest(w, err)
		return
	}

	if err := c.Decoder.Decode(r.Body, &reqBody); err != nil {
		c.Logger.Error("update car bad request", zap.Error(err))
		c.Responder.ErrorBadRequest(w, err)
		return
	}

	if err := c.Carer.UpdateCar(r.Context(), id, reqBody); err != nil {
		c.Logger.Error("update car internal error", zap.Error(err))
		c.Responder.ErrorInternal(w, err)
		return
	}
	c.Responder.OutputJSON(w, http.StatusOK)
}

// @Summary	CreateCar By RegNum
// @Tags		car
// @Accept		json
// @Produce	json
// @Param		request	body	CreateRequest	true "Creating"
// @Success	200
// @Failure      400
// @Failure      500
// @Router		/car [post]
func (c Car) CreateCars(w http.ResponseWriter, r *http.Request) {
	var reqBody CreateRequest

	if err := c.Decoder.Decode(r.Body, &reqBody); err != nil {
		c.Logger.Error("create car bad request", zap.Error(err))
		c.Responder.ErrorBadRequest(w, err)
		return
	}

	if err := c.Carer.CreateCars(r.Context(), reqBody.RegNums); err != nil {
		c.Logger.Error("create car internal error", zap.Error(err))
		c.Responder.ErrorInternal(w, err)
		return
	}
	c.Responder.OutputJSON(w, http.StatusOK)
}
