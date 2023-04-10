package handlers

import (
	"drunklish/internal/service/word/dto"
)

type CreateWordFn func(word dto.CreateWordRequest) (*dto.ResponseFromCreateWord, error)

func CreateWordHandler(service WordService) CreateWordFn {
	return func(word dto.CreateWordRequest) (*dto.ResponseFromCreateWord, error) {
		createWord, err := service.CreateWord(word)
		if err != nil {
			return nil, err
		}

		return createWord, nil
	}
}
