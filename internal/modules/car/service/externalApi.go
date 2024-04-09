package service

import (
	"encoding/json"
	"goTest/internal/models"
	"io"
	"net/http"

	"go.uber.org/zap"
)

type Api interface {
	Get(regNum string) (models.Car, error)
}

type externalApi struct {
	apiURL  string
	loggger *zap.Logger
}

func (e externalApi) Get(regNum string) (models.Car, error) {
	var car models.Car

	client := &http.Client{}
	req, err := http.NewRequest("GET", e.apiURL+"/info?regNum="+regNum, nil)
	if err != nil {

		return models.Car{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return models.Car{}, err
	}

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.Car{}, err
	}
	err = json.Unmarshal(bodyText, &car)
	if err != nil {
		return models.Car{}, err
	}
	return car, nil
}
