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
	Code   string   `json:"code,omitempty"`
	Title  string   `json:"title,omitempty"`
	Detail string   `json:"detail,omitempty"`
	Meta   struct{} `json:"meta,omitempty"`
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	msg := "Hello, world!"
	fmt.Fprint(w, msg)
}
