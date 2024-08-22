package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/zettadam/adamz-api-go/internal/config"
	"github.com/zettadam/adamz-api-go/internal/models"
)

func CalendarRouter(app *config.Application) http.Handler {
	r := chi.NewRouter()

	r.Get("/", handleReadLatestEvents(app))
	r.Post("/", handleCreateEvent(app))

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", handleReadEvent(app))
		r.Put("/", handleUpdateEvent(app))
		r.Delete("/", handleDeleteEvent(app))
	})

	return r
}

func handleReadLatestEvents(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := app.EventStore.ReadLatest(10)
		WriteJSONResponse(w, err, http.StatusOK, data)
	}
}

func handleCreateEvent(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p models.EventRequest
		ReadJSONRequest(w, r, &p)
		// TODO: Validate payload

		data, err := app.EventStore.CreateOne(p)
		WriteJSONResponse(w, err, http.StatusCreated, data)
	}
}

// ----------------------------------------------------------------------------
// Handlers in specific code snippet context
// ----------------------------------------------------------------------------

func handleReadEvent(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := ParseId(w, chi.URLParam(r, "id"))
		data, err := app.EventStore.ReadOne(id)
		WriteJSONResponse(w, err, http.StatusOK, data)
	}
}

func handleUpdateEvent(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := ParseId(w, chi.URLParam(r, "id"))
		var p models.EventRequest
		ReadJSONRequest(w, r, &p)
		// TODO: Validate payload

		data, err := app.EventStore.UpdateOne(id, p)
		WriteJSONResponse(w, err, http.StatusOK, data)
	}
}

func handleDeleteEvent(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := ParseId(w, chi.URLParam(r, "id"))
		data, err := app.EventStore.DeleteOne(id)
		WriteJSONResponse(w, err, http.StatusNoContent, data)
	}
}
