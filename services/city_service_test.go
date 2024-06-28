package services

import (
	"context"
	"errors"
	"net/http"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/gen"

	"github.com/NikolNikolaeva/project_weather/config"
	api "github.com/NikolNikolaeva/project_weather/generated/api/project-weather/rest"
	"github.com/NikolNikolaeva/project_weather/generated/api/weatherapi"
	"github.com/NikolNikolaeva/project_weather/generated/dao/model"
	mock_repositories "github.com/NikolNikolaeva/project_weather/generated/go-mocks/repositories"
	mock_resources "github.com/NikolNikolaeva/project_weather/generated/go-mocks/resources"
	mock_services "github.com/NikolNikolaeva/project_weather/generated/go-mocks/services"
)

func Test_GetById_CityNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repositories.NewMockCityRepo(ctrl)
	mockConverter := mock_resources.NewMockConverterI(ctrl)
	mockHandler := mock_services.NewMockWeatherAPIClient(ctrl)

	service := NewAPIService(mockRepo, mockConverter, mockHandler, &config.ApplicationConfiguration{})

	cityID := "1"

	mockRepo.EXPECT().FindByID(cityID).Return(nil, errors.New("city not found"))

	response, err := service.GetById(context.Background(), cityID)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusInternalServerError, response.Code)
	assert.Equal(t, "city not found", response.Body)
}

func Test_GetById_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repositories.NewMockCityRepo(ctrl)
	mockConverter := mock_resources.NewMockConverterI(ctrl)
	mockHandler := mock_services.NewMockWeatherAPIClient(ctrl)

	service := NewAPIService(mockRepo, mockConverter, mockHandler, &config.ApplicationConfiguration{})

	cityID := "1"
	modelCity := &model.City{
		ID:        cityID,
		Name:      "Test City",
		Latitude:  "0",
		Longitude: "0",
	}
	apiCity := &api.City{
		Id:        cityID,
		Name:      "Test City",
		Latitude:  "0",
		Longitude: "0",
	}

	mockRepo.EXPECT().FindByID(cityID).Return(modelCity, nil)
	mockConverter.EXPECT().ConvertModelCityToApiCity(modelCity).Return(apiCity)

	response, err := service.GetById(context.Background(), cityID)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, apiCity, response.Body)
}

func Test_GetById_RequireId(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repositories.NewMockCityRepo(ctrl)
	mockConverter := mock_resources.NewMockConverterI(ctrl)
	mockHandler := mock_services.NewMockWeatherAPIClient(ctrl)

	service := NewAPIService(mockRepo, mockConverter, mockHandler, &config.ApplicationConfiguration{})

	response, err := service.GetById(context.Background(), "")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
	assert.Equal(t, "id is required", response.Body)
}

func Test_DeleteById_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repositories.NewMockCityRepo(ctrl)
	mockConverter := mock_resources.NewMockConverterI(ctrl)
	mockHandler := mock_services.NewMockWeatherAPIClient(ctrl)

	service := NewAPIService(mockRepo, mockConverter, mockHandler, &config.ApplicationConfiguration{})

	cityID := "1"
	modelCity := &model.City{
		ID:        cityID,
		Name:      "Test City",
		Latitude:  "0",
		Longitude: "0",
	}
	apiCity := &api.City{
		Id:        cityID,
		Name:      "Test City",
		Latitude:  "0",
		Longitude: "0",
	}

	mockRepo.EXPECT().FindByID(cityID).Return(modelCity, nil)
	mockRepo.EXPECT().DeleteByID(cityID).Return(gen.ResultInfo{}, nil)
	mockConverter.EXPECT().ConvertModelCityToApiCity(modelCity).Return(apiCity)

	response, err := service.DeleteById(context.Background(), cityID)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, apiCity, response.Body)
}

func Test_DeleteById_CityNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repositories.NewMockCityRepo(ctrl)
	mockConverter := mock_resources.NewMockConverterI(ctrl)
	mockHandler := mock_services.NewMockWeatherAPIClient(ctrl)

	service := NewAPIService(mockRepo, mockConverter, mockHandler, &config.ApplicationConfiguration{})

	cityID := "1"

	mockRepo.EXPECT().FindByID(cityID).Return(nil, errors.New("city not found"))

	response, err := service.DeleteById(context.Background(), cityID)
	assert.Error(t, err)
	assert.Equal(t, http.StatusNotFound, response.Code)
	assert.Equal(t, "City not found", response.Body)
}

func Test_DeleteById_RequireId(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repositories.NewMockCityRepo(ctrl)
	mockConverter := mock_resources.NewMockConverterI(ctrl)
	mockHandler := mock_services.NewMockWeatherAPIClient(ctrl)

	service := NewAPIService(mockRepo, mockConverter, mockHandler, &config.ApplicationConfiguration{})

	response, err := service.DeleteById(context.Background(), "")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, response.Code)
	assert.Equal(t, "id is required", response.Body)
}

func Test_GetAll_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repositories.NewMockCityRepo(ctrl)
	mockConverter := mock_resources.NewMockConverterI(ctrl)
	mockHandler := mock_services.NewMockWeatherAPIClient(ctrl)

	service := NewAPIService(mockRepo, mockConverter, mockHandler, &config.ApplicationConfiguration{})

	modelCities := []*model.City{
		{
			ID:        "1",
			Name:      "City 1",
			Latitude:  "0",
			Longitude: "0",
		},
		{
			ID:        "2",
			Name:      "City 2",
			Latitude:  "1",
			Longitude: "1",
		},
	}
	apiCities := []*api.City{
		{
			Id:        "1",
			Name:      "City 1",
			Latitude:  "0",
			Longitude: "0",
		},
		{
			Id:        "2",
			Name:      "City 2",
			Latitude:  "1",
			Longitude: "1",
		},
	}

	mockRepo.EXPECT().GetAll().Return(modelCities, nil)

	for i, city := range modelCities {
		mockConverter.EXPECT().ConvertModelCityToApiCity(city).Return(apiCities[i])
	}

	response, err := service.GetAll(context.Background())
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, apiCities, response.Body)
}

