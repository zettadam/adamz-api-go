package web

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/zettadam/adamz-api-go/internal/config"
	"github.com/zettadam/adamz-api-go/internal/models"
)

func PostsRouter(app *config.Application) http.Handler {
	r := chi.NewRouter()

	r.Get("/", handleReadLatestPosts(app))
	r.Post("/", handleCreatePost(app))

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", handleReadPost(app))
		r.Put("/", handleUpdatePost(app))
		r.Delete("/", handleDeletePost(app))
	})

	return r
}

func handleReadLatestPosts(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := app.PostStore.ReadLatest(10)
		if err != nil {
			slog.Error("Error fetching latest posts", err)
		}
		json.NewEncoder(w).Encode(data)
	}
}

func handleCreatePost(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var post *models.Post
		json.NewDecoder(r.Body).Decode(&post)

		fmt.Fprintf(w, "Created post %#v", post)
	}
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

func handleReadPost(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.Context().Value("id")

		fmt.Fprintf(w, "ReadPost (%s)", id)
	}
}

func handleUpdatePost(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.Context().Value("id")

		fmt.Fprintf(w, "UpdatePost (%s)", id)
	}
}

func handleDeletePost(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.Context().Value("id")

		fmt.Fprintf(w, "DeletePost (%s)", id)
	}
}
