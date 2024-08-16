package web

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func TasksRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", handleReadLatestTasks)
	r.Post("/", handleCreateTask)

	r.Route("/{id}", func(r chi.Router) {
		r.Use(taskCtx)
		r.Get("/", handleReadTask)
		r.Put("/", handleUpdateTask)
		r.Delete("/", handleDeleteTask)
	})

	return r
}

func handleReadLatestTasks(w http.ResponseWriter, r *http.Request) {
	msg := "ReadLatestTasks"
	fmt.Fprint(w, msg)
}

func handleCreateTask(w http.ResponseWriter, r *http.Request) {
	msg := "CreateTask"
	fmt.Fprint(w, msg)
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

func handleReadTask(w http.ResponseWriter, r *http.Request) {
	msg := "ReadTask"
	fmt.Fprint(w, msg)
}

func handleUpdateTask(w http.ResponseWriter, r *http.Request) {
	msg := "UpdateTask"
	fmt.Fprint(w, msg)
}

func handleDeleteTask(w http.ResponseWriter, r *http.Request) {
	msg := "DeleteTask"
	fmt.Fprint(w, msg)
}
