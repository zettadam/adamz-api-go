package web

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
)

type ApiError struct {
	Status  int               `json:"status"`
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors"`
}

func (e ApiError) Error() string {
	return fmt.Sprintf("API error: %d", e.Status)
}

func ParseId(w http.ResponseWriter, input string) int64 {
	id, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		msg := "Bad request"
		slog.Error(msg, slog.String("err", err.Error()))
		writeJSON(w, http.StatusBadRequest, ApiError{
			Message: msg + ". Unable to parse URL path parameter (id)",
		})
	}
	return id
}

func ReadJSONRequest(w http.ResponseWriter, r *http.Request, data any) {
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		msg := "Bad request"
		slog.Error(msg, slog.String("err", err.Error()))
		writeJSON(w, http.StatusBadRequest, ApiError{
			Message: msg + ". Invalid JSON request body",
		})
		return
	}
}

func WriteJSONResponse(w http.ResponseWriter, err error, statusCode int, data any) {
	if err != nil {
		msg := "Internal server error"
		slog.Error(msg, slog.String("error", err.Error()))
		writeJSON(w, http.StatusInternalServerError, ApiError{
			Message: msg,
		})
		return
	}

	writeJSON(w, statusCode, data)
	return
}

func writeJSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
