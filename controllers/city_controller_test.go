package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	mock_repositories "github.com/NikolNikolaeva/project_weather/generated/go-mocks/repositories"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/NikolNikolaeva/project_weather/generated/dao/model"
)

func Test_GetCityById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCityRepo := mock_repositories.NewMockCityRepo(ctrl)
	controller := NewCityController(mockCityRepo)

	testID := "1"
	testCity := &model.City{ID: testID, Name: "Sofia", Country: "Bulgaria"}

	mockCityRepo.EXPECT().FindCityByID(testID).Return(testCity, nil)

	app := fiber.New()
	app.Get("/cities/:id", controller.GetCityById)

	req := httptest.NewRequest(http.MethodGet, "/cities/1", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var responseBody struct {
		Status string      `json:"status"`
		Data   *model.City `json:"data"`
	}

	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	assert.NoError(t, err)
	assert.Equal(t, "success", responseBody.Status)
	assert.Equal(t, testCity, responseBody.Data)
}

func Test_DeleteCity(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCityRepo := mock_repositories.NewMockCityRepo(ctrl)
	controller := NewCityController(mockCityRepo)

	testID := "1"
	testCity := &model.City{ID: testID, Name: "Sofia", Country: "Bulgaria"}

	mockCityRepo.EXPECT().FindCityByID(testID).Return(testCity, nil)
	mockCityRepo.EXPECT().DeleteCityByID(testID).Return(nil)

	app := fiber.New()
	app.Delete("/cities/:id", controller.DeleteCity)

	req := httptest.NewRequest(http.MethodDelete, "/cities/1", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var responseBody struct {
		Status string      `json:"status"`
		Data   *model.City `json:"data"`
	}

	err = json.NewDecoder(resp.Body).Decode(&responseBody)

	assert.NoError(t, err)
	assert.Equal(t, "success", responseBody.Status)
	assert.Equal(t, testCity, responseBody.Data)
}

func TestRegisterCity(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCityRepo := mock_repositories.NewMockCityRepo(ctrl)
	controller := NewCityController(mockCityRepo)

	testCity := &model.City{Name: "Sofia", Country: "Bulgaria"}
	mockCityRepo.EXPECT().RegisterCity(gomock.Any()).Return(testCity, nil)

	app := fiber.New()
	app.Post("/cities/", controller.RegisterCity)

	reqBody := strings.NewReader(`{"name": "Sofia", "country": "Bulgaria"}`)
	req := httptest.NewRequest(http.MethodPost, "/cities/", reqBody)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var responseBody struct {
		Status bool        `json:"success"`
		Data   *model.City `json:"data"`
	}

	err = json.NewDecoder(resp.Body).Decode(&responseBody)

	assert.NoError(t, err)
	assert.Equal(t, true, responseBody.Status)
	assert.Equal(t, testCity, responseBody.Data)
}

func TestUpdateCity(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCityRepo := mock_repositories.NewMockCityRepo(ctrl)
	controller := NewCityController(mockCityRepo)

	testID := "1"
	testCity := &model.City{Name: "Sofia", Country: "Bulgaria"}
	mockCityRepo.EXPECT().UpdateCityByID(testID, gomock.Any()).Return(testCity, nil)

	app := fiber.New()
	app.Put("/cities/:id", controller.UpdateCity)

	reqBody := strings.NewReader(`{"name": "Sofia", "country": "Bulgaria"}`)
	req := httptest.NewRequest(http.MethodPut, "/cities/1", reqBody)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestGetAllCities(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCityRepo := mock_repositories.NewMockCityRepo(ctrl)
	controller := NewCityController(mockCityRepo)

	testCities := []*model.City{
		{ID: "1", Name: "Sofia", Country: "Bulgaria"},
		{ID: "2", Name: "Plovdiv", Country: "Bulgaria"},
	}

	mockCityRepo.EXPECT().GetAllCity().Return(testCities, nil)

	app := fiber.New()
	app.Get("/cities/", controller.GetAllCities)

	req := httptest.NewRequest(http.MethodGet, "/cities/", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var responseBody struct {
		Status string        `json:"status"`
		Data   []*model.City `json:"data"`
	}

	err = json.NewDecoder(resp.Body).Decode(&responseBody)

	assert.NoError(t, err)
	assert.Equal(t, len(testCities), len(responseBody.Data))
	assert.Equal(t, "success", responseBody.Status)
	assert.Equal(t, testCities, responseBody.Data)
}
