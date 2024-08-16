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

	r.Get("/", readLatestPosts)
	r.Post("/", createPost)

	r.Route("/{id}", func(r chi.Router) {
		r.Use(postCtx)
		r.Get("/", readPost)
		r.Put("/", updatePost)
		r.Delete("/", deletePost)
	})

	return r
}

func readLatestPosts(w http.ResponseWriter, r *http.Request) {
	msg := "ReadLatestPosts"
	fmt.Fprint(w, msg)
}

func createPost(w http.ResponseWriter, r *http.Request) {
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

func readPost(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id")

	fmt.Fprintf(w, "ReadPost (%s)", id)
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id")

	fmt.Fprintf(w, "UpdatePost (%s)", id)
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id")

	fmt.Fprintf(w, "DeletePost (%s)", id)
}
