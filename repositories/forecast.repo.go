package repositories

import (
	"errors"
	"gorm.io/gorm"
	"project_weather/generated/dao"
	"project_weather/generated/dao/model"
)

type ForecastRepo struct {
	q *dao.Query
}

func NewForecastRepo(query *dao.Query) ForecastRepo {
	return ForecastRepo{
		q: query,
	}
}

func (self *ForecastRepo) FindByID(id string) (*model.Forecast, error) {
	forecast, err := self.q.Forecast.Where(
		self.q.Forecast.ID.Eq(id),
	).First()
	if err != nil {
		return nil, err
	}
	return forecast, nil
}

func (self *ForecastRepo) Create(forecast *model.Forecast) error {
	return self.q.Forecast.Create(forecast)
}

func (self *ForecastRepo) Update(id string, forecast *model.Forecast) error {
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

func (self *ForecastRepo) Delete(id string) error {
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

func (self *ForecastRepo) FindAll() ([]*model.Forecast, error) {
	forecasts, err := self.q.Forecast.Find()
	if err != nil {
		return nil, err
	}
	return forecasts, nil
}
