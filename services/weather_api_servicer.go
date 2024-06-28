package services

import (
	"log"

	"github.com/NikolNikolaeva/project_weather/config"
	"github.com/NikolNikolaeva/project_weather/repositories"
)

//go:generate mockgen --build_flags=--mod=mod -destination ../generated/go-mocks/services/mock_weather_api_service.go . WeatherService
type WeatherService interface {
	StartFetching()
}

type weatherService struct {
	dataGetter    WeatherDataGetter
	cityRepo      repositories.CityRepo
	forecastRepo  repositories.ForecastRepo
	handler       WeatherAPIClient
	configuration *config.ApplicationConfiguration
}

func NewWeatherService(dataGetter WeatherDataGetter, cityRepo repositories.CityRepo, forecastRepo repositories.ForecastRepo, configuration *config.ApplicationConfiguration, handler WeatherAPIClient) WeatherService {
	return &weatherService{
		dataGetter:    dataGetter,
		cityRepo:      cityRepo,
		forecastRepo:  forecastRepo,
		configuration: configuration,
		handler:       handler,
	}
}

func (ws *weatherService) StartFetching() {
	ws.fetchAndStoreWeatherData()
}

func (ws *weatherService) fetchAndStoreWeatherData() {
	cities, err := ws.cityRepo.GetAll()
	if err != nil {
		log.Printf("Error fetching cities")
		return
	}

	for _, city := range cities {
		ws.updateCityForecast(city.Name)
	}
}

func (ws *weatherService) updateCityForecast(cityName string) {

	_, err := ws.handler.HandleForecast(cityName, 30, ws.configuration.CredFile)
	if err != nil {
		log.Printf("Error fetching forecast for city %s: %v", cityName, err)
	}
}
