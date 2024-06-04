package services

import (
	swagger "github.com/weatherapicom/go"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/NikolNikolaeva/project_weather/generated/dao/model"
	mocks "github.com/NikolNikolaeva/project_weather/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_WeatherHandler_Handle_CityRegister(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockCityRepo := mocks.NewMockCityRepo(controller)
	mockCityRepo.EXPECT().RegisterCity(&model.City{
		Name:      "Sofia",
		Country:   "Bulgaria",
		Latitude:  "42.098000",
		Longitude: "43.787000",
	}).Return(&model.City{
		ID: "some_id",
	}, nil)

	mockForeCast := mocks.NewMockForecastRepo(controller)

	mockGetter := mocks.NewMockWeatherDataGetter(controller)
	mockGetter.EXPECT().GetData(gomock.Any()).Return(&swagger.InlineResponse2001{
		Current: &swagger.Current{
			LastUpdated: "2007-01-02 15:04",
		},
		Location: &swagger.Location{
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

	mockCityRepo := mocks.NewMockCityRepo(controller)
	mockForecastRepo := mocks.NewMockForecastRepo(controller)
	mockGetter := mocks.NewMockWeatherDataGetter(controller)

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
		mockWeatherData *swagger.InlineResponse2001
		expectedError   error
	}{
		{
			description: "Handle current weather data",
			url:         "www.weather.com",
			period:      "current",
			mockCity: &model.City{
				ID: "some_id",
			},
			mockWeatherData: &swagger.InlineResponse2001{
				Current: &swagger.Current{
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
			mockWeatherData: &swagger.InlineResponse2001{
				Forecast: &swagger.Forecast{
					Forecastday: []swagger.ForecastForecastday{},
				},
			},
			expectedError: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			mockCityRepo.EXPECT().RegisterCity(gomock.Any()).Return(testCase.mockCity, nil)
			mockGetter.EXPECT().GetData(testCase.url).Return(testCase.mockWeatherData, nil)

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