func Test_GetAll_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repositories.NewMockCityRepo(ctrl)
	mockConverter := mock_resources.NewMockConverterI(ctrl)
	mockHandler := mock_services.NewMockWeatherAPIClient(ctrl)

	service := NewAPIService(mockRepo, mockConverter, mockHandler, &config.ApplicationConfiguration{})

	mockRepo.EXPECT().GetAll().Return(nil, errors.New("db error"))

	response, err := service.GetAll(context.Background())
	assert.Error(t, err, "db error")
	assert.Equal(t, http.StatusInternalServerError, response.Code)
	assert.Nil(t, response.Body)
}

func Test_Register_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repositories.NewMockCityRepo(ctrl)
	mockConverter := mock_resources.NewMockConverterI(ctrl)
	mockHandler := mock_services.NewMockWeatherAPIClient(ctrl)

	service := NewAPIService(mockRepo, mockConverter, mockHandler, &config.ApplicationConfiguration{})

	var wg sync.WaitGroup
	wg.Add(1)

	apiCity := &api.City{
		Name: "Test City",
	}
	modelCity := &model.City{
		Name:      "Test City",
		Latitude:  "0",
		Longitude: "0",
	}
	registeredCity := modelCity
	registeredCity.ID = "1"

	location := &weatherapi.Location{
		Lat: 0,
		Lon: 0,
	}
	forecast := &weatherapi.Forecast{}

	mockConverter.EXPECT().ConvertApiCityToModelCity(apiCity).Return(modelCity)
	mockHandler.EXPECT().HandleCityData(modelCity.Name, gomock.Any()).Return(location)
	mockRepo.EXPECT().Register(modelCity).Return(registeredCity, nil)
	mockConverter.EXPECT().ConvertModelCityToApiCity(registeredCity).Return(apiCity)
	mockHandler.EXPECT().HandleForecast(modelCity.Name, int32(30), gomock.Any()).Return(forecast, nil).Do(func(string, int32, string) {
		wg.Done()
	}).Times(1)

	response, err := service.Register(context.Background(), *apiCity)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, response.Code)
	assert.Equal(t, apiCity, response.Body)

	wg.Wait()
}

func Test_Register_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repositories.NewMockCityRepo(ctrl)
	mockConverter := mock_resources.NewMockConverterI(ctrl)
	mockHandler := mock_services.NewMockWeatherAPIClient(ctrl)

	service := NewAPIService(mockRepo, mockConverter, mockHandler, &config.ApplicationConfiguration{})

	apiCity := &api.City{
		Name: "Test City",
	}
	modelCity := &model.City{
		Name:      "Test City",
		Latitude:  "0",
		Longitude: "0",
	}
	location := &weatherapi.Location{
		Lat: 0,
		Lon: 0,
	}

	mockConverter.EXPECT().ConvertApiCityToModelCity(apiCity).Return(modelCity).Times(1)
	mockHandler.EXPECT().HandleCityData(modelCity.Name, gomock.Any()).Return(location).Times(1)
	mockRepo.EXPECT().Register(modelCity).Return(&model.City{}, errors.New("db error")).Times(1)

	response, err := service.Register(context.Background(), *apiCity)
	assert.Error(t, err, "db error")
	assert.Equal(t, http.StatusInternalServerError, response.Code)
	assert.Nil(t, response.Body)
}

func Test_UpdateByID_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repositories.NewMockCityRepo(ctrl)
	mockConverter := mock_resources.NewMockConverterI(ctrl)
	mockHandler := mock_services.NewMockWeatherAPIClient(ctrl)

	service := NewAPIService(mockRepo, mockConverter, mockHandler, &config.ApplicationConfiguration{})

	cityID := "1"
	apiCity := &api.City{
		Name: "Updated City",
	}
	modelCity := &model.City{
		Name: "Updated City",
	}
	updatedCity := modelCity
	updatedCity.ID = cityID

	mockConverter.EXPECT().ConvertApiCityToModelCity(apiCity).Return(modelCity)
	mockRepo.EXPECT().UpdateByID(cityID, modelCity).Return(updatedCity, nil)
	mockConverter.EXPECT().ConvertModelCityToApiCity(updatedCity).Return(apiCity)

	response, err := service.UpdateByID(context.Background(), cityID, *apiCity)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, apiCity, response.Body)
}

func Test_UpdateByID_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repositories.NewMockCityRepo(ctrl)
	mockConverter := mock_resources.NewMockConverterI(ctrl)
	mockHandler := mock_services.NewMockWeatherAPIClient(ctrl)

	service := NewAPIService(mockRepo, mockConverter, mockHandler, &config.ApplicationConfiguration{})

	cityID := "1"
	apiCity := &api.City{
		Name: "Updated City",
	}
	modelCity := &model.City{
		Name: "Updated City",
	}

	mockConverter.EXPECT().ConvertApiCityToModelCity(apiCity).Return(modelCity)
	mockRepo.EXPECT().UpdateByID(cityID, modelCity).Return(nil, errors.New("db error"))

	response, err := service.UpdateByID(context.Background(), cityID, *apiCity)
	assert.Error(t, err, "db error")
	assert.Equal(t, http.StatusInternalServerError, response.Code)
	assert.Nil(t, response.Body)
}
