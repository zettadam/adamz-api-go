package web

import (
	"fmt"
	"net/http"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	msg := "Hello, world!"
	fmt.Fprint(w, msg)
}
