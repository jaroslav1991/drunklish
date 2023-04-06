package handlers

import (
	"drunklish/internal/service/word/dto"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetWordsHandler_Positive(t *testing.T) {
	service := &mockService{
		fnG: func(word dto.RequestForGettingWord) ([]*dto.ResponseWord, error) {
			assert.Equal(t, int64(1), word.UserId)

			return []*dto.ResponseWord{{
				Word:      "qwe",
				Translate: "qwe",
			}}, nil
		},
	}

	word := dto.RequestForGettingWord{UserId: int64(1)}

	expectedResponse := []*dto.ResponseWord{{
		Word:      "qwe",
		Translate: "qwe",
	}}

	handler := GetWordsHandler(service)

	actualResponse, actualErr := handler(word)
	assert.NoError(t, actualErr)
	assert.Equal(t, expectedResponse, actualResponse)
}

func TestGetWordsHandler_Negative(t *testing.T) {
	service := &mockService{
		fnG: func(word dto.RequestForGettingWord) ([]*dto.ResponseWord, error) {
			assert.Equal(t, int64(1), word.UserId)

			return nil, errors.New("fuck up")
		},
	}

	word := dto.RequestForGettingWord{UserId: int64(1)}

	handler := GetWordsHandler(service)

	_, actualErr := handler(word)
	assert.Error(t, actualErr)
}
