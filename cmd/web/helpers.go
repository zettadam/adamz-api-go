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

func NewApiError(status int, err error) ApiError {
	return ApiError{
		Status:  status,
		Message: err.Error(),
	}
}

func InvalidRequestData(errors map[string]string) ApiError {
	return ApiError{
		Status:  http.StatusUnprocessableEntity,
		Message: "Validation Error",
		Errors:  errors,
	}
}

func InvalidJSON() ApiError {
	return NewApiError(http.StatusBadRequest, fmt.Errorf("Invalid JSON request data"))
}

func ReadJSONRequest(w http.ResponseWriter, r *http.Request, data any) {
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, ApiError{
			Message: "Bad Request. Invalid JSON request body",
		})
	}
}

func WriteJSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

type ApiHandler func(w http.ResponseWriter, r *http.Request) error

func HandleApiErrors(h ApiHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			if apiErr, ok := err.(ApiError); ok {
				WriteJSON(w, apiErr.Status, apiErr)
			} else {
				apiError := map[string]any{
					"status":  http.StatusInternalServerError,
					"message": "Internal Server Error",
				}
				WriteJSON(w, http.StatusInternalServerError, apiError)
			}
			slog.Error("API Error", "err", err.Error(), "path", r.URL.Path)
		}
	}
}

func ParseId(w http.ResponseWriter, input string) int64 {
	id, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		slog.Error("Bad request", slog.String("err", err.Error()))
		WriteJSON(w, http.StatusBadRequest, ApiError{
			Message: "Bad request. Unable to parse URL path parameter (id)",
		})
	}
	return id
}

func WriteResponse(w http.ResponseWriter, data any, err error) {
	if err != nil {
		slog.Error("DB", slog.String("error", err.Error()))
		WriteJSON(w, http.StatusInternalServerError, ApiError{
			Message: "Internal Server Error",
		})
	}

	WriteJSON(w, http.StatusOK, data)
}
