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
				errorHandler(w, http.StatusUnprocessableEntity, err)
				log.Println("can't sign up user, lox", err)
				return
			}

			respondHandler(w, http.StatusOK, respUser)
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
				errorHandler(w, http.StatusUnauthorized, err)
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

			respondHandler(w, http.StatusOK, respUser)
		}
	}
}

func CreateWordHandler(wd *word.Word) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userWord model.Word
		if r.Method == http.MethodPost {
			data, err := io.ReadAll(r.Body)
			if err != nil {
				log.Println("can't read data from createWord", err)
				return
			}

			defer r.Body.Close()

			if err := json.Unmarshal(data, &userWord); err != nil {
				log.Println("can't unmarshal data from createWord", err)
				return
			}

			respWord, err := wd.CreateWord(&userWord)
			if err != nil {
				log.Println("can't create word, lox", err)
				return
			}

			result, err := json.Marshal(respWord)
			if err != nil {
				log.Println("can't marshal data from createWord", err)
				return
			}

			log.Println(string(result))
		}
	}
}

func GetWordsHandler(wd *word.Word) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userWord model.Word

		if r.Method == http.MethodGet {
			data, err := io.ReadAll(r.Body)
			if err != nil {
				log.Println("can't read data from getWords", err)
				return
			}

			defer r.Body.Close()

			if err := json.Unmarshal(data, &userWord); err != nil {
				log.Println("can't unmarshal data from getWords", err)
				return
			}

			words, err := wd.GetWordsByUserId(userWord.UserId)
			if err != nil {
				log.Println("can't get words, lox", err)
				return
			}

			result, err := json.Marshal(words)
			if err != nil {
				log.Println("can't marshal get words, lox", err)
				return
			}

			fmt.Println(string(result))
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
		fmt.Println(string(response))
	}
}
