package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/NikolNikolaeva/project_weather/resources/swagger"
)

type WeatherDataGetter interface {
	GetData(url string) (*swagger.WeatherDTO, error)
}

type weatherDataGetter struct {
	client *http.Client
}

func NewWeatherDataGetter(client *http.Client) WeatherDataGetter {
	return &weatherDataGetter{
		client: client,
	}
}

func (self *weatherDataGetter) GetData(url string) (*swagger.WeatherDTO, error) {
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