package web

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

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

func handleReadPost(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_id := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(_id, 10, 64)
		if err != nil {
			slog.Error("Unable to convert parameter to int64", id, err)
		}

		data, err := app.PostStore.ReadOne(id)
		if err != nil {
			slog.Error("Error fetching post",
				slog.String("id", _id),
				slog.Any("err", err),
			)
		}
		json.NewEncoder(w).Encode(data)
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
