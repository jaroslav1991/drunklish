package handlers

import "drunklish/internal/service/word/dto"

type GetStatisticFn func(statistic dto.RequestStatistic) (*dto.ResponseStatistic, error)

func GetStatisticHandler(service WordService) GetStatisticFn {
	return func(statistic dto.RequestStatistic) (*dto.ResponseStatistic, error) {
		stat, err := service.GetStatistic(statistic)
		if err != nil {
			return nil, err
		}

		return stat, nil
	}
}
