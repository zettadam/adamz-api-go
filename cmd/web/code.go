package web

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/zettadam/adamz-api-go/internal/types"
)

func (app *Application) handleReadLatestCodeSnippets(w http.ResponseWriter, r *http.Request) {
	data, err := app.Service.CodeSnippets.ReadLatest(10)
	WriteJSONResponse(w, err, http.StatusOK, data)
}

func (app *Application) handleCreateCodeSnippet(w http.ResponseWriter, r *http.Request) {
	var p types.CodeSnippetRequest
	ReadJSONRequest(w, r, &p)

	err := validator.New().Struct(p)
	validationErrors := err.(validator.ValidationErrors)
	if err != nil {
		slog.Error("Code snippet validation errors", slog.Any("details", validationErrors))
		WriteJSONResponse(w, err, http.StatusNotAcceptable, p)
		return
	}

	data, err := app.Service.CodeSnippets.CreateOne(p)
	WriteJSONResponse(w, err, http.StatusCreated, data)
}

// ----------------------------------------------------------------------------
// Handlers in specific code snippet context
// ----------------------------------------------------------------------------

func (app *Application) handleReadCodeSnippet(w http.ResponseWriter, r *http.Request) {
	id := ParseId(w, chi.URLParam(r, "id"))
	data, err := app.Service.CodeSnippets.ReadOne(id)
	WriteJSONResponse(w, err, http.StatusOK, data)
}

func (app *Application) handleUpdateCodeSnippet(w http.ResponseWriter, r *http.Request) {
	id := ParseId(w, chi.URLParam(r, "id"))
	var p types.CodeSnippetRequest
	ReadJSONRequest(w, r, &p)

	err := validator.New().Struct(p)
	validationErrors := err.(validator.ValidationErrors)
	if err != nil {
		slog.Error("Code snippet validation errors", slog.Any("details", validationErrors))
		WriteJSONResponse(w, err, http.StatusNotAcceptable, p)
		return
	}

	data, err := app.Service.CodeSnippets.UpdateOne(id, p)
	WriteJSONResponse(w, err, http.StatusOK, data)
}

func (app *Application) handleDeleteCodeSnippet(w http.ResponseWriter, r *http.Request) {
	id := ParseId(w, chi.URLParam(r, "id"))
	data, err := app.Service.CodeSnippets.DeleteOne(id)
	WriteJSONResponse(w, err, http.StatusNoContent, data)
}
