package bootstrap

import (
	"os"

	"github.com/NikolNikolaeva/project_weather/config"
	"go.uber.org/fx"
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
