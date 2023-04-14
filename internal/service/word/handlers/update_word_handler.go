package handlers

import "drunklish/internal/service/word/dto"

type UpdateFn func(word dto.RequestForUpdateWord) (*dto.ResponseWord, error)

func UpdateWordHandler(service WordService) UpdateFn {
	return func(word dto.RequestForUpdateWord) (*dto.ResponseWord, error) {
		updateWord, err := service.UpdateWord(word)
		if err != nil {
			return nil, err
		}

		return updateWord, nil
	}
}
