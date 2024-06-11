package bootstrap

import (
	"os"

	"go.uber.org/fx"

	"github.com/NikolNikolaeva/project_weather/config"
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
