package web

import (
	"fmt"
	"net/http"
)

type JSONResponse struct {
	Status string `json:"status"`
	Data   any    `json:"data"`
}

type JSONError struct {
	Status int      `json:"status"`
	Code   string   `json:"code omitable"`
	Title  string   `json:"title omitable"`
	Detail string   `json:"detail omitable"`
	Meta   struct{} `json:"meta omitable"`
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	msg := "Hello, world!"
	fmt.Fprint(w, msg)
}
