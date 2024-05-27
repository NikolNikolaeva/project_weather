package services

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"net/http"
	"project_weather/controllers"
	"project_weather/generated/dao/model"
	"project_weather/resources/dto"
	"time"
)

const (
	forecastUrl         = "http://api.weatherapi.com/v1/forecast.json?key=%s&q=%s&days=%v&aqi=no&alerts=no"
	currentTimeUrl      = "http://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no"
	templateDate        = "2006-01-02"
	templateDateAndTime = "2006-01-02 15:04"
)

var periodToDays = map[string]int{
	"daily":   1,
	"weekly":  7,
	"monthly": 30,
	"current": 0,
}

type WeatherApiService struct {
	client *http.Client
	apiKey string
	db     *gorm.DB
}

func NewWeatherService(client *http.Client, apiKey string, db *gorm.DB) *WeatherApiService {
	return &WeatherApiService{
		client: client,
		apiKey: apiKey,
		db:     db,
	}
}

func (self *WeatherApiService) GetRoutes() []controllers.Route {
	return []controllers.Route{
		{
			Method:  http.MethodGet,
			Path:    "/weather/:city/:period",
			Handler: self.GetWeatherByCity,
		},
		{
			Method:  http.MethodPost,
			Path:    "/weather/city",
			Handler: self.SaveCity,
		},
		{
			Method:  http.MethodPost,
			Path:    "/weather/forecast",
			Handler: self.SaveForecast,
		},
	}
}

func (self *WeatherApiService) GetWeatherByCity(ctx *fiber.Ctx) error {
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
		url = fmt.Sprintf(currentTimeUrl, self.apiKey, cityName)
	} else {
		url = fmt.Sprintf(forecastUrl, self.apiKey, cityName, days)
	}

	weatherData, err := self.fetchWeatherData(url)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	city, err := self.saveCityData(weatherData)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	formattedData := self.formatWeatherData(weatherData, period, city.ID)
	return ctx.Status(http.StatusOK).JSON(formattedData)
}

func (self *WeatherApiService) fetchWeatherData(url string) (*dto.WeatherDTO, error) {
	response, err := self.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch weather data: status code %d", response.StatusCode)
	}

	var weatherData dto.WeatherDTO
	if err := json.NewDecoder(response.Body).Decode(&weatherData); err != nil {
		return nil, fmt.Errorf("failed to parse weather data: %v", err)
	}

	return &weatherData, nil
}

func (self *WeatherApiService) saveCityData(weatherData *dto.WeatherDTO) (*model.City, error) {
	var city model.City
	err := self.db.Where("name = ?", weatherData.Location.Name).First(&city).Error
	fmt.Println(err)
	if err == nil {
		// City already exists
		return &city, nil
	}

	// Create new city
	city = model.City{
		Name:      weatherData.Location.Name,
		Country:   weatherData.Location.Country,
		Latitude:  fmt.Sprintf("%f", weatherData.Location.Lat),
		Longitude: fmt.Sprintf("%f", weatherData.Location.Lon),
	}

	if err := self.db.Create(&city).Error; err != nil {
		return nil, err
	}

	return &city, nil
}

func (self *WeatherApiService) formatWeatherData(weatherData *dto.WeatherDTO, period string, cityID string) interface{} {
	if period == "current" {
		// Format current weather data
		lastUpdated, _ := time.Parse(templateDateAndTime, weatherData.Current.LastUpdated)
		currentData := dto.ForecastDTO{
			CityID:       cityID,
			ForecastDate: lastUpdated,
			Temperature:  fmt.Sprintf("%.1f", weatherData.Current.TempC),
			Condition:    weatherData.Current.Condition.Text,
		}
		self.saveForecastData(currentData)
		return currentData
	}

	// Format forecast weather data
	var forecasts []dto.ForecastDTO
	for _, day := range weatherData.Forecast.ForecastDays {
		date, _ := time.Parse(templateDate, day.Date)
		forecast := dto.ForecastDTO{
			CityID:       cityID,
			ForecastDate: date,
			Temperature:  fmt.Sprintf("%.1f", day.Day.AvgTempC),
			Condition:    day.Day.Condition.Text,
		}
		self.saveForecastData(forecast)
		forecasts = append(forecasts, forecast)
	}

	return forecasts
}

func (self *WeatherApiService) saveForecastData(forecast dto.ForecastDTO) {
	forecastModel := model.Forecast{
		CityID:       forecast.CityID,
		ForecastDate: forecast.ForecastDate,
		Temperature:  forecast.Temperature,
		Condition:    forecast.Condition,
	}
	self.db.Create(&forecastModel)
}

func (self *WeatherApiService) SaveCity(ctx *fiber.Ctx) error {
	var cityDTO dto.CityDTO
	if err := ctx.BodyParser(&cityDTO); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	city := model.City{
		ID:        cityDTO.ID,
		Name:      cityDTO.Name,
		Country:   cityDTO.Country,
		Latitude:  cityDTO.Latitude,
		Longitude: cityDTO.Longitude,
		CreatedAt: cityDTO.CreatedAt,
		UpdatedAt: cityDTO.UpdatedAt,
	}

	if err := self.db.Create(&city).Error; err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save city",
		})
	}

	return ctx.Status(http.StatusOK).JSON(city)
}

func (self *WeatherApiService) SaveForecast(ctx *fiber.Ctx) error {
	var forecastDTO dto.ForecastDTO
	if err := ctx.BodyParser(&forecastDTO); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	forecast := model.Forecast{
		ID:           forecastDTO.ID,
		CityID:       forecastDTO.CityID,
		ForecastDate: forecastDTO.ForecastDate,
		Temperature:  forecastDTO.Temperature,
		Condition:    forecastDTO.Condition,
		CreatedAt:    forecastDTO.CreatedAt,
		UpdatedAt:    forecastDTO.UpdatedAt,
	}

	if err := self.db.Create(&forecast).Error; err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save forecast",
		})
	}

	return ctx.Status(http.StatusOK).JSON(forecast)
}
