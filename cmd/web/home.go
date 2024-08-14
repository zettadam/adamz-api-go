package web

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	msg := "Hello, world!"
	fmt.Fprint(w, msg)
}
