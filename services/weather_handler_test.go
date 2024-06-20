package services

//
//import (
//	"io/ioutil"
//	"os"
//	"testing"
//	"time"
//
//	api "github.com/NikolNikolaeva/project_weather/generated/api/weatherapi"
//
//	"github.com/stretchr/testify/assert"
//	"go.uber.org/mock/gomock"
//
//	"github.com/NikolNikolaeva/project_weather/generated/dao/model"
//	mock_repositories "github.com/NikolNikolaeva/project_weather/generated/go-mocks/repositories"
//	mock_services "github.com/NikolNikolaeva/project_weather/generated/go-mocks/services"
//)
//
//func Test_WeatherHandler_Handle_CityRegister(t *testing.T) {
//	controller := gomock.NewController(t)
//	defer controller.Finish()
//
//	city := &model.City{
//		Name:      "Sofia",
//		Country:   "Bulgaria",
//		Latitude:  "42.098000",
//		Longitude: "43.786999",
//	}
//	mockCityRepo := mock_repositories.NewMockCityRepo(controller)
//	mockCityRepo.EXPECT().Register(city).Return(&model.City{
//		ID: "some_id",
//	}, nil)
//
//	mockForeCast := mock_repositories.NewMockForecastRepo(controller)
//
//	mockGetter := mock_services.NewMockWeatherDataGetter(controller)
//	mockGetter.EXPECT().GetCurrentData("Sofia", "test_api_key").Return(
//		&api.Current{
//			LastUpdated: time.Now().Format(time.DateOnly),
//		},
//		&api.Location{
//			Name:    "Sofia",
//			Country: "Bulgaria",
//			Lat:     42.098,
//			Lon:     43.787,
//		}, nil)
//
//	weatherHandler := NewWeatherAPIClient(mockCityRepo, mockForeCast, mockGetter)
//
//	credContent := `{"apiKey": "test_api_key"}`
//	tmpFile, err := ioutil.TempFile("", "temp.json")
//	assert.NoError(t, err)
//	defer os.Remove(tmpFile.Name())
//
//	_, err = tmpFile.Write([]byte(credContent))
//	assert.NoError(t, err)
//	err = tmpFile.Close()
//	assert.NoError(t, err)
//
//	curr, err := weatherHandler.HandleCurrantData("Sofia", tmpFile.Name())
//	assert.Equal(t, curr.LastUpdated, time.Now().Format(time.DateOnly))
//
//	assert.Nil(t, err)
//}
//
//func Test_WeatherHandler_Handle(t *testing.T) {
//	controller := gomock.NewController(t)
//	defer controller.Finish()
//
//	mockCityRepo := mock_repositories.NewMockCityRepo(controller)
//	mockForecastRepo := mock_repositories.NewMockForecastRepo(controller)
//	mockGetter := mock_services.NewMockWeatherDataGetter(controller)
//
//	weatherHandler := NewWeatherAPIClient(mockCityRepo, mockForecastRepo, mockGetter)
//
//	credContent := `{"apiKey": "test_api_key"}`
//	tmpFile, err := ioutil.TempFile("", "temp.json")
//	assert.NoError(t, err)
//	defer os.Remove(tmpFile.Name())
//
//	_, err = tmpFile.Write([]byte(credContent))
//	assert.NoError(t, err)
//	err = tmpFile.Close()
//	assert.NoError(t, err)
//
//	id := "some_id"
//	testCases := []struct {
//		description      string
//		url              string
//		period           string
//		mockCity         *model.City
//		mockCurrData     *api.Current
//		mockForecastData *api.Forecast
//		city             *api.Location
//		expectedError    error
//	}{
//		{
//			description: "Handle current weather data",
//			url:         "www.weather.com",
//			period:      "current",
//			mockCity: &model.City{
//				ID: id,
//			},
//			mockCurrData: &api.Current{
//				LastUpdated: time.Now().Format(templateDateAndTime),
//			},
//			city: &api.Location{
//				Name:    "Sofia",
//				Country: "Bulgaria",
//				Lat:     42.098,
//				Lon:     43.787,
//			},
//			expectedError: nil,
//		},
//		{
//			description: "Handle daily weather data",
//			url:         "www.weather.com",
//			period:      "daily",
//			mockCity: &model.City{
//				ID: id,
//			},
//			mockForecastData: &api.Forecast{
//				Forecastday: []api.ForecastForecastday{},
//			},
//			city: &api.Location{
//				Name:    "Sofia",
//				Country: "Bulgaria",
//				Lat:     42.098,
//				Lon:     43.787,
//			},
//			expectedError: nil,
//		},
//	}
//
//	for _, testCase := range testCases {
//		t.Run(testCase.description, func(t *testing.T) {
//			mockCityRepo.EXPECT().Register(gomock.Any()).Return(testCase.mockCity, nil)
//
//			if testCase.period == "current" {
//				mockGetter.EXPECT().GetCurrentData("Sofia", "test_api_key").Return(testCase.mockCurrData, testCase.city, nil)
//				data, err := weatherHandler.HandleCurrantData("Sofia", tmpFile.Name())
//
//				assert.Equal(t, testCase.mockCurrData.LastUpdated, data.LastUpdated)
//				assert.Equal(t, testCase.expectedError, err)
//			} else {
//				mockGetter.EXPECT().GetForecastData("Sofia", int32(1), "test_api_key").Return(testCase.mockForecastData, testCase.city, nil)
//				data, err := weatherHandler.HandleForecast("Sofia", int32(1), tmpFile.Name())
//				assert.Equal(t, testCase.mockForecastData, data)
//				assert.Equal(t, testCase.expectedError, err)
//			}
//
//		})
//	}
//}
