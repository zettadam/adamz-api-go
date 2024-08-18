package web

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/zettadam/adamz-api-go/internal/config"
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
		msg := "CreateLink"
		fmt.Fprint(w, msg)
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
		msg := "UpdateLink"
		fmt.Fprint(w, msg)
	}
}

func handleDeleteLink(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := ParseId(w, chi.URLParam(r, "id"))
		data, err := app.LinkStore.DeleteOne(id)
		WriteJSONResponse(w, err, http.StatusNoContent, data)
	}
}
