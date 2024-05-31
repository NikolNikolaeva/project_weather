package controllers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/NikolNikolaeva/project_weather/generated/dao/model"
	"github.com/NikolNikolaeva/project_weather/mocks"
)

func TestCreateForecast(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockForecastRepo := repositories.NewMockForecastRepo(ctrl)
	controller := NewForecastController(mockForecastRepo)

	mockForecastRepo.EXPECT().Create(gomock.Any()).Return(nil)

	app := fiber.New()
	app.Post("/forecasts", controller.CreateForecast)

	reqBody := strings.NewReader(`{"CityID": "1", "Temperature": "22.5", "Condition": "Sunny"}`)
	req := httptest.NewRequest(http.MethodPost, "/forecasts", reqBody)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}

func TestGetAllForecasts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockForecastRepo := repositories.NewMockForecastRepo(ctrl)
	controller := NewForecastController(mockForecastRepo)

	testForecasts := []*model.Forecast{
		{CityID: "1", Temperature: "22.5", Condition: "Sunny"},
		{CityID: "2", Temperature: "18.0", Condition: "Cloudy"},
	}

	mockForecastRepo.EXPECT().FindAll().Return(testForecasts, nil)

	app := fiber.New()
	app.Get("/forecasts", controller.GetAllForecasts)

	req := httptest.NewRequest(http.MethodGet, "/forecasts", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestGetForecastByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockForecastRepo := repositories.NewMockForecastRepo(ctrl)
	controller := NewForecastController(mockForecastRepo)

	testID := "1"
	testForecast := &model.Forecast{CityID: testID, Temperature: "22.5", Condition: "Sunny"}

	mockForecastRepo.EXPECT().FindByID(testID).Return(testForecast, nil)

	app := fiber.New()
	app.Get("/forecasts/:id", controller.GetForecastByID)

	req := httptest.NewRequest(http.MethodGet, "/forecasts/1", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var responseBody struct {
		Err error `json:"err"`
	}

	err = json.NewDecoder(resp.Body).Decode(&responseBody)

	assert.NoError(t, err)
	assert.Equal(t, nil, responseBody.Err)
}

func TestUpdateForecast(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockForecastRepo := repositories.NewMockForecastRepo(ctrl)
	controller := NewForecastController(mockForecastRepo)

	testID := "1"

	mockForecastRepo.EXPECT().Update(testID, gomock.Any()).Return(nil)

	app := fiber.New()
	app.Put("/forecasts/:id", controller.UpdateForecast)

	reqBody := strings.NewReader(`{"CityID": "1", "Temperature": "22.5", "Condition": "Sunny"}`)
	req := httptest.NewRequest(http.MethodPut, "/forecasts/1", reqBody)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var responseBody struct {
		Err error `json:"err"`
	}

	err = json.NewDecoder(resp.Body).Decode(&responseBody)

	assert.NoError(t, err)
	assert.Equal(t, nil, responseBody.Err)
}

func TestDeleteForecast(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockForecastRepo := repositories.NewMockForecastRepo(ctrl)
	controller := NewForecastController(mockForecastRepo)

	testID := "1"

	mockForecastRepo.EXPECT().Delete(testID).Return(nil)

	app := fiber.New()
	app.Delete("/forecasts/:id", controller.DeleteForecast)

	req := httptest.NewRequest(http.MethodDelete, "/forecasts/1", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

}
