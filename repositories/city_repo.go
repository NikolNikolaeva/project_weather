package repositories

import (
	"errors"

	"gorm.io/gorm/clause"

	"gorm.io/gen"
	"gorm.io/gorm"

	"github.com/NikolNikolaeva/project_weather/generated/dao"
	"github.com/NikolNikolaeva/project_weather/generated/dao/model"
)

//go:generate mockgen --build_flags=--mod=mod -destination ../generated/go-mocks/repositories/mock_city_repo.go . CityRepo
type CityRepo interface {
	FindByID(id string) (*model.City, error)
	Register(city *model.City) (*model.City, error)
	UpdateByID(id string, city *model.City) (*model.City, error)
	DeleteByID(id string) (gen.ResultInfo, error)
	GetAll() ([]*model.City, error)
	FindByNameAndCountry(name string, country string) (*model.City, error)
}

type cityRepo struct {
	q *dao.Query
}

func NewRepo(query *dao.Query) CityRepo {
	return &cityRepo{q: query}
}

func (self *cityRepo) FindByID(id string) (*model.City, error) {
	return self.q.City.Where(
		self.q.City.ID.Eq(id),
	).First()
}

func (self *cityRepo) DeleteByID(id string) (gen.ResultInfo, error) {
	return self.q.City.Where(self.q.City.ID.Eq(id)).Delete()

}

func (self *cityRepo) GetAll() ([]*model.City, error) {
	return self.q.City.Find()
}

func (self *cityRepo) Register(city *model.City) (*model.City, error) {
	return city, self.q.City.Clauses(clause.OnConflict{
		OnConstraint: "unique_City",
		UpdateAll:    true,
	}).Create(city)
}

func (self *cityRepo) FindByNameAndCountry(name string, country string) (*model.City, error) {
	return self.q.City.Where(
		self.q.City.Name.Eq(name),
		self.q.City.Country.Eq(country),
	).First()
}

func (self *cityRepo) UpdateByID(id string, city *model.City) (*model.City, error) {

	_, err := self.q.City.Where(
		self.q.City.ID.Eq(city.ID)).
		Updates(&city)
	if errors.Is(err, gorm.ErrForeignKeyViolated) {
		return nil, gorm.ErrForeignKeyViolated
	}
	return city, err
}
