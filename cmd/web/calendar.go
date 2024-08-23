package web

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/zettadam/adamz-api-go/internal/types"
)

func (app *Application) handleReadLatestEvents(w http.ResponseWriter, r *http.Request) {
	data, err := app.Service.Events.ReadLatest(10)
	WriteJSONResponse(w, err, http.StatusOK, data)
}

func (app *Application) handleCreateEvent(w http.ResponseWriter, r *http.Request) {
	var p types.EventRequest
	ReadJSONRequest(w, r, &p)

	err := validator.New().Struct(p)
	validationErrors := err.(validator.ValidationErrors)
	if err != nil {
		slog.Error("Event validation errors", slog.Any("details", validationErrors))
		WriteJSONResponse(w, err, http.StatusNotAcceptable, p)
		return
	}

	data, err := app.Service.Events.CreateOne(p)
	WriteJSONResponse(w, err, http.StatusCreated, data)
}

// ----------------------------------------------------------------------------
// Handlers in specific code snippet context
// ----------------------------------------------------------------------------

func (app *Application) handleReadEvent(w http.ResponseWriter, r *http.Request) {
	id := ParseId(w, chi.URLParam(r, "id"))
	data, err := app.Service.Events.ReadOne(id)
	WriteJSONResponse(w, err, http.StatusOK, data)
}

func (app *Application) handleUpdateEvent(w http.ResponseWriter, r *http.Request) {
	id := ParseId(w, chi.URLParam(r, "id"))
	var p types.EventRequest
	ReadJSONRequest(w, r, &p)

	err := validator.New().Struct(p)
	validationErrors := err.(validator.ValidationErrors)
	if err != nil {
		slog.Error("Event validation errors", slog.Any("details", validationErrors))
		WriteJSONResponse(w, err, http.StatusNotAcceptable, p)
		return
	}

	data, err := app.Service.Events.UpdateOne(id, p)
	WriteJSONResponse(w, err, http.StatusOK, data)
}

func (app *Application) handleDeleteEvent(w http.ResponseWriter, r *http.Request) {
	id := ParseId(w, chi.URLParam(r, "id"))
	data, err := app.Service.Events.DeleteOne(id)
	WriteJSONResponse(w, err, http.StatusNoContent, data)
}
