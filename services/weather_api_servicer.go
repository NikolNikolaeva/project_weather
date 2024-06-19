package services

import (
	"log"

	"github.com/NikolNikolaeva/project_weather/config"
	"github.com/NikolNikolaeva/project_weather/repositories"
)

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
	//ticker := time.NewTicker(interval)
	//defer ticker.Stop()
	//
	//for range ticker.C {
	ws.fetchAndStoreWeatherData()
	//}
}

func (ws *weatherService) fetchAndStoreWeatherData() {
	cities, err := ws.cityRepo.GetAll()
	if err != nil {
		log.Printf("Error fetching cities: %v", err)
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
