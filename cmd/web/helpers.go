package web

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/zettadam/adamz-api-go/internal/stores"
)

type JSONResponse struct {
	Status string `json:"status"`
	Data   any    `json:"data"`
}

type vError struct {
	field   string
	message string
	tag     string
}

type JSONError struct {
	Status int               `json:"status"`
	Code   string            `json:"code,omitempty"`
	Title  string            `json:"title,omitempty"`
	Detail string            `json:"detail,omitempty"`
	Errors map[string]string `json:"errors,omitempty"`
	Meta   struct{}          `json:"meta,omitempty"`
}

func ParseId(w http.ResponseWriter, input string) int64 {
	id, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		msg := "Bad request"
		slog.Error(msg, slog.String("err", err.Error()))
		writeJSON(w, http.StatusBadRequest, JSONError{
			Status: http.StatusBadRequest,
			Detail: "Unable to parse URL path parameter (id)",
		}, nil)
	}
	return id
}

func ReadJSONRequest(w http.ResponseWriter, r *http.Request, dest any) {
	err := json.NewDecoder(r.Body).Decode(&dest)
	if err != nil {
		msg := "Bad request"
		slog.Error(msg, slog.String("err", err.Error()))
		writeJSON(w, http.StatusBadRequest, JSONError{
			Status: http.StatusBadRequest,
			Detail: "Invalid JSON request body",
		}, nil)
	}
	return
}

func WriteJSONResponse(w http.ResponseWriter, err error, statusCode int, data any) {
	if stores.IsNotFound(err) != false {
		writeJSON(w, http.StatusOK, JSONResponse{
			Status: "success",
			Data:   nil,
		}, nil)
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
		}, nil)
		return
	}

	writeJSON(w, statusCode, JSONResponse{
		Status: "success",
		Data:   data,
	}, nil)
	return
}

func WriteValidationErrors(w http.ResponseWriter, payload validator.ValidationErrors) {
	errs := make(map[string]string)

	for i := range payload {
		errs[payload[i].Field()] = payload[i].Error()
	}

	writeJSON(w, http.StatusExpectationFailed, JSONError{
		Status: http.StatusExpectationFailed,
		Title:  "Validation failed",
		Detail: "We found some errors when looking at the data sent",
		Errors: errs,
	}, nil)
	return
}

func writeJSON(w http.ResponseWriter, statusCode int, data any, headers http.Header) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}
	js = append(js, '\n')

	for k, v := range headers {
		w.Header()[k] = v
	}

	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(statusCode)
	w.Write(js)

	return nil
}
