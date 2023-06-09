package handlers

import (
	"drunklish/internal/service/word/dto"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteWordHandler_Positive(t *testing.T) {
	service := &mockService{
		fnD: func(word dto.RequestForDeletingWord) (*dto.ResponseFromDeleting, error) {
			assert.Equal(t, "qwerty123", word.Token)
			assert.Equal(t, int64(1), word.Id)

			return &dto.ResponseFromDeleting{Answer: "deleting success"}, nil
		},
	}

	word := dto.RequestForDeletingWord{
		Token: "qwerty123",
		Id:    int64(1),
	}

	handler := DeleteWordHandler(service)

	actualAnswer, actualErr := handler(word)
	assert.NoError(t, actualErr)
	assert.Equal(t, &dto.ResponseFromDeleting{Answer: "deleting success"}, actualAnswer)
}

func TestDeleteWordHandler_Negative(t *testing.T) {
	service := &mockService{
		fnD: func(word dto.RequestForDeletingWord) (*dto.ResponseFromDeleting, error) {
			assert.Equal(t, "qwerty123", word.Token)
			assert.Equal(t, int64(1), word.Id)

			return nil, errors.New("fuck up")
		},
	}

	word := dto.RequestForDeletingWord{
		Token: "qwerty123",
		Id:    int64(1),
	}

	handler := DeleteWordHandler(service)

	_, actualErr := handler(word)
	assert.Error(t, actualErr)
}
