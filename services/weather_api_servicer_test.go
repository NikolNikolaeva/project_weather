package services

import (
	"bytes"
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/NikolNikolaeva/project_weather/config"
	"github.com/NikolNikolaeva/project_weather/generated/dao/model"
	mock_repositories "github.com/NikolNikolaeva/project_weather/generated/go-mocks/repositories"
	mock_services "github.com/NikolNikolaeva/project_weather/generated/go-mocks/services"
)

func TestWeatherService_StartFetching(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDataGetter := mock_services.NewMockWeatherDataGetter(ctrl)
	mockCityRepo := mock_repositories.NewMockCityRepo(ctrl)
	mockForecastRepo := mock_repositories.NewMockForecastRepo(ctrl)
	configuration := &config.ApplicationConfiguration{
		CredFile: "credentials.json",
	}
	mockHandler := mock_services.NewMockWeatherAPIClient(ctrl)

	service := NewWeatherService(mockDataGetter, mockCityRepo, mockForecastRepo, configuration, mockHandler)

	t.Run("Error fetching cities", func(t *testing.T) {
		mockCityRepo.EXPECT().GetAll().Return(nil, fmt.Errorf("failed to fetch cities")).Times(1)

		var buf bytes.Buffer
		log.SetOutput(&buf)
		defer func() {
			log.SetOutput(nil)
		}()

		service.StartFetching()

		logOutput := buf.String()
		assert.Contains(t, logOutput, "Error fetching cities")
	})

	t.Run("Error fetching forecast", func(t *testing.T) {
		mockCityRepo.EXPECT().GetAll().Return([]*model.City{
			{
				Name: "City test",
			},
		}, nil).Times(1)
		mockHandler.EXPECT().HandleForecast("City test", int32(30), configuration.CredFile).
			Return(nil, fmt.Errorf("failed to fetch forecast")).Times(1)

		var buf bytes.Buffer
		log.SetOutput(&buf)
		defer func() {
			log.SetOutput(nil)
		}()

		service.StartFetching()

		logOutput := buf.String()
		assert.Contains(t, logOutput, "Error fetching forecast for city")
	})
}
