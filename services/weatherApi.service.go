package services

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"net/http"
	"project_weather/config"
	"project_weather/controllers"
	"project_weather/generated/dao/model"
	"project_weather/repositories"
	"project_weather/resources/swagger"
	"time"
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
	cityRepo     repositories.CityRepo
	forecastRepo repositories.ForecastRepo
	client       *http.Client
	apiKey       string
	db           *gorm.DB
	config       *config.ApplicationConfiguration
}

func NewWeatherService(client *http.Client, apiKey string, db *gorm.DB, cityDB repositories.CityRepo, forecastDB repositories.ForecastRepo, config *config.ApplicationConfiguration) WeatherApiService {
	return &weatherApiService{
		client:       client,
		apiKey:       apiKey,
		db:           db,
		cityRepo:     cityDB,
		forecastRepo: forecastDB,
		config:       config,
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

	var url string
	if period == "current" {
		url = fmt.Sprintf(self.config.CurrentTimeUrl, self.apiKey, cityName)
	} else {
		url = fmt.Sprintf(self.config.ForecastUrl, self.apiKey, cityName, days)
	}

	weatherData, err := self.fetchWeatherData(url)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	city := &model.City{
		Name:      weatherData.Location.Name,
		Country:   weatherData.Location.Country,
		Latitude:  fmt.Sprintf("%f", weatherData.Location.Lat),
		Longitude: fmt.Sprintf("%f", weatherData.Location.Lon),
	}
	city, err = self.cityRepo.RegisterCity(city)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	formattedData, err := self.formatWeatherData(weatherData, period, city.ID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(http.StatusOK).JSON(formattedData)
}

func (self *weatherApiService) fetchWeatherData(url string) (*swagger.WeatherDTO, error) {
	response, err := self.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch weather data: status code %d", response.StatusCode)
	}

	var weatherData swagger.WeatherDTO
	if err := json.NewDecoder(response.Body).Decode(&weatherData); err != nil {
		return nil, fmt.Errorf("failed to parse weather data: %v", err)
	}

	return &weatherData, nil
}

func (self *weatherApiService) formatWeatherData(weatherData *swagger.WeatherDTO, period string, cityID string) ([]swagger.ForecastDTO, error) {

	var forecasts []swagger.ForecastDTO

	if period == "current" {
		// Format current weather data
		lastUpdated, _ := time.Parse(templateDateAndTime, weatherData.Current.LastUpdated)
		currentData := swagger.ForecastDTO{
			CityID:       cityID,
			ForecastDate: lastUpdated,
			Temperature:  fmt.Sprintf("%.1f", weatherData.Current.TempC),
			Condition:    weatherData.Current.Condition.Text,
		}
		forecasts = append(forecasts, currentData)
		return forecasts, nil
	}

	// Format forecast weather data
	for _, day := range weatherData.Forecast.ForecastDays {
		date, err := time.Parse(templateDate, day.Date)

		if err != nil {
			return nil, fmt.Errorf("failed to parse forecast date: %v", err)
		}
		forecast := swagger.ForecastDTO{
			CityID:       cityID,
			ForecastDate: date,
			Temperature:  fmt.Sprintf("%.1f", day.Day.AvgTempC),
			Condition:    day.Day.Condition.Text,
		}

		forecastToSave := &model.Forecast{
			CityID:       cityID,
			ForecastDate: date,
			Temperature:  fmt.Sprintf("%.1f", day.Day.AvgTempC),
			Condition:    day.Day.Condition.Text,
		}
		err = self.forecastRepo.Create(forecastToSave)
		if err != nil {
			return nil, err
		}

		forecasts = append(forecasts, forecast)
	}

	return forecasts, nil
}
