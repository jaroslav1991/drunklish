package handlers

import (
	"drunklish/internal/model"
	"drunklish/internal/pkg/httputils"
	"drunklish/internal/service/auth/token"
	dto "drunklish/internal/service/word/dto"
	"encoding/json"
	"fmt"
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
				httputils.WriteErrorResponse(w, fmt.Errorf("%w: %v", httputils.ReadBodyError, err))
				return
			}

			defer r.Body.Close()

			if err := json.Unmarshal(data, &userWord); err != nil {
				httputils.WriteErrorResponse(w, fmt.Errorf("%w: %v", httputils.UnmarshalError, err))
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

			respWord, err := wd.CreateWord(dto.CreateWordRequest{
				Word:      userWord.Word,
				Translate: userWord.Translate,
				UserId:    userId,
			})
			if err != nil {
				httputils.WriteErrorResponse(w, err)
				log.Println(err)
				return
			}

			httputils.WriteSuccessResponse(w, http.StatusOK, respWord)
		}
	}
}
