package web

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

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
		data, err := app.NoteStore.ReadLatest(10)
		if err != nil {
			slog.Error("Error fetching notes", err)
		}
		WriteJSON(w, 200, data)
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

func handleReadNote(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_id := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(_id, 10, 64)
		if err != nil {
			slog.Error("Unable to convert parameter to int64", id, err)
		}

		data, err := app.NoteStore.ReadOne(id)
		if err != nil {
			slog.Error("Error fetching note",
				slog.String("id", _id),
				slog.Any("err", err),
			)
		}
		WriteJSON(w, 200, data)
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
