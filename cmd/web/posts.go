package web

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/zettadam/adamz-api-go/internal/models"
)

func PostsRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", handleReadLatestPosts)
	r.Post("/", handleCreatePost)

	r.Route("/{id}", func(r chi.Router) {
		r.Use(postCtx)
		r.Get("/", handleReadPost)
		r.Put("/", handleUpdatePost)
		r.Delete("/", handleDeletePost)
	})

	return r
}

func handleReadLatestPosts(w http.ResponseWriter, r *http.Request) {
	msg := "ReadLatestPosts"
	fmt.Fprint(w, msg)
}

func handleCreatePost(w http.ResponseWriter, r *http.Request) {
	var post *models.Post
	json.NewDecoder(r.Body).Decode(&post)

	fmt.Fprintf(w, "Created post %#v", post)
}

// ----------------------------------------------------------------------------
// Handlers in specific post context
// ----------------------------------------------------------------------------

func postCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		slog.Info("URL params", slog.Any("id", id))

		if id != "" {
			// get post by ID
		} else {
			fmt.Fprint(w, "Post not found")
			return
		}

		ctx := context.WithValue(r.Context(), "id", id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func handleReadPost(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id")

	fmt.Fprintf(w, "ReadPost (%s)", id)
}

func handleUpdatePost(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id")

	fmt.Fprintf(w, "UpdatePost (%s)", id)
}

func handleDeletePost(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id")

	fmt.Fprintf(w, "DeletePost (%s)", id)
}
