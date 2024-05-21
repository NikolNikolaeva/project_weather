package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

const TableNameCity = "city"

type City struct {
	gorm.Model

	ID        string  `json:"id" gorm:"primaryKey"`
	Name      string  `json:"name" gorm:"size:100;not null"`
	Country   string  `json:"country" gorm:"size:100;not null"`
	Latitude  float64 `json:"latitude" gorm:"type:decimal(9,6);not null"`
	Longitude float64 `json:"longitude" gorm:"type:decimal(9,6);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (*City) TableName() string {
	return TableNameCity
}

func (self *City) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	self.ID = uuid.NewString()
	return
}
