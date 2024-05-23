package bootstrap

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"log"
	"net/http"
	"project_weather/config"
	"project_weather/controllers"
	"project_weather/repositories"
	"project_weather/services"
	"slices"
)

var FXModule_HTTPServer = fx.Module(
	"http-server",
	fx.Provide(
		createFiberApp,
		createAPIRoutes,
		repositories.NewCityRepo,
		controllers.NewCityController,
		repositories.NewForecastRepo,
		controllers.NewForecastController,
		createWeatherApiService,
		services.NewCityService,
		services.NewForecastService,
	),
	fx.Invoke(
		configureAPIRoutes,
		registerServerStartHook,
	),
)

func createFiberApp() *fiber.App {
	return fiber.New()
}

func createAPIRoutes(cities *controllers.CityController, forecasts *controllers.ForecastController, weather *services.WeatherApiService) []controllers.Route {
	return slices.Concat(
		cities.GetRoutes(),
		forecasts.GetRoutes(),
		weather.GetRoutes(),
	)
}

func createWeatherApiService(config *config.ApplicationConfiguration) *services.WeatherApiService {
	client := &http.Client{}
	return services.NewWeatherController(client, config.ApiKeyWeatherApi)
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
