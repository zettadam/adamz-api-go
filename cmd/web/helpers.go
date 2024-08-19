package web

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/zettadam/adamz-api-go/internal/stores"
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
		writeJSON(w, http.StatusBadRequest, JSONError{
			Status: http.StatusBadRequest,
			Detail: "Unable to parse URL path parameter (id)",
		})
	}
	return id
}

func ReadJSONRequest(w http.ResponseWriter, r *http.Request, data any) {
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		msg := "Bad request"
		slog.Error(msg, slog.String("err", err.Error()))
		writeJSON(w, http.StatusBadRequest, JSONError{
			Status: http.StatusBadRequest,
			Detail: "Invalid JSON request body",
		})
		return
	}
}

func WriteJSONResponse(w http.ResponseWriter, err error, statusCode int, data any) {
	if stores.IsNotFound(err) != false {
		writeJSON(w, http.StatusOK, JSONResponse{
			Status: "success",
			Data:   nil,
		})
		return
	}
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			slog.Error("DB Error",
				slog.Group("details",
					slog.String("code", pgErr.Code),
					slog.String("error", pgErr.Message),
				),
			)
		} else {
			slog.Error("Other", slog.String("details", err.Error()))
		}

		writeJSON(w, http.StatusInternalServerError, JSONError{
			Status: http.StatusInternalServerError,
			Title:  "Internal Serever Error",
			Detail: "Something went wrong on the server.",
		})
		return
	}

	writeJSON(w, statusCode, JSONResponse{
		Status: "success",
		Data:   data,
	})
	return
}

func writeJSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Add("content-type", "application/vnd.api+json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
