package services

import (
	"fmt"
	"log"
	"time"

	"github.com/NikolNikolaeva/project_weather/resources"

	api "github.com/NikolNikolaeva/project_weather/generated/api/weatherapi"
	"github.com/NikolNikolaeva/project_weather/generated/dao/model"
	"github.com/NikolNikolaeva/project_weather/repositories"
)

const (
	templateDate        = "2006-01-02"
	templateDateAndTime = "2006-01-02 15:04 "
)

type WeatherAPIClient interface {
	HandleCurrantData(q string, cred string) (*api.Current, error)
	HandleForecast(q string, days int32, cred string) (*api.Forecast, error)
	HandleCityData(name string, cred string) *api.Location
}

type weatherAPIClient struct {
	cityRepo          repositories.CityRepo
	forecastRepo      repositories.ForecastRepo
	weatherDataGetter WeatherDataGetter
}

func NewWeatherAPIClient(cityRepo repositories.CityRepo, foreCastRepo repositories.ForecastRepo, weatherDataGetter WeatherDataGetter) WeatherAPIClient {
	return &weatherAPIClient{
		cityRepo:          cityRepo,
		forecastRepo:      foreCastRepo,
		weatherDataGetter: weatherDataGetter,
	}
}

func (self *weatherAPIClient) HandleCurrantData(q string, cred string) (*api.Current, error) {
	key, err := resources.GetApiKey(cred)

	if err != nil {
		return nil, err
	}
	currentData, location, err := self.weatherDataGetter.GetCurrentData(q, key)
	if err != nil {
		return nil, err
	}

	_, err = self.handleCity(location)
	if err != nil {
		return nil, err
	}

	output := self.formCurrentData(currentData)

	return output, nil
}

func (self *weatherAPIClient) HandleForecast(q string, days int32, cred string) (*api.Forecast, error) {
	key, err := resources.GetApiKey(cred)
	if err != nil {
		return nil, err
	}

	forecast, location, err := self.weatherDataGetter.GetForecastData(q, days, key)
	if err != nil {
		return nil, err
	}

	city, err := self.handleCity(location)
	if err != nil {
		return nil, err
	}

	forecastOutput := api.Forecast{}
	forecastOutput.Forecastday = []api.ForecastForecastday{}
	for _, day := range forecast.Forecastday {

		date, err := time.Parse(templateDate, day.Date)
		if err != nil {
			return nil, fmt.Errorf("failed to parse forecast date: %v", err)
		}

		forecastToUpdate := &model.Forecast{
			CityID:       city.ID,
			ForecastDate: date,
			Temperature:  fmt.Sprintf("%.1f", day.Day.AvgtempC),
			Condition:    day.Day.Condition.Text,
		}

		err = self.forecastRepo.Create(forecastToUpdate)

		if err != nil {
			return nil, err
		}

		data := self.formForecastData(&day)
		forecastOutput.Forecastday = append(forecastOutput.Forecastday, *data)
	}

	self.deleteOutdatedForecasts()

	return &forecastOutput, nil
}

func (self *weatherAPIClient) handleCity(location *api.Location) (*model.City, error) {

	city := &model.City{
		Name:      location.Name,
		Country:   location.Country,
		Latitude:  fmt.Sprintf("%f", location.Lat),
		Longitude: fmt.Sprintf("%f", location.Lon),
	}

	return self.cityRepo.Register(city)
}

func (self *weatherAPIClient) deleteOutdatedForecasts() {
	err := self.forecastRepo.DeleteByPastDate()
	if err != nil {
		log.Printf("Error deleting outdated forecasts: %v", err)
	}
}

func (self *weatherAPIClient) formCurrentData(current *api.Current) *api.Current {

	outputData := &api.Current{
		LastUpdated: current.LastUpdated,
		TempC:       current.TempC,
		Condition:   current.Condition,
		Cloud:       current.Cloud,
		IsDay:       current.IsDay,
	}

	return outputData
}

func (self *weatherAPIClient) formForecastData(data *api.ForecastForecastday) *api.ForecastForecastday {

	day := &api.ForecastDay{
		Condition:         data.Day.Condition,
		AvgtempC:          data.Day.AvgtempC,
		MaxtempC:          data.Day.MaxtempC,
		MintempC:          data.Day.MintempC,
		DailyChanceOfRain: data.Day.DailyChanceOfRain,
	}

	var hour []api.ForecastHour

	for _, h := range data.Hour {
		x := &api.ForecastHour{
			Time:  h.Time,
			TempC: h.TempC,
			IsDay: h.IsDay,
		}
		hour = append(hour, *x)
	}

	outputData := &api.ForecastForecastday{
		Date: data.Date,
		Day:  day,
		Hour: hour,
	}

	return outputData
}

func (self *weatherAPIClient) HandleCityData(name string, cred string) *api.Location {
	key, err := resources.GetApiKey(cred)
	location, err := self.weatherDataGetter.GetLocation(name, key)
	if err != nil {
		log.Printf("Error getting current data: %v", err)
	}
	return location
}
