package handlers

import "drunklish/internal/service/word/dto"

type DeleteWordFn func(word dto.RequestForDeletingWord) error

func DeleteWordHandler(service WordService) DeleteWordFn {
	return func(word dto.RequestForDeletingWord) error {
		if err := service.DeleteWordByWord(word); err != nil {
			return err
		}

		return nil
	}
}
