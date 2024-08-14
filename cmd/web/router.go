package web

import (
	"context"
	"net/http"
	"regexp"
	"strings"

	"github.com/zettadam/adamz-api-go/internal/config"
)

type route struct {
	method  string
	regex   *regexp.Regexp
	handler http.HandlerFunc
}

type ctxKey struct{}

var routes = []route{
	register("GET", "/", Home),

	// Posts
	register("GET", "/posts", ReadLatestPosts),
	register("POST", "/posts/new", CreatePost),
	register("GET", "/posts/([^/]+)", ReadPostDetail),
	register("PUT", "/posts/([^/]+)", UpdatePost),
	register("DELETE", "/posts/([^/]+)", DeletePost),

	// Notes
	register("GET", "/notes", ReadLatestNotes),
	register("POST", "/notes/new", CreateNote),
	register("GET", "/notes/([^/]+)", ReadNoteDetail),
	register("PUT", "/notes/([^/]+)", UpdateNote),
	register("DELETE", "/notes/([^/])", DeleteNote),

	// Code Snippets
	register("GET", "/code", ReadLatestCodeSnippets),
	register("POST", "/code/new", CreateCodeSnippet),
	register("GET", "/code/([^/]+)", ReadCodeSnippetDetail),
	register("PUT", "/code/([^/]+)", UpdateCodeSnippet),
	register("DELETE", "/code/([^/]+)", DeleteCodeSnippet),

	// Links
	register("GET", "/links", ReadLatestLinks),
	register("POST", "/links/new", CreateLink),
	register("GET", "/links/([^/]+)", ReadLinkDetail),
	register("PUT", "/links/([^/]+)", UpdateLink),
	register("DELETE", "/links/([^/]+)", DeleteLink),

	// Tasks
	register("GET", "/tasks", ReadLatestTasks),
	register("POST", "/tasks/new", CreateTask),
	register("GET", "/tasks/([^/]+)", ReadTaskDetail),
	register("PUT", "/tasks/([^/]+)", UpdateTask),
	register("DELETE", "/tasks/([^/]+)", DeleteTask),

	// Calendar
	register("GET", "/calendar", Calendar),
}

func register(method, pattern string, handler http.HandlerFunc) route {
	return route{method, regexp.MustCompile("^" + pattern + "$"), handler}
}

func Router(app *config.Application) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var allow []string

		app.InfoLog.Printf("Request: %s", r.URL)

		for _, route := range routes {
			matches := route.regex.FindStringSubmatch(r.URL.Path)
			if len(matches) > 0 {
				if r.Method != route.method {
					allow = append(allow, route.method)
					continue
				}
				ctx := context.WithValue(r.Context(), ctxKey{}, matches[1:])
				route.handler(w, r.WithContext(ctx))
				return
			}
		}

		if len(allow) > 0 {
			w.Header().Set("Allow", strings.Join(allow, ", "))
			app.ErrorLog.Printf("Method %s not allowed", r.Method)
			http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)

			return
		}

		app.ErrorLog.Printf("Route not found: %s", r.URL)
		http.NotFound(w, r)
	})
}

func getField(r *http.Request, index int) string {
	fields := r.Context().Value(ctxKey{}).([]string)
	return fields[index]
}
