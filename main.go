package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type application struct {
	logger *slog.Logger
	client *http.Client
}

func main() {
	addr := flag.String("addr", ":8080", "HTTP network address")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	app := &application{
		logger: logger,
		client: client,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", app.GetHandler)
	mux.HandleFunc("GET /{id}", app.GetHandler)

	logger.Info("starting server", slog.Any("addr", ":8080"))

	err := http.ListenAndServe(*addr, mux)

	logger.Error(err.Error())

	os.Exit(1)
}
