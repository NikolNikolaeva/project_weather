package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/NikolNikolaeva/project_weather/generated/dao/model"
	"github.com/NikolNikolaeva/project_weather/repositories"
)

type ForecastController interface {
	GetRoutes() []Route
	CreateForecast(ctx *fiber.Ctx) error
	GetAllForecasts(ctx *fiber.Ctx) error
	GetForecastByID(ctx *fiber.Ctx) error
	UpdateForecast(ctx *fiber.Ctx) error
	DeleteForecast(ctx *fiber.Ctx) error
}

type forecastController struct {
	DB repositories.ForecastRepo
}

func NewForecastController(db repositories.ForecastRepo) ForecastController {
	return &forecastController{
		DB: db,
	}
}

func (c *forecastController) GetRoutes() []Route {
	return []Route{
		{
			Method:  http.MethodPost,
			Path:    "/forecasts",
			Handler: c.CreateForecast,
		},
		{
			Method:  http.MethodGet,
			Path:    "/forecasts",
			Handler: c.GetAllForecasts,
		},
		{
			Method:  http.MethodGet,
			Path:    "/forecasts/:id",
			Handler: c.GetForecastByID,
		},
		{
			Method:  http.MethodPut,
			Path:    "/forecasts/:id",
			Handler: c.UpdateForecast,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/forecasts/:id",
			Handler: c.DeleteForecast,
		},
	}
}

// CreateForecast Create forecast
//
//	@Summary		Create forecast
//	@Description	Create forecast
//	@Tags			Forecast
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	string
//
//	@Failure		400	{string}	string
//	@Created		201 {string} string
//
//	@Router			/forecasts [post]
func (c *forecastController) CreateForecast(ctx *fiber.Ctx) error {
	forecast := new(model.Forecast)
	if err := ctx.BodyParser(forecast); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	if err := c.DB.Create(forecast); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusCreated).JSON(forecast)
}

// GetAllForecasts Get all forecast
//
//	@Summary		Get all forecast
//	@Description	Get all forecast details
//	@Tags			Forecast
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	string
//	@Failure		500	{string}	string
//	@Router			/forecasts [get]
func (c *forecastController) GetAllForecasts(ctx *fiber.Ctx) error {
	forecasts, err := c.DB.FindAll()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(forecasts)
}

// GetForecastByID Get forecast by id
//
//	@Summary		Get forecast by id
//	@Description	Get forecast by id
//	@Tags			Forecast
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	" get forecast by id"
//	@Success		200	{string}	string
//	@Failure		404	{string}	string
//	@Router			/forecasts/{id} [get]
func (c *forecastController) GetForecastByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	forecast, err := c.DB.FindByID(id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Forecast not found",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(forecast)
}

// UpdateForecast Update forecast by id
//
//	@Summary		Update forecast by id
//	@Description	Update forecast by id
//	@Tags			Forecast
//	@Accept			json
//	@Produce		json
//	@Param			id			path		string	true	" Update forecast by id"
//	@Param			forecast	body		string	true	" Update forecast by id"
//	@Success		200			{string}	string
//	@Failure		400			{string}	string
//	@Failure		500			{string}	string
//	@Router			/forecasts/{id} [put]
func (c *forecastController) UpdateForecast(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	forecast := new(model.Forecast)
	if err := ctx.BodyParser(forecast); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	if err := c.DB.Update(id, forecast); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(forecast)
}

// DeleteForecast Delete forecast by id
//
//	@Summary		Delete forecast by id
//	@Description	Delete forecast by id
//	@Tags			Forecast
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	" Delete forecast by id"
//	@Success		200	{string}	string
//	@Failure		500	{string}	string
//	@Router			/forecasts/{id} [delete]
func (c *forecastController) DeleteForecast(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if err := c.DB.Delete(id); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).SendString("Forecast deleted")
}
