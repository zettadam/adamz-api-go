package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/zettadam/adamz-api-go/internal/config"
	"github.com/zettadam/adamz-api-go/internal/models"
)

func TasksRouter(app *config.Application) http.Handler {
	r := chi.NewRouter()

	r.Get("/", handleReadLatestTasks(app))
	r.Post("/", handleCreateTask(app))

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", handleReadTask(app))
		r.Put("/", handleUpdateTask(app))
		r.Delete("/", handleDeleteTask(app))
	})

	return r
}

func handleReadLatestTasks(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := app.TaskStore.ReadLatest(10)
		WriteJSONResponse(w, err, http.StatusOK, data)
	}
}

func handleCreateTask(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p models.TaskRequest
		ReadJSONRequest(w, r, &p)
		// TODO: Validate payload

		data, err := app.TaskStore.CreateOne(p)
		WriteJSONResponse(w, err, http.StatusCreated, data)
	}
}

// ----------------------------------------------------------------------------
// Handlers in specific task context
// ----------------------------------------------------------------------------

func handleReadTask(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := ParseId(w, chi.URLParam(r, "id"))

		data, err := app.TaskStore.ReadOne(id)
		WriteJSONResponse(w, err, http.StatusOK, data)
	}
}

func handleUpdateTask(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := ParseId(w, chi.URLParam(r, "id"))

		var p models.TaskRequest
		ReadJSONRequest(w, r, &p)
		// TODO: Validate payload

		data, err := app.TaskStore.UpdateOne(id, p)
		WriteJSONResponse(w, err, http.StatusOK, data)
	}
}

func handleDeleteTask(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := ParseId(w, chi.URLParam(r, "id"))

		data, err := app.TaskStore.DeleteOne(id)
		WriteJSONResponse(w, err, http.StatusNoContent, data)
	}
}
