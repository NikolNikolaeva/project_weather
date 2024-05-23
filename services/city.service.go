package services

import (
	"fmt"
	"project_weather/generated/dao/model"
)

type CityService struct{}

func NewCityService() *CityService {
	return &CityService{}
}

func (s *CityService) FormatWeatherDataForCity(weatherData map[string]interface{}) *model.City {
	name := weatherData["city"].(map[string]interface{})["name"].(string)
	country := weatherData["city"].(map[string]interface{})["country"].(string)
	latitude := weatherData["city"].(map[string]interface{})["coord"].(map[string]interface{})["lat"].(float64)
	longitude := weatherData["city"].(map[string]interface{})["coord"].(map[string]interface{})["lon"].(float64)

	city := &model.City{
		Name:      name,
		Country:   country,
		Latitude:  fmt.Sprintf("%.2f", latitude),
		Longitude: fmt.Sprintf("%.2f", longitude),
	}

	return city
}
