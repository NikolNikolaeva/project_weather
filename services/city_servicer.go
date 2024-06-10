package services

import (
	"context"
	"log"
	"net/http"

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
}

// NewCityAPIService creates a default api service
func NewCityAPIService(db repositories.CityRepo, convert resources.ConverterI) api.CityAPIServicer {
	return &CityAPIService{
		DB:      db,
		Convert: convert,
	}
}

// DeleteCityById -
func (s *CityAPIService) DeleteCityById(ctx context.Context, id string) (api.ImplResponse, error) {
	city, err := s.DB.FindCityByID(id)

	if err != nil {
		return api.Response(http.StatusNotFound, "City not found"), err
	}

	if id == "" {
		return api.Response(http.StatusBadRequest, "id is required"), nil
	}
	err = s.DB.DeleteCityByID(id)
	if err != nil {
		return api.Response(http.StatusInternalServerError, "City not found"), err
	}

	cityApi := s.Convert.ConvertModelCityToApiCity(city) //Todo

	return api.Response(http.StatusOK, cityApi), nil
}

// GetAllCities -
func (s *CityAPIService) GetAllCities(ctx context.Context) (api.ImplResponse, error) {
	cities, err := s.DB.GetAllCity()

	if err != nil {
		return api.Response(http.StatusInternalServerError, nil), err
	}

	var citiesApi []*api.City
	for _, city := range cities {
		cityApi := s.Convert.ConvertModelCityToApiCity(city) //Todo
		citiesApi = append(citiesApi, cityApi)
	}

	return api.Response(http.StatusOK, citiesApi), nil
}

// GetCityById -
func (s *CityAPIService) GetCityById(ctx context.Context, id string) (api.ImplResponse, error) {
	if id == "" {
		return api.Response(http.StatusUnprocessableEntity, "id is required"), nil
	}

	city, err := s.DB.FindCityByID(id)
	if err != nil {
		log.Print(err.Error())
		return api.Response(http.StatusInternalServerError, err.Error()), nil
	}

	cityApi := s.Convert.ConvertModelCityToApiCity(city) //Todo

	return api.Response(http.StatusOK, cityApi), nil
}

// RegisterCity -
func (s *CityAPIService) RegisterCity(ctx context.Context, city api.City) (api.ImplResponse, error) {
	cityModel := s.Convert.ConvertApiCityToModelCity(&city) //Todo

	cityNew, err := s.DB.RegisterCity(cityModel)
	if err != nil {
		return api.Response(http.StatusInternalServerError, nil), err
	}

	cityApi := s.Convert.ConvertModelCityToApiCity(cityNew) //Todo

	return api.Response(http.StatusCreated, cityApi), nil
}

// UpdateCityById -
func (s *CityAPIService) UpdateCityById(ctx context.Context, id string, city api.City) (api.ImplResponse, error) {
	cityModel := s.Convert.ConvertApiCityToModelCity(&city) //Todo

	updatedCity, err := s.DB.UpdateCityByID(id, cityModel)
	if err != nil {
		return api.Response(http.StatusInternalServerError, nil), err
	}

	cityApi := s.Convert.ConvertModelCityToApiCity(updatedCity) //Todo

	return api.Response(http.StatusOK, cityApi), nil
}
