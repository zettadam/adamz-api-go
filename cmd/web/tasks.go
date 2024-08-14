package web

import (
	"fmt"
	"net/http"
)

func ReadLatestTasks(w http.ResponseWriter, r *http.Request) {
	msg := "ReadLatestTasks"
	fmt.Fprint(w, msg)
}

func ReadTasks(w http.ResponseWriter, r *http.Request) {
	msg := "ReadTasks"
	fmt.Fprint(w, msg)
}

func ReadTaskDetail(w http.ResponseWriter, r *http.Request) {
	msg := "ReadtaskDetail"
	fmt.Fprint(w, msg)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	msg := "CreateTask"
	fmt.Fprint(w, msg)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	msg := "UpdateTask"
	fmt.Fprint(w, msg)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	msg := "DeleteTask"
	fmt.Fprint(w, msg)
}
