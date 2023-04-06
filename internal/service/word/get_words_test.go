package word

import (
	"drunklish/internal/pkg/httputils"
	"drunklish/internal/service/word/dto"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWord_GetWordsByUserId_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	word := dto.RequestForGettingWord{UserId: int64(1)}

	wordsFromGet := dto.ResponseWords{Words: []dto.ResponseWord{{
		Word:      "boogaga1",
		Translate: "boo1",
	}, {
		Word:      "boogaga2",
		Translate: "boo",
	}}}

	repository := NewMockRepository(ctrl)

	repository.EXPECT().GetWords(word).Return(&wordsFromGet, nil)

	service := NewWordService(repository)

	actualWords, err := service.GetWordsByUserId(word)
	assert.NoError(t, err)

	assert.Equal(t, &dto.ResponseWords{Words: []dto.ResponseWord{{
		Word:      "boogaga1",
		Translate: "boo1",
	}, {
		Word:      "boogaga2",
		Translate: "boo",
	}}}, actualWords)
}

func TestWord_GetWordsByUserId_NegativeFailGetWord(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	word := dto.RequestForGettingWord{UserId: int64(1)}

	repository := NewMockRepository(ctrl)

	repository.EXPECT().GetWords(word).Return(nil, errors.New("fail create word"))

	service := NewWordService(repository)

	_, err := service.GetWordsByUserId(word)
	assert.ErrorIs(t, err, httputils.ErrInternalServer)

}
