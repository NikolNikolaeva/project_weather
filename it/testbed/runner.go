package testbed

import (
	"net/url"
	"testing"

	"github.com/NikolNikolaeva/project_weather/generated/dao"
	appcontext "github.com/NikolNikolaeva/project_weather/it/testbed/internal/app_context"

	"github.com/pkg/errors"

	"github.com/NikolNikolaeva/project_weather/generated/dao/model"
	"github.com/NikolNikolaeva/project_weather/it/testbed/generated/client"

	. "github.com/onsi/gomega"
)

type Runner struct {
	t *testing.T

	FakeWeatherApi FakeWeatherApi
	RootURL        *url.URL
	Dao            *dao.Query
	APIClient      *client.APIClient
	Context        appcontext.Context
}

func (self *Runner) Reset() *Runner {
	self.Dao.City.UnderlyingDB().Where("1=1").Delete(&model.City{})
	self.Dao.Forecast.UnderlyingDB().Where("1=1").Delete(&model.Forecast{})
	self.FakeWeatherApi.Reset()

	return self
}

func (self *Runner) Case(name string, test func(t *testing.T)) {
	self.t.Run(name, func(t *testing.T) {
		t.Cleanup(func() { self.Reset() })
		test(t)
	})
}

func (self *Runner) CreateCity(city *model.City) *model.City {
	err := self.Dao.City.WithContext(self.Context).Create(city)

	Expect(err).ShouldNot(HaveOccurred())

	return city
}

func (self *Runner) CreateForecast(forecast *model.Forecast) *model.Forecast {
	err := self.Dao.Forecast.WithContext(self.Context).Create(forecast)
	Expect(err).ShouldNot(HaveOccurred())

	Expect(
		self.Dao.Forecast.WithContext(self.Context).Create(forecast),
	).ShouldNot(HaveOccurred())

	return forecast
}

func (self *Runner) FailIf(err error) {
	if err != nil {
		self.Context.Cancel(errors.Wrap(err, "unexpected error received"))
	}
}
