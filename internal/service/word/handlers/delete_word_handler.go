package handlers

import "drunklish/internal/service/word/dto"

type DeleteWordFn func(word dto.RequestForDeletingWord) (*dto.ResponseFromDeleting, error)

func DeleteWordHandler(service WordService) DeleteWordFn {
	return func(word dto.RequestForDeletingWord) (*dto.ResponseFromDeleting, error) {
		answer, err := service.DeleteWordByWord(word)
		if err != nil {
			return nil, err
		}

		return answer, nil
	}
}
