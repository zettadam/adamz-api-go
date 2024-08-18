package web

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/zettadam/adamz-api-go/internal/config"
)

type PostRequest struct {
	Title        string   `json:"title"`
	Slug         string   `json:"slug"`
	Abstract     string   `json:"abstract"`
	Significance int      `json:"significance"`
	Body         string   `json:"body"`
	PublishedAt  string   `json:"published_at"`
	Tags         []string `json:"tags"`
}

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
		WriteJSONResponse(w, err, http.StatusOK, data)
	}
}

func handleCreatePost(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var post *PostRequest
		ReadJSONRequest(w, r, &post)
		fmt.Fprintf(w, "Created post %#v", post)
	}
}

// ----------------------------------------------------------------------------
// Handlers in specific post context
// ----------------------------------------------------------------------------

func handleReadPost(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := ParseId(w, chi.URLParam(r, "id"))
		data, err := app.PostStore.ReadOne(id)
		WriteJSONResponse(w, err, http.StatusOK, data)
	}
}

func handleUpdatePost(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := ParseId(w, chi.URLParam(r, "id"))
		fmt.Fprintf(w, "UpdatePost (%s)", id)
	}
}

func handleDeletePost(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := ParseId(w, chi.URLParam(r, "id"))
		data, err := app.PostStore.DeleteOne(id)
		WriteJSONResponse(w, err, http.StatusNoContent, data)
	}
}
