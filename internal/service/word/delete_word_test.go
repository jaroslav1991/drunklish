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

	repository := NewMockRepository(ctrl)
	repository.EXPECT().DeleteWord(wordRequest).Return(nil)

	service := NewWordService(repository)

	actualErr := service.DeleteWordByWord(wordRequest)
	assert.NoError(t, actualErr)
}

func TestWord_DeleteWordByWord_Negative(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	wordRequest := dto.RequestForDeletingWord{
		Word:   "boogaga",
		UserId: int64(1),
	}

	repository := NewMockRepository(ctrl)
	repository.EXPECT().DeleteWord(wordRequest).Return(errors.New("fail deleting"))

	service := NewWordService(repository)

	actualErr := service.DeleteWordByWord(wordRequest)
	assert.Error(t, actualErr)
	assert.ErrorIs(t, actualErr, httputils.ErrWordNotExist)
}
