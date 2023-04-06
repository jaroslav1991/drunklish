package handlers

import "drunklish/internal/service/word/dto"

type GetAllFn func(word dto.RequestForGettingWord) ([]*dto.ResponseWord, error)

func GetWordsHandler(service WordService) GetAllFn {
	return func(word dto.RequestForGettingWord) ([]*dto.ResponseWord, error) {
		allWords, err := service.GetWordsByUserId(word)
		if err != nil {
			return nil, err
		}

		return allWords, nil
	}
}
