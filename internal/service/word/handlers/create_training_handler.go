package handlers

import "drunklish/internal/service/word/dto"

type CreateTrainingFn func(training dto.RequestForTraining) (*dto.ResponseForTraining, error)

func CreateTrainingHandler(service WordService) CreateTrainingFn {
	return func(training dto.RequestForTraining) (*dto.ResponseForTraining, error) {
		createTraining, err := service.CreateTrainingWords(training)
		if err != nil {
			return nil, err
		}

		return createTraining, nil
	}
}
