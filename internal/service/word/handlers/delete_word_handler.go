package handlers

import (
	"drunklish/internal/model"
	"drunklish/internal/pkg/httputils"
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
				httputils.WriteErrorResponse(w, fmt.Errorf("%w: %v", httputils.ReadBodyError, err))
				return
			}

			defer r.Body.Close()

			if err := json.Unmarshal(data, &userWord); err != nil {
				httputils.WriteErrorResponse(w, fmt.Errorf("%w: %v", httputils.UnmarshalError, err))
				log.Println(err)
				return
			}

			var userId int64

			jwtCookies, err := r.Cookie("jwt")
			if err != nil {
				return
			}

			if jwtCookies == nil {
				httputils.WriteErrorResponse(w, fmt.Errorf("%w: %v", httputils.ErrValidation, err))
				return
			}

			authClaims, err := token.ParseToken(jwtCookies.Value)
			if err != nil {
				return
			}
			userId = authClaims.UserId

			if err := wd.DeleteWordByWord(userWord.Word, userId); err != nil {
				httputils.WriteErrorResponse(w, err)
				return
			}

			httputils.WriteSuccessResponse(w, http.StatusOK, fmt.Sprintf("deleting word - %s -  success", userWord.Word))
		}
	}
}
