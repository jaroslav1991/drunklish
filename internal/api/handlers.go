package api

import (
	"drunklish/internal/model"
	"drunklish/internal/service/auth"
	"drunklish/internal/service/word"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func SignUpHandler(a *auth.Auth) http.HandlerFunc {
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
				log.Println("can't unmarshal data from user", err)
				return
			}

			respUser, err := a.SignUp(&user)
			if err != nil {
				log.Println("can't sign up user, lox", err)
				return
			}

			result, err := json.Marshal(respUser)
			if err != nil {
				log.Println("can't marshal data from user", err)
				return
			}
			fmt.Println(string(result))
		}
	}
}

func SignInHandler(a *auth.Auth) http.HandlerFunc {
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
				log.Println("can't unmarshal data from user", err)
				return
			}
			respUser, err := a.SignIn(&user)
			if err != nil {
				log.Println("can't sign in user, lox", err)
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

			result, err := json.Marshal(respUser)
			if err != nil {
				log.Println("can't marshal data from user", err)
				return
			}
			fmt.Println(string(result))
		}
	}
}

func CreateWordHandler(wd *word.Word) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userWord model.Word
		if r.Method == http.MethodPost {
			data, err := io.ReadAll(r.Body)
			if err != nil {
				log.Println("can't read data from word", err)
				return
			}

			defer r.Body.Close()

			if err := json.Unmarshal(data, &userWord); err != nil {
				log.Println("can't unmarshal data from word", err)
				return
			}

			respWord, err := wd.CreateWord(&userWord)
			if err != nil {
				log.Println("can't create word, lox", err)
				return
			}

			result, err := json.Marshal(respWord)
			if err != nil {
				log.Println("can't marshal data from word", err)
				return
			}

			log.Println(string(result))
		}
	}
}
