package web

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NotesRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", readLatestNotes)
	r.Route("/{noteID}", func(r chi.Router) {
		r.Get("/", readNoteDetail)
		r.Post("/", createNote)
		r.Put("/", updateNote)
		r.Delete("/", deleteNote)
	})

	return r
}

func readLatestNotes(w http.ResponseWriter, r *http.Request) {
	msg := "ReadLatestNotes"
	fmt.Fprint(w, msg)
}

func readNoteDetail(w http.ResponseWriter, r *http.Request) {
	msg := "ReadPostDetail"
	fmt.Fprint(w, msg)
}

func createNote(w http.ResponseWriter, r *http.Request) {
	msg := "CreateNote"
	fmt.Fprint(w, msg)
}

func updateNote(w http.ResponseWriter, r *http.Request) {
	msg := "UpdateNote"
	fmt.Fprint(w, msg)
}

func deleteNote(w http.ResponseWriter, r *http.Request) {
	msg := "DeleteNote"
	fmt.Fprint(w, msg)
}
