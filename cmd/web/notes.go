package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/zettadam/adamz-api-go/internal/types"
)

func (app *Application) handleReadLatestNotes(w http.ResponseWriter, r *http.Request) {
	data, err := app.Service.Notes.ReadLatest(10)
	WriteJSONResponse(w, err, http.StatusOK, data)
}

func (app *Application) handleCreateNote(w http.ResponseWriter, r *http.Request) {
	var p types.NoteRequest
	ReadJSONRequest(w, r, &p)
	// TODO: Validate payload

	data, err := app.Service.Notes.CreateOne(p)
	WriteJSONResponse(w, err, http.StatusCreated, data)
}

// ----------------------------------------------------------------------------
// Handlers in specific note context
// ----------------------------------------------------------------------------

func (app *Application) handleReadNote(w http.ResponseWriter, r *http.Request) {
	id := ParseId(w, chi.URLParam(r, "id"))
	data, err := app.Service.Notes.ReadOne(id)
	WriteJSONResponse(w, err, http.StatusOK, data)
}

func (app *Application) handleUpdateNote(w http.ResponseWriter, r *http.Request) {
	id := ParseId(w, chi.URLParam(r, "id"))
	var p types.NoteRequest
	ReadJSONRequest(w, r, &p)
	// TODO: Validate payload

	data, err := app.Service.Notes.UpdateOne(id, p)
	WriteJSONResponse(w, err, http.StatusOK, data)
}

func (app *Application) handleDeleteNote(w http.ResponseWriter, r *http.Request) {
	id := ParseId(w, chi.URLParam(r, "id"))
	data, err := app.Service.Notes.DeleteOne(id)
	WriteJSONResponse(w, err, http.StatusNoContent, data)
}
