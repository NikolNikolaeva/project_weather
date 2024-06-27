package testbed

import (
	"fmt"
	"time"

	api "github.com/NikolNikolaeva/project_weather/generated/api/weatherapi"

	"github.com/NikolNikolaeva/project_weather/generated/dao/model"
	"github.com/NikolNikolaeva/project_weather/it/testbed/generated/client"
)

func NewTestCity() *model.City {

	return &model.City{
		ID:        "1",
		Name:      "Sofia",
		Country:   "Bulgaria",
		Longitude: "",
		Latitude:  "",
	}
}

func AsCityDto(city *model.City) client.City {
	if city == nil {
		return client.City{}
	}

	return client.City{
		Id:      &city.ID,
		Name:    city.Name,
		Country: city.Country,
	}
}

func NewTestForecast() *model.Forecast {
	return &model.Forecast{
		ID:           "1",
		CityID:       "1",
		Temperature:  "20.0",
		Condition:    "Sunny",
		ForecastDate: time.Date(2006, 01, 02, 0, 0, 0, 0, time.UTC),
	}
}

func AsForecastDto(forecast *model.Forecast) client.Forecast {
	if forecast == nil {
		return client.Forecast{}
	}

	return client.Forecast{
		Id:           &forecast.ID,
		CityId:       forecast.CityID,
		Condition:    forecast.Condition,
		Temperature:  forecast.Temperature,
		ForecastDate: fmt.Sprintf("%v", forecast.ForecastDate),
	}
}

func NewTestCurrentWeather() *api.Current {
	return &api.Current{
		TempC: 15.0,
		Condition: &api.CurrentCondition{
			Text: "Sunny",
		},
	}
}

func NewTestForecastDayWeather() *api.Forecast {
	return &api.Forecast{
		Forecastday: []api.ForecastForecastday{
			{Date: "2006-01-01",
				Day: &api.ForecastDay{
					Condition: &api.ForecastDayCondition{
						Text: "Sunny",
					},
				}},
		},
	}
}
