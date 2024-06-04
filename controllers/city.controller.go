package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/NikolNikolaeva/project_weather/generated/dao/model"
	"github.com/NikolNikolaeva/project_weather/repositories"
	"github.com/gofiber/fiber/v2"
)

type CityController interface {
	GetRoutes() []Route
	GetCityById(ctx *fiber.Ctx) error
	DeleteCity(ctx *fiber.Ctx) error
	RegisterCity(ctx *fiber.Ctx) error
	UpdateCity(ctx *fiber.Ctx) error
	GetAllCities(ctx *fiber.Ctx) error
}

type cityController struct {
	DB repositories.CityRepo
}

func NewCityController(db repositories.CityRepo) CityController {
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

// GetCityById Get city by id
//
//	@Summary		Get city by id
//	@Description	Get city details by id
//	@Tags			City
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"city search by an id"
//	@Success		200	{string}	string
//	@Router			/cities/{id} [get]
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

// DeleteCity Delete city by id
//
//	@Summary		Delete city by id
//	@Description	Delete city by id
//	@Tags			City
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"city search by an id"
//	@Success		200	{string}	string
//	@Router			/cities/{id} [delete]
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

// RegisterCity Register city by id
//
//	@Summary		Register city by id
//	@Description	Register city by id
//	@Tags			City
//	@Accept			json
//	@Produce		json
//	@Param			city	body		string	true	" register city"
//	@Success		200		{string}	string
//
//	@Failure		400		{string}	string
//
//	@Failure		500		{string}	string
//	@Router			/cities [post]
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

// UpdateCity Update city by id
//
//	@Summary		Update city by id
//	@Description	Update city by id
//	@Tags			City
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string	true	" update city by id"
//	@Param			city	body		string	true	" update city"
//	@Success		200		{string}	string
//
//	@Failure		400		{string}	string
//
//	@Failure		500		{string}	string
//	@Router			/cities/{id} [put]
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

// GetAllCities Get all cities
//
//	@Summary		Get all cities
//	@Description	Get all cities details
//	@Tags			City
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	string
//	@Failure		500	{string}	string
//	@Router			/cities [get]
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
