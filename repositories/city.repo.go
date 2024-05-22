package repositories

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"project_weather/generated/dao"
	"project_weather/generated/dao/model"
)

type CityRepo struct {
	q *dao.Query
}

func NewCityRepo(query *dao.Query) CityRepo {
	return CityRepo{q: query}
}

func (self *CityRepo) FindCityByID(id string) (*model.City, error) {
	city, err := self.q.City.Where(
		self.q.City.ID.Eq(id),
	).First()
	if err != nil {
		return nil, err
	}
	return city, nil
}

func (self *CityRepo) DeleteCityByID(id string) error {
	city, err := self.q.City.Where(
		self.q.City.ID.Eq(id),
	).First()
	if err != nil {
		return err
	}
	_, err = self.q.City.Delete(city)
	if err != nil {
		return err
	}
	return nil
}

func (self *CityRepo) GetAllCity() ([]*model.City, error) {
	cities, err := self.q.City.Find()
	if err != nil {
		return nil, err
	}
	return cities, nil
}

func (self *CityRepo) RegisterCity(city *model.City) (*model.City, error) {

	existCity, err := self.q.City.Where(
		self.q.City.Name.Eq(city.Name),
	).First()

	if err == nil && existCity.ID != "" {
		fmt.Println(existCity.ID)
		return existCity, err
	}

	if err == gorm.ErrRecordNotFound {

		err = self.q.City.Create(city)

		if err != nil {
			return nil, err
		}
	}
	return city, nil

}

func (self *CityRepo) UpdateCityByID(id string, city *model.City) (*model.City, error) {
	cityExist, err := self.q.City.Where(
		self.q.City.ID.Eq(id),
	).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("record not found")
		}
		return nil, err
	}

	if _, err := self.q.City.Where(self.q.City.ID.Eq(id)).Updates(city); err != nil {
		return nil, err
	}

	return cityExist, nil
}
