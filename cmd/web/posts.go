package web

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Post struct {
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	Body        string `json:"body"`
	PublishedAt string `json:"publishedAt"`
}

func ReadLatestPosts(w http.ResponseWriter, r *http.Request) {
	msg := "ReadLatestPosts"
	fmt.Fprint(w, msg)
}

func ReadPosts(w http.ResponseWriter, r *http.Request) {
	msg := "ReadPosts"
	fmt.Fprint(w, msg)
}

func ReadPostDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	post := Post{
		Title:       "First post",
		Slug:        "first-post",
		Body:        "<p>This is my first post.</p>",
		PublishedAt: "2024-08-15T12:34:00Z",
	}
	json.NewEncoder(w).Encode(post)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post Post
	json.NewDecoder(r.Body).Decode(&post)

	fmt.Fprintf(w, "Created post %#v", post)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	id := GetPathParams(r, 0)
	fmt.Fprintf(w, "UpdatePost (%s)", id)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	id := GetPathParams(r, 0)
	fmt.Fprintf(w, "DeletePost (%s)", id)
}
