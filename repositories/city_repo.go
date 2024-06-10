package repositories

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/NikolNikolaeva/project_weather/generated/dao"
	"github.com/NikolNikolaeva/project_weather/generated/dao/model"
)

//go:generate mockgen --build_flags=--mod=mod -destination ../generated/go-mocks/repositories/mock_city_repo.go . CityRepo
type CityRepo interface {
	FindCityByID(id string) (*model.City, error)
	RegisterCity(city *model.City) (*model.City, error)
	UpdateCityByID(id string, city *model.City) (*model.City, error)
	DeleteCityByID(id string) error
	GetAllCity() ([]*model.City, error)
}

type cityRepo struct {
	q *dao.Query
}

func NewCityRepo(query *dao.Query) CityRepo {
	return &cityRepo{q: query}
}

func (self *cityRepo) FindCityByID(id string) (*model.City, error) {
	city, err := self.q.City.Where(
		self.q.City.ID.Eq(id),
	).First()
	if err != nil {
		return nil, err
	}
	return city, nil
}

func (self *cityRepo) DeleteCityByID(id string) error {
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

func (self *cityRepo) GetAllCity() ([]*model.City, error) {
	cities, err := self.q.City.Find()
	if err != nil {
		return nil, err
	}
	return cities, nil
}

func (self *cityRepo) RegisterCity(city *model.City) (*model.City, error) {
	existCity, err := self.q.City.Where(
		self.q.City.Name.Eq(city.Name),
	).First()

	if err == nil && existCity.ID != "" {
		return existCity, err
	}
	fmt.Printf("city: %#v", city)
	fmt.Printf("existCity: %#v, err: %v", existCity, err)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = self.q.City.Create(city)
		fmt.Printf("created city: %#v, err: %v", city, err)
		if err != nil {
			return nil, err
		}
	}

	return city, nil

}

func (self *cityRepo) UpdateCityByID(id string, city *model.City) (*model.City, error) {
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
