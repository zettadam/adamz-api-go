package web

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/zettadam/adamz-api-go/internal/config"
)

func NotesRouter(app *config.Application) http.Handler {
	r := chi.NewRouter()

	r.Get("/", handleReadLatestNotes(app))
	r.Post("/", handleCreateNote(app))

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", handleReadNote(app))
		r.Put("/", handleUpdateNote(app))
		r.Delete("/", handleDeleteNote(app))
	})

	return r
}

func handleReadLatestNotes(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		msg := "ReadLatestNotes"
		fmt.Fprint(w, msg)
	}
}

func handleCreateNote(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		msg := "CreateNote"
		fmt.Fprint(w, msg)
	}
}

// ----------------------------------------------------------------------------
// Handlers in specific note context
// ----------------------------------------------------------------------------

func noteCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		slog.Info("URL params", slog.Any("id", id))

		if id != "" {
			// get post by ID
		} else {
			fmt.Fprint(w, "Note not found")
			return
		}

		ctx := context.WithValue(r.Context(), "id", id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func handleReadNote(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		msg := "ReadNote"
		fmt.Fprint(w, msg)
	}
}

func handleUpdateNote(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		msg := "UpdateNote"
		fmt.Fprint(w, msg)
	}
}

func handleDeleteNote(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		msg := "DeleteNote"
		fmt.Fprint(w, msg)
	}
}
