package main

import (
	"context"
	"flag"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"

	"github.com/zettadam/adamz-api-go/cmd/web"
	"github.com/zettadam/adamz-api-go/internal/config"
	"github.com/zettadam/adamz-api-go/internal/stores"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Unable to determine working directory", err)
	}

	var cfg config.Configuration

	// --------------------------------------------------------------------------
	// CLI FLAGS -> CONFIGURATION
	// --------------------------------------------------------------------------

	flag.StringVar(&cfg.Addr, "addr", ":3000", "HTTP network address")
	flag.StringVar(&cfg.LogLevel, "logLevel", slog.LevelError.String(), "Logging level: INFO, WARN, ERROR, DEBUG")
	flag.DurationVar(&cfg.ReadTimeout, "rtimeout", 5*time.Second, "Read timeout (in seconds)")
	flag.DurationVar(&cfg.WriteTimeout, "wtimeout", 5*time.Second, "Write timeout (in seconds)")
	flag.Parse()

	// --------------------------------------------------------------------------
	// LOGGING
	// --------------------------------------------------------------------------
	lvl := new(slog.LevelVar)
	opts := &slog.HandlerOptions{
		Level: lvl,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.SourceKey {
				s := a.Value.Any().(*slog.Source)
				if f, ok := strings.CutPrefix(s.File, wd); ok {
					s.File = f
				}
			}
			return a
		},
	}

	if cfg.LogLevel == "DEBUG" {
		lvl.Set(slog.LevelDebug)
		opts.AddSource = true
	}

	l := slog.New(slog.NewTextHandler(os.Stderr, opts))

	slog.SetDefault(l)

	// --------------------------------------------------------------------------
	// DATABASE
	// --------------------------------------------------------------------------

	dbPool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Unable to connect to database", err)
	}
	defer dbPool.Close()

	app := &web.Application{
		Service: &web.Service{
			CodeSnippets: &stores.CodeSnippetStore{DB: dbPool},
			Events:       &stores.EventStore{DB: dbPool},
			Links:        &stores.LinkStore{DB: dbPool},
			Notes:        &stores.NoteStore{DB: dbPool},
			Posts:        &stores.PostStore{DB: dbPool},
			Tasks:        &stores.TaskStore{DB: dbPool},
		},
	}

	// --------------------------------------------------------------------------
	// SERVER
	server := &http.Server{
		Addr:         cfg.Addr,
		Handler:      app.SetupRouter(),
		ReadTimeout:  cfg.ReadTimeout * time.Second,
		WriteTimeout: cfg.WriteTimeout * time.Second,
	}

	slog.Info(
		"Server created with configuration",
		slog.Any("cfg", cfg),
	)

	// GO!
	slog.Info(
		"Server is listening",
		slog.String("addr", cfg.Addr),
	)

	server.ListenAndServe()
}
