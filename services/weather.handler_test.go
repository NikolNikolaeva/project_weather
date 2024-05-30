package services

import (
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/NikolNikolaeva/project_weather/generated/dao/model"
	repositories "github.com/NikolNikolaeva/project_weather/mocks"
	"github.com/NikolNikolaeva/project_weather/resources/swagger"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_WeatherHandler_Handle_CityRegister(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockCityRepo := repositories.NewMockCityRepo(controller)
	mockCityRepo.EXPECT().RegisterCity(&model.City{
		Name:      "Sofia",
		Country:   "Bulgaria",
		Latitude:  "42.098000",
		Longitude: "43.787000",
	}).Return(&model.City{
		ID: "some_id",
	}, nil)

	mockForeCast := repositories.NewMockForecastRepo(controller)

	mockGetter := repositories.NewMockWeatherDataGetter(controller)
	mockGetter.EXPECT().GetData(gomock.Any()).Return(&swagger.WeatherDTO{
		Current: swagger.CurrentDTO{
			LastUpdated: "2007-01-02 15:04",
		},
		Location: swagger.LocationDTO{
			Name:    "Sofia",
			Country: "Bulgaria",
			Lat:     42.098,
			Lon:     43.787,
		},
	}, nil)

	weatherHandler := NewWeatherHandler(mockCityRepo, mockForeCast, mockGetter)

	_, err := weatherHandler.Handle("www.weather.com", "current")

	assert.NoError(t, err)
}

func Test_WeatherHandler_Handle(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockCityRepo := repositories.NewMockCityRepo(controller)
	mockForecastRepo := repositories.NewMockForecastRepo(controller)
	mockGetter := repositories.NewMockWeatherDataGetter(controller)

	weatherHandler := NewWeatherHandler(mockCityRepo, mockForecastRepo, mockGetter)

	testCases := []struct {
		description     string
		url             string
		period          string
		mockCity        *model.City
		mockWeatherData *swagger.WeatherDTO
		expectedError   error
	}{
		{
			description: "Handle current weather data",
			url:         "www.weather.com",
			period:      "current",
			mockCity: &model.City{
				ID: "some_id",
			},
			mockWeatherData: &swagger.WeatherDTO{
				Current: swagger.CurrentDTO{
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
			mockWeatherData: &swagger.WeatherDTO{
				Forecast: swagger.ForecastsDTO{
					ForecastDays: []swagger.ForecastDayDTO{},
				},
			},
			expectedError: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			mockCityRepo.EXPECT().RegisterCity(gomock.Any()).Return(testCase.mockCity, nil)
			mockGetter.EXPECT().GetData(testCase.url).Return(testCase.mockWeatherData, nil)

			_, err := weatherHandler.Handle(testCase.url, testCase.period)
			assert.Equal(t, testCase.expectedError, err)
		})
	}
}

func Test_WeatherHandler_GetUrlForWeatherApi(t *testing.T) {
	handler := &weatherHandler{}

	testCases := []struct {
		description    string
		period         string
		apiKey         string
		city           string
		days           int
		ForecastUrl    string
		CurrentTimeUrl string
		expectedUrl    string
	}{
		{
			description:    "Get current weather URL",
			period:         "current",
			apiKey:         "test_api_key",
			city:           "Sofia",
			CurrentTimeUrl: "https://api.weather.com/v1/current?apikey=%s&city=%s",
			expectedUrl:    "https://api.weather.com/v1/current?apikey=test_api_key&city=Sofia",
		},
		{
			description: "Get forecast weather URL",
			period:      "daily",
			apiKey:      "test_api_key",
			city:        "Sofia",
			days:        1,
			ForecastUrl: "https://api.weather.com/v1/forecast?apikey=%s&city=%s&days=%d",
			expectedUrl: "https://api.weather.com/v1/forecast?apikey=test_api_key&city=Sofia&days=1",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			actualUrl := handler.GetUrlForWeatherApi(testCase.period, testCase.apiKey, testCase.city, testCase.days, testCase.ForecastUrl, testCase.CurrentTimeUrl)
			assert.Equal(t, testCase.expectedUrl, actualUrl)
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

func Test_WeatherHandler_formatWeatherForecastData(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockForecastRepo := repositories.NewMockForecastRepo(controller)
	weatherHandler := &weatherHandler{
		forecastRepo: mockForecastRepo,
	}

	cityID := "some_id"
	weatherData := &swagger.WeatherDTO{
		Forecast: swagger.ForecastsDTO{
			ForecastDays: []swagger.ForecastDayDTO{
				{
					Date: time.Now().Format(templateDate),
					Day: swagger.DayDTO{
						AvgTempC: 15.5,
						Condition: swagger.ConditionDTO{
							Text: "Partly Cloudy",
						},
					},
				},
			},
		},
	}

	mockForecastRepo.EXPECT().Create(gomock.Any()).Return(nil)

	forecasts, err := weatherHandler.formatWeatherForecastData(weatherData, cityID)
	assert.NoError(t, err)
	assert.Len(t, forecasts, 1)
	assert.Equal(t, cityID, forecasts[0].CityID)
	assert.Equal(t, "Partly Cloudy", forecasts[0].Condition)
}

func Test_WeatherHandler_formatWeatherCurrentData(t *testing.T) {
	weatherHandler := &weatherHandler{}

	lastUpdated := time.Now().Format(templateDateAndTime)
	cityID := "some_id"
	weatherData := &swagger.WeatherDTO{
		Current: swagger.CurrentDTO{
			LastUpdated: lastUpdated,
			TempC:       20.5,
			Condition:   swagger.ConditionDTO{Text: "Sunny"},
		},
	}

	currentData, err := weatherHandler.formatWeatherCurrentData(weatherData, cityID)
	assert.NoError(t, err)
	assert.Equal(t, cityID, currentData.CityID)
	assert.Equal(t, lastUpdated, currentData.ForecastDate.Format(templateDateAndTime))
	assert.Equal(t, "Sunny", currentData.Condition)
	assert.Equal(t, "20.5", currentData.Temperature)
}
