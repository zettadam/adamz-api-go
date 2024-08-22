package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/zettadam/adamz-api-go/internal/config"
	"github.com/zettadam/adamz-api-go/internal/models"
)

func PostsRouter(app *config.Application) http.Handler {
	r := chi.NewRouter()

	r.Get("/", handleReadLatestPosts(app))
	r.Post("/", handleCreatePost(app))

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", handleReadPost(app))
		r.Put("/", handleUpdatePost(app))
		r.Delete("/", handleDeletePost(app))
	})

	return r
}

func handleReadLatestPosts(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := app.PostStore.ReadLatest(10)
		WriteJSONResponse(w, err, http.StatusOK, data)
	}
}

func handleCreatePost(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p models.PostRequest
		ReadJSONRequest(w, r, &p)
		// TODO: Validate payload

		data, err := app.PostStore.CreateOne(p)
		WriteJSONResponse(w, err, http.StatusCreated, data)
	}
}

// ----------------------------------------------------------------------------
// Handlers in specific post context
// ----------------------------------------------------------------------------

func handleReadPost(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := ParseId(w, chi.URLParam(r, "id"))
		data, err := app.PostStore.ReadOne(id)
		WriteJSONResponse(w, err, http.StatusOK, data)
	}
}

func handleUpdatePost(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := ParseId(w, chi.URLParam(r, "id"))
		var p models.PostRequest
		ReadJSONRequest(w, r, &p)
		// TODO: Validate payload

		data, err := app.PostStore.UpdateOne(id, p)
		WriteJSONResponse(w, err, http.StatusOK, data)
	}
}

func handleDeletePost(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := ParseId(w, chi.URLParam(r, "id"))
		data, err := app.PostStore.DeleteOne(id)
		WriteJSONResponse(w, err, http.StatusNoContent, data)
	}
}

// ----------------------------------------------------------------------------
// Validation
// ----------------------------------------------------------------------------

func validatePost(input models.Post) (models.Post, error) {
	return input, nil
}
