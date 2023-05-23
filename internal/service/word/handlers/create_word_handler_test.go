package handlers

import (
	"drunklish/internal/service/word/dto"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type mockService struct {
	fnC    func(word dto.CreateWordRequest) (*dto.ResponseFromCreateWord, error)
	fnG    func(word dto.RequestForGettingWord) (*dto.ResponseWords, error)
	fnGP   func(period dto.RequestForGetByPeriod) (*dto.ResponseWords, error)
	fnD    func(word dto.RequestForDeletingWord) (*dto.ResponseFromDeleting, error)
	fnU    func(word dto.RequestForUpdateWord) (*dto.ResponseWord, error)
	fnTr   func(training dto.RequestForTraining) (*dto.ResponseForTraining, error)
	fnSt   func(statistic dto.RequestStatistic) (*dto.ResponseStatistic, error)
	fnCrSt func(statistic dto.RequestCreateStatistic) (*dto.ResponseCreateStatistic, error)
}

func (m *mockService) CreateWord(word dto.CreateWordRequest) (*dto.ResponseFromCreateWord, error) {
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

func (m *mockService) UpdateWord(word dto.RequestForUpdateWord) (*dto.ResponseWord, error) {
	return m.fnU(word)
}

func (m *mockService) CreateTrainingWords(training dto.RequestForTraining) (*dto.ResponseForTraining, error) {
	return m.fnTr(training)
}

func (m *mockService) GetStatistic(statistic dto.RequestStatistic) (*dto.ResponseStatistic, error) {
	return m.fnSt(statistic)
}

func (m *mockService) CreateStatistic(statistic dto.RequestCreateStatistic) (*dto.ResponseCreateStatistic, error) {
	return m.fnCrSt(statistic)
}

func TestCreateWordHandler_Positive(t *testing.T) {
	service := &mockService{
		fnC: func(word dto.CreateWordRequest) (*dto.ResponseFromCreateWord, error) {
			assert.Equal(t, "qwe", word.Word)
			assert.Equal(t, "qwe", word.Translate)
			assert.Equal(t, "qwerty123", word.Token)

			return &dto.ResponseFromCreateWord{
				Word:      "qwe",
				Translate: "qwe",
			}, nil
		},
	}

	expectedResponse := &dto.ResponseFromCreateWord{
		Word:      "qwe",
		Translate: "qwe",
	}

	handler := CreateWordHandler(service)

	actualResponse, actualErr := handler(dto.CreateWordRequest{
		Word:      "qwe",
		Translate: "qwe",
		CreatedAt: time.Now(),
		Token:     "qwerty123",
	})
	assert.NoError(t, actualErr)
	assert.Equal(t, expectedResponse, actualResponse)
}

func TestCreateWordHandler_Negative(t *testing.T) {
	service := &mockService{
		fnC: func(word dto.CreateWordRequest) (*dto.ResponseFromCreateWord, error) {
			assert.Equal(t, "qwe", word.Word)
			assert.Equal(t, "qwe", word.Translate)
			assert.Equal(t, "qwerty123", word.Token)

			return nil, errors.New("fuck up")
		},
	}

	handler := CreateWordHandler(service)

	_, actualErr := handler(dto.CreateWordRequest{
		Word:      "qwe",
		Translate: "qwe",
		CreatedAt: time.Now(),
		Token:     "qwerty123",
	})
	assert.Error(t, actualErr)
}
