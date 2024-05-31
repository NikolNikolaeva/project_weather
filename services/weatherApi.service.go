package services

import (
	"github.com/NikolNikolaeva/project_weather/config"
	"github.com/NikolNikolaeva/project_weather/controllers"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type WeatherApiService interface {
	GetRoutes() []controllers.Route
	GetWeatherByCity(ctx *fiber.Ctx) error
}

var periodToDays = map[string]int{
	"daily":   1,
	"weekly":  7,
	"monthly": 30,
	"current": 0,
}

type weatherApiService struct {
	weatherHandler WeatherHandler
	config         *config.ApplicationConfiguration
	credFile       string
}

func NewWeatherService(
	xHandler WeatherHandler,
	config *config.ApplicationConfiguration,
	credFile string,
) WeatherApiService {

	return &weatherApiService{
		weatherHandler: xHandler,
		config:         config,
		credFile:       credFile,
	}
}

func (self *weatherApiService) GetRoutes() []controllers.Route {
	return []controllers.Route{
		{
			Method:  http.MethodGet,
			Path:    "/weather/:city/:period",
			Handler: self.GetWeatherByCity,
		},
	}
}

func (self *weatherApiService) GetWeatherByCity(ctx *fiber.Ctx) error {
	cityName := ctx.Params("city")
	period := ctx.Params("period")

	days, exists := periodToDays[period]
	if !exists {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid period. Please specify 'current', 'daily', 'weekly', or 'monthly'.",
		})
	}

	url, err := self.weatherHandler.GetUrlForWeatherApi(period, self.credFile, cityName, days, self.config.ForecastUrl, self.config.CurrentTimeUrl)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "error getting weather url, missing cred file with key",
		})
	}

	res, err := self.weatherHandler.Handle(url, period)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
	}

	return ctx.Status(http.StatusOK).JSON(res)
}
