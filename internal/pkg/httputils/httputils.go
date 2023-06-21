package httputils

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func WriteErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	response := map[string]any{"error": "internal server error", "path": r.URL.Path}
	code := http.StatusInternalServerError

	if errors.Is(err, ErrValidation) {
		code = http.StatusBadRequest
		response["error"] = err.Error()
	}
	if errors.Is(err, ErrExistEmail) {
		code = http.StatusUnprocessableEntity
		response["error"] = err.Error()
	}

	if errors.Is(err, UnmarshalError) {
		code = http.StatusBadRequest
		response["error"] = err.Error()
	}

	if errors.Is(err, ReadBodyError) {
		code = http.StatusBadRequest
		response["error"] = err.Error()
	}

	if errors.Is(err, ErrWordNotExist) {
		code = http.StatusForbidden
		response["error"] = err.Error()
	}

	log.Println(err, r.URL, code)
	WriteSuccessResponse(w, r, code, response)
}

func WriteSuccessResponse(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if data != nil {
		response, err := json.Marshal(data)
		if err != nil {
			return
		}
		w.Write(response)
	}
}
