package bootstrap

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/fx"

	"github.com/NikolNikolaeva/project_weather/config"
	api "github.com/NikolNikolaeva/project_weather/generated/api/project-weather/rest"
	"github.com/NikolNikolaeva/project_weather/generated/dao"
	"github.com/NikolNikolaeva/project_weather/repositories"
	"github.com/NikolNikolaeva/project_weather/resources"
	"github.com/NikolNikolaeva/project_weather/services"
	"github.com/NikolNikolaeva/project_weather/utils"
)

var FXModule_HTTPServer = fx.Module(
	"http-server",
	fx.Provide(
		createAPIRoutes,
		createCityRepo,
		createCityController,
		createForecastRepo,
		createForecastController,
		createHTTPClient,
		createWeatherHandler,
		createWeatherDataRetriever,
		createWeatherController,
		createHttpServer,
		createConverter,
		createMuxRouter,
	),
	fx.Invoke(
		configureAPIRoutes,
		registerServerStartHook,
	),
)

func createAPIRoutes(cities *api.CityAPIController, forecasts *api.ForecastAPIController, weather *api.WeatherAPIController) api.Routes {
	return utils.Merge(
		cities.Routes(),
		weather.Routes(),
		forecasts.Routes(),
	)
}

func createWeatherDataRetriever(client *http.Client) services.WeatherDataGetter {
	return services.NewWeatherDataRetriever(client)
}

func createWeatherHandler(cityRepo repositories.CityRepo, forecastRepo repositories.ForecastRepo, getter services.WeatherDataGetter) services.WeatherHandler {
	return services.NewWeatherHandler(cityRepo, forecastRepo, getter)
}

func createHTTPClient() *http.Client {
	return services.NewHTTPClient()
}

func createCityRepo(q *dao.Query) repositories.CityRepo {
	return repositories.NewCityRepo(q)
}

func createForecastRepo(q *dao.Query) repositories.ForecastRepo {
	return repositories.NewForecastRepo(q)
}

func createCityController(db repositories.CityRepo, convert resources.ConverterI) *api.CityAPIController {
	return api.NewCityAPIController(services.NewCityAPIService(db, convert))
}

func createForecastController(db repositories.ForecastRepo, convert resources.ConverterI) *api.ForecastAPIController {
	return api.NewForecastAPIController(services.NewForecastAPIService(db, convert))
}

func createWeatherController(handler services.WeatherHandler, config *config.ApplicationConfiguration) *api.WeatherAPIController {
	return api.NewWeatherAPIController(services.NewWeatherAPIService(handler, config))
}

func createMuxRouter() *mux.Router {
	return mux.NewRouter()

}

func configureAPIRoutes(app *mux.Router, routes api.Routes) {
	for _, route := range routes {
		log.Printf("Registering route: %s %s", route.Method, route.Pattern)
		app.HandleFunc(route.Pattern, route.HandlerFunc)
	}
	http.Handle("/", app)
}

func createHttpServer(routes api.Routes, config *config.ApplicationConfiguration, router *mux.Router) *http.Server {
	return &http.Server{
		Addr:    ":" + config.HTTPPort,
		Handler: router,
	}
}

func createConverter() resources.ConverterI {
	return resources.NewConverter()
}

func registerServerStartHook(lc fx.Lifecycle, server *http.Server) {
	lc.Append(fx.StartStopHook(
		func() {
			go func() {
				if err := server.ListenAndServe(); err != nil {
					log.Fatalf("failed to start server: %v", err)
				}
			}()
		},
		func() {
			if err := server.Shutdown(context.Background()); err != nil {
				log.Fatalf("failed to shutdown server: %v", err)
			}
		},
	))
}
