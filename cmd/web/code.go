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

func CodeSnippetsRouter(app *config.Application) http.Handler {
	r := chi.NewRouter()

	r.Get("/", handleReadLatestCodeSnippets(app))
	r.Post("/", handleCreateCodeSnippet(app))

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", handleReadCodeSnippet(app))
		r.Put("/", handleUpdateCodeSnippet(app))
		r.Delete("/", handleDeleteCodeSnippet(app))
	})

	return r
}

func handleReadLatestCodeSnippets(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := app.CodeSnippetStore.ReadLatest(10)
		if err != nil {
			slog.Error("Error fetching code snippets", err)
		}
		json.NewEncoder(w).Encode(data)
	}
}

func handleCreateCodeSnippet(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		msg := "CodeSnippets: CreateOne"
		fmt.Fprint(w, msg)
	}
}

// ----------------------------------------------------------------------------
// Handlers in specific code snippet context
// ----------------------------------------------------------------------------

func handleReadCodeSnippet(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_id := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(_id, 10, 64)
		if err != nil {
			slog.Error("Unable to convert parameter to int64", id, err)
		}

		data, err := app.CodeSnippetStore.ReadOne(id)
		if err != nil {
			slog.Error("Error fetching code snippet",
				slog.String("id", _id),
				slog.Any("err", err),
			)
		}
		json.NewEncoder(w).Encode(data)
	}
}

func handleUpdateCodeSnippet(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		msg := "CodeSnippets: UpdateOne"
		fmt.Fprint(w, msg)
	}
}

func handleDeleteCodeSnippet(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		msg := "CodeSnippets: DeleteOne"
		fmt.Fprint(w, msg)
	}
}
