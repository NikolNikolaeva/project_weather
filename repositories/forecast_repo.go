package repositories

import (
	"errors"
	"time"

	"gorm.io/gorm/clause"

	"gorm.io/gorm"

	"github.com/NikolNikolaeva/project_weather/generated/dao"
	"github.com/NikolNikolaeva/project_weather/generated/dao/model"
)

//go:generate mockgen --build_flags=--mod=mod -destination ../generated/go-mocks/repositories/mock_forecast_repo.go . ForecastRepo
type ForecastRepo interface {
	FindByID(id string) (*model.Forecast, error)
	Create(forecast *model.Forecast) error
	Update(cityId string, forecast *model.Forecast) error
	Delete(id string) error
	FindAll() ([]*model.Forecast, error)
	FindByCityId(cityId string) ([]*model.Forecast, error)
	FindByCityIdAndPeriodDays(cityId string, days int) ([]*model.Forecast, error)
	DeleteByPastDate() error
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
	return self.q.Forecast.Where(
		self.q.Forecast.ID.Eq(id),
	).First()
}

func (self *forecastRepo) FindByCityId(cityId string) ([]*model.Forecast, error) {
	return self.q.Forecast.Where(
		self.q.Forecast.CityID.Eq(cityId),
	).Find()
}

func (self *forecastRepo) Create(forecast *model.Forecast) error {
	return self.q.Forecast.Clauses(clause.OnConflict{
		OnConstraint: "unique_Forecast",
		UpdateAll:    true,
	}).Create(forecast)
}

func (self *forecastRepo) Update(cityId string, forecast *model.Forecast) error {

	_, err := self.q.Forecast.Where(
		self.q.Forecast.CityID.Eq(cityId),
		self.q.Forecast.ForecastDate.Eq(forecast.ForecastDate)).
		Updates(forecast) //?

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return gorm.ErrRecordNotFound
	}
	if errors.Is(err, gorm.ErrForeignKeyViolated) {
		return gorm.ErrForeignKeyViolated
	}
	return err
}

func (self *forecastRepo) Delete(id string) error {

	_, err := self.q.Forecast.Where(
		self.q.Forecast.ID.Eq(id),
	).Delete()

	return err
}

func (self *forecastRepo) FindAll() ([]*model.Forecast, error) {
	return self.q.Forecast.Find()
}

func (self *forecastRepo) FindByCityIdAndPeriodDays(cityId string, days int) ([]*model.Forecast, error) {
	return self.q.Forecast.Where(
		self.q.Forecast.CityID.Eq(cityId),
		self.q.Forecast.ForecastDate.Between(time.Now(), time.Now().AddDate(0, 0, days-1)),
	).Find()

}

func (self *forecastRepo) DeleteByPastDate() error {
	_, err := self.q.Forecast.Where(
		self.q.Forecast.ForecastDate.Lt(time.Now()),
	).Delete()
	return err
}
