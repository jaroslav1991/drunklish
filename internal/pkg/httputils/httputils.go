package httputils

import (
	"encoding/json"
	"errors"
	"go.uber.org/zap"
	"net/http"
)

func WriteErrorResponse(logger *zap.Logger, w http.ResponseWriter, r *http.Request, err error) {
	response := map[string]any{"error": "internal server error"}
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

	WriteSuccessResponse(logger, w, r, code, response)
}

func WriteSuccessResponse(logger *zap.Logger, w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if data != nil {
		response, err := json.Marshal(data)
		if err != nil {
			return
		}
		w.Write(response)

		if code > 199 && code <= 299 {
			logger.Debug("success fetch", zap.String("path", r.URL.Path), zap.Int("code", code))
		} else {
			logger.Error("failed fetch", zap.String("path", r.URL.Path), zap.Int("code", code), zap.String("error", string(response)))
		}
	}
}
