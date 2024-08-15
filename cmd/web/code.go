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

	r.Get("/", readLatestCodeSnippets)
	r.Post("/", createCodeSnippet)

	r.Route("/{snippetID}", func(r chi.Router) {
		r.Use(codeSnippetCtx)
		r.Get("/", readCodeSnippet)
		r.Put("/", updateCodeSnippet)
		r.Delete("/", deleteCodeSnippet)
	})

	return r
}

func readLatestCodeSnippets(w http.ResponseWriter, r *http.Request) {
	msg := "ReadLatestCodeSnippets"
	fmt.Fprint(w, msg)
}

func createCodeSnippet(w http.ResponseWriter, r *http.Request) {
	msg := "CreateCodeSnippet"
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

func readCodeSnippet(w http.ResponseWriter, r *http.Request) {
	msg := "ReadCodeSnippet"
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
