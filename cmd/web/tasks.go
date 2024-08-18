package web

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/zettadam/adamz-api-go/internal/config"
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
		WriteResponse(w, data, err)
	}
}

func handleCreateTask(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		msg := "CreateTask"
		fmt.Fprint(w, msg)
	}
}

// ----------------------------------------------------------------------------
// Handlers in specific task context
// ----------------------------------------------------------------------------

func handleReadTask(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := ParseId(w, chi.URLParam(r, "id"))
		data, err := app.TaskStore.ReadOne(id)
		WriteResponse(w, data, err)
	}
}

func handleUpdateTask(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		msg := "UpdateTask"
		fmt.Fprint(w, msg)
	}
}

func handleDeleteTask(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		msg := "DeleteTask"
		fmt.Fprint(w, msg)
	}
}
