package word

import (
	"drunklish/internal/pkg/httputils"
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
		Word:   "boogaga",
		UserId: int64(1),
	}

	expected := dto.ResponseFromDeleting{Answer: "deleting success"}

	repository := NewMockRepository(ctrl)
	repository.EXPECT().DeleteWord(wordRequest).Return(&expected, nil)

	service := NewWordService(repository)

	actualAnswer, actualErr := service.DeleteWordByWord(wordRequest)
	assert.NoError(t, actualErr)
	assert.Equal(t, &expected, actualAnswer)
}

func TestWord_DeleteWordByWord_Negative(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	wordRequest := dto.RequestForDeletingWord{
		Word:   "boogaga",
		UserId: int64(1),
	}

	repository := NewMockRepository(ctrl)
	repository.EXPECT().DeleteWord(wordRequest).Return(nil, errors.New("fail deleting"))

	service := NewWordService(repository)

	_, actualErr := service.DeleteWordByWord(wordRequest)
	assert.Error(t, actualErr)
	assert.ErrorIs(t, actualErr, httputils.ErrWordNotExist)
}
