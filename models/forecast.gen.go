package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

const TableNameForecast = "forecast"

type Forecast struct {
	gorm.Model

	ID           string    `gorm:"primaryKey" json:"id"`
	CityID       uint      `gorm:"not null" json:"city_id"`
	ForecastDate time.Time `gorm:"not null" json:"forecast_date"`
	Temperature  float64   `gorm:"type:decimal(5,2)" json:"temperature"`
	Condition    string    `gorm:"size:100" json:"condition"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (*Forecast) TableName() string {
	return TableNameForecast
}

func (self *Forecast) BeforeCreate(tx *gorm.DB) (err error) {
	self.ID = uuid.NewString()
	return
}
