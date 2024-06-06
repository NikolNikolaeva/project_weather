package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	api "github.com/NikolNikolaeva/project_weather/generated/api/weatherapi"
	mock_services "github.com/NikolNikolaeva/project_weather/generated/go-mocks/services"
)

func Test_GetCurrentData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_services.NewMockWeatherDataGetter(ctrl)

	q := "London"
	key := "test_key"

	expectedCurrent := &api.Current{TempC: 20.0}
	expectedLocation := &api.Location{Name: "London"}

	mockClient.EXPECT().
		GetCurrentData(q, key).
		Return(expectedCurrent, expectedLocation, nil)

	current, location, err := mockClient.GetCurrentData(q, key)
	assert.NoError(t, err)
	assert.Equal(t, expectedCurrent, current)
	assert.Equal(t, expectedLocation, location)
}

func Test_GetForecastData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_services.NewMockWeatherDataGetter(ctrl)

	q := "London"
	days := int32(3)
	key := "test_key"

	expectedForecast := &api.Forecast{}
	expectedLocation := &api.Location{Name: "London"}

	mockClient.EXPECT().
		GetForecastData(q, days, key).
		Return(expectedForecast, expectedLocation, nil)

	forecast, location, err := mockClient.GetForecastData(q, days, key)
	assert.NoError(t, err)
	assert.Equal(t, expectedForecast, forecast)
	assert.Equal(t, expectedLocation, location)
}
