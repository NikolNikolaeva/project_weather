package services

import (
	"fmt"
	"time"

	"github.com/NikolNikolaeva/project_weather/generated/dao/model"
	"github.com/NikolNikolaeva/project_weather/repositories"
	"github.com/NikolNikolaeva/project_weather/resources/swagger"
)

type WeatherHandler interface {
	Handle(url, period string) ([]swagger.ForecastDTO, error)
}

type weatherHandler struct {
	cityRepo repositories.CityRepo
	forecastRepo repositories.ForecastRepo
	weatherDataGetter WeatherDataGetter
}

func NewWeatherHandler(cityRepo repositories.CityRepo, foreCastRepo repositories.ForecastRepo, weatherDataGetter WeatherDataGetter) WeatherHandler {
	return &weatherHandler{
		cityRepo:     cityRepo,
		forecastRepo: foreCastRepo,
		weatherDataGetter: weatherDataGetter,
	}
}

func (self *weatherHandler) Handle(url, period string) ([]swagger.ForecastDTO, error) {
	weatherData, err := self.weatherDataGetter.GetData(url)
	if err != nil {
		return nil, err
	}

	city := &model.City{
		Name:      weatherData.Location.Name,
		Country:   weatherData.Location.Country,
		Latitude:  fmt.Sprintf("%f", weatherData.Location.Lat),
		Longitude: fmt.Sprintf("%f", weatherData.Location.Lon),
	}
	city, err = self.cityRepo.RegisterCity(city)
	if err != nil {
		return nil, err
	}

	if period == "current" {
		currentWeather, err := self.formatWeatherCurrentData(weatherData, city.ID)
		if err != nil {
			return nil, err
		}

		return []swagger.ForecastDTO{currentWeather}, nil
	}

	forecast, err := self.formatWeatherForecastData(weatherData, city.ID)
	if err != nil {
		return nil, err
	}

	return forecast, nil
}

func (self *weatherHandler) formatWeatherForecastData(weatherData *swagger.WeatherDTO, cityID string) ([]swagger.ForecastDTO, error) {
	// Format forecast weather data
	var forecasts []swagger.ForecastDTO
	for _, day := range weatherData.Forecast.ForecastDays {
		date, err := time.Parse(templateDate, day.Date)
		fmt.Println(date, err)
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

func (self *weatherHandler) formatWeatherCurrentData(weatherData *swagger.WeatherDTO, cityID string) (swagger.ForecastDTO, error) {
	// Format current weather data
	lastUpdated, err := time.Parse(templateDateAndTime, weatherData.Current.LastUpdated)
	if err != nil {
		return swagger.ForecastDTO{}, err
	}

	currentData := swagger.ForecastDTO{
		CityID:       cityID,
		ForecastDate: lastUpdated,
		Temperature:  fmt.Sprintf("%.1f", weatherData.Current.TempC),
		Condition:    weatherData.Current.Condition.Text,
	}

	return currentData, nil
}
