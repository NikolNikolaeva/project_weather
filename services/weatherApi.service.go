package services

import (
	"fmt"
	"net/http"

	"github.com/NikolNikolaeva/project_weather/config"
	"github.com/NikolNikolaeva/project_weather/controllers"

	"github.com/gofiber/fiber/v2"
)

const (
	templateDate        = "2006-01-02"
	templateDateAndTime = "2006-01-02 15:04"
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
	xHandler WeatherHandler
	apiKey       string
	config       *config.ApplicationConfiguration
}

func NewWeatherService(
	apiKey string,
	xHandler WeatherHandler,
	config       *config.ApplicationConfiguration,
	) WeatherApiService {

	return &weatherApiService{
		apiKey:       apiKey,
		xHandler: xHandler,
		config: config,
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

	// TODO: Move to the new class
	var url string
	if period == "current" {
		url = fmt.Sprintf(self.config.CurrentTimeUrl, self.apiKey, cityName)
	} else {
		url = fmt.Sprintf(self.config.ForecastUrl, self.apiKey, cityName, days)
	}

	res, err := self.xHandler.Handle(url, period)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
	}

	return ctx.Status(http.StatusOK).JSON(res)
}
