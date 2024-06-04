package services

import (
	"io/ioutil"
	"os"
	"testing"
	"time"

	api "github.com/NikolNikolaeva/project_weather/generated/api/weatherapi"

	"github.com/NikolNikolaeva/project_weather/generated/dao/model"
	"github.com/NikolNikolaeva/project_weather/generated/go-mocks/repositories"
	"github.com/NikolNikolaeva/project_weather/generated/go-mocks/services"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_WeatherHandler_Handle_CityRegister(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockCityRepo := mock_repositories.NewMockCityRepo(controller)
	mockCityRepo.EXPECT().RegisterCity(&model.City{
		Name:      "Sofia",
		Country:   "Bulgaria",
		Latitude:  "42.098000",
		Longitude: "43.787000",
	}).Return(&model.City{
		ID: "some_id",
	}, nil)

	mockForeCast := mock_repositories.NewMockForecastRepo(controller)

	mockGetter := mock_services.NewMockWeatherDataGetter(controller)
	mockGetter.EXPECT().GetCurrentData(gomock.Any(), gomock.Any()).Return(&api.InlineResponse2001{
		Current: &api.Current{
			LastUpdated: "2007-01-02 15:04",
		},
		Location: &api.Location{
			Name:    "Sofia",
			Country: "Bulgaria",
			Lat:     42.098,
			Lon:     43.787,
		},
	}, nil)

	weatherHandler := NewWeatherHandler(mockCityRepo, mockForeCast, mockGetter)

	credContent := `{"apiKey": "test_api_key"}`
	tmpFile, err := ioutil.TempFile("", "temp.json")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.Write([]byte(credContent))
	assert.NoError(t, err)
	err = tmpFile.Close()
	assert.NoError(t, err)

	_, err = weatherHandler.HandleCurrantData("www.weather.com", tmpFile.Name())

	assert.NoError(t, err)
}

func Test_WeatherHandler_Handle(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockCityRepo := mock_repositories.NewMockCityRepo(controller)
	mockForecastRepo := mock_repositories.NewMockForecastRepo(controller)
	mockGetter := mock_services.NewMockWeatherDataGetter(controller)

	weatherHandler := NewWeatherHandler(mockCityRepo, mockForecastRepo, mockGetter)

	credContent := `{"apiKey": "test_api_key"}`
	tmpFile, err := ioutil.TempFile("", "temp.json")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.Write([]byte(credContent))
	assert.NoError(t, err)
	err = tmpFile.Close()
	assert.NoError(t, err)

	testCases := []struct {
		description     string
		url             string
		period          string
		mockCity        *model.City
		mockWeatherData *api.InlineResponse2001
		expectedError   error
	}{
		{
			description: "Handle current weather data",
			url:         "www.weather.com",
			period:      "current",
			mockCity: &model.City{
				ID: "some_id",
			},
			mockWeatherData: &api.InlineResponse2001{
				Current: &api.Current{
					LastUpdated: time.Now().Format(templateDateAndTime),
				},
			},
			expectedError: nil,
		},
		{
			description: "Handle daily weather data",
			url:         "www.weather.com",
			period:      "daily",
			mockCity: &model.City{
				ID: "some_id",
			},
			mockWeatherData: &api.InlineResponse2001{
				Forecast: &api.Forecast{
					Forecastday: []api.ForecastForecastday{},
				},
			},
			expectedError: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			mockCityRepo.EXPECT().RegisterCity(gomock.Any()).Return(testCase.mockCity, nil)
			mockGetter.EXPECT().GetCurrentData(testCase.url, gomock.Any()).Return(testCase.mockWeatherData, nil)

			var err error
			if testCase.period == "current" {
				_, err = weatherHandler.HandleCurrantData(testCase.url, tmpFile.Name())
			} else {
				_, err = weatherHandler.HandleCurrantData(testCase.url, tmpFile.Name())
			}
			assert.Equal(t, testCase.expectedError, err)
		})
	}
}

func Test_WeatherHandler_getApiKey(t *testing.T) {

	// Prepare a temporary credentials file for testing
	credContent := `{"apiKey": "test_api_key"}`
	tmpFile, err := ioutil.TempFile("", "temp.json")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.Write([]byte(credContent))
	assert.NoError(t, err)
	err = tmpFile.Close()
	assert.NoError(t, err)

	handler := &weatherHandler{}

	apiKey, err := handler.getApiKey(tmpFile.Name())
	assert.NoError(t, err)
	assert.Equal(t, "test_api_key", apiKey)
}
