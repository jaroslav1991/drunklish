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

func TestWord_DeleteWordByWord_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	wordRequest := dto.RequestForDeletingWord{
		Token: "qwerty123",
		Id:    int64(1),
	}

	expected := dto.ResponseFromDeleting{Answer: "deleting success"}

	repository := NewMockRepository(ctrl)
	repository.EXPECT().DeleteWord(int64(1), wordRequest.Id).Return(&expected, nil)

	service := NewWordService(repository)
	service.parseTokenFn = func(tokenString string) (*token.AuthClaims, error) {
		return &token.AuthClaims{UserId: int64(1)}, nil
	}

	actualAnswer, actualErr := service.DeleteWordByWord(wordRequest)
	assert.NoError(t, actualErr)
	assert.Equal(t, &expected, actualAnswer)
}

func TestWord_DeleteWordByWord_Negative(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	wordRequest := dto.RequestForDeletingWord{
		Token: "qwerty123",
		Id:    int64(1),
	}

	repository := NewMockRepository(ctrl)
	repository.EXPECT().DeleteWord(int64(1), wordRequest.Id).Return(nil, errors.New("fail deleting"))

	service := NewWordService(repository)
	service.parseTokenFn = func(tokenString string) (*token.AuthClaims, error) {
		return &token.AuthClaims{UserId: int64(1)}, nil
	}

	_, actualErr := service.DeleteWordByWord(wordRequest)
	assert.Error(t, actualErr)
	assert.ErrorIs(t, actualErr, httputils.ErrWordNotExist)
}

func TestWord_DeleteWordByWord_NegativeFailToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	wordRequest := dto.RequestForDeletingWord{
		Token: "qwerty123",
		Id:    int64(1),
	}

	repository := NewMockRepository(ctrl)

	service := NewWordService(repository)
	service.parseTokenFn = func(tokenString string) (*token.AuthClaims, error) {
		return nil, errors.New("fuck up")
	}

	_, actualErr := service.DeleteWordByWord(wordRequest)
	assert.Error(t, actualErr)
	assert.ErrorIs(t, actualErr, httputils.ErrValidation)
}
