package internal

import (
	"context"
	"database/sql"
	"os"

	"github.com/NikolNikolaeva/project_weather/bootstrap"
	"github.com/NikolNikolaeva/project_weather/it/internal/gomisc/lang"
	arrays "github.com/NikolNikolaeva/project_weather/it/internal/gomisc/lang/array"
	"github.com/NikolNikolaeva/project_weather/it/internal/gomisc/logs"

	"go.uber.org/fx"
	gpg "gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/magiconair/properties"
)

const DatabaseDriver = "postgres"
const DatabaseURL = "postgres://postgres:postgres@127.0.0.1:5432/postgres?sslmode=disable"

func asFinalizer(_ int, key string) func() {
	original, exists := os.LookupEnv(key)

	return func() {
		_ = os.Unsetenv(key)
		if exists {
			_ = os.Setenv(key, original)
		}
	}
}

func NewGormDatabase() (*gorm.DB, func()) {
	connection := lang.Must(sql.Open(DatabaseDriver, DatabaseURL))

	return lang.Must(gorm.Open(
		gpg.New(gpg.Config{Conn: connection}),
		&gorm.Config{SkipDefaultTransaction: false},
	)), func() { _ = connection.Close() }
}

func NewFXApplication(ctx context.Context) (*fx.App, func()) {
	app := fx.New(
		fx.NopLogger,
		fx.Module(
			"",
			fx.Provide(
				func() logs.Logger { return logs.NopLogger },
			),
		),

		bootstrap.FXModule_Core,
		bootstrap.FXModule_Persistence,
		bootstrap.FXModule_HTTPServer,
	)

	return lang.Must(app, app.Start(ctx)), func() { _ = app.Stop(ctx) }
}

func RunWithEnvironment(path string, callback func() int) int {
	environment := properties.MustLoadFile(path, properties.UTF8)
	defer func(finalizers []func()) {
		for _, finalizer := range finalizers {
			finalizer()
		}
	}(arrays.Map(environment.Keys(), asFinalizer))

	for _, key := range environment.Keys() {
		_ = os.Setenv(key, environment.MustGet(key))
	}

	return callback()
}
