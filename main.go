package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/zettadam/adamz-api-go/cmd/web"
	"github.com/zettadam/adamz-api-go/internal/config"
)

type configuration struct {
	addr     string
	rtimeout time.Duration
	wtimeout time.Duration
}

func main() {
	// --------------------------------------------------------------------------
	// Logging

	app := &config.Application{
		InfoLog:  log.New(os.Stdout, "INFO:\t", log.LstdFlags|log.Ldate|log.Ltime|log.Lshortfile),
		ErrorLog: log.New(os.Stderr, "ERROR:\t", log.LstdFlags|log.Ldate|log.Ltime|log.Lshortfile),
	}

	err := start(app)
	if err != nil {
		app.ErrorLog.Fatal(err)
	}
}

func start(app *config.Application) error {
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
		Handler:      web.Router(app),
		ReadTimeout:  cfg.rtimeout * time.Second,
		WriteTimeout: cfg.wtimeout * time.Second,
	}

	app.InfoLog.Printf("Server created with flags %s, %s, %s", cfg.addr, cfg.rtimeout, cfg.wtimeout)
	app.InfoLog.Printf("Server is listening on %s", cfg.addr)

	return server.ListenAndServe()
}
