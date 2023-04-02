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

	userId := int64(1)

	wordsFromGet := []*dto.ResponseWord{{
		Word:      "boogaga1",
		Translate: "boo1",
	}, {
		Word:      "boogaga2",
		Translate: "boo2",
	}}

	repository := NewMockRepository(ctrl)

	repository.EXPECT().GetWords(userId).Return(wordsFromGet, nil)

	service := NewWordService(repository)

	actualWords, err := service.GetWordsByUserId(userId)
	assert.NoError(t, err)

	assert.Equal(t, []*dto.ResponseWord{{
		Word:      "boogaga1",
		Translate: "boo1",
	}, {
		Word:      "boogaga2",
		Translate: "boo2",
	}}, actualWords)
}

func TestWord_GetWordsByUserId_NegativeFailGetWord(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userId := int64(1)

	repository := NewMockRepository(ctrl)

	repository.EXPECT().GetWords(userId).Return(nil, errors.New("fail create word"))

	service := NewWordService(repository)

	_, err := service.GetWordsByUserId(userId)
	assert.ErrorIs(t, err, httputils.ErrInternalServer)

}
