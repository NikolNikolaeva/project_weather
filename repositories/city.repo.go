package repositories

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"project_weather/generated/dao"
	"project_weather/generated/dao/model"
)

type CityRepo struct {
	Db *gorm.DB
	q  *dao.Query
}

func NewCityRepo(db *gorm.DB, query *dao.Query) CityRepo {
	return CityRepo{Db: db, q: query}
}

func (self *CityRepo) FindCityByID(id string) (*model.City, error) {
	var city model.City
	result := self.Db.Where("id = ?", id).First(&city)
	if result.Error != nil {
		return nil, result.Error
	}
	return &city, nil
}

func (self *CityRepo) DeleteCityByID(id string) error {
	city := model.City{}
	self.Db.Find(&city, "id = ?", id)

	if city.ID == "" {
		fmt.Println(city.ID)
		return errors.New("Record not found")
	}
	return self.Db.Where("id = ?", id).Delete(&city).Error
}

func (self *CityRepo) GetAllCity() ([]*model.City, error) {
	//var cities []model.City
	cities, err := self.q.City.Find()
	if err != nil {
		return nil, err
	}
	//cities := self.Db.Model(model.City{})
	//result := self.Db.Find(&cities)
	//if result.Error != nil {
	//	return nil, result.Error
	//}
	return cities, nil
}

func (self *CityRepo) RegisterCity(city *model.City) (*model.City, error) {
	cityExist := model.City{}
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

func (self *CityRepo) UpdateCityByID(id string, city *model.City) (*model.City, error) {
	existingCity := model.City{}
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
