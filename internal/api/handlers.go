package api

import (
	"drunklish/internal/storage"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func SignUpHandler(s *storage.Storage) http.HandlerFunc {
	var user storage.User
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

			respUser, err := s.SignUp(&user)
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

func SignInHandler(s *storage.Storage) http.HandlerFunc {
	var user storage.User
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

			respUser, err := s.SignIn(&user)
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

func CreateWordHandler(s *storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var word storage.Word
		if r.Method == http.MethodPost {
			data, err := io.ReadAll(r.Body)
			if err != nil {
				log.Println("can't read data from word", err)
				return
			}

			defer r.Body.Close()

			if err := json.Unmarshal(data, &word); err != nil {
				log.Println("can't unmarshal data from word", err)
				return
			}

			respWord, err := s.CreateWord(&word)
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
