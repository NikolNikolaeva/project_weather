package controllers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"project_weather/generated/dao/model"
	"project_weather/repositories"
)

type CityDBController interface {
	GetRoutes() []Route
	GetCityById(ctx *fiber.Ctx) error
	DeleteCity(ctx *fiber.Ctx) error
	RegisterCity(ctx *fiber.Ctx) error
	UpdateCity(ctx *fiber.Ctx) error
	GetAllCities(ctx *fiber.Ctx) error
}

type cityController struct {
	DB repositories.CityDB
}

func NewCityController(db repositories.CityDB) CityDBController {
	return &cityController{
		DB: db,
	}
}

func (r *cityController) GetRoutes() []Route {
	return []Route{
		{
			Method:  http.MethodPost,
			Path:    "/cities/",
			Handler: r.RegisterCity,
		},
		{
			Method:  http.MethodGet,
			Path:    "/cities/",
			Handler: r.GetAllCities,
		},
		{
			Method:  http.MethodGet,
			Path:    "/cities/:id",
			Handler: r.GetCityById,
		},
		{
			Method:  http.MethodPut,
			Path:    "/cities/:id",
			Handler: r.UpdateCity,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/cities/:id",
			Handler: r.DeleteCity,
		},
	}
}

func (r *cityController) GetCityById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"status":  "error",
			"message": "Id can't be empty",
		})
		return errors.New("id can't be empty")
	}

	city, err := r.DB.FindCityByID(id)
	if err != nil {
		ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"status": "error",
		})
		log.Printf(err.Error())
		return err
	}
	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "City found",
		"data":    city,
	})
	return nil
}

func (r *cityController) DeleteCity(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	city, err := r.DB.FindCityByID(id)

	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"status":  "error",
			"message": "City not found",
		})
	}

	if id == "" {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"status":  "error",
			"message": "Invalid id",
		})
		return errors.New("Invalid id")
	}
	err = r.DB.DeleteCityByID(id)
	if err != nil {
		ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "error",
			"message": err,
		})
		return err
	}
	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "City deleted",
		"data":    city,
	})
	return nil
}

func (r *cityController) RegisterCity(ctx *fiber.Ctx) error {
	city := &model.City{}

	err := ctx.BodyParser(&city)
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
		return err
	}

	city, err = r.DB.RegisterCity(city)
	if err != nil {
		ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
		return err
	}

	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"success": true,
		"data":    city,
		"message": "City created",
	})
	return nil
}

func (r *cityController) UpdateCity(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	city := new(model.City)
	if err := ctx.BodyParser(city); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	updatedCity, err := r.DB.UpdateCityByID(id, city)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(updatedCity)
}

func (r *cityController) GetAllCities(ctx *fiber.Ctx) error {

	cities, err := r.DB.GetAllCity()
	if err != nil {
		ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "error",
			"message": err,
		})
		return err
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Success",
		"data":    cities,
	})
}
