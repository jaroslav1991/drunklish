package word

import (
	"drunklish/internal/pkg/httputils"
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
	}
	wordFromCreated := dto.ResponseFromCreateWor{
		Word:      "boogaga",
		Translate: "смешняшка",
	}

	repository := NewMockRepository(ctrl)
	repository.EXPECT().Create(wordForCreateWord).Return(&wordFromCreated, nil)

	service := NewWordService(repository)

	actualWord, err := service.CreateWord(wordForCreateWord)
	assert.NoError(t, err)

	assert.Equal(t, &dto.ResponseFromCreateWor{
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
	}

	repository := NewMockRepository(ctrl)
	repository.EXPECT().Create(wordForCreateWord).Return(nil, errors.New("fail create word"))

	service := NewWordService(repository)

	_, err := service.CreateWord(wordForCreateWord)
	assert.Error(t, err)
}

func TestWord_CreateWord_NegativeFailLengthWord(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	wordForCreateWord := dto.CreateWordRequest{
		Word:      "",
		Translate: "смешняшка",
	}

	repository := NewMockRepository(ctrl)

	service := NewWordService(repository)

	_, err := service.CreateWord(wordForCreateWord)
	assert.ErrorIs(t, err, httputils.ErrValidation)
}

func TestWord_CreateWord_NegativeFailLengthTranslate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	wordForCreateWord := dto.CreateWordRequest{
		Word:      "boogaga",
		Translate: "",
	}

	repository := NewMockRepository(ctrl)

	service := NewWordService(repository)

	_, err := service.CreateWord(wordForCreateWord)
	assert.ErrorIs(t, err, httputils.ErrValidation)
}
