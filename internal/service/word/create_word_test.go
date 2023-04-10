package word

import (
	"drunklish/internal/pkg/httputils"
	"drunklish/internal/service/auth/token"
	"drunklish/internal/service/word/dto"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWord_CreateWord_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	wordForCreateWord := dto.CreateWordRequest{
		Word:      "boogaga",
		Translate: "смешняшка",
		Token:     "qwerty123",
	}
	wordFromCreated := dto.ResponseFromCreateWord{
		Word:      "boogaga",
		Translate: "смешняшка",
	}

	repository := NewMockRepository(ctrl)
	repository.EXPECT().Create(
		wordForCreateWord.Word,
		wordForCreateWord.Translate,
		wordForCreateWord.CreatedAt,
		int64(1)).Return(&wordFromCreated, nil)

	service := NewWordService(repository)
	service.parseTokenFn = func(tokenString string) (*token.AuthClaims, error) {
		return &token.AuthClaims{UserId: int64(1)}, nil
	}

	actualWord, err := service.CreateWord(wordForCreateWord)
	assert.NoError(t, err)

	assert.Equal(t, &dto.ResponseFromCreateWord{
		Word:      "boogaga",
		Translate: "смешняшка",
	}, actualWord)
}

func TestWord_CreateWord_NegativeFailCreateWord(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	wordForCreateWord := dto.CreateWordRequest{
		Word:      "boogaga",
		Translate: "смешняшка",
		Token:     "qwerty123",
	}

	repository := NewMockRepository(ctrl)
	repository.EXPECT().Create(
		wordForCreateWord.Word,
		wordForCreateWord.Translate,
		wordForCreateWord.CreatedAt,
		int64(1)).Return(nil, errors.New("fail create word"))

	service := NewWordService(repository)
	service.parseTokenFn = func(tokenString string) (*token.AuthClaims, error) {
		return &token.AuthClaims{UserId: int64(1)}, nil
	}

	_, err := service.CreateWord(wordForCreateWord)
	assert.Error(t, err)
}

func TestWord_CreateWord_NegativeFailLengthWord(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	wordForCreateWord := dto.CreateWordRequest{
		Word:      "",
		Translate: "смешняшка",
		Token:     "qwerty123",
	}

	repository := NewMockRepository(ctrl)

	service := NewWordService(repository)
	service.parseTokenFn = func(tokenString string) (*token.AuthClaims, error) {
		return &token.AuthClaims{UserId: int64(1)}, nil
	}

	_, err := service.CreateWord(wordForCreateWord)
	assert.ErrorIs(t, err, httputils.ErrValidation)
}

func TestWord_CreateWord_NegativeFailLengthTranslate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	wordForCreateWord := dto.CreateWordRequest{
		Word:      "boogaga",
		Translate: "",
		Token:     "qwerty123",
	}

	repository := NewMockRepository(ctrl)

	service := NewWordService(repository)
	service.parseTokenFn = func(tokenString string) (*token.AuthClaims, error) {
		return &token.AuthClaims{UserId: int64(1)}, nil
	}

	_, err := service.CreateWord(wordForCreateWord)
	assert.ErrorIs(t, err, httputils.ErrValidation)
}

func TestWord_CreateWord_NegativeFailToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	wordForCreateWord := dto.CreateWordRequest{
		Word:      "boogaga",
		Translate: "boogaga",
		Token:     "",
	}

	repository := NewMockRepository(ctrl)

	service := NewWordService(repository)
	service.parseTokenFn = func(tokenString string) (*token.AuthClaims, error) {
		return nil, errors.New("fuck up")
	}

	_, err := service.CreateWord(wordForCreateWord)
	assert.ErrorIs(t, err, httputils.ErrValidation)
}
