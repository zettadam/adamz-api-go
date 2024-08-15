package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Post struct {
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	Body        string `json:"body"`
	PublishedAt string `json:"publishedAt"`
}

func PostsRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", readLatestPosts)
	r.Route("/{postID}", func(r chi.Router) {
		r.Get("/", readPostDetail)
		r.Post("/", createPost)
		r.Put("/", updatePost)
		r.Delete("/", deletePost)
	})

	return r
}

func readLatestPosts(w http.ResponseWriter, r *http.Request) {
	msg := "ReadLatestPosts"
	fmt.Fprint(w, msg)
}

func readPostDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	post := Post{
		Title:       "First post",
		Slug:        "first-post",
		Body:        "<p>This is my first post.</p>",
		PublishedAt: "2024-08-15T12:34:00Z",
	}
	json.NewEncoder(w).Encode(post)
}

func createPost(w http.ResponseWriter, r *http.Request) {
	var post Post
	json.NewDecoder(r.Body).Decode(&post)

	fmt.Fprintf(w, "Created post %#v", post)
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	id := 123
	fmt.Fprintf(w, "UpdatePost (%d)", id)
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	id := 123
	fmt.Fprintf(w, "DeletePost (%d)", id)
}
