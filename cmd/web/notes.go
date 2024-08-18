package web

import (
	"fmt"
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
		data, err := app.NoteStore.ReadLatest(10)
		WriteJSONResponse(w, err, http.StatusOK, data)
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
		id := ParseId(w, chi.URLParam(r, "id"))
		data, err := app.NoteStore.ReadOne(id)
		WriteJSONResponse(w, err, http.StatusOK, data)
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
		id := ParseId(w, chi.URLParam(r, "id"))
		data, err := app.NoteStore.DeleteOne(id)
		WriteJSONResponse(w, err, http.StatusNoContent, data)
	}
}
