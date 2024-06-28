package services

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/NikolNikolaeva/project_weather/config"

	api "github.com/NikolNikolaeva/project_weather/generated/api/project-weather/rest"
	"github.com/NikolNikolaeva/project_weather/repositories"
	"github.com/NikolNikolaeva/project_weather/resources"
)

// CityAPIService is a service that implements the logic for the CityAPIService
// This service should implement the business logic for every endpoint for the CityAPI API.
// Include any external packages or services that will be required by this service.
type CityAPIService struct {
	DB      repositories.CityRepo
	Convert resources.ConverterI
	handler WeatherAPIClient
	config  *config.ApplicationConfiguration
}

// NewAPIService creates a default api service
func NewAPIService(db repositories.CityRepo, convert resources.ConverterI, handler WeatherAPIClient, config *config.ApplicationConfiguration) api.CityAPIServicer {
	return &CityAPIService{
		DB:      db,
		Convert: convert,
		handler: handler,
		config:  config,
	}
}

// DeleteById -
func (s *CityAPIService) DeleteById(ctx context.Context, id string) (api.ImplResponse, error) {

	if id == "" {
		return api.Response(http.StatusBadRequest, "id is required"), nil
	}
	city, err := s.DB.FindByID(id)

	if err != nil {
		return api.Response(http.StatusNotFound, "City not found"), err
	}
	_, err = s.DB.DeleteByID(id)
	if err != nil {
		return api.Response(http.StatusInternalServerError, "City not found"), err
	}

	cityApi := s.Convert.ConvertModelCityToApiCity(city)

	return api.Response(http.StatusOK, cityApi), nil
}

// GetAll -
func (s *CityAPIService) GetAll(ctx context.Context) (api.ImplResponse, error) {
	cities, err := s.DB.GetAll()

	if err != nil {
		return api.Response(http.StatusInternalServerError, nil), err
	}

	var citiesApi []*api.City
	for _, city := range cities {
		cityApi := s.Convert.ConvertModelCityToApiCity(city)
		citiesApi = append(citiesApi, cityApi)
	}

	return api.Response(http.StatusOK, citiesApi), nil
}

// GetById -
func (s *CityAPIService) GetById(ctx context.Context, id string) (api.ImplResponse, error) {
	if id == "" {
		return api.Response(http.StatusUnprocessableEntity, "id is required"), nil
	}

	city, err := s.DB.FindByID(id)
	if err != nil {
		log.Print(err.Error())
		return api.Response(http.StatusInternalServerError, err.Error()), nil
	}

	cityApi := s.Convert.ConvertModelCityToApiCity(city)

	return api.Response(http.StatusOK, cityApi), nil
}

// Register -
func (s *CityAPIService) Register(ctx context.Context, city api.City) (api.ImplResponse, error) {
	cityModel := s.Convert.ConvertApiCityToModelCity(&city)

	location := s.handler.HandleCityData(cityModel.Name, s.config.CredFile)

	cityModel.Latitude = fmt.Sprintf("%f", location.Lat)
	cityModel.Longitude = fmt.Sprintf("%f", location.Lon)

	cityNew, err := s.DB.Register(cityModel)
	if err != nil {
		return api.Response(http.StatusInternalServerError, nil), err
	}

	go func() {
		_, _ = s.handler.HandleForecast(cityModel.Name, 30, s.config.CredFile)
	}()

	cityApi := s.Convert.ConvertModelCityToApiCity(cityNew)

	return api.Response(http.StatusCreated, cityApi), nil
}

// UpdateByID -
func (s *CityAPIService) UpdateByID(ctx context.Context, id string, city api.City) (api.ImplResponse, error) {
	cityModel := s.Convert.ConvertApiCityToModelCity(&city)

	updatedCity, err := s.DB.UpdateByID(id, cityModel)
	if err != nil {
		return api.Response(http.StatusInternalServerError, nil), err
	}

	cityApi := s.Convert.ConvertModelCityToApiCity(updatedCity)

	return api.Response(http.StatusOK, cityApi), nil
}
