package web

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func CodeSnippetsRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", readLatestCodeSnippets)
	r.Route("/{snippetID}", func(r chi.Router) {
		r.Get("/", readCodeSnippetDetail)
		r.Post("/", createCodeSnippet)
		r.Put("/", updateCodeSnippet)
		r.Delete("/", deleteCodeSnippet)
	})

	return r
}

func readLatestCodeSnippets(w http.ResponseWriter, r *http.Request) {
	msg := "ReadLatestCodeSnippets"
	fmt.Fprint(w, msg)
}

func readCodeSnippetDetail(w http.ResponseWriter, r *http.Request) {
	msg := "ReadCodeSnippetDetail"
	fmt.Fprint(w, msg)
}

func createCodeSnippet(w http.ResponseWriter, r *http.Request) {
	msg := "CreateCodeSnippet"
	fmt.Fprint(w, msg)
}

func updateCodeSnippet(w http.ResponseWriter, r *http.Request) {
	msg := "UpdateCodeSnippet"
	fmt.Fprint(w, msg)
}

func deleteCodeSnippet(w http.ResponseWriter, r *http.Request) {
	msg := "DeleteCodeSnippet"
	fmt.Fprint(w, msg)
}
