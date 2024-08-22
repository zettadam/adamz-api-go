package web

import (
	"fmt"
	"net/http"

	chi "github.com/go-chi/chi/v5"
	middleware "github.com/go-chi/chi/v5/middleware"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"

	"github.com/zettadam/adamz-api-go/internal/config"
	"github.com/zettadam/adamz-api-go/internal/stores"
)

type Service struct {
	CodeSnippets *stores.CodeSnippetStore
	Events       *stores.EventStore
	Links        *stores.LinkStore
	Notes        *stores.NoteStore
	Posts        *stores.PostStore
	Tasks        *stores.TaskStore
}

type Application struct {
	Config  *config.Configuration
	Service *Service
}

func (app *Application) SetupRouter() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Root handler
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		msg := "Hello, world!"
		fmt.Fprint(w, msg)
	})

	// Calendar routes
	r.Get("/calendar", app.handleReadLatestEvents)
	r.Post("/calendar", app.handleCreateEvent)
	r.Get("/calendar/{id}", app.handleReadEvent)
	r.Put("/calendar/{id}", app.handleUpdateEvent)
	r.Delete("/calendar/{id}", app.handleDeleteEvent)

	// Code (snippets) routes
	r.Get("/code", app.handleReadLatestCodeSnippets)
	r.Post("/code", app.handleCreateCodeSnippet)
	r.Get("/code/{id}", app.handleReadCodeSnippet)
	r.Put("/code/{id}", app.handleUpdateCodeSnippet)
	r.Delete("/code/{id}", app.handleDeleteCodeSnippet)

	// Links routes
	r.Get("/links", app.handleReadLatestLinks)
	r.Post("/links", app.handleCreateLink)
	r.Get("/links/{id}", app.handleReadLink)
	r.Put("/links/{id}", app.handleUpdateLink)
	r.Delete("/links/{id}", app.handleDeleteLink)

	// Notes routes
	r.Get("/notes", app.handleReadLatestNotes)
	r.Post("/notes", app.handleCreateNote)
	r.Get("/notes/{id}", app.handleReadNote)
	r.Put("/notes/{id}", app.handleUpdateNote)
	r.Delete("/notes/{id}", app.handleDeleteNote)

	// Posts routes
	r.Get("/posts", app.handleReadLatestPosts)
	r.Post("/posts", app.handleCreatePost)
	r.Get("/posts/{id}", app.handleReadPost)
	r.Put("/posts/{id}", app.handleUpdatePost)
	r.Delete("/posts/{id}", app.handleDeletePost)

	// Tasks routes
	r.Get("/tasks", app.handleReadLatestTasks)
	r.Post("/tasks", app.handleCreateTask)
	r.Get("/tasks/{id}", app.handleReadTask)
	r.Put("/tasks/{id}", app.handleUpdateTask)
	r.Delete("/tasks/{id}", app.handleDeleteTask)

	return r
}
