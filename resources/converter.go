package resources

import (
	"log"
	"time"

	api "github.com/NikolNikolaeva/project_weather/generated/api/project-weather/rest"
	"github.com/NikolNikolaeva/project_weather/generated/dao/model"
)

//go:generate mockgen --build_flags=--mod=mod -destination ../generated/go-mocks/resources/mock_converter.go . ConverterI
type ConverterI interface {
	ConvertApiCityToModelCity(cityApi *api.City) *model.City
	ConvertModelCityToApiCity(cityModel *model.City) *api.City
	ConvertModelForecastToApiForecast(forecastModel *model.Forecast) *api.Forecast
	ConvertApiForecastToModelForecast(forecastModel *api.Forecast) *model.Forecast
}

type converter struct {
}

func NewConverter() ConverterI {
	return &converter{}
}

func (self *converter) ConvertApiCityToModelCity(cityApi *api.City) *model.City {

	return &model.City{
		ID:        cityApi.Id,
		Name:      cityApi.Name,
		Country:   cityApi.Country,
		Latitude:  cityApi.Latitude,
		Longitude: cityApi.Longitude,
	}
}

func (self *converter) ConvertModelCityToApiCity(cityModel *model.City) *api.City {

	return &api.City{
		Id:        cityModel.ID,
		Name:      cityModel.Name,
		Country:   cityModel.Country,
		Latitude:  cityModel.Latitude,
		Longitude: cityModel.Longitude,
	}
}

func (self *converter) ConvertModelForecastToApiForecast(forecastModel *model.Forecast) *api.Forecast {

	return &api.Forecast{
		Id:           forecastModel.ID,
		CityId:       forecastModel.CityID,
		ForecastDate: forecastModel.ForecastDate.Format(time.DateOnly),
		Temperature:  forecastModel.Temperature,
		Condition:    forecastModel.Condition,
	}
}

func (self *converter) ConvertApiForecastToModelForecast(forecastModel *api.Forecast) *model.Forecast {

	templateDate := "2006-01-02"
	day, err := time.Parse(templateDate, forecastModel.ForecastDate)
	if err != nil {
		log.Println("Error parsing forecast date, err: ", err)
		return nil
	}
	return &model.Forecast{
		ID:           forecastModel.Id,
		CityID:       forecastModel.CityId,
		ForecastDate: day,
		Temperature:  forecastModel.Temperature,
		Condition:    forecastModel.Condition,
	}
}
