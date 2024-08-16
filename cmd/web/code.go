package web

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func CodeSnippetsRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", handleReadLatestCodeSnippets)
	r.Post("/", handleCreateCodeSnippet)

	r.Route("/{snippetID}", func(r chi.Router) {
		r.Use(codeSnippetCtx)
		r.Get("/", handleReadCodeSnippet)
		r.Put("/", handleUpdateCodeSnippet)
		r.Delete("/", handleDeleteCodeSnippet)
	})

	return r
}

func handleReadLatestCodeSnippets(w http.ResponseWriter, r *http.Request) {
	msg := "CodeSnippets: ReadLatest"
	fmt.Fprint(w, msg)
}

func handleCreateCodeSnippet(w http.ResponseWriter, r *http.Request) {
	msg := "CodeSnippets: CreateOne"
	fmt.Fprint(w, msg)
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

func handleReadCodeSnippet(w http.ResponseWriter, r *http.Request) {
	msg := "Code Snippets: ReadOne"
	fmt.Fprint(w, msg)
}

func handleUpdateCodeSnippet(w http.ResponseWriter, r *http.Request) {
	msg := "CodeSnippets: UpdateOne"
	fmt.Fprint(w, msg)
}

func handleDeleteCodeSnippet(w http.ResponseWriter, r *http.Request) {
	msg := "CodeSnippets: DeleteOne"
	fmt.Fprint(w, msg)
}
