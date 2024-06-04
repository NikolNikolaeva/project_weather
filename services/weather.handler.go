package services

import (
	"encoding/json"
	"fmt"
	"github.com/NikolNikolaeva/project_weather/generated/dao/model"
	"github.com/NikolNikolaeva/project_weather/repositories"
	swagger "github.com/weatherapicom/go"
	"io/ioutil"
	"os"
	"time"
)

const (
	templateDate        = "2006-01-02"
	templateDateAndTime = "2006-01-02 15:04 "
)

type Cred struct {
	ApiKey string `json:"apiKey"`
}

type WeatherHandler interface {
	HandleCurrantData(q string, cred string) (*swagger.Current, error)
	HandleForecast(q string, days int32, cred string) (*swagger.Forecast, error)
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

func (self *weatherHandler) HandleCurrantData(q string, cred string) (*swagger.Current, error) {
	key, err := self.getApiKey(cred)
	if err != nil {
		return nil, err
	}
	currentData, location, err := self.weatherDataGetter.GetCurrentData(q, key)
	if err != nil {
		return nil, err
	}

	city := &model.City{
		Name:      location.Name,
		Country:   location.Country,
		Latitude:  fmt.Sprintf("%f", location.Lat),
		Longitude: fmt.Sprintf("%f", location.Lon),
	}

	_, err = self.cityRepo.RegisterCity(city)

	if err != nil {
		return nil, err
	}

	output := self.formCurrentData(currentData)
	return output, nil
}

func (self *weatherHandler) HandleForecast(q string, days int32, cred string) (*swagger.Forecast, error) {
	key, err := self.getApiKey(cred)
	if err != nil {
		return nil, err
	}
	forecast, location, err := self.weatherDataGetter.GetForecastData(q, days, key)
	if err != nil {
		return nil, err
	}

	city := &model.City{
		Name:      location.Name,
		Country:   location.Country,
		Latitude:  fmt.Sprintf("%f", location.Lat),
		Longitude: fmt.Sprintf("%f", location.Lon),
	}

	city, err = self.cityRepo.RegisterCity(city)

	if err != nil {
		return nil, err
	}

	for _, day := range forecast.Forecastday {
		date, err := time.Parse(templateDate, day.Date)

		if err != nil {
			return nil, fmt.Errorf("failed to parse forecast date: %v", err)
		}

		forecastToSave := &model.Forecast{
			CityID:       city.ID,
			ForecastDate: date,
			Temperature:  fmt.Sprintf("%.1f", day.Day.AvgtempC),
			Condition:    day.Day.Condition.Text,
		}
		err = self.forecastRepo.Create(forecastToSave)
		if err != nil {
			return nil, err
		}

	}

	return forecast, nil
}

func (self *weatherHandler) formCurrentData(current *swagger.Current) *swagger.Current {

	var outputData *swagger.Current

	outputData.TempC = current.TempC
	outputData.Condition = current.Condition
	outputData.LastUpdated = current.LastUpdated
	outputData.Cloud = current.Cloud
	outputData.IsDay = current.IsDay

	return outputData
}

//func (self *weatherHandler) formForecastData(current *swagger.Forecast) *swagger.Current {
//
//	var outputData *swagger.Fo
//
//	outputData.TempC = current.TempC
//	outputData.Condition = current.Condition
//	outputData.LastUpdated = current.LastUpdated
//	outputData.Cloud = current.Cloud
//	outputData.IsDay = current.IsDay
//
//	return outputData
//}
