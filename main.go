package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/zettadam/adamz-api-go/cmd/web"
)

type configuration struct {
	addr     string
	rtimeout time.Duration
	wtimeout time.Duration
}

func main() {
	// --------------------------------------------------------------------------
	// Logging
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))

	slog.SetDefault(logger)

	err := start()
	if err != nil {
		slog.Error("Error starting application.")
	}
}

func start() error {
	var cfg configuration

	// --------------------------------------------------------------------------
	// CLI Flags

	flag.StringVar(&cfg.addr, "addr", ":3000", "HTTP network address")
	flag.DurationVar(&cfg.rtimeout, "rtimeout", 5*time.Second, "Read timeout (in seconds)")
	flag.DurationVar(&cfg.wtimeout, "wtimeout", 5*time.Second, "Write timeout (in seconds)")

	flag.Parse()

	// --------------------------------------------------------------------------
	// Server

	server := &http.Server{
		Addr:         cfg.addr,
		Handler:      web.Router(),
		ReadTimeout:  cfg.rtimeout * time.Second,
		WriteTimeout: cfg.wtimeout * time.Second,
	}

	slog.Info("Server created with flags", "cfg", cfg)
	slog.Info("Server is listening", "addr", cfg.addr)

	return server.ListenAndServe()
}
