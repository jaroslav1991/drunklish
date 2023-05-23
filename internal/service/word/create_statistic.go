package word

import (
	"drunklish/internal/pkg/httputils"
	"drunklish/internal/service/word/dto"
	"fmt"
)

func (w *Word) CreateStatistic(statistic dto.RequestCreateStatistic) (*dto.ResponseCreateStatistic, error) {
	token, err := w.parseTokenFn(statistic.Token)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", httputils.ErrValidation)
	}

	stat, err := w.repo.CreateStatisticByTrainingId(statistic.TrainingId, statistic.CorrectAnswers.Words, statistic.WrongAnswers.Words, token.UserId)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", httputils.ErrInternalServer, err)
	}

	return stat, nil
}
