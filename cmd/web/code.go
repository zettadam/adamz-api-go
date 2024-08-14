package web

import (
	"fmt"
	"net/http"
)

func ReadLatestCodeSnippets(w http.ResponseWriter, r *http.Request) {
	msg := "ReadLatestCodeSnippets"
	fmt.Fprint(w, msg)
}

func ReadCodeSnippets(w http.ResponseWriter, r *http.Request) {
	msg := "ReadCodeSnippets"
	fmt.Fprint(w, msg)
}

func ReadCodeSnippetDetail(w http.ResponseWriter, r *http.Request) {
	msg := "ReadCodeSnippetDetail"
	fmt.Fprint(w, msg)
}

func CreateCodeSnippet(w http.ResponseWriter, r *http.Request) {
	msg := "CreateCodeSnippet"
	fmt.Fprint(w, msg)
}

func UpdateCodeSnippet(w http.ResponseWriter, r *http.Request) {
	msg := "UpdateCodeSnippet"
	fmt.Fprint(w, msg)
}

func DeleteCodeSnippet(w http.ResponseWriter, r *http.Request) {
	msg := "DeleteCodeSnippet"
	fmt.Fprint(w, msg)
}
