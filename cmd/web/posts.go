package web

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"

	"github.com/zettadam/adamz-api-go/internal/types"
)

func (app *Application) handleReadLatestPosts(w http.ResponseWriter, r *http.Request) {
	data, err := app.Service.Posts.ReadLatest(10)
	WriteJSONResponse(w, err, http.StatusOK, data)
}

func (app *Application) handleCreatePost(w http.ResponseWriter, r *http.Request) {
	var p types.PostRequest
	ReadJSONRequest(w, r, &p)

	err := validator.New().Struct(p)
	validationErrors := err.(validator.ValidationErrors)
	if err != nil {
		slog.Error("Post validation errors", slog.Any("details", validationErrors))
		WriteValidationErrors(w, validationErrors)
		return
	}

	data, err := app.Service.Posts.CreateOne(p)
	WriteJSONResponse(w, err, http.StatusCreated, data)
}

// ----------------------------------------------------------------------------
// Handlers in specific post context
// ----------------------------------------------------------------------------

func (app *Application) handleReadPost(w http.ResponseWriter, r *http.Request) {
	id := ParseId(w, chi.URLParam(r, "id"))
	data, err := app.Service.Posts.ReadOne(id)
	WriteJSONResponse(w, err, http.StatusOK, data)
}

func (app *Application) handleUpdatePost(w http.ResponseWriter, r *http.Request) {
	id := ParseId(w, chi.URLParam(r, "id"))

	var p types.PostRequest
	ReadJSONRequest(w, r, &p)

	err := validator.New().Struct(p)
	validationErrors := err.(validator.ValidationErrors)
	if err != nil {
		slog.Error("Post validation errors", slog.Any("details", validationErrors))
		WriteValidationErrors(w, validationErrors)
		return
	}

	data, err := app.Service.Posts.UpdateOne(id, p)
	WriteJSONResponse(w, validationErrors, http.StatusOK, data)
}

func (app *Application) handleDeletePost(w http.ResponseWriter, r *http.Request) {
	id := ParseId(w, chi.URLParam(r, "id"))
	data, err := app.Service.Posts.DeleteOne(id)
	WriteJSONResponse(w, err, http.StatusNoContent, data)
}
