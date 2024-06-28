package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	mock_repositories "github.com/NikolNikolaeva/project_weather/generated/go-mocks/repositories"
	mock_services "github.com/NikolNikolaeva/project_weather/generated/go-mocks/services"
)

func Test_HandleCurrantData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCityRepo := mock_repositories.NewMockCityRepo(ctrl)
	mockForecastRepo := mock_repositories.NewMockForecastRepo(ctrl)
	mockWeatherDataGetter := mock_services.NewMockWeatherDataGetter(ctrl)
	client := NewWeatherAPIClient(mockCityRepo, mockForecastRepo, mockWeatherDataGetter)

	q := "test_query"

	t.Run("Error getting API key", func(t *testing.T) {
		output, err := client.HandleCurrantData(q, "invalid_cred")
		assert.Error(t, err)
		assert.Nil(t, output)
	})
}

func TestWeatherAPIClient_HandleForecast(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCityRepo := mock_repositories.NewMockCityRepo(ctrl)
	mockForecastRepo := mock_repositories.NewMockForecastRepo(ctrl)
	mockWeatherDataGetter := mock_services.NewMockWeatherDataGetter(ctrl)
	client := NewWeatherAPIClient(mockCityRepo, mockForecastRepo, mockWeatherDataGetter)

	q := "test_query"
	days := int32(5)

	t.Run("Error getting API key", func(t *testing.T) {
		output, err := client.HandleForecast(q, days, "invalid_cred")
		assert.Error(t, err, "error fetching data")
		assert.Nil(t, output)
	})
}

func TestWeatherAPIClient_HandleCityData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockWeatherDataGetter := mock_services.NewMockWeatherDataGetter(ctrl)
	client := &weatherAPIClient{weatherDataGetter: mockWeatherDataGetter}

	name := "Test City"

	t.Run("Error getting API key", func(t *testing.T) {
		output := client.HandleCityData(name, "invalid_cred")
		assert.Nil(t, output)
	})

}
