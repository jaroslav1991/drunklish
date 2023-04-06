package handlers

import (
	"drunklish/internal/service/word/dto"
)

type CreateWordFn func(word dto.CreateWordRequest) (*dto.ResponseFromCreateWor, error)

func CreateWordHandler(service WordService) CreateWordFn {
	return func(word dto.CreateWordRequest) (*dto.ResponseFromCreateWor, error) {
		createWord, err := service.CreateWord(word)
		if err != nil {
			return nil, err
		}

		return createWord, nil
	}
}
