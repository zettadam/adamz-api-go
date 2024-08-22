package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/zettadam/adamz-api-go/internal/config"
	"github.com/zettadam/adamz-api-go/internal/models"
)

func LinksRouter(app *config.Application) http.Handler {
	r := chi.NewRouter()

	r.Get("/", handleReadLatestLinks(app))
	r.Post("/", handleCreateLink(app))

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", handleReadLink(app))
		r.Put("/", handleUpdateLink(app))
		r.Delete("/", handleDeleteLink(app))
	})

	return r
}

func handleReadLatestLinks(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := app.LinkStore.ReadLatest(10)
		WriteJSONResponse(w, err, http.StatusOK, data)
	}
}

func handleCreateLink(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p models.LinkRequest
		ReadJSONRequest(w, r, &p)
		// TODO: Validate payload

		data, err := app.LinkStore.CreateOne(p)
		WriteJSONResponse(w, err, http.StatusCreated, data)
	}
}

// ----------------------------------------------------------------------------
// Handlers in specific link context
// ----------------------------------------------------------------------------

func handleReadLink(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := ParseId(w, chi.URLParam(r, "id"))
		data, err := app.LinkStore.ReadOne(id)
		WriteJSONResponse(w, err, http.StatusOK, data)
	}
}

func handleUpdateLink(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := ParseId(w, chi.URLParam(r, "id"))
		var p models.LinkRequest
		ReadJSONRequest(w, r, &p)
		// TODO: Validate payload

		data, err := app.LinkStore.UpdateOne(id, p)
		WriteJSONResponse(w, err, http.StatusOK, data)
	}
}

func handleDeleteLink(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := ParseId(w, chi.URLParam(r, "id"))
		data, err := app.LinkStore.DeleteOne(id)
		WriteJSONResponse(w, err, http.StatusNoContent, data)
	}
}
