package main

import (
	"flag"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/zettadam/adamz-api-go/cmd/web"
)

type Configuration struct {
	addr     string
	logLevel string
	rtimeout time.Duration
	wtimeout time.Duration
}

func (c Configuration) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("addr", c.addr),
		slog.String("logLevel", c.logLevel),
		slog.Any("rtimeout", c.rtimeout),
		slog.Any("wtimeout", c.wtimeout),
	)
}

func main() {
	wd, err0 := os.Getwd()
	if err0 != nil {
		log.Fatal("Unable to determine working directory")
	}

	var cfg Configuration

	// --------------------------------------------------------------------------
	// CLI Flags -> Configuration

	flag.StringVar(&cfg.addr, "addr", ":3000", "HTTP network address")
	flag.StringVar(&cfg.logLevel, "logLevel", slog.LevelError.String(), "Logging level: INFO, WARN, ERROR, DEBUG")
	flag.DurationVar(&cfg.rtimeout, "rtimeout", 5*time.Second, "Read timeout (in seconds)")
	flag.DurationVar(&cfg.wtimeout, "wtimeout", 5*time.Second, "Write timeout (in seconds)")

	flag.Parse()

	// --------------------------------------------------------------------------
	// Logging

	logLevel := new(slog.LevelVar)
	loggerOptions := &slog.HandlerOptions{
		Level: logLevel,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.SourceKey {
				source := a.Value.Any().(*slog.Source)
				// remove current working directory and only leave the relative path to the program
				if file, ok := strings.CutPrefix(source.File, wd); ok {
					source.File = file
				}
			}
			return a
		},
	}

	if cfg.logLevel == "DEBUG" {
		logLevel.Set(slog.LevelDebug)
		loggerOptions.AddSource = true
	}

	logger := slog.New(slog.NewTextHandler(os.Stderr, loggerOptions))

	slog.SetDefault(logger)

	// --------------------------------------------------------------------------
	// Start
	err := start(cfg)
	if err != nil {
		slog.Error("Error starting server", slog.String("err", err.Error()))
	}
}

func start(cfg Configuration) error {
	// --------------------------------------------------------------------------
	// Server

	server := &http.Server{
		Addr:         cfg.addr,
		Handler:      web.Router(),
		ReadTimeout:  cfg.rtimeout * time.Second,
		WriteTimeout: cfg.wtimeout * time.Second,
	}

	slog.Info("Server created with configuration", slog.Any("cfg", cfg))
	slog.Info("Server is listening", slog.String("addr", cfg.addr))

	return server.ListenAndServe()
}
