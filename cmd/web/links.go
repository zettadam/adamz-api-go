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

func LinksRouter(app *config.Application) http.Handler {
	r := chi.NewRouter()

	r.Get("/", handleReadLatestLinks(app))
	r.Post("/", handleCreateLink(app))

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", handleReadLink(app))
		r.Put("/", handleUpdateLink(app))
		r.Delete("/", handleDeleteLink(app))
	})

	return r
}

func handleReadLatestLinks(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := app.LinkStore.ReadLatest(10)
		if err != nil {
			slog.Error("Error fetching links", err)
		}
		json.NewEncoder(w).Encode(data)
	}
}

func handleCreateLink(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		msg := "CreateLink"
		fmt.Fprint(w, msg)
	}
}

// ----------------------------------------------------------------------------
// Handlers in specific link context
// ----------------------------------------------------------------------------

func handleReadLink(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_id := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(_id, 10, 64)
		if err != nil {
			slog.Error("Unable to convert parameter to int64", id, err)
		}

		data, err := app.LinkStore.ReadOne(id)
		if err != nil {
			slog.Error("Error fetching link",
				slog.String("id", _id),
				slog.Any("err", err),
			)
		}
		json.NewEncoder(w).Encode(data)
	}
}

func handleUpdateLink(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		msg := "UpdateLink"
		fmt.Fprint(w, msg)
	}
}

func handleDeleteLink(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		msg := "DeleteLink"
		fmt.Fprint(w, msg)
	}
}
