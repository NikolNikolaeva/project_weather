package services

import (
	"fmt"
	"project_weather/generated/dao/model"
	"time"
)

type ForecastService struct{}

func NewForecastService() *ForecastService {
	return &ForecastService{}
}

func (s *ForecastService) FormatDailyForecasts(cityID string, dailyData []map[string]interface{}) []*model.Forecast {
	var forecasts []*model.Forecast

	for _, day := range dailyData {
		forecastDateStr := day["date"].(string)
		forecastDate, _ := time.Parse("2006-01-02", forecastDateStr)
		temperature := day["day"].(map[string]interface{})["avgtemp_c"].(float64)
		condition := day["day"].(map[string]interface{})["condition"].(map[string]interface{})["text"].(string)

		forecast := &model.Forecast{
			CityID:       cityID,
			ForecastDate: forecastDate,
			Temperature:  fmt.Sprintf("%.1f", temperature),
			Condition:    condition,
		}

		forecasts = append(forecasts, forecast)
	}

	return forecasts
}
