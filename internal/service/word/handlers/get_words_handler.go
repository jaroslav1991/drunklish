package handlers

import (
	"drunklish/internal/model"
	"drunklish/internal/pkg/httputils"
	"drunklish/internal/service/auth/token"
	"drunklish/internal/service/word/dto"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type GetAllWords interface {
	GetWordsByUserId(userId int64) ([]*dto.ResponseWord, error)
}

type GetWordsByTime interface {
	GetWordsByCreatedAt(userId int64, createdAt time.Time) (*model.Word, error)
}

func GetWordsHandler(wd GetAllWords) http.HandlerFunc {
	var userWord model.Word
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodGet {
			data, err := io.ReadAll(r.Body)
			if err != nil {
				httputils.WriteErrorResponse(w, fmt.Errorf("%w: %v", httputils.ReadBodyError, err))
				return
			}

			defer r.Body.Close()

			if err := json.Unmarshal(data, &userWord); err != nil {
				httputils.WriteErrorResponse(w, fmt.Errorf("%w: %v", httputils.UnmarshalError, err))
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

			words, err := wd.GetWordsByUserId(userId)
			if err != nil {
				httputils.WriteErrorResponse(w, err)
				log.Println(err)
				return
			}
			httputils.WriteSuccessResponse(w, http.StatusOK, words)
		}
	}
}
