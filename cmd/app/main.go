package main

import (
	"log/slog"
	"os"
)

var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}

func run() error {
	logger.Info("starting app")

	return nil
}
