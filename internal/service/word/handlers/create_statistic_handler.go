package handlers

import "drunklish/internal/service/word/dto"

type CreateStatisticFn func(statistic dto.RequestCreateStatistic) (*dto.ResponseCreateStatistic, error)

func CreateStatisticHandler(service WordService) CreateStatisticFn {
	return func(statistic dto.RequestCreateStatistic) (*dto.ResponseCreateStatistic, error) {
		stat, err := service.CreateStatistic(statistic)
		if err != nil {
			return nil, err
		}

		return stat, nil
	}
}
