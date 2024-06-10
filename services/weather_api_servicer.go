package services

import (
	"context"
	"errors"
	"net/http"

	"github.com/NikolNikolaeva/project_weather/config"
	api "github.com/NikolNikolaeva/project_weather/generated/api/project-weather/rest"
)

var periodToDays = map[string]int{
	"daily":   1,
	"weekly":  7,
	"monthly": 30,
	"current": 0,
}

type WeatherAPIService struct {
	weatherHandler WeatherHandler
	config         *config.ApplicationConfiguration
}

// NewWeatherAPIService creates a default api service
func NewWeatherAPIService(xHandler WeatherHandler, config *config.ApplicationConfiguration) api.WeatherAPIServicer {
	return &WeatherAPIService{
		weatherHandler: xHandler,
		config:         config,
	}
}

// GetWeatherByCity -
func (s *WeatherAPIService) GetWeatherByCity(ctx context.Context, city string, period string) (api.ImplResponse, error) {

	days, exists := periodToDays[period]
	if !exists {
		return api.Response(http.StatusNotFound, nil), errors.New("Invalid period. Please specify 'current', 'daily', 'weekly', or 'monthly'.")
	}

	if period == "current" {
		current, err := s.weatherHandler.HandleCurrantData(city, s.config.CredFile)
		if err != nil {
			return api.Response(http.StatusBadRequest, nil), err
		}
		return api.Response(http.StatusOK, current), nil

	} else {
		forecast, err := s.weatherHandler.HandleForecast(city, int32(days), s.config.CredFile)
		if err != nil {
			return api.Response(http.StatusBadRequest, nil), err
		}
		return api.Response(http.StatusOK, forecast), nil

	}
}
