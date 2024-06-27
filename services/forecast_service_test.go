package services

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"go.uber.org/mock/gomock"

	"github.com/NikolNikolaeva/project_weather/generated/api/weatherapi"
	"github.com/NikolNikolaeva/project_weather/generated/dao/model"
	mock_repositories "github.com/NikolNikolaeva/project_weather/generated/go-mocks/repositories"
	mock_resources "github.com/NikolNikolaeva/project_weather/generated/go-mocks/resources"
	mock_services "github.com/NikolNikolaeva/project_weather/generated/go-mocks/services"

	"github.com/stretchr/testify/assert"

	"github.com/NikolNikolaeva/project_weather/config"
)

func TestForecastAPIService_GetByCityIdAndPeriod(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mock_repositories.NewMockForecastRepo(ctrl)
	mockDBCity := mock_repositories.NewMockCityRepo(ctrl)
	mockConvert := mock_resources.NewMockConverterI(ctrl)
	mockHandler := mock_services.NewMockWeatherAPIClient(ctrl)
	mockConfig := &config.ApplicationConfiguration{}

	service := NewForecastAPIService(mockDB, mockConvert, mockHandler, mockConfig, mockDBCity)

	ctx := context.Background()

	t.Run("City Id required", func(t *testing.T) {
		resp, err := service.GetByCityIdAndPeriod(ctx, "", "")
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotImplemented, resp.Code)
		assert.Equal(t, "City id is required", resp.Body)
	})

	t.Run("City not found", func(t *testing.T) {
		mockDBCity.EXPECT().FindByID("123").Return(nil, nil)

		resp, err := service.GetByCityIdAndPeriod(ctx, "123", "daily")
		assert.Error(t, err)
		assert.Equal(t, http.StatusNotImplemented, resp.Code)
		assert.Equal(t, "City not found", err.Error())
	})

	t.Run("Invalid period", func(t *testing.T) {
		mockDBCity.EXPECT().FindByID("123").Return(&model.City{Name: "Test City"}, nil)

		resp, err := service.GetByCityIdAndPeriod(ctx, "123", "invalid")
		assert.Equal(t, http.StatusNotImplemented, resp.Code)
		assert.Equal(t, nil, resp.Body)
		assert.Error(t, err, "Invalid period. Please specify 'current', 'daily', 'weekly', or 'monthly'.")
	})

	t.Run("Valid period", func(t *testing.T) {
		mockDBCity.EXPECT().FindByID("123").Return(&model.City{Name: "Test City"}, nil)
		mockDB.EXPECT().FindByCityIdAndPeriodDays("123", 7).Return([]*model.Forecast{}, nil)

		resp, err := service.GetByCityIdAndPeriod(ctx, "123", "weekly")
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, resp.Code)
		assert.Equal(t, []*model.Forecast{}, resp.Body)
	})

	t.Run("Error from FindByID", func(t *testing.T) {
		mockDBCity.EXPECT().FindByID("123").Return(nil, errors.New("database error"))

		resp, err := service.GetByCityIdAndPeriod(ctx, "123", "daily")
		assert.Error(t, err)
		assert.Equal(t, http.StatusInternalServerError, resp.Code)
		assert.Error(t, err, "database error")
	})

	t.Run("Error from HandleCurrantData", func(t *testing.T) {
		mockDBCity.EXPECT().FindByID("123").Return(&model.City{Name: "Test City"}, nil)
		mockHandler.EXPECT().HandleCurrantData("Test City", gomock.Any()).Return(&weatherapi.Current{}, errors.New("handler error"))

		resp, err := service.GetByCityIdAndPeriod(ctx, "123", "current")
		assert.Error(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.Code)
		assert.Error(t, err, "handler error")
	})

	t.Run("Error from FindByCityIdAndPeriodDays", func(t *testing.T) {
		mockDBCity.EXPECT().FindByID("123").Return(&model.City{Name: "Test City"}, nil)
		mockDB.EXPECT().FindByCityIdAndPeriodDays("123", 7).Return(nil, errors.New("db error"))

		resp, err := service.GetByCityIdAndPeriod(ctx, "123", "weekly")
		assert.Error(t, err)
		assert.Equal(t, http.StatusInternalServerError, resp.Code)
		assert.Error(t, err, "db error")
	})
}
