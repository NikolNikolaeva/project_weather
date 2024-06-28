package services

import (
	"context"
	"errors"
	"net/http"

	"github.com/NikolNikolaeva/project_weather/config"
	api "github.com/NikolNikolaeva/project_weather/generated/api/project-weather/rest"
	"github.com/NikolNikolaeva/project_weather/repositories"
	"github.com/NikolNikolaeva/project_weather/resources"
)

// ForecastAPIService is a service that implements the logic for the ForecastAPIServicer
// This service should implement the business logic for every endpoint for the ForecastAPI API.
// Include any external packages or services that will be required by this service.
type ForecastAPIService struct {
	DB      repositories.ForecastRepo
	DBCity  repositories.CityRepo
	Convert resources.ConverterI
	handler WeatherAPIClient
	config  *config.ApplicationConfiguration
}

// NewForecastAPIService creates a default api service
func NewForecastAPIService(DB repositories.ForecastRepo, convert resources.ConverterI, handler WeatherAPIClient, config *config.ApplicationConfiguration, DBCity repositories.CityRepo) api.ForecastAPIServicer {
	return &ForecastAPIService{
		DB:      DB,
		Convert: convert,
		handler: handler,
		config:  config,
		DBCity:  DBCity,
	}
}

func (self *ForecastAPIService) getForecastsByCityId(ctx context.Context, cityId string) (api.ImplResponse, error) {

	if cityId == "" {
		return api.Response(http.StatusNotImplemented, "City id is required"), nil
	}
	forecasts, err := self.DB.FindByCityId(cityId)

	if err != nil {
		return api.Response(http.StatusInternalServerError, nil), err
	}

	var forecastsApi []api.Forecast

	for _, forecast := range forecasts {
		forecastApi := self.Convert.ConvertModelForecastToApiForecast(forecast)
		forecastsApi = append(forecastsApi, *forecastApi)
	}

	return api.Response(http.StatusOK, forecastsApi), nil
}

func (self *ForecastAPIService) getDays(period string) int {
	switch period {
	case "daily":
		return 1
	case "weekly":
		return 7
	case "monthly":

		return 14
	default:
		return 0
	}
}

func (self *ForecastAPIService) GetByCityIdAndPeriod(ctx context.Context, cityId string, period string) (api.ImplResponse, error) {
	if period == "" {
		return self.getForecastsByCityId(ctx, cityId)
	}

	cityExist, err := self.DBCity.FindByID(cityId)
	if err != nil {
		return api.Response(http.StatusInternalServerError, nil), err
	}
	if cityExist == nil {
		return api.Response(http.StatusNotImplemented, nil), errors.New("City not found")
	}

	if period == "current" {

		current, err := self.handler.HandleCurrantData(cityExist.Name, self.config.CredFile)
		if err != nil {
			return api.Response(http.StatusBadRequest, nil), err
		}
		return api.Response(http.StatusOK, current), nil

	}

	days := self.getDays(period)

	if days == 0 {
		return api.Response(http.StatusNotImplemented, nil), errors.New("Invalid period. Please specify 'current', 'daily', 'weekly', or 'monthly'.")
	}

	forecasts, err := self.DB.FindByCityIdAndPeriodDays(cityId, days)

	if err != nil {
		return api.Response(http.StatusInternalServerError, nil), err
	}

	return api.Response(http.StatusOK, forecasts), nil
}
