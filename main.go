package main

import (
	"go.uber.org/fx"
	"project_weather/bootstrap"
)

func main() {
	fx.New(
		bootstrap.FXModule_Core,
		bootstrap.FXModule_HTTPServer,
		bootstrap.FXModule_Persistence,
	).Run()
}
