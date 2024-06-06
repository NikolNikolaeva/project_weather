package repositories

import (
	"errors"

	"gorm.io/gorm"

	"github.com/NikolNikolaeva/project_weather/generated/dao"
	"github.com/NikolNikolaeva/project_weather/generated/dao/model"
)

//go:generate mockgen --build_flags=--mod=mod -destination ../generated/go-mocks/repositories/mock_forecast_repo.go . ForecastRepo
type ForecastRepo interface {
	FindByID(id string) (*model.Forecast, error)
	Create(forecast *model.Forecast) error
	Update(id string, forecast *model.Forecast) error
	Delete(id string) error
	FindAll() ([]*model.Forecast, error)
}

type forecastRepo struct {
	q *dao.Query
}

func NewForecastRepo(query *dao.Query) ForecastRepo {
	return &forecastRepo{
		q: query,
	}
}

func (self *forecastRepo) FindByID(id string) (*model.Forecast, error) {
	forecast, err := self.q.Forecast.Where(
		self.q.Forecast.ID.Eq(id),
	).First()
	if err != nil {
		return nil, err
	}
	return forecast, nil
}

func (self *forecastRepo) FindByCityId(cityId string) ([]*model.Forecast, error) {
	forecast, err := self.q.Forecast.Where(
		self.q.Forecast.CityID.Eq(cityId),
	).Find()
	if err != nil {
		return nil, err
	}
	return forecast, nil
}

func (self *forecastRepo) Create(forecast *model.Forecast) error {

	err := self.q.Forecast.Create(forecast)
	return err
}

func (self *forecastRepo) Update(id string, forecast *model.Forecast) error {
	_, err := self.q.Forecast.Where(
		self.q.Forecast.ID.Eq(id),
	).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("record not found")
		}
		return err
	}

	if _, err := self.q.Forecast.Where(self.q.City.ID.Eq(id)).Updates(forecast); err != nil {
		return err
	}

	return nil
}

func (self *forecastRepo) Delete(id string) error {
	forecast, err := self.q.Forecast.Where(
		self.q.Forecast.ID.Eq(id),
	).First()
	if err != nil {
		return err
	}
	_, err = self.q.Forecast.Delete(forecast)
	if err != nil {
		return err
	}
	return nil
}

func (self *forecastRepo) FindAll() ([]*model.Forecast, error) {
	forecasts, err := self.q.Forecast.Find()
	if err != nil {
		return nil, err
	}
	return forecasts, nil
}
