package bootstrap

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.uber.org/fx"

	"github.com/NikolNikolaeva/project_weather/config"
	api "github.com/NikolNikolaeva/project_weather/generated/api/project-weather/rest"
	"github.com/NikolNikolaeva/project_weather/utils"

	"github.com/NikolNikolaeva/project_weather/repositories"
	"github.com/NikolNikolaeva/project_weather/resources"
	"github.com/NikolNikolaeva/project_weather/services"

	gocron "github.com/go-co-op/gocron/v2"
)

var FXModule_HTTPServer = fx.Module(
	"http-server",
	fx.Provide(
		createAPIRoutes,
		createCityController,
		createForecastController,
		createHTTPClient,
		createWeatherHandler,
		createWeatherDataRetriever,
		createWeatherService,
		createHttpServer,
		createConverter,
		createMuxRouter,
		createScheduler,
	),
	fx.Invoke(
		registerServerStartHook,
		registerSchedulerStartStopHook,
		registerFetchWeatherDataCronJob,
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

func createWeatherHandler(cityRepo repositories.CityRepo, forecastRepo repositories.ForecastRepo, getter services.WeatherDataGetter) services.WeatherAPIClient {
	return services.NewWeatherAPIClient(cityRepo, forecastRepo, getter)
}

func createHTTPClient() *http.Client {
	return services.NewHTTPClient()
}

func createCityController(db repositories.CityRepo, convert resources.ConverterI, client services.WeatherAPIClient, configuration *config.ApplicationConfiguration) *api.CityAPIController {
	return api.NewCityAPIController(services.NewAPIService(db, convert, client, configuration))
}

func createForecastController(db repositories.ForecastRepo, convert resources.ConverterI, handler services.WeatherAPIClient, config *config.ApplicationConfiguration, DBCity repositories.CityRepo) *api.ForecastAPIController {
	return api.NewForecastAPIController(services.NewForecastAPIService(db, convert, handler, config, DBCity))
}

func createWeatherService(dataGetter services.WeatherDataGetter,
	cityRepo repositories.CityRepo,
	forecastRepo repositories.ForecastRepo,
	configuration *config.ApplicationConfiguration,
	handler services.WeatherAPIClient) services.WeatherService {
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

func createHttpServer(config *config.ApplicationConfiguration, router *mux.Router) *http.Server {
	return &http.Server{
		Addr:    ":" + config.HTTPPort,
		Handler: router,
	}
}

func createConverter() resources.ConverterI {
	return resources.NewConverter()
}

func createScheduler() (gocron.Scheduler, error) {
	return gocron.NewScheduler()
}

func registerFetchWeatherDataCronJob(service services.WeatherService, s gocron.Scheduler) error {
	_, err := s.NewJob(
		gocron.DurationJob(
			30*time.Second,
		),
		gocron.NewTask(
			service.StartFetching,
		),
	)

	return err
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

func registerSchedulerStartStopHook(lc fx.Lifecycle, scheduler gocron.Scheduler) {
	lc.Append(fx.StartStopHook(
		func() {
			log.Println("starting scheduler...")
			scheduler.Start()
		},
		func() {
			log.Println("shutting scheduler down...")
			_ = scheduler.Shutdown()
		},
	))
}
