package word

import (
	"drunklish/internal/pkg/httputils"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWord_DeleteWordByWord_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	word := "boogaga"
	userId := int64(1)

	repository := NewMockRepository(ctrl)
	repository.EXPECT().DeleteWord(word, userId).Return(nil)

	service := NewWordService(repository)

	actualErr := service.DeleteWordByWord(word, userId)
	assert.NoError(t, actualErr)
}

func TestWord_DeleteWordByWord_Negative(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	word := "boogaga"
	userId := int64(1)

	repository := NewMockRepository(ctrl)
	repository.EXPECT().DeleteWord(word, userId).Return(errors.New("fail deleting"))

	service := NewWordService(repository)

	actualErr := service.DeleteWordByWord(word, userId)
	assert.Error(t, actualErr)
	assert.ErrorIs(t, actualErr, httputils.ErrWordNotExist)
}
