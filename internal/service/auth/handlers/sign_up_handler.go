package handlers

import (
	"drunklish/internal/model"
	"drunklish/internal/service/auth/dto"
	"encoding/json"
	"io"
	"log"
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
				log.Println("can't read data from user", err)
				return
			}

			defer r.Body.Close()

			if err := json.Unmarshal(data, &user); err != nil {
				errorHandler(w, http.StatusBadRequest, nil)
				log.Println(err)
				return
			}

			_, err = a.SignUp(user)
			if err != nil {
				errorHandler(w, http.StatusUnprocessableEntity, err)
				log.Println(err)
				return
			}

			respondHandler(w, http.StatusCreated, map[string]string{"INFO": "success created"})
		}
	}
}

func errorHandler(w http.ResponseWriter, code int, err error) {
	respondHandler(w, code, map[string]string{"error": err.Error()})
}

func respondHandler(w http.ResponseWriter, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		response, err := json.Marshal(data)
		if err != nil {
			return
		}
		w.Write(response)
	}
}
