package testbed

import (
	"fmt"
	"github.com/NikolNikolaeva/project_weather/generated/dao/model"
	"github.com/NikolNikolaeva/project_weather/it/testbed/generated/client"
	"time"
)

func NewTestCity() *model.City {

	return &model.City{
		ID:        "1",
		Name:      "Sofia",
		Country:   "Bulgaria",
		Longitude: "0",
		Latitude:  "0",
	}
}

func AsCityDto(city *model.City) client.City {
	if city == nil {
		return client.City{}
	}

	return client.City{
		Id:        &city.ID,
		Name:      city.Name,
		Country:   city.Country,
		Longitude: &city.Longitude,
		Latitude:  &city.Latitude,
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
