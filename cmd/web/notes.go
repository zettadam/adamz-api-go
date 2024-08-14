package web

import (
	"fmt"
	"net/http"
)

func ReadLatestNotes(w http.ResponseWriter, r *http.Request) {
	msg := "ReadLatestNotes"
	fmt.Fprint(w, msg)
}

func ReadNotes(w http.ResponseWriter, r *http.Request) {
	msg := "ReadNotes"
	fmt.Fprint(w, msg)
}

func ReadNoteDetail(w http.ResponseWriter, r *http.Request) {
	msg := "ReadPostDetail"
	fmt.Fprint(w, msg)
}

func CreateNote(w http.ResponseWriter, r *http.Request) {
	msg := "CreateNote"
	fmt.Fprint(w, msg)
}

func UpdateNote(w http.ResponseWriter, r *http.Request) {
	msg := "UpdateNote"
	fmt.Fprint(w, msg)
}

func DeleteNote(w http.ResponseWriter, r *http.Request) {
	msg := "DeleteNote"
	fmt.Fprint(w, msg)
}
