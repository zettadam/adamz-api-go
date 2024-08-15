package web

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NotesRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", readLatestNotes)
	r.Post("/", createNote)

	r.Route("/{id}", func(r chi.Router) {
		r.Use(noteCtx)
		r.Get("/", readNote)
		r.Put("/", updateNote)
		r.Delete("/", deleteNote)
	})

	return r
}

func readLatestNotes(w http.ResponseWriter, r *http.Request) {
	msg := "ReadLatestNotes"
	fmt.Fprint(w, msg)
}

func createNote(w http.ResponseWriter, r *http.Request) {
	msg := "CreateNote"
	fmt.Fprint(w, msg)
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

func readNote(w http.ResponseWriter, r *http.Request) {
	msg := "ReadNote"
	fmt.Fprint(w, msg)
}

func updateNote(w http.ResponseWriter, r *http.Request) {
	msg := "UpdateNote"
	fmt.Fprint(w, msg)
}

func deleteNote(w http.ResponseWriter, r *http.Request) {
	msg := "DeleteNote"
	fmt.Fprint(w, msg)
}
