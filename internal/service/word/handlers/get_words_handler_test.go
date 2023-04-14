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
			assert.Equal(t, "qwerty123", word.Token)

			return &dto.ResponseWords{Words: []dto.ResponseWord{{
				Id:        int64(1),
				Word:      "qwe",
				Translate: "qwe",
			}}}, nil
		},
	}

	word := dto.RequestForGettingWord{Token: "qwerty123"}

	expectedResponse := &dto.ResponseWords{Words: []dto.ResponseWord{{
		Id:        int64(1),
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
			assert.Equal(t, "qwerty123", word.Token)

			return nil, errors.New("fuck up")
		},
	}

	word := dto.RequestForGettingWord{Token: "qwerty123"}

	handler := GetWordsHandler(service)

	_, actualErr := handler(word)
	assert.Error(t, actualErr)
}

func TestGetWordByPeriodHandler_Positive(t *testing.T) {
	service := &mockService{
		fnGP: func(period dto.RequestForGetByPeriod) (*dto.ResponseWords, error) {
			now := time.Now()
			assert.Equal(t, "", period.Token)
			assert.Equal(t, now, period.CreatedAt.FirstDate)
			assert.Equal(t, now, period.CreatedAt.SecondDate)

			return &dto.ResponseWords{Words: []dto.ResponseWord{{
				Id:        int64(1),
				Word:      "qwe",
				Translate: "qwe",
			}}}, nil
		},
	}
	var period dto.RequestForGetByPeriod

	requestFromPeriod := dto.RequestForGetByPeriod{
		Token: period.Token,
		CreatedAt: dto.Period{
			FirstDate:  time.Now(),
			SecondDate: time.Now(),
		},
	}

	expectedResponse := &dto.ResponseWords{Words: []dto.ResponseWord{{
		Id:        int64(1),
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
			now := time.Now()
			assert.Equal(t, "", period.Token)
			assert.Equal(t, now, period.CreatedAt.FirstDate)
			assert.Equal(t, now, period.CreatedAt.SecondDate)

			return nil, errors.New("fuck up")
		},
	}

	var period dto.RequestForGetByPeriod

	requestFromPeriod := dto.RequestForGetByPeriod{
		Token: period.Token,
		CreatedAt: dto.Period{
			FirstDate:  time.Now(),
			SecondDate: time.Now(),
		},
	}

	handler := GetWordByPeriodHandler(service)

	_, actualErr := handler(requestFromPeriod)
	assert.Error(t, actualErr)
}
