package handlers

import (
	"drunklish/internal/service/word/dto"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateTrainingHandler_Positive(t *testing.T) {
	service := &mockService{
		fnTr: func(training dto.RequestForTraining) (*dto.ResponseForTraining, error) {
			assert.Equal(t, dto.ResponseWordList{dto.ResponseWord{Id: 1, Word: "test", Translate: "test"}}, training.Words.Words)
			assert.Equal(t, dto.ResponseWordList{dto.ResponseWord{Id: 1, Word: "test", Translate: "test"}}, training.Answers.Words)
			assert.Equal(t, int64(1), training.WordsTotal)
			assert.Equal(t, "token", training.Token)

			return &dto.ResponseForTraining{Id: int64(1), Words: []dto.ResponseWord{{
				Id:        1,
				Word:      "test",
				Translate: "test",
			}}}, nil
		},
	}

	expectedResponse := &dto.ResponseForTraining{
		Id: int64(1),
		Words: []dto.ResponseWord{{
			Id:        1,
			Word:      "test",
			Translate: "test",
		}},
	}

	handler := CreateTrainingHandler(service)

	actualResponse, actualErr := handler(dto.RequestForTraining{
		Words: dto.ResponseWords{Words: []dto.ResponseWord{{
			Id:        1,
			Word:      "test",
			Translate: "test",
		}}},
		Answers: dto.ResponseWords{Words: []dto.ResponseWord{{
			Id:        1,
			Word:      "test",
			Translate: "test",
		}}},
		WordsTotal: 1,
		Token:      "token",
	})
	assert.NoError(t, actualErr)
	assert.Equal(t, expectedResponse, actualResponse)
}

func TestCreateTrainingHandler_Negative(t *testing.T) {
	service := &mockService{
		fnTr: func(training dto.RequestForTraining) (*dto.ResponseForTraining, error) {
			assert.Equal(t, dto.ResponseWordList{dto.ResponseWord{Id: 1, Word: "test", Translate: "test"}}, training.Words.Words)
			assert.Equal(t, dto.ResponseWordList{dto.ResponseWord{Id: 1, Word: "test", Translate: "test"}}, training.Answers.Words)
			assert.Equal(t, int64(1), training.WordsTotal)
			assert.Equal(t, "token", training.Token)

			return nil, errors.New("fuck up")
		},
	}

	handler := CreateTrainingHandler(service)

	_, actualErr := handler(dto.RequestForTraining{
		Words: dto.ResponseWords{Words: []dto.ResponseWord{{
			Id:        1,
			Word:      "test",
			Translate: "test",
		}}},
		Answers: dto.ResponseWords{Words: []dto.ResponseWord{{
			Id:        1,
			Word:      "test",
			Translate: "test",
		}}},
		WordsTotal: 1,
		Token:      "token",
	})
	assert.Error(t, actualErr)
}
