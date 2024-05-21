package bootstrap

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"log"
	"project_weather/config"
	"project_weather/controllers"
	"project_weather/repositories"
	"slices"
)

var FXModule_HTTPServer = fx.Module(
	"http-server",
	fx.Provide(
		createFiberApp,
		createAPIRoutes,
		createCityController,
	),
	fx.Invoke(
		configureAPIRoutes,
		registerServerStartHook,
	),
)

func createFiberApp() *fiber.App {
	return fiber.New()
}

func createCityController(db *gorm.DB) *controllers.CityController {
	return &controllers.CityController{
		DB: repositories.CityRepo{
			Db: db,
		},
	}
}

func createAPIRoutes(cities *controllers.CityController) []controllers.Route {
	return slices.Concat(
		cities.GetRoutes(),
	)
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
