package main

import (
	"fmt"
	"github.com/happise/pixelwars/config"
	"github.com/happise/pixelwars/container"
	"github.com/happise/pixelwars/repository"
	"github.com/happise/pixelwars/router"
	"log/slog"
	"os"
)

func main() {
	logger := initLogger("local")
	conf := config.LoadConfig(logger)
	db, err := repository.NewDatabase(logger, conf)
	if err != nil {
		logger.Error("error connecting to database", "error", err)
		os.Exit(1)
	}
	cont := container.NewContainer(conf, logger, db)
	r := router.Init(cont)
	err = r.Start(fmt.Sprintf(":%d", conf.Port))
	if err != nil {
		logger.Error("error starting server", err)
		os.Exit(1)
	}
}

func initLogger(env string) *slog.Logger {
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	var handler slog.Handler
	handler = slog.NewTextHandler(os.Stdout, opts)
	if env == "production" {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	}
	return slog.New(handler)
}
