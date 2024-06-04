package services

import (
	"context"
	"fmt"
	"net/http"

	swagger "github.com/weatherapicom/go"
)

type WeatherDataGetter interface {
	GetCurrentData(q string, key string) (*swagger.Current, *swagger.Location, error)
	GetForecastData(q string, days int32, key string) (*swagger.Forecast, *swagger.Location, error)
}

type weatherDataGetter struct {
	client *http.Client
}

func NewWeatherDataGetter(client *http.Client) WeatherDataGetter {
	return &weatherDataGetter{
		client: client,
	}
}

func (self *weatherDataGetter) GetCurrentData(q string, key string) (*swagger.Current, *swagger.Location, error) {

	config := swagger.NewConfiguration()
	api := swagger.NewAPIClient(config)

	ctx := context.WithValue(context.Background(), swagger.ContextAPIKey, swagger.APIKey{Key: key})

	weather, resp, err := api.APIsApi.RealtimeWeather(ctx, q, nil)

	if err != nil {
		return nil, nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, nil, fmt.Errorf("Error calling realtime weather: %s", resp.Status)
	}

	return weather.Current, weather.Location, nil
}

func (self *weatherDataGetter) GetForecastData(q string, days int32, key string) (*swagger.Forecast, *swagger.Location, error) {

	config := swagger.NewConfiguration()
	api := swagger.NewAPIClient(config)

	ctx := context.WithValue(context.Background(), swagger.ContextAPIKey, swagger.APIKey{Key: key})

	weather, resp, err := api.APIsApi.ForecastWeather(ctx, q, days, nil)

	if err != nil {
		return nil, nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, nil, fmt.Errorf("Error calling realtime weather: %s", resp.Status)
	}

	return weather.Forecast, weather.Location, nil
}
