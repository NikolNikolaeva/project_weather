package services

import (
	"context"
	"fmt"
	"net/http"

	api "github.com/NikolNikolaeva/project_weather/generated/api/weatherapi"
)

//go:generate mockgen --build_flags=--mod=mod -destination ../generated/go-mocks/services/mock_weather_data_getter.go . WeatherDataGetter
type WeatherDataGetter interface {
	GetCurrentData(q string, key string) (*api.Current, *api.Location, error)
	GetForecastData(q string, days int32, key string) (*api.Forecast, *api.Location, error)
}

type weatherDataGetter struct {
	client *http.Client
}

func NewWeatherDataRetriever(client *http.Client) WeatherDataGetter {
	return &weatherDataGetter{
		client: client,
	}
}

func (self *weatherDataGetter) GetCurrentData(q string, key string) (*api.Current, *api.Location, error) {

	config := api.NewConfiguration()
	client := api.NewAPIClient(config)

	ctx := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{Key: key})

	weather, resp, err := client.APIsApi.RealtimeWeather(ctx, q, nil)

	if err != nil {
		return nil, nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, nil, fmt.Errorf("Error calling realtime weather: %s", resp.Status)
	}

	return weather.Current, weather.Location, nil

}

func (self *weatherDataGetter) GetForecastData(q string, days int32, key string) (*api.Forecast, *api.Location, error) {

	config := api.NewConfiguration()
	client := api.NewAPIClient(config)

	ctx := context.WithValue(context.Background(), api.ContextAPIKey, api.APIKey{Key: key})

	weather, resp, err := client.APIsApi.ForecastWeather(ctx, q, days, nil)

	if err != nil {
		return nil, nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, nil, fmt.Errorf("Error calling realtime weather: %s", resp.Status)
	}

	return weather.Forecast, weather.Location, nil
}
