package testbed

import (
	"fmt"
	"net/url"
	"testing"
	"time"

	"go.uber.org/fx"

	. "github.com/onsi/gomega"

	"github.com/NikolNikolaeva/project_weather/generated/dao"
	"github.com/NikolNikolaeva/project_weather/it/testbed/generated/client"
	internal "github.com/NikolNikolaeva/project_weather/it/testbed/internal"
	appcontext "github.com/NikolNikolaeva/project_weather/it/testbed/internal/app_context"
	"github.com/NikolNikolaeva/project_weather/it/testbed/internal/gomisc/lang"
)

var instance *testbed

func Bootstrap(m *testing.M) int {
	return internal.RunWithEnvironment(
		"../../resources/environments/integration-test.env",
		func() int {
			defer func() { instance.destroy() }()

			ctx := appcontext.New()
			FakeWeatherApi, FakeWeatherApicancel := NewFakeWeatherApi(
				lang.GetEnv("FAKE_CFS_URL", "http://localhost:8081"),
			)
			application, applicationCancel := internal.NewFXApplication(ctx) // the FX app has an implicit dependency on the Fake CFS!!
			database, cancel := internal.NewGormDatabase()                   // the database has an implicit dependency on the FX app!!

			instance = &testbed{
				Context:        ctx,
				FakeWeatherApi: FakeWeatherApi,
				App:            application,
				Dao:            dao.Use(database),
				RootURL: lang.First(url.Parse(
					lang.GetEnv("APPLICATION_ROOT_URL", "http://localhost:8080"),
				)),
				Destroyers: []func(){
					cancel,
					applicationCancel,
					FakeWeatherApicancel,
					func() { ctx.Cancel(nil) },
				},
			}

			return m.Run()
		},
	)
}

type testbed struct {
	App            *fx.App
	FakeWeatherApi FakeWeatherApi
	Destroyers     []func()
	RootURL        *url.URL
	Dao            *dao.Query
	Context        appcontext.Context
}

func (self *testbed) destroy() {
	for _, destroyer := range self.Destroyers {
		destroyer()
	}
}

func (self *testbed) runner(t *testing.T, timeout time.Duration) *Runner {
	return &Runner{
		t: t,

		Dao:            self.Dao,
		FakeWeatherApi: self.FakeWeatherApi,
		RootURL:        self.RootURL,
		Context:        self.Context.Child(timeout),
		APIClient: client.NewAPIClient(&client.Configuration{
			Servers: client.ServerConfigurations{{URL: fmt.Sprintf(
				"%s/api",
				lang.GetEnv("APPLICATION_ROOT_URL", "http://localhost:8080"),
			)}},
		}),
	}
}

func (self *testbed) timeboxed(t *testing.T, timeout time.Duration, test func(t *testing.T, tb *Runner)) {
	timer := time.AfterFunc(timeout, func() {
		self.destroy()
		panic(fmt.Sprintf("%v failed to finish in %v", t.Name(), timeout))
	})
	defer timer.Stop()

	test(t, instance.runner(t, timeout).Reset())
}

func Prepare(t *testing.T, timeout time.Duration, test func(t *testing.T, tb *Runner)) {
	RegisterTestingT(t)

	instance.timeboxed(t, timeout, test)
}
