package controllers

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"project_weather/generated/dao/model"
	"project_weather/repositories"
)

type ForecastController struct {
	DB repositories.ForecastRepo
}

func NewForecastController(db repositories.ForecastRepo) *ForecastController {
	return &ForecastController{
		DB: db,
	}
}

func (c *ForecastController) GetRoutes() []Route {
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

func (c *ForecastController) CreateForecast(ctx *fiber.Ctx) error {
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

func (c *ForecastController) GetAllForecasts(ctx *fiber.Ctx) error {
	forecasts, err := c.DB.FindAll()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(forecasts)
}

func (c *ForecastController) GetForecastByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	forecast, err := c.DB.FindByID(id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Forecast not found",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(forecast)
}

func (c *ForecastController) UpdateForecast(ctx *fiber.Ctx) error {
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

func (c *ForecastController) DeleteForecast(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if err := c.DB.Delete(id); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).SendString("Forecast deleted")
}
