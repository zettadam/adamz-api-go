package web

import (
	"fmt"
	"net/http"
)

func ReadLatestLinks(w http.ResponseWriter, r *http.Request) {
	msg := "ReadLatestLinks"
	fmt.Fprint(w, msg)
}

func ReadLinks(w http.ResponseWriter, r *http.Request) {
	msg := "ReadLinks"
	fmt.Fprint(w, msg)
}

func ReadLinkDetail(w http.ResponseWriter, r *http.Request) {
	msg := "ReadLinkDetail"
	fmt.Fprint(w, msg)
}

func CreateLink(w http.ResponseWriter, r *http.Request) {
	msg := "CreateLink"
	fmt.Fprint(w, msg)
}

func UpdateLink(w http.ResponseWriter, r *http.Request) {
	msg := "UpdateLink"
	fmt.Fprint(w, msg)
}

func DeleteLink(w http.ResponseWriter, r *http.Request) {
	msg := "DeleteLink"
	fmt.Fprint(w, msg)
}
