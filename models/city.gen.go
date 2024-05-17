package models

import "gorm.io/gorm"

const TableNameCity = "city"

type City struct {
	ID        string  `json:"id" gorm:"primary_key;auto_increment"`
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func (self *City) TableName() string {
	return TableNameCity
}

func MigrateCity(db *gorm.DB) error {
	err := db.AutoMigrate(&City{})
	return err
}
