package handlers

import (
	"drunklish/internal/model"
	"drunklish/internal/service/auth/token"
	dto "drunklish/internal/service/word/dto"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type CreateWord interface {
	CreateWord(word dto.CreateWordRequest) (*model.Word, error)
}

func CreateWordHandler(wd CreateWord) http.HandlerFunc {
	var userWord dto.CreateWordRequest
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			data, err := io.ReadAll(r.Body)
			if err != nil {
				log.Println("can't read data from createWord", err)
				return
			}

			defer r.Body.Close()

			if err := json.Unmarshal(data, &userWord); err != nil {
				errorHandler(w, http.StatusBadRequest, nil)
				log.Println(err)
				return
			}

			var userId int64

			jwtCookies, _ := r.Cookie("jwt")
			if jwtCookies != nil {
				authClaims, err := token.ParseToken(jwtCookies.Value)
				if err != nil {
					return
				}
				userId = authClaims.UserId
			}

			respWord, err := wd.CreateWord(dto.CreateWordRequest{
				Word:      userWord.Word,
				Translate: userWord.Translate,
				UserId:    userId,
			})
			if err != nil {
				errorHandler(w, http.StatusUnauthorized, err)
				log.Println(err)
				return
			}

			respondHandler(w, http.StatusCreated, respWord)
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
