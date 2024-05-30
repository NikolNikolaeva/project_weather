package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/NikolNikolaeva/project_weather/generated/dao/model"
	"github.com/NikolNikolaeva/project_weather/repositories"
	"github.com/NikolNikolaeva/project_weather/resources/swagger"
)

const (
	templateDate        = "2006-01-02"
	templateDateAndTime = "2006-01-02 15:04 "
)

type Cred struct {
	ApiKey string `json:"apiKey"`
}

type WeatherHandler interface {
	Handle(url, period string) ([]swagger.ForecastDTO, error)
	GetUrlForWeatherApi(period string, apiKey string, city string, days int, ForecastUrl string, CurrentTimeUrl string) string
	getApiKey(credFile string) (string, error)
}

type weatherHandler struct {
	cityRepo          repositories.CityRepo
	forecastRepo      repositories.ForecastRepo
	weatherDataGetter WeatherDataGetter
}

func NewWeatherHandler(cityRepo repositories.CityRepo, foreCastRepo repositories.ForecastRepo, weatherDataGetter WeatherDataGetter) WeatherHandler {
	return &weatherHandler{
		cityRepo:          cityRepo,
		forecastRepo:      foreCastRepo,
		weatherDataGetter: weatherDataGetter,
	}
}

func (self *weatherHandler) getApiKey(credFile string) (string, error) {
	jsonFile, err := os.Open(credFile)
	if err != nil {
		return "", err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return "", err
	}

	var cred Cred = Cred{}
	json.Unmarshal(byteValue, &cred)

	if err != nil {
		return "", err
	}
	return cred.ApiKey, nil
}

func (self *weatherHandler) GetUrlForWeatherApi(period string, apiKey string, city string, days int, ForecastUrl string, CurrentTimeUrl string) string {
	if period == "current" {
		return fmt.Sprintf(CurrentTimeUrl, apiKey, city)
	} else {
		return fmt.Sprintf(ForecastUrl, apiKey, city, days)
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
	fmt.Println(lastUpdated)
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
