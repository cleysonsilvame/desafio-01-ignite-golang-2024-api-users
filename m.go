package main

import (
	"log/slog"
	"net/http"
	"time"

	"api-users/api"
)

func main() {
	if err := run(); err != nil {
		slog.Error("failed to run code", "error", err)
		return
	}

	slog.Info("All systems offline!")
}

func run() error {
	app := api.NewApplication()

	api := api.NewHanlder(app)

	s := http.Server{
		Addr:         ":8080",
		Handler:      api,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  time.Minute,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
