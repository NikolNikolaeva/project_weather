package repositories

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"net/http"
	"project_weather/models"
)

type CityRepository struct {
	DB *gorm.DB
}

func (r *CityRepository) GetCityById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	city := &models.City{}
	if id == "" {
		ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "error",
			"message": "Id can't be empty",
		})
		return errors.New("Id can't be empty")
	}

	fmt.Println("ID: ", id)
	err := r.DB.Where("id = ?", id).First(city).Error
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

func (r *CityRepository) DeleteCity(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	city := &models.City{ID: id}

	if id == "" {
		ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "error",
			"message": "Invalid id",
		})
		return errors.New("Invalid id")
	}
	err := r.DB.Delete(&city)
	if err.Error != nil {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"status":  "error",
			"message": err,
		})
		return err.Error
	}
	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "City deleted",
		"data":    city,
	})
	return nil
}

func (r *CityRepository) RegisterCity(ctx *fiber.Ctx) error {
	city := &models.City{}

	err := ctx.BodyParser(&city)
	if err != nil {
		ctx.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
		})
		return err
	}

	err = r.DB.Create(&city).Error
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

func (r *CityRepository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/register_city", r.RegisterCity)
	api.Delete("/delete_city/:id", r.DeleteCity)
	api.Get("/get_city/:id", r.GetCityById)
}
