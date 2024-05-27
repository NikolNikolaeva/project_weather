package dto

import (
	"time"
)

// CityDTO represents the data transfer object for the City model
type CityDTO struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Country   string    `json:"country"`
	Latitude  string    `json:"latitude"`
	Longitude string    `json:"longitude"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// WeatherDTO represents the top-level weather data response from the API
type WeatherDTO struct {
	Location LocationDTO  `json:"location"`
	Current  CurrentDTO   `json:"current"`
	Forecast ForecastsDTO `json:"forecast"`
}

// LocationDTO represents the location data
type LocationDTO struct {
	Name    string  `json:"name"`
	Region  string  `json:"region"`
	Country string  `json:"country"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
}

// CurrentDTO represents the current weather data
type CurrentDTO struct {
	Location    LocationDTO  `json:"location"`
	LastUpdated string       `json:"last_updated"`
	TempC       float64      `json:"temp_c"`
	Condition   ConditionDTO `json:"condition"`
}

// ForecastsDTO holds the forecast data
type ForecastsDTO struct {
	ForecastDays []ForecastDayDTO `json:"forecastday"`
}

// ForecastDayDTO represents a single day's forecast data
type ForecastDayDTO struct {
	Date string `json:"date"`
	Day  DayDTO `json:"day"`
}

// DayDTO represents the detailed day information
type DayDTO struct {
	AvgTempC  float64      `json:"avgtemp_c"`
	Condition ConditionDTO `json:"condition"`
}

// ConditionDTO represents the weather condition
type ConditionDTO struct {
	Text string `json:"text"`
}

// ForecastDTO represents the data transfer object for the Forecast model
type ForecastDTO struct {
	ID           string    `json:"id"`
	CityID       string    `json:"city_id"`
	ForecastDate time.Time `json:"forecast_date"`
	Temperature  string    `json:"temperature"`
	Condition    string    `json:"condition"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
