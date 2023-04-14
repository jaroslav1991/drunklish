package handlers

import (
	"drunklish/internal/service/word/dto"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpdateWordHandler_Positive(t *testing.T) {
	service := &mockService{
		fnU: func(word dto.RequestForUpdateWord) (*dto.ResponseWord, error) {
			assert.Equal(t, "qwe", word.Word)
			assert.Equal(t, "qwe", word.Translate)
			assert.Equal(t, int64(1), word.Id)

			return &dto.ResponseWord{
				Id:        1,
				Word:      "qwe",
				Translate: "qwe",
			}, nil
		},
	}

	word := dto.RequestForUpdateWord{
		Word:      "qwe",
		Translate: "qwe",
		Id:        int64(1),
		Token:     "qwerty",
	}

	expectResponse := &dto.ResponseWord{
		Id:        int64(1),
		Word:      "qwe",
		Translate: "qwe",
	}

	handler := UpdateWordHandler(service)

	actualResponse, err := handler(word)
	assert.NoError(t, err)
	assert.Equal(t, expectResponse, actualResponse)
}

func TestUpdateWordHandler_Negative(t *testing.T) {
	service := &mockService{
		fnU: func(word dto.RequestForUpdateWord) (*dto.ResponseWord, error) {
			assert.Equal(t, "qwe", word.Word)
			assert.Equal(t, "qwe", word.Translate)
			assert.Equal(t, int64(1), word.Id)

			return nil, errors.New("fuck up")
		},
	}

	word := dto.RequestForUpdateWord{
		Word:      "qwe",
		Translate: "qwe",
		Id:        int64(1),
		Token:     "qwerty",
	}

	handler := UpdateWordHandler(service)

	_, err := handler(word)
	assert.Error(t, err)
}
