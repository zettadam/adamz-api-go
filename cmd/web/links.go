package web

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/zettadam/adamz-api-go/internal/types"
)

func (app *Application) handleReadLatestLinks(w http.ResponseWriter, r *http.Request) {
	data, err := app.Service.Links.ReadLatest(10)
	WriteJSONResponse(w, err, http.StatusOK, data)
}

func (app *Application) handleCreateLink(w http.ResponseWriter, r *http.Request) {
	var p types.LinkRequest
	ReadJSONRequest(w, r, &p)

	err := validator.New().Struct(p)
	validationErrors := err.(validator.ValidationErrors)
	if err != nil {
		slog.Error("Link validation errors", slog.Any("details", validationErrors))
		WriteJSONResponse(w, err, http.StatusNotAcceptable, p)
		return
	}

	data, err := app.Service.Links.CreateOne(p)
	WriteJSONResponse(w, err, http.StatusCreated, data)
}

// ----------------------------------------------------------------------------
// Handlers in specific link context
// ----------------------------------------------------------------------------

func (app *Application) handleReadLink(w http.ResponseWriter, r *http.Request) {
	id := ParseId(w, chi.URLParam(r, "id"))
	data, err := app.Service.Links.ReadOne(id)
	WriteJSONResponse(w, err, http.StatusOK, data)
}

func (app *Application) handleUpdateLink(w http.ResponseWriter, r *http.Request) {
	id := ParseId(w, chi.URLParam(r, "id"))

	var p types.LinkRequest
	ReadJSONRequest(w, r, &p)

	err := validator.New().Struct(p)
	validationErrors := err.(validator.ValidationErrors)
	if err != nil {
		slog.Error("Link validation errors", slog.Any("details", validationErrors))
		WriteJSONResponse(w, err, http.StatusNotAcceptable, p)
		return
	}

	data, err := app.Service.Links.UpdateOne(id, p)
	WriteJSONResponse(w, err, http.StatusOK, data)
}

func (app *Application) handleDeleteLink(w http.ResponseWriter, r *http.Request) {
	id := ParseId(w, chi.URLParam(r, "id"))
	data, err := app.Service.Links.DeleteOne(id)
	WriteJSONResponse(w, err, http.StatusNoContent, data)
}
