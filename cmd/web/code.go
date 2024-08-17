package web

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/zettadam/adamz-api-go/internal/config"
)

func CodeSnippetsRouter(app *config.Application) http.Handler {
	r := chi.NewRouter()

	r.Get("/", handleReadLatestCodeSnippets(app))
	r.Post("/", handleCreateCodeSnippet(app))

	r.Route("/{snippetID}", func(r chi.Router) {
		r.Get("/", handleReadCodeSnippet(app))
		r.Put("/", handleUpdateCodeSnippet(app))
		r.Delete("/", handleDeleteCodeSnippet(app))
	})

	return r
}

func handleReadLatestCodeSnippets(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		msg := "CodeSnippets: ReadLatest"
		fmt.Fprint(w, msg)
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

func codeSnippetCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		slog.Info("URL params", slog.Any("id", id))

		if id != "" {
			// get post by ID
		} else {
			fmt.Fprint(w, "Code snippet not found")
			return
		}

		ctx := context.WithValue(r.Context(), "id", id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func handleReadCodeSnippet(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		msg := "Code Snippets: ReadOne"
		fmt.Fprint(w, msg)
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
