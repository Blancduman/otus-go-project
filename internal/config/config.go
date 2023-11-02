package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

type Config struct {
	App   App
	Mongo Mongo
	HTTP  HTTP
	Kafka Kafka
	GRPC  GRPC
}

type App struct {
	Name     string `envconfig:"APP_NAME" default:"app"`
	LogLevel string `envconfig:"LOG_LEVEL" default:"debug"`
}

type HTTP struct {
	Port    int32    `envconfig:"HTTP_PORT" default:"8080"`
	Schemes []string `envconfig:"HTTP_SCHEMES" default:"http"`
}

func Load() (Config, error) {
	cnf := Config{} //nolint:exhaustruct

	if err := godotenv.Load(".env"); err != nil && !errors.Is(err, os.ErrNotExist) {
		return cnf, errors.Wrap(err, "read .env file")
	}

	if err := envconfig.Process("", &cnf); err != nil {
		return cnf, errors.Wrap(err, "read environment")
	}

	return cnf, nil
}

func (c *Config) LogLevel() (zerolog.Level, error) {
	lvl, err := zerolog.ParseLevel(c.App.LogLevel)
	if err != nil {
		return 0, errors.Wrapf(err, "loading log level from config value %q", c.App.LogLevel)
	}

	return lvl, nil
}

func (c *Config) HTTPAddr() string {
	return fmt.Sprintf(":%d", c.HTTP.Port)
}
