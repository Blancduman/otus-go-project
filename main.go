package main

import (
	"context"
	"github.com/Blancduman/banners-rotation/cmd"
	"github.com/Blancduman/banners-rotation/internal/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	conf, err := config.Load()
	if err != nil {
		panic(err)
	}

	logLevel, err := conf.LogLevel()
	if err != nil {
		panic(err)
	}

	zerolog.SetGlobalLevel(logLevel)
	exitCode := 0

	err = cmd.Run(context.Background(), conf)
	if err != nil {
		log.Err(err).Send()

		exitCode = 1
	}

	os.Exit(exitCode)
}
