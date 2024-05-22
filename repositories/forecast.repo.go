package repositories

import (
	"gorm.io/gorm"
	"project_weather/generated/dao/model"
)

type ForecastRepo struct {
	Db *gorm.DB
}

func NewForecastRepo(db *gorm.DB) ForecastRepo {
	return ForecastRepo{
		Db: db,
	}
}

func (r *ForecastRepo) FindByID(id string) (*model.Forecast, error) {
	var forecast model.Forecast
	if err := r.Db.First(&forecast, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &forecast, nil
}

func (r *ForecastRepo) Create(forecast *model.Forecast) error {
	return r.Db.Create(forecast).Error
}

func (r *ForecastRepo) Update(id string, forecast *model.Forecast) error {
	return r.Db.Model(&model.Forecast{}).Where("id = ?", id).Updates(forecast).Error
}

func (r *ForecastRepo) Delete(id string) error {
	return r.Db.Delete(&model.Forecast{}, "id = ?", id).Error
}

func (r *ForecastRepo) FindAll() ([]model.Forecast, error) {
	var forecasts []model.Forecast
	if err := r.Db.Find(&forecasts).Error; err != nil {
		return nil, err
	}
	return forecasts, nil
}
