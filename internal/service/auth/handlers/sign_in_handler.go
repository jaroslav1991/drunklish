package handlers

import (
	"drunklish/internal/model"
	"drunklish/internal/service/auth/dto"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type SignIn interface {
	SignIn(req model.User) (*dto.ResponseUser, error)
}

func SignInHandler(a SignIn) http.HandlerFunc {
	var user model.User
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
			respUser, err := a.SignIn(user)
			if err != nil {
				errorHandler(w, http.StatusUnauthorized, err)
				log.Println(err)
				return
			}

			http.SetCookie(w, &http.Cookie{
				Name:     "jwt",
				Value:    respUser.Token,
				Path:     "/",
				Expires:  time.Now().Add(100 * time.Hour),
				Secure:   true,
				HttpOnly: true,
			})

			respondHandler(w, http.StatusOK, respUser)
		}
	}
}
