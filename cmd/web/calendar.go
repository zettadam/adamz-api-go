package web

import (
	"fmt"
	"net/http"
)

func Calendar(w http.ResponseWriter, r *http.Request) {
	msg := "Calendar"
	fmt.Fprint(w, msg)
}
