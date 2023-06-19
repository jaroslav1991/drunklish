package handlers

import (
	"drunklish/internal/service/word/dto"
)

type GetAllFn func(word dto.RequestForGettingWord) (*dto.ResponseWords, error)

type GetAllByPeriodFn func(period dto.RequestForGetByPeriod) (*dto.ResponseWords, error)

func GetWordsHandler(service WordService) GetAllFn {
	return func(word dto.RequestForGettingWord) (*dto.ResponseWords, error) {
		allWords, err := service.GetWordsByUserId(word)
		if err != nil {
			return nil, err
		}

		return allWords, nil
	}
}

func GetWordByPeriodHandler(service WordService) GetAllByPeriodFn {
	return func(period dto.RequestForGetByPeriod) (*dto.ResponseWords, error) {
		allWords, err := service.GetWordsByCreatedAt(period)
		if err != nil {
			return nil, err
		}

		return allWords, nil
	}
}
