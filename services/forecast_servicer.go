package services

import (
	"context"
	"fmt"
	"net/http"

	api "github.com/NikolNikolaeva/project_weather/generated/api/project-weather/rest"
	"github.com/NikolNikolaeva/project_weather/repositories"
	"github.com/NikolNikolaeva/project_weather/resources"
)

// ForecastAPIService is a service that implements the logic for the ForecastAPIServicer
// This service should implement the business logic for every endpoint for the ForecastAPI API.
// Include any external packages or services that will be required by this service.
type ForecastAPIService struct {
	DB      repositories.ForecastRepo
	Convert resources.ConverterI
}

// NewForecastAPIService creates a default api service
func NewForecastAPIService(DB repositories.ForecastRepo, convert resources.ConverterI) api.ForecastAPIServicer {
	return &ForecastAPIService{
		DB:      DB,
		Convert: convert,
	}
}

// CreateForecast -
func (s *ForecastAPIService) CreateForecast(ctx context.Context, forecast api.Forecast) (api.ImplResponse, error) {

	forecastModel := s.Convert.ConvertApiForecastToModelForecast(&forecast) //todo

	if err := s.DB.Create(forecastModel); err != nil {
		return api.Response(http.StatusInternalServerError, nil), err
	}

	return api.Response(http.StatusCreated, forecast), nil
}

// DeleteForecastById -
func (s *ForecastAPIService) DeleteForecastById(ctx context.Context, id string) (api.ImplResponse, error) {
	if err := s.DB.Delete(id); err != nil {
		return api.Response(http.StatusInternalServerError, nil), err
	}

	return api.Response(http.StatusNoContent, "Forecast deleted!"), nil
}

// GetAllForecasts -
func (s *ForecastAPIService) GetAllForecasts(ctx context.Context) (api.ImplResponse, error) {
	forecasts, err := s.DB.FindAll()
	if err != nil {
		return api.Response(http.StatusInternalServerError, nil), err
	}
	var forecastsApi []*api.Forecast
	for _, forecast := range forecasts {
		fmt.Printf("%#v\n", forecast)
		forecastApi := s.Convert.ConvertModelForecastToApiForecast(forecast) //Todo
		forecastsApi = append(forecastsApi, forecastApi)
	}

	return api.Response(http.StatusOK, forecastsApi), nil
}

// GetForecastById -
func (s *ForecastAPIService) GetForecastById(ctx context.Context, id string) (api.ImplResponse, error) {
	forecast, err := s.DB.FindByID(id)
	if err != nil {
		return api.Response(http.StatusInternalServerError, nil), err
	}

	forecastApi := s.Convert.ConvertModelForecastToApiForecast(forecast) //Todo

	return api.Response(http.StatusOK, forecastApi), nil
}

// UpdateForecast -
func (s *ForecastAPIService) UpdateForecast(ctx context.Context, id string, forecast api.Forecast) (api.ImplResponse, error) {
	forecastModel := s.Convert.ConvertApiForecastToModelForecast(&forecast) //todo

	if err := s.DB.Update(id, forecastModel); err != nil {
		return api.Response(http.StatusInternalServerError, nil), err
	}

	return api.Response(http.StatusOK, forecast), nil
}
