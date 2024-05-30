package bootstrap

import (
	"fmt"
	"log"
	"net/http"
	"slices"

	"github.com/NikolNikolaeva/project_weather/config"
	"github.com/NikolNikolaeva/project_weather/controllers"
	"github.com/NikolNikolaeva/project_weather/generated/dao"
	"github.com/NikolNikolaeva/project_weather/repositories"
	"github.com/NikolNikolaeva/project_weather/services"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

var FXModule_HTTPServer = fx.Module(
	"http-server",
	fx.Provide(
		createFiberApp,
		createAPIRoutes,
		createCityRepo,
		createCityController,
		createForecastRepo,
		createForecastController,
		createWeatherApiService,
		createHTTPClient,
		createWetherHandler,
		createWeatherDataGetter,
	),
	fx.Invoke(
		configureAPIRoutes,
		registerServerStartHook,
	),
)

func createFiberApp() *fiber.App {
	return fiber.New()
}

func createAPIRoutes(cities controllers.CityController, forecasts controllers.ForecastController, weather services.WeatherApiService) []controllers.Route {
	return slices.Concat(
		cities.GetRoutes(),
		forecasts.GetRoutes(),
		weather.GetRoutes(),
	)
}

func createWeatherDataGetter(client *http.Client) services.WeatherDataGetter {
	return services.NewWeatherDataGetter(client)
}

func createWetherHandler(cityRepo repositories.CityRepo, forecastRepo repositories.ForecastRepo, getter services.WeatherDataGetter) services.WeatherHandler {
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

func createForecastController(db repositories.ForecastRepo) controllers.ForecastController {
	return controllers.NewForecastController(db)
}

func createCityController(db repositories.CityRepo) controllers.CityController {
	return controllers.NewCityController(db)
}

func createWeatherApiService(weatherHandler services.WeatherHandler, config *config.ApplicationConfiguration) services.WeatherApiService {

	return services.NewWeatherService(weatherHandler, config, config.CredFile)
}

func configureAPIRoutes(app *fiber.App, routes []controllers.Route) {
	for _, route := range routes {
		log.Printf("Registering route: %s %s", route.Method, route.Path)
		app.Add(route.Method, route.Path, route.Handler)
	}
}

func registerServerStartHook(lc fx.Lifecycle, app *fiber.App, config *config.ApplicationConfiguration) {
	lc.Append(fx.StartStopHook(
		func() {
			go func() {
				err := app.Listen(fmt.Sprintf(":%s", config.HTTPPort))
				if err != nil {
					log.Fatalf("Failed to start server: %v", err)
				}
				log.Printf("Server is starting on port %s", config.HTTPPort)
			}()
		},
		func() {
			if err := app.Shutdown(); err != nil {
				log.Printf("Error shutting down server: %v", err)
			}
			log.Println("Server stopped successfully")
		},
	))
}
