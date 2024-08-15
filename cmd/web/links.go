package web

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func LinksRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", readLatestLinks)
	r.Route("/{linkID}", func(r chi.Router) {
		r.Get("/", readLinkDetail)
		r.Post("/", createLink)
		r.Put("/", updateLink)
		r.Delete("/", deleteLink)
	})

	return r
}

func readLatestLinks(w http.ResponseWriter, r *http.Request) {
	msg := "ReadLatestLinks"
	fmt.Fprint(w, msg)
}

func readLinkDetail(w http.ResponseWriter, r *http.Request) {
	msg := "ReadLinkDetail"
	fmt.Fprint(w, msg)
}

func createLink(w http.ResponseWriter, r *http.Request) {
	msg := "CreateLink"
	fmt.Fprint(w, msg)
}

func updateLink(w http.ResponseWriter, r *http.Request) {
	msg := "UpdateLink"
	fmt.Fprint(w, msg)
}

func deleteLink(w http.ResponseWriter, r *http.Request) {
	msg := "DeleteLink"
	fmt.Fprint(w, msg)
}
