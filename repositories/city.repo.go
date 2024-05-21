package repositories

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"project_weather/models"
)

type CityRepo struct {
	Db *gorm.DB
}

func (self *CityRepo) FindCityByID(id string) (*models.City, error) {
	var city models.City
	result := self.Db.Where("id = ?", id).First(&city)
	if result.Error != nil {
		return nil, result.Error
	}
	return &city, nil
}

func (self *CityRepo) DeleteCityByID(id string) error {
	city := models.City{}
	self.Db.Find(&city, "id = ?", id)

	if city.ID == "" {
		fmt.Println(city.ID)
		return errors.New("Record not found")
	}
	return self.Db.Where("id = ?", id).Delete(&city).Error
}

func (self *CityRepo) GetAllCity() (*[]models.City, error) {
	var cities []models.City
	result := self.Db.Find(&cities)
	if result.Error != nil {
		return nil, result.Error
	}
	return &cities, nil
}

func (self *CityRepo) RegisterCity(city *models.City) (*models.City, error) {
	cityExist := models.City{}
	self.Db.Find(&cityExist, "name = ?", city.Name)

	if cityExist.ID != "" {
		return &cityExist, nil
	}

	result := self.Db.Create(&city)
	if result.Error != nil {
		return nil, result.Error
	}
	return city, nil
}

func (self *CityRepo) UpdateCityByID(id string, city *models.City) (*models.City, error) {
	existingCity := models.City{}
	if err := self.Db.First(&existingCity, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("record not found")
		}
		return nil, err
	}

	// Update the city fields
	if err := self.Db.Model(&existingCity).Updates(city).Error; err != nil {
		return nil, err
	}

	return &existingCity, nil
}
