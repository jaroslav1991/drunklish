package word

import (
	"drunklish/internal/service/auth/token"
	"drunklish/internal/service/word/dto"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestWord_CreateListFromUpload(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	requestWords := dto.RequestFromUpload{Words: dto.ResponseWords{Words: dto.ResponseWordList{{
		Word:      "test",
		Translate: "тест",
	}, {
		Word:      "text",
		Translate: "текст",
	}}}}

	responseWords := []dto.ResponseFromCreateWord{{
		Word:      "test",
		Translate: "тест",
	}, {
		Word:      "text",
		Translate: "текст",
	}}

	createAt := time.Now()

	repository := NewMockRepository(ctrl)
	for i := 0; i < len(requestWords.Words.Words); i++ {
		repository.EXPECT().CreateFromUpload(requestWords.Words.Words[i].Word, requestWords.Words.Words[i].Translate, createAt, int64(1)).Return(&responseWords[i], nil)
	}
	service := NewWordService(repository)

	service.parseTokenFn = func(tokenString string) (*token.AuthClaims, error) {
		return &token.AuthClaims{UserId: int64(1)}, nil
	}

	tkn := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsIkVtYWlsIjoicXdlQGdtYWlsLmNvbSIsImV4cCI6MTY4OTMzNjEyMH0.wdi8RB1ufiN67sA-QYl5Wd9BmbLVMsDlVVqmamnIw9k"

	actualWords, err := service.CreateListFromUpload(requestWords, tkn)
	assert.NoError(t, err)

	assert.Equal(t, &[]dto.ResponseFromCreateWord{{
		Word:      "test",
		Translate: "тест",
	}, {
		Word:      "text",
		Translate: "текст",
	}}, actualWords)
}
