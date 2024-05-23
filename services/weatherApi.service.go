package services

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"project_weather/controllers"
)

const (
	dailyUrl  = "http://api.weatherapi.com/v1/forecast.json?key=%s&q=%s&days=1&aqi=no&alerts=no"
	weeklyUrl = "http://api.weatherapi.com/v1/forecast.json?key=%s&q=%s&days=7&aqi=no&alerts=no"
	monthly   = "http://api.weatherapi.com/v1/forecast.json?key=%s&q=%s&days=30&aqi=no&alerts=no"
)

type WeatherApiService struct {
	client *http.Client
	apiKey string
}

func NewWeatherController(client *http.Client, apiKey string) *WeatherApiService {
	return &WeatherApiService{
		client: client,
		apiKey: apiKey,
	}
}

func (self *WeatherApiService) GetRoutes() []controllers.Route {
	return []controllers.Route{
		{
			Method:  http.MethodGet,
			Path:    "/weather/:city/:period",
			Handler: self.GetWeatherByCity,
		},
	}
}

func (self *WeatherApiService) GetWeatherByCity(ctx *fiber.Ctx) error {
	city := ctx.Params("city")
	period := ctx.Params("period") // period can be "daily", "weekly", or "monthly"

	var url string
	switch period {
	case "daily":
		url = fmt.Sprintf(dailyUrl, self.apiKey, city)
	case "weekly":
		url = fmt.Sprintf(weeklyUrl, self.apiKey, city)
	case "monthly":
		url = fmt.Sprintf(monthly, self.apiKey, city)
	default:
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid period. Please specify 'daily', 'weekly', or 'monthly'.",
		})
	}

	weatherData, err := self.fetchWeatherData(url)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(weatherData)
}

func (self *WeatherApiService) fetchWeatherData(url string) (interface{}, error) {
	response, err := self.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch weather data: status code %d", response.StatusCode)
	}

	var weatherData interface{}
	if err := json.NewDecoder(response.Body).Decode(&weatherData); err != nil {
		return nil, fmt.Errorf("failed to parse weather data: %v", err)
	}

	return weatherData, nil
}
