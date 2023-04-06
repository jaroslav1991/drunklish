package handlers

import (
	"drunklish/internal/service/word/dto"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteWordHandler_Positive(t *testing.T) {
	service := &mockService{
		fnD: func(word dto.RequestForDeletingWord) error {
			assert.Equal(t, "qwe", word.Word)
			assert.Equal(t, int64(1), word.UserId)

			return nil
		},
	}

	word := dto.RequestForDeletingWord{
		Word:   "qwe",
		UserId: int64(1),
	}

	handler := DeleteWordHandler(service)

	actualErr := handler(word)
	assert.NoError(t, actualErr)
}

func TestDeleteWordHandler_Negative(t *testing.T) {
	service := &mockService{
		fnD: func(word dto.RequestForDeletingWord) error {
			assert.Equal(t, "qwe", word.Word)
			assert.Equal(t, int64(1), word.UserId)

			return errors.New("fuck up")
		},
	}

	word := dto.RequestForDeletingWord{
		Word:   "qwe",
		UserId: int64(1),
	}

	handler := DeleteWordHandler(service)

	actualErr := handler(word)
	assert.Error(t, actualErr)
}
