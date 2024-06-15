package services

import (
	"log"
	"time"

	"github.com/NikolNikolaeva/project_weather/config"
	"github.com/NikolNikolaeva/project_weather/generated/dao/model"
	"github.com/NikolNikolaeva/project_weather/repositories"
)

type WeatherService interface {
	StartFetching(interval time.Duration)
}

type weatherService struct {
	dataGetter    WeatherDataGetter
	cityRepo      repositories.CityRepo
	forecastRepo  repositories.ForecastRepo
	handler       WeatherHandler
	configuration *config.ApplicationConfiguration
}

func NewWeatherService(dataGetter WeatherDataGetter, cityRepo repositories.CityRepo, forecastRepo repositories.ForecastRepo, configuration *config.ApplicationConfiguration, handler WeatherHandler) WeatherService {
	return &weatherService{
		dataGetter:    dataGetter,
		cityRepo:      cityRepo,
		forecastRepo:  forecastRepo,
		configuration: configuration,
		handler:       handler,
	}
}

func (ws *weatherService) StartFetching(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		ws.fetchAndStoreWeatherData()
	}
}

func (ws *weatherService) fetchAndStoreWeatherData() {
	cities, err := ws.cityRepo.GetAllCity()
	if err != nil {
		log.Printf("Error fetching cities: %v", err)
		return
	}

	for _, city := range cities {
		ws.updateCityForecast(city)
	}
}

func (ws *weatherService) updateCityForecast(city *model.City) {

	_, err := ws.handler.HandleForecast(city.Name, 30, ws.configuration.CredFile)
	if err != nil {
		log.Printf("Error fetching forecast for city %s: %v", city.Name, err)
		return
	}

}
