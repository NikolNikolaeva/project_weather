package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

const TableNameCity = "city"

type City struct {
	ID        string  `json:"id" gorm:"primary_key;type:varchar(36);not null"`
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func (*City) TableName() string {
	return TableNameCity
}

func MigrateCity(db *gorm.DB) error {
	err := db.AutoMigrate(&City{})
	return err
}

func (user *City) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	user.ID = uuid.NewString()
	return
}
