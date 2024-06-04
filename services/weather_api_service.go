package services

import (
	"net/http"

	"github.com/NikolNikolaeva/project_weather/config"
	"github.com/NikolNikolaeva/project_weather/controllers"

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

// GetWeatherByCity Get weather data
//
//	@Summary		Get weather data
//	@Description	Get weather data details
//	@Tags			Weather
//	@Accept			json
//	@Produce		json
//	@Param			city	path		string	true	"weather search by a city"
//	@Param			period	path		string	true	"weather search by a period"
//	@Success		200		{string}	string
//	@Router			/weather/{city}/{period} [get]
func (self *weatherApiService) GetWeatherByCity(ctx *fiber.Ctx) error {
	cityName := ctx.Params("city")
	period := ctx.Params("period")

	days, exists := periodToDays[period]
	if !exists {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid period. Please specify 'current', 'daily', 'weekly', or 'monthly'.",
		})
	}

	if period == "current" {
		current, err := self.weatherHandler.HandleCurrantData(cityName, self.credFile)
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return ctx.Status(http.StatusOK).JSON(fiber.Map{
			"data": current,
		})
	} else {
		forecast, err := self.weatherHandler.HandleForecast(cityName, int32(days), self.credFile)
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return ctx.Status(http.StatusOK).JSON(fiber.Map{
			"data": forecast,
		})
	}
}
