package services

import (
	"testing"

	"github.com/NikolNikolaeva/project_weather/generated/dao/model"
	repositories "github.com/NikolNikolaeva/project_weather/mocks"
	"github.com/NikolNikolaeva/project_weather/resources/swagger"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_WeatherHandler_Handle_Current(t *testing.T) {
	controller := gomock.NewController(t)
	mockCityRepo := repositories.NewMockCityDB(controller)
	mockCityRepo.EXPECT().RegisterCity(gomock.Any()).Return(&model.City{
		ID:        "some_id",
	}, nil)

	mockForeCast := repositories.NewMockForecastDB(controller)

	mockGetter := repositories.NewMockWeatherDataGetter(controller)
	mockGetter.EXPECT().GetData(gomock.Any()).Return(&swagger.WeatherDTO{
		Current:  swagger.CurrentDTO{
			LastUpdated: "2006-01-02 15:04",
		},
	}, nil)

	xHandler := NewWeatherHandler(mockCityRepo, mockForeCast, mockGetter)

	_, err := xHandler.Handle("www.weather.com", "current")

	assert.NoError(t, err)
	assert.True(t, controller.Satisfied())
}

func Test_WeatherHandler_Handle_CityRegister(t *testing.T) {
	controller := gomock.NewController(t)
	mockCityRepo := repositories.NewMockCityDB(controller)
	mockCityRepo.EXPECT().RegisterCity(&model.City{
		Name:      "Sofia",
		Country:   "Bulgaria",
		Latitude:  "42.098000",
		Longitude: "43.787000",
	}).Return(&model.City{
		ID:        "some_id",
	}, nil)

	mockForeCast := repositories.NewMockForecastDB(controller)

	mockGetter := repositories.NewMockWeatherDataGetter(controller)
	mockGetter.EXPECT().GetData(gomock.Any()).Return(&swagger.WeatherDTO{
		Current:  swagger.CurrentDTO{
			LastUpdated: "2006-01-02 15:04",
		},
		Location: swagger.LocationDTO{
			Name:    "Sofia",
			Country: "Bulgaria",
			Lat:     42.098,
			Lon:     43.787,
		},
	}, nil)

	xHandler := NewWeatherHandler(mockCityRepo, mockForeCast, mockGetter)

	_, err := xHandler.Handle("www.weather.com", "current")

	assert.NoError(t, err)
	assert.True(t, controller.Satisfied())
}
