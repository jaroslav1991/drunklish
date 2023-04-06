package handlers

import (
	"drunklish/internal/service/word/dto"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type mockService struct {
	fnC  func(word dto.CreateWordRequest) (*dto.ResponseFromCreateWor, error)
	fnG  func(word dto.RequestForGettingWord) (*dto.ResponseWords, error)
	fnGP func(period dto.RequestForGetByPeriod) (*dto.ResponseWords, error)
	fnD  func(word dto.RequestForDeletingWord) (*dto.ResponseFromDeleting, error)
}

func (m *mockService) CreateWord(word dto.CreateWordRequest) (*dto.ResponseFromCreateWor, error) {
	return m.fnC(word)
}

func (m *mockService) GetWordsByUserId(word dto.RequestForGettingWord) (*dto.ResponseWords, error) {
	return m.fnG(word)
}

func (m *mockService) DeleteWordByWord(word dto.RequestForDeletingWord) (*dto.ResponseFromDeleting, error) {
	return m.fnD(word)
}

func (m *mockService) GetWordsByCreatedAt(period dto.RequestForGetByPeriod) (*dto.ResponseWords, error) {
	return m.fnGP(period)
}

func TestCreateWordHandler_Positive(t *testing.T) {
	service := &mockService{
		fnC: func(word dto.CreateWordRequest) (*dto.ResponseFromCreateWor, error) {
			assert.Equal(t, "qwe", word.Word)
			assert.Equal(t, "qwe", word.Translate)

			return &dto.ResponseFromCreateWor{
				Word:      "qwe",
				Translate: "qwe",
			}, nil
		},
	}

	expectedResponse := &dto.ResponseFromCreateWor{
		Word:      "qwe",
		Translate: "qwe",
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
		fnC: func(word dto.CreateWordRequest) (*dto.ResponseFromCreateWor, error) {
			assert.Equal(t, "qwe", word.Word)
			assert.Equal(t, "qwe", word.Translate)

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
