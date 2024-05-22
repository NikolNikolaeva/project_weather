package config

import (
	"github.com/lpernett/godotenv"
	"time"
)
import "github.com/kelseyhightower/envconfig"

func NewApplicationConfiguration(envfile string) (*ApplicationConfiguration, error) {
	if envfile != "" {
		if err := godotenv.Load(envfile); err != nil {
			return nil, err
		}
	}

	result := &ApplicationConfiguration{}

	return result, envconfig.Process("", result)
}

type ApplicationConfiguration struct {
	LogLevel string `envconfig:"LOG_LEVEL" default:"debug"`

	HTTPPort               string        `envconfig:"HTTP_PORT" default:"8080"`
	HTTPServerReadTimeout  time.Duration `envconfig:"HTTP_SERVER_READ_TIMEOUT" default:"5s"`
	HTTPServerWriteTimeout time.Duration `envconfig:"HTTP_SERVER_WRITE_TIMEOUT" default:"5s"`

	DBHost                 string `envconfig:"DB_HOST"`
	DBPort                 string `envconfig:"DB_PORT"`
	DBName                 string `envconfig:"DB_NAME"`
	DBUsername             string `envconfig:"DB_USERNAME"`
	DBPassword             string `envconfig:"DB_PASSWORD"`
	DatabaseMigrationsPath string `envconfig:"DATABASE_MIGRATIONS_PATH"`
}
