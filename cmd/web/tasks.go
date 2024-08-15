package web

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func TasksRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", readLatestTasks)
	r.Route("/{taskID}", func(r chi.Router) {
		r.Get("/", readTaskDetail)
		r.Post("/", createTask)
		r.Put("/", updateTask)
		r.Delete("/", deleteTask)
	})

	return r
}

func readLatestTasks(w http.ResponseWriter, r *http.Request) {
	msg := "ReadLatestTasks"
	fmt.Fprint(w, msg)
}

func readTaskDetail(w http.ResponseWriter, r *http.Request) {
	msg := "ReadtaskDetail"
	fmt.Fprint(w, msg)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	msg := "CreateTask"
	fmt.Fprint(w, msg)
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	msg := "UpdateTask"
	fmt.Fprint(w, msg)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	msg := "DeleteTask"
	fmt.Fprint(w, msg)
}
