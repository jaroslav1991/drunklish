package handlers

import (
	"drunklish/internal/model"
	"drunklish/internal/service/auth/token"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type DeleteWord interface {
	DeleteWordByWord(word string, userId int64) error
}

func DeleteWordHandler(wd DeleteWord) http.HandlerFunc {
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
				errorHandler(w, http.StatusBadRequest, nil)
				log.Println(err)
				return
			}

			var userId int64

			jwtCookies, err := r.Cookie("jwt")
			if err != nil {
				return
			}

			if jwtCookies == nil {
				respondHandler(w, http.StatusForbidden, map[string]string{"error": "forbidden"})
				log.Println(err)
				return
			}

			authClaims, err := token.ParseToken(jwtCookies.Value)
			if err != nil {
				return
			}
			userId = authClaims.UserId

			if err := wd.DeleteWordByWord(userWord.Word, userId); err != nil {
				errorHandler(w, http.StatusNotFound, err)
				log.Println(err)
				return
			}

			respondHandler(w, http.StatusOK, fmt.Sprintf("deleting word - %s -  success", userWord.Word))
		}
	}
}
