package handlers

import (
	"drunklish/internal/service/word/dto"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateStatisticHandler_Positive(t *testing.T) {
	service := &mockService{
		fnCrSt: func(statistic dto.RequestCreateStatistic) (*dto.ResponseCreateStatistic, error) {
			assert.Equal(t, int64(1), statistic.TrainingId)
			assert.Equal(t, dto.ResponseWordList{dto.ResponseWord{Id: 1, Word: "test", Translate: "test"}}, statistic.CorrectAnswers.Words)
			assert.Equal(t, dto.ResponseWordList{dto.ResponseWord{Id: 2, Word: "test2", Translate: "test"}}, statistic.WrongAnswers.Words)
			assert.Equal(t, "token", statistic.Token)

			return &dto.ResponseCreateStatistic{Id: int64(1)}, nil
		},
	}

	expectedResponse := &dto.ResponseCreateStatistic{
		Id: int64(1),
	}

	handler := CreateStatisticHandler(service)

	actualResponse, actualErr := handler(dto.RequestCreateStatistic{
		TrainingId: int64(1),
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
		Token: "token",
	},
	)
	assert.NoError(t, actualErr)
	assert.Equal(t, expectedResponse, actualResponse)
}

func TestCreateStatisticHandler_Negative(t *testing.T) {
	service := &mockService{
		fnCrSt: func(statistic dto.RequestCreateStatistic) (*dto.ResponseCreateStatistic, error) {
			assert.Equal(t, int64(1), statistic.TrainingId)
			assert.Equal(t, dto.ResponseWordList{dto.ResponseWord{Id: 1, Word: "test", Translate: "test"}}, statistic.CorrectAnswers.Words)
			assert.Equal(t, dto.ResponseWordList{dto.ResponseWord{Id: 2, Word: "test2", Translate: "test"}}, statistic.WrongAnswers.Words)
			assert.Equal(t, "token", statistic.Token)

			return nil, errors.New("fuck up")
		},
	}

	handler := CreateStatisticHandler(service)

	_, actualErr := handler(dto.RequestCreateStatistic{
		TrainingId: int64(1),
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
		Token: "token",
	},
	)
	assert.Error(t, actualErr)
}
