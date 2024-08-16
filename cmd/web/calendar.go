package web

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func CalendarRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", handleReadLatestCalendar)

	return r
}

func handleReadLatestCalendar(w http.ResponseWriter, r *http.Request) {
	msg := "Calendar"
	fmt.Fprint(w, msg)
}
