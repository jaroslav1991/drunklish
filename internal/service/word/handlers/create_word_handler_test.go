package handlers

import (
	"drunklish/internal/model"
	"drunklish/internal/service/word/dto"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type mockService struct {
	fnC func(word dto.CreateWordRequest) (*model.Word, error)
	fnG func(word dto.RequestForGettingWord) ([]*dto.ResponseWord, error)
	fnD func(word dto.RequestForDeletingWord) error
}

func (m *mockService) CreateWord(word dto.CreateWordRequest) (*model.Word, error) {
	return m.fnC(word)
}

func (m *mockService) GetWordsByUserId(word dto.RequestForGettingWord) ([]*dto.ResponseWord, error) {
	return m.fnG(word)
}

func (m *mockService) DeleteWordByWord(word dto.RequestForDeletingWord) error {
	return m.fnD(word)
}

func TestCreateWordHandler_Positive(t *testing.T) {
	service := &mockService{
		fnC: func(word dto.CreateWordRequest) (*model.Word, error) {
			assert.Equal(t, "qwe", word.Word)
			assert.Equal(t, "qwe", word.Translate)
			assert.Equal(t, time.Now(), word.CreatedAt)
			assert.Equal(t, int64(1), word.UserId)

			return &model.Word{
				Word:      "qwe",
				Translate: "qwe",
				CreatedAt: word.CreatedAt,
				UserId:    word.UserId,
			}, nil
		},
	}

	expectedResponse := &model.Word{
		Word:      "qwe",
		Translate: "qwe",
		CreatedAt: time.Now(),
		UserId:    int64(1),
	}

	handler := CreateWordHandler(service)

	actualResponse, actualErr := handler(dto.CreateWordRequest{
		Word:      "qwe",
		Translate: "qwe",
		CreatedAt: time.Now(),
		UserId:    int64(1),
	})
	assert.NoError(t, actualErr)
	assert.Equal(t, expectedResponse, actualResponse)
}

func TestCreateWordHandler_Negative(t *testing.T) {
	service := &mockService{
		fnC: func(word dto.CreateWordRequest) (*model.Word, error) {
			assert.Equal(t, "qwe", word.Word)
			assert.Equal(t, "qwe", word.Translate)
			assert.Equal(t, word.CreatedAt, word.CreatedAt)
			assert.Equal(t, int64(1), word.UserId)

			return nil, errors.New("fuck up")
		},
	}

	handler := CreateWordHandler(service)

	_, actualErr := handler(dto.CreateWordRequest{
		Word:      "qwe",
		Translate: "qwe",
		CreatedAt: time.Now(),
		UserId:    int64(1),
	})
	assert.Error(t, actualErr)
}
