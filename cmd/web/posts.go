package web

import (
	"fmt"
	"net/http"
)

func ReadLatestPosts(w http.ResponseWriter, r *http.Request) {
	msg := "LatestPosts"
	fmt.Fprint(w, msg)
}

func ReadPosts(w http.ResponseWriter, r *http.Request) {
	msg := "ReadPosts"
	fmt.Fprint(w, msg)
}

func ReadPostDetail(w http.ResponseWriter, r *http.Request) {
	id := GetPathParams(r, 0)
	fmt.Fprintf(w, "PostDetail (%s)", id)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	msg := "CreatePost"
	fmt.Fprint(w, msg)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	id := GetPathParams(r, 0)
	fmt.Fprintf(w, "UpdatePost (%s)", id)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	id := GetPathParams(r, 0)
	fmt.Fprintf(w, "DeletePost (%s)", id)
}
