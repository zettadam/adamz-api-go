package web

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

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
		if err != nil {
			slog.Error("Error fetching tasks", err)
		}
		json.NewEncoder(w).Encode(data)
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
		_id := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(_id, 10, 64)
		if err != nil {
			slog.Error("Unable to convert parameter to int64", id, err)
		}

		data, err := app.TaskStore.ReadOne(id)
		if err != nil {
			slog.Error("Error fetching task",
				slog.String("id", _id),
				slog.Any("err", err),
			)
		}
		json.NewEncoder(w).Encode(data)
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
