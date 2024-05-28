package swagger

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

//------------------------------------------
/*
package swagger

type Current struct {
	LastUpdatedEpoch int32                 `json:"last_updated_epoch,omitempty"`
	LastUpdated      string                `json:"last_updated,omitempty"`
	TempC            float32               `json:"temp_c,omitempty"`
	TempF            float32               `json:"temp_f,omitempty"`
	IsDay            int32                 `json:"is_day,omitempty"`
	Condition        *ForecastDayCondition `json:"condition,omitempty"`
}

type Forecast struct {
	Forecastday []ForecastForecastday `json:"forecastday,omitempty"`
}

type ForecastForecastday struct {
	Date      string         `json:"date,omitempty"`
	DateEpoch int32          `json:"date_epoch,omitempty"`
	Day       *ForecastDay   `json:"day,omitempty"`
	Hour      []ForecastHour `json:"hour,omitempty"`
}

type ForecastDay struct {
	MaxtempC   float32               `json:"maxtemp_c,omitempty"`
	MaxtempF   float32               `json:"maxtemp_f,omitempty"`
	MintempC   float32               `json:"mintemp_c,omitempty"`
	MintempF   float32               `json:"mintemp_f,omitempty"`
	AvgtempC   float32               `json:"avgtemp_c,omitempty"`
	AvgtempF   float32               `json:"avgtemp_f,omitempty"`
	MaxwindMph float32               `json:"maxwind_mph,omitempty"`
	MaxwindKph float32               `json:"maxwind_kph,omitempty"`
	Condition  *ForecastDayCondition `json:"condition,omitempty"`
}

type ForecastDayCondition struct {
	Text string `json:"text,omitempty"`
	Icon string `json:"icon,omitempty"`
	Code int32  `json:"code,omitempty"`
}

type ForecastHour struct {
	TimeEpoch    int32                 `json:"time_epoch,omitempty"`
	Time         string                `json:"time,omitempty"`
	TempC        float32               `json:"temp_c,omitempty"`
	TempF        float32               `json:"temp_f,omitempty"`
	IsDay        int32                 `json:"is_day,omitempty"`
	Condition    *ForecastDayCondition `json:"condition,omitempty"`
	WindMph      float32               `json:"wind_mph,omitempty"`
	WindKph      float32               `json:"wind_kph,omitempty"`
	WindDegree   float32               `json:"wind_degree,omitempty"`
	WindDir      string                `json:"wind_dir,omitempty"`
	PressureMb   float32               `json:"pressure_mb,omitempty"`
	PressureIn   float32               `json:"pressure_in,omitempty"`
	PrecipMm     float32               `json:"precip_mm,omitempty"`
	PrecipIn     float32               `json:"precip_in,omitempty"`
	Humidity     float32               `json:"humidity,omitempty"`
	Cloud        float32               `json:"cloud,omitempty"`
	FeelslikeC   float32               `json:"feelslike_c,omitempty"`
	FeelslikeF   float32               `json:"feelslike_f,omitempty"`
	WindchillC   float32               `json:"windchill_c,omitempty"`
	WindchillF   float32               `json:"windchill_f,omitempty"`
	HeatindexC   float32               `json:"heatindex_c,omitempty"`
	HeatindexF   float32               `json:"heatindex_f,omitempty"`
	DewpointC    float32               `json:"dewpoint_c,omitempty"`
	DewpointF    float32               `json:"dewpoint_f,omitempty"`
	WillItRain   int32                 `json:"will_it_rain,omitempty"`
	ChanceOfRain float32               `json:"chance_of_rain,omitempty"`
	WillItSnow   int32                 `json:"will_it_snow,omitempty"`
	ChanceOfSnow float32               `json:"chance_of_snow,omitempty"`
	VisKm        float32               `json:"vis_km,omitempty"`
	VisMiles     float32               `json:"vis_miles,omitempty"`
	GustMph      float32               `json:"gust_mph,omitempty"`
	GustKph      float32               `json:"gust_kph,omitempty"`
	Uv           int32                 `json:"uv,omitempty"`
}

*/
