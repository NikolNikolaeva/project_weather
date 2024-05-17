package models

import "gorm.io/gorm"

const TableNameForecast = "forecast"

type Forecast struct {
	ID     string `json:"id" gorm:"primary_key;auto_increment"`
	CityID uint   `json:"city_id"`
	//period *string `json:"period"`
}

func (*Forecast) TableName() string {
	return TableNameForecast
}

func MigrateForecast(db *gorm.DB) error {
	err := db.AutoMigrate(&Forecast{})
	return err
}
