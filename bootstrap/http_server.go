package bootstrap

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/NikolNikolaeva/project_weather/config"
	api "github.com/NikolNikolaeva/project_weather/generated/api/project-weather/rest"
	"github.com/NikolNikolaeva/project_weather/utils"

	"github.com/gorilla/mux"
	"go.uber.org/fx"

	"github.com/NikolNikolaeva/project_weather/generated/dao"
	"github.com/NikolNikolaeva/project_weather/repositories"
	"github.com/NikolNikolaeva/project_weather/resources"
	"github.com/NikolNikolaeva/project_weather/services"
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
		createWeatherService,
		createHttpServer,
		createConverter,
		createMuxRouter,
	),
	fx.Invoke(
		registerServerStartHook,
		fetchWeatherData,
	),
)

func createAPIRoutes(cities *api.CityAPIController, forecasts *api.ForecastAPIController) api.Routes {
	return utils.Merge(
		cities.Routes(),
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

func createForecastController(db repositories.ForecastRepo, convert resources.ConverterI, handler services.WeatherHandler, config *config.ApplicationConfiguration, DBCity repositories.CityRepo) *api.ForecastAPIController {
	return api.NewForecastAPIController(services.NewForecastAPIService(db, convert, handler, config, DBCity))
}

func createWeatherService(dataGetter services.WeatherDataGetter,
	cityRepo repositories.CityRepo,
	forecastRepo repositories.ForecastRepo,
	configuration *config.ApplicationConfiguration,
	handler services.WeatherHandler) services.WeatherService {
	return services.NewWeatherService(dataGetter, cityRepo, forecastRepo, configuration, handler)
}

func createMuxRouter(routes api.Routes) *mux.Router {
	router := mux.NewRouter()

	for _, route := range routes {
		log.Printf("Registering route: %s %s", route.Method, route.Pattern)
		router.HandleFunc(route.Pattern, route.HandlerFunc).Methods(route.Method)
	}
	return router
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

func fetchWeatherData(service services.WeatherService) {
	go service.StartFetching(time.Second * 30)

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
