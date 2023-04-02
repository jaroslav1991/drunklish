package handlers

import (
	"drunklish/internal/model"
	"drunklish/internal/pkg/httputils"
	"drunklish/internal/service/auth/dto"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type SignUp interface {
	SignUp(req dto.SignUpRequest) (*model.User, error)
}

func SignUpHandler(a SignUp) http.HandlerFunc {
	var user dto.SignUpRequest
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			data, err := io.ReadAll(r.Body)
			if err != nil {
				httputils.WriteErrorResponse(w, fmt.Errorf("%w: %v", httputils.ReadBodyError, err))
				return
			}

			defer r.Body.Close()

			if err := json.Unmarshal(data, &user); err != nil {
				httputils.WriteErrorResponse(w, fmt.Errorf("%w: %v", httputils.UnmarshalError, err))
				return
			}

			_, err = a.SignUp(user)
			if err != nil {
				httputils.WriteErrorResponse(w, err)
				return
			}

			httputils.WriteSuccessResponse(w, http.StatusOK, map[string]string{"INFO": "success created"})
		}
	}
}
