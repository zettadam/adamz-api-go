package web

import (
	"fmt"
	"net/http"

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
		WriteJSONResponse(w, err, http.StatusOK, data)
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
		id := ParseId(w, chi.URLParam(r, "id"))
		data, err := app.CodeSnippetStore.ReadOne(id)
		WriteJSONResponse(w, err, http.StatusOK, data)
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
		id := ParseId(w, chi.URLParam(r, "id"))
		data, err := app.CodeSnippetStore.DeleteOne(id)
		WriteJSONResponse(w, err, http.StatusNoContent, data)
	}
}
