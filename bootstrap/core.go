package bootstrap

import (
	"go.uber.org/fx"
	"os"
	"project_weather/config"
)

var FXModule_Core = fx.Module(
	"core",
	fx.Provide(
		createApplicationConfiguration,
	),
)

func createApplicationConfiguration() (*config.ApplicationConfiguration, error) {
	return config.NewApplicationConfiguration(os.Getenv("ENVIRONMENT"))
}
