package main

import (
	"log/slog"
	"os"
	"sysqueue/cmd/web"
)

var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}

func run() error {
	logger.Info("starting web")

	router := web.NewRouter()
	server := web.NewServer(":9999", router)
	server.Start()
	return nil
}
