package handlers

import (
	"drunklish/internal/service/word/dto"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetWordsHandler_Positive(t *testing.T) {
	service := &mockService{
		fnG: func(word dto.RequestForGettingWord) (*dto.ResponseWords, error) {
			assert.Equal(t, int64(1), word.UserId)

			return &dto.ResponseWords{Words: []dto.ResponseWord{{
				Word:      "qwe",
				Translate: "qwe",
			}}}, nil
		},
	}

	word := dto.RequestForGettingWord{UserId: int64(1)}

	expectedResponse := &dto.ResponseWords{Words: []dto.ResponseWord{{
		Word:      "qwe",
		Translate: "qwe",
	}}}

	handler := GetWordsHandler(service)

	actualResponse, actualErr := handler(word)
	assert.NoError(t, actualErr)
	assert.Equal(t, expectedResponse, actualResponse)
}

func TestGetWordsHandler_Negative(t *testing.T) {
	service := &mockService{
		fnG: func(word dto.RequestForGettingWord) (*dto.ResponseWords, error) {
			assert.Equal(t, int64(1), word.UserId)

			return nil, errors.New("fuck up")
		},
	}

	word := dto.RequestForGettingWord{UserId: int64(1)}

	handler := GetWordsHandler(service)

	_, actualErr := handler(word)
	assert.Error(t, actualErr)
}

func TestGetWordByPeriodHandler_Positive(t *testing.T) {
	service := &mockService{
		fnGP: func(period dto.RequestForGetByPeriod) (*dto.ResponseWords, error) {
			assert.Equal(t, int64(1), period.UserId)
			assert.Equal(t, time.Now(), period.CreatedAt.FirstDate)
			assert.Equal(t, time.Now(), period.CreatedAt.SecondDate)

			return &dto.ResponseWords{Words: []dto.ResponseWord{{
				Word:      "qwe",
				Translate: "qwe",
			}}}, nil
		},
	}

	requestFromPeriod := dto.RequestForGetByPeriod{
		UserId: 1,
		CreatedAt: dto.Period{
			FirstDate:  time.Now(),
			SecondDate: time.Now(),
		},
	}

	expectedResponse := &dto.ResponseWords{Words: []dto.ResponseWord{{
		Word:      "qwe",
		Translate: "qwe",
	}}}

	handler := GetWordByPeriodHandler(service)

	actualResponse, actualErr := handler(requestFromPeriod)
	assert.NoError(t, actualErr)
	assert.Equal(t, expectedResponse, actualResponse)
}

func TestGetWordByPeriodHandler_Negative(t *testing.T) {
	service := &mockService{
		fnGP: func(period dto.RequestForGetByPeriod) (*dto.ResponseWords, error) {
			assert.Equal(t, int64(1), period.UserId)
			assert.Equal(t, time.Now(), period.CreatedAt.FirstDate)
			assert.Equal(t, time.Now(), period.CreatedAt.SecondDate)

			return nil, errors.New("fuck up")
		},
	}

	requestFromPeriod := dto.RequestForGetByPeriod{
		UserId: 1,
		CreatedAt: dto.Period{
			FirstDate:  time.Now(),
			SecondDate: time.Now(),
		},
	}

	handler := GetWordByPeriodHandler(service)

	_, actualErr := handler(requestFromPeriod)
	assert.Error(t, actualErr)
}
