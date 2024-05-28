package main

import (
	"github.com/NikolNikolaeva/project_weather/bootstrap"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		bootstrap.FXModule_Core,
		bootstrap.FXModule_Persistence,
		bootstrap.FXModule_HTTPServer,
	).Run()
}
