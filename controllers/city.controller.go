package controllers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"project_weather/models"
	"project_weather/repositories"
)

type CityDB interface {
	FindCityByID(id string) (*models.City, error)
	RegisterCity(city *models.City) error
	UpdateCity(city *models.City) error
	DeleteCityByID(id string) error
	GetAllCity() ([]models.City, error)
}

type CityController struct {
	DB repositories.CityRepo
}

func (r *CityController) GetCityById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "error",
			"message": "Id can't be empty",
		})
		return errors.New("Id can't be empty")
	}

	city, err := r.DB.FindCityByID(id)
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"status":  "error",
			"message": "City not found",
		})
		return err
	}
	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "City found",
		"data":    city,
	})
	return nil
}

func (r *CityController) DeleteCity(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	city, err := r.DB.FindCityByID(id)

	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"status":  "error",
			"message": "City not found",
		})
	}

	if id == "" {
		ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "error",
			"message": "Invalid id",
		})
		return errors.New("Invalid id")
	}
	err = r.DB.DeleteCityByID(id)
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
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

func (r *CityController) RegisterCity(ctx *fiber.Ctx) error {
	city := &models.City{}

	err := ctx.BodyParser(&city)
	if err != nil {
		ctx.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
		return err
	}

	city, err = r.DB.RegisterCity(city)
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
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

func (r *CityController) UpdateCity(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	city, err := r.DB.FindCityByID(id)
	// If no such note present return an error
	if id == "" {
		return ctx.Status(404).JSON(fiber.Map{"status": "error", "message": "No note present", "data": nil})
	}

	var updatedCity *models.City
	err = ctx.BodyParser(&updatedCity)

	if err != nil {
		ctx.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
		return err
	}

	city.Name = updatedCity.Name
	city.Longitude = updatedCity.Longitude
	city.Latitude = updatedCity.Latitude

	city, err = r.DB.RegisterCity(city)
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
		return err
	}
	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"success": true,
		"data":    city,
		"message": "City updated",
	})
	return nil
}

func (r *CityController) GetAllCities(ctx *fiber.Ctx) error {

	cities, err := r.DB.GetAllCity()
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"status":  "error",
			"message": err,
		})
		return err
	}
	if len(*cities) == 0 {
		return ctx.Status(404).JSON(fiber.Map{"status": "error", "message": "No notes present", "data": nil})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Success",
		"data":    cities,
	})
}

func (r *CityController) SetupRoutes(app *fiber.App) {
	api := app.Group("/cities")
	api.Post("/", r.RegisterCity)
	api.Delete("/:id", r.DeleteCity)
	api.Get("/:id", r.GetCityById)
	api.Put("/:id", r.UpdateCity)
	api.Get("/", r.GetAllCities)
}
