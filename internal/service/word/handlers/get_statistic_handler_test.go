package handlers

import (
	"drunklish/internal/service/word/dto"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetStatisticHandler_Positive(t *testing.T) {
	service := &mockService{
		fnSt: func(statistic dto.RequestStatistic) (*dto.ResponseStatistic, error) {
			assert.Equal(t, int64(1), statistic.TrainingId)
			assert.Equal(t, "token", statistic.Token)

			return &dto.ResponseStatistic{
				CorrectAnswers: dto.ResponseWords{Words: []dto.ResponseWord{{
					Id:        1,
					Word:      "test",
					Translate: "test",
				}}},
				WrongAnswers: dto.ResponseWords{Words: []dto.ResponseWord{{
					Id:        2,
					Word:      "test2",
					Translate: "test",
				}}},
			}, nil
		},
	}

	expectedResponse := &dto.ResponseStatistic{
		CorrectAnswers: dto.ResponseWords{Words: []dto.ResponseWord{{
			Id:        1,
			Word:      "test",
			Translate: "test",
		}}},
		WrongAnswers: dto.ResponseWords{Words: []dto.ResponseWord{{
			Id:        2,
			Word:      "test2",
			Translate: "test",
		}}},
	}

	handler := GetStatisticHandler(service)

	actualResponse, actualErr := handler(dto.RequestStatistic{
		TrainingId: 1,
		Token:      "token",
	})
	assert.NoError(t, actualErr)
	assert.Equal(t, expectedResponse, actualResponse)
}

func TestGetStatisticHandler_Negative(t *testing.T) {
	service := &mockService{
		fnSt: func(statistic dto.RequestStatistic) (*dto.ResponseStatistic, error) {
			assert.Equal(t, int64(1), statistic.TrainingId)
			assert.Equal(t, "token", statistic.Token)

			return nil, errors.New("fuck up")
		},
	}

	handler := GetStatisticHandler(service)

	_, actualErr := handler(dto.RequestStatistic{
		TrainingId: 1,
		Token:      "token",
	})
	assert.Error(t, actualErr)
}
