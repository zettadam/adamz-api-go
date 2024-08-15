package web

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func LinksRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", readLatestLinks)
	r.Post("/", createLink)

	r.Route("/{id}", func(r chi.Router) {
		r.Use(linkCtx)
		r.Get("/", readLink)
		r.Put("/", updateLink)
		r.Delete("/", deleteLink)
	})

	return r
}

func readLatestLinks(w http.ResponseWriter, r *http.Request) {
	msg := "ReadLatestLinks"
	fmt.Fprint(w, msg)
}

func createLink(w http.ResponseWriter, r *http.Request) {
	msg := "CreateLink"
	fmt.Fprint(w, msg)
}

// ----------------------------------------------------------------------------
// Handlers in specific link context
// ----------------------------------------------------------------------------

func linkCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		slog.Info("URL params", slog.Any("id", id))

		if id != "" {
			// get post by ID
		} else {
			fmt.Fprint(w, "Link not found")
			return
		}

		ctx := context.WithValue(r.Context(), "id", id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func readLink(w http.ResponseWriter, r *http.Request) {
	msg := "ReadLink"
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
