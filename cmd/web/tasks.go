package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/zettadam/adamz-api-go/internal/types"
)

func (app *Application) handleReadLatestTasks(w http.ResponseWriter, r *http.Request) {
	data, err := app.Service.Tasks.ReadLatest(10)
	WriteJSONResponse(w, err, http.StatusOK, data)
}

func (app *Application) handleCreateTask(w http.ResponseWriter, r *http.Request) {
	var p types.TaskRequest
	ReadJSONRequest(w, r, &p)
	// TODO: Validate payload

	data, err := app.Service.Tasks.CreateOne(p)
	WriteJSONResponse(w, err, http.StatusCreated, data)
}

// ----------------------------------------------------------------------------
// Handlers in specific task context
// ----------------------------------------------------------------------------

func (app *Application) handleReadTask(w http.ResponseWriter, r *http.Request) {
	id := ParseId(w, chi.URLParam(r, "id"))

	data, err := app.Service.Tasks.ReadOne(id)
	WriteJSONResponse(w, err, http.StatusOK, data)
}

func (app *Application) handleUpdateTask(w http.ResponseWriter, r *http.Request) {
	id := ParseId(w, chi.URLParam(r, "id"))

	var p types.TaskRequest
	ReadJSONRequest(w, r, &p)
	// TODO: Validate payload

	data, err := app.Service.Tasks.UpdateOne(id, p)
	WriteJSONResponse(w, err, http.StatusOK, data)
}

func (app *Application) handleDeleteTask(w http.ResponseWriter, r *http.Request) {
	id := ParseId(w, chi.URLParam(r, "id"))

	data, err := app.Service.Tasks.DeleteOne(id)
	WriteJSONResponse(w, err, http.StatusNoContent, data)
}
