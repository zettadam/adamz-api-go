package web

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/zettadam/adamz-api-go/internal/config"
)

func LinksRouter(app *config.Application) http.Handler {
	r := chi.NewRouter()

	r.Get("/", handleReadLatestLinks(app))
	r.Post("/", handleCreateLink(app))

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", handleReadLink(app))
		r.Put("/", handleUpdateLink(app))
		r.Delete("/", handleDeleteLink(app))
	})

	return r
}

func handleReadLatestLinks(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		msg := "ReadLatestLinks"
		fmt.Fprint(w, msg)
	}
}

func handleCreateLink(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		msg := "CreateLink"
		fmt.Fprint(w, msg)
	}
}

// ----------------------------------------------------------------------------
// Handlers in specific link context
// ----------------------------------------------------------------------------

func linkCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		slog.Info("URL params", slog.Any("id", id))

		if id != "" {
			// get post by ID
		} else {
			fmt.Fprint(w, "Link not found")
			return
		}

		ctx := context.WithValue(r.Context(), "id", id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func handleReadLink(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		msg := "ReadLink"
		fmt.Fprint(w, msg)
	}
}

func handleUpdateLink(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		msg := "UpdateLink"
		fmt.Fprint(w, msg)
	}
}

func handleDeleteLink(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		msg := "DeleteLink"
		fmt.Fprint(w, msg)
	}
}
