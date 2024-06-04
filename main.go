package main

import (
	"go.uber.org/fx"

	"github.com/NikolNikolaeva/project_weather/bootstrap"
)

//	@title			Weather API
//	@version		1.0
//	@description	This is a weather api
//	@termsOfService	http://swagger.io/terms/

func main() {
	fx.New(
		bootstrap.FXModule_Core,
		bootstrap.FXModule_Persistence,
		bootstrap.FXModule_HTTPServer,
	).Run()
}
