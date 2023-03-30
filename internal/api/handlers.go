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
	var userWord model.Word
	return func(w http.ResponseWriter, r *http.Request) {
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

			//var userId int64
			//
			//jwtCookies, _ := r.Cookie("jwt")
			//if jwtCookies != nil {
			//	authClaims, err := users.ParseToken(jwtCookies.Value)
			//	if err != nil {
			//		return
			//	}
			//	userId = authClaims.UserId
			//}

			respWord, err := wd.CreateWord(&userWord)
			if err != nil {
				errorHandler(w, http.StatusUnauthorized, err)
				log.Println("can't create word, lox", err)
				return
			}

			//result, err := json.Marshal(respWord)
			//if err != nil {
			//	log.Println("can't marshal data from createWord", err)
			//	return
			//}
			//
			//log.Println(string(result))
			respondHandler(w, http.StatusCreated, respWord)
		}
	}
}

func GetWordsHandler(wd *word.Word) http.HandlerFunc {
	var userWord model.Word
	return func(w http.ResponseWriter, r *http.Request) {

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

			//var userId int64
			//
			//jwtCookies, _ := r.Cookie("jwt")
			//if jwtCookies != nil {
			//	authClaims, err := users.ParseToken(jwtCookies.Value)
			//	if err != nil {
			//		return
			//	}
			//	userId = authClaims.UserId
			//}

			words, err := wd.GetWordsByUserId(userWord.UserId)
			if err != nil {
				errorHandler(w, http.StatusUnauthorized, err)
				log.Println("can't get words, lox", err)
				return
			}

			//result, err := json.Marshal(words)
			//if err != nil {
			//	log.Println("can't marshal get words, lox", err)
			//	return
			//}
			//
			//fmt.Println(string(result))

			respondHandler(w, http.StatusOK, words)
		}
	}
}

// todo: deleting issue with success response, if another user try to delete word, but word stay in db

func DeleteWordHandler(wd *word.Word) http.HandlerFunc {
	var userWord model.Word
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			data, err := io.ReadAll(r.Body)
			if err != nil {
				log.Println("can't read data from delete word", err)
				return
			}

			defer r.Body.Close()

			if err := json.Unmarshal(data, &userWord); err != nil {
				log.Println("can't unmarshal data from getWords", err)
				return
			}

			//var userId int64 ------> when would be implement front
			//var email string ------> when would be implement front

			//jwtCookies, _ := r.Cookie("jwt")
			//if jwtCookies != nil {
			//	authClaims, err := users.ParseToken(jwtCookies.Value)
			//	if err != nil {
			//		return
			//	}
			//	email = authClaims.Email ----------> when would be implement front
			//	userId = authClaims.UserId
			//}

			fmt.Println(userWord.UserId)
			fmt.Println(userWord.Word)
			if err := wd.DeleteWordByWord(userWord.Word, userWord.UserId); err != nil {
				errorHandler(w, http.StatusNotFound, err)
				return
			}

			respondHandler(w, http.StatusOK, fmt.Sprintf("deleting word - %s -  success", userWord.Word))
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
