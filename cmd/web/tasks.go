package web

import (
	"context"
	"fmt"
	"log/slog"
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
		msg := "ReadLatestTasks"
		fmt.Fprint(w, msg)
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

func taskCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		slog.Info("URL params", slog.Any("id", id))

		if id != "" {
			// get post by ID
		} else {
			fmt.Fprint(w, "Task not found")
			return
		}

		ctx := context.WithValue(r.Context(), "id", id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func handleReadTask(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		msg := "ReadTask"
		fmt.Fprint(w, msg)
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
