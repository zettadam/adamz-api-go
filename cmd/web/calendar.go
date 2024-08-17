package web

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/zettadam/adamz-api-go/internal/config"
)

func CalendarRouter(app *config.Application) http.Handler {
	r := chi.NewRouter()

	r.Get("/", handleReadLatestCalendar(app))

	return r
}

func handleReadLatestCalendar(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		msg := "Calendar"
		fmt.Fprint(w, msg)
	}
}
