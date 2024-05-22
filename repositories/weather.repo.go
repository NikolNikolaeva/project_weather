package repositories

import "net/http"

type WeatherRepository struct {
	c *http.Client
}

func NewWeatherRepository() WeatherRepository {

}
