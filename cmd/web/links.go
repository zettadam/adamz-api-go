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

	r.Get("/", handleReadLatestLinks)
	r.Post("/", handleCreateLink)

	r.Route("/{id}", func(r chi.Router) {
		r.Use(linkCtx)
		r.Get("/", handleReadLink)
		r.Put("/", handleUpdateLink)
		r.Delete("/", handleDeleteLink)
	})

	return r
}

func handleReadLatestLinks(w http.ResponseWriter, r *http.Request) {
	msg := "ReadLatestLinks"
	fmt.Fprint(w, msg)
}

func handleCreateLink(w http.ResponseWriter, r *http.Request) {
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

func handleReadLink(w http.ResponseWriter, r *http.Request) {
	msg := "ReadLink"
	fmt.Fprint(w, msg)
}

func handleUpdateLink(w http.ResponseWriter, r *http.Request) {
	msg := "UpdateLink"
	fmt.Fprint(w, msg)
}

func handleDeleteLink(w http.ResponseWriter, r *http.Request) {
	msg := "DeleteLink"
	fmt.Fprint(w, msg)
}
