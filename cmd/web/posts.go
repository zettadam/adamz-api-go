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
	msg := "PostDetail"
	fmt.Fprint(w, msg)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	msg := "CreatePost"
	fmt.Fprint(w, msg)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	msg := "UpdatePost"
	fmt.Fprint(w, msg)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	msg := "DeletePost"
	fmt.Fprint(w, msg)
}
