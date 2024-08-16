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

	chi "github.com/go-chi/chi/v5"
	middleware "github.com/go-chi/chi/v5/middleware"
	pgxpool "github.com/jackc/pgx/v5/pgxpool"

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
	err := start()
	if err != nil {
		log.Fatal("Error starting server", err)
	}
}

func start() error {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Unable to determine working directory", err)
	}

	var cfg Configuration

	// --------------------------------------------------------------------------
	// CLI FLAGS -> CONFIGURATION

	flag.StringVar(&cfg.addr, "addr", ":3000", "HTTP network address")
	flag.StringVar(&cfg.logLevel, "logLevel", slog.LevelError.String(), "Logging level: INFO, WARN, ERROR, DEBUG")
	flag.DurationVar(&cfg.rtimeout, "rtimeout", 5*time.Second, "Read timeout (in seconds)")
	flag.DurationVar(&cfg.wtimeout, "wtimeout", 5*time.Second, "Write timeout (in seconds)")
	flag.Parse()

	// --------------------------------------------------------------------------
	// LOGGING
	setupLogging(cfg, wd)

	// --------------------------------------------------------------------------
	// DATABASE
	dbpool, err := connectDB()
	if err != nil {
		log.Fatal("Unable to connect to database", err)
	}
	defer dbpool.Close()

	// --------------------------------------------------------------------------
	// ROUTING
	r := setupRouter(cfg)

	// --------------------------------------------------------------------------
	// SERVER
	server := setupServer(cfg, r)

	// GO!
	slog.Info("Server is listening", slog.String("addr", cfg.addr))
	return server.ListenAndServe()
}

func setupLogging(cfg Configuration, wd string) {
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

	if cfg.logLevel == "DEBUG" {
		lvl.Set(slog.LevelDebug)
		opts.AddSource = true
	}

	l := slog.New(slog.NewTextHandler(os.Stderr, opts))

	slog.SetDefault(l)
}

func connectDB() (*pgxpool.Pool, error) {
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	return dbpool, nil
}

func setupRouter(cfg Configuration) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(cfg.rtimeout * time.Second))

	r.Get("/", web.Home)
	r.Mount("/posts", web.PostsRouter())
	r.Mount("/notes", web.NotesRouter())
	r.Mount("/code", web.CodeSnippetsRouter())
	r.Mount("/links", web.LinksRouter())
	r.Mount("/tasks", web.TasksRouter())
	r.Mount("/calendar", web.CalendarRouter())

	return r
}

func setupServer(cfg Configuration, handler http.Handler) *http.Server {
	server := &http.Server{
		Addr:         cfg.addr,
		Handler:      handler,
		ReadTimeout:  cfg.rtimeout * time.Second,
		WriteTimeout: cfg.wtimeout * time.Second,
	}

	slog.Info("Server created with configuration", slog.Any("cfg", cfg))

	return server
}
