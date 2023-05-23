package word

import (
	"drunklish/internal/pkg/httputils"
	"drunklish/internal/service/word/dto"
	"fmt"
)

func (w *Word) GetStatistic(statistic dto.RequestStatistic) (*dto.ResponseStatistic, error) {
	token, err := w.parseTokenFn(statistic.Token)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", httputils.ErrValidation)
	}

	trainingInfo, err := w.repo.GetTrainingInfoById(statistic.TrainingId, token.UserId)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", httputils.ErrInternalServer, err)
	}

	var statisticWords dto.ResponseStatistic
	for i := 0; i < len(trainingInfo.Words); i++ {
		if trainingInfo.Answers[i].Translate != trainingInfo.Words[i].Translate {
			statisticWords.WrongAnswers.Words = append(statisticWords.WrongAnswers.Words, trainingInfo.Answers[i])
		} else {
			statisticWords.CorrectAnswers.Words = append(statisticWords.CorrectAnswers.Words, trainingInfo.Answers[i])
		}
	}

	stat, err := w.repo.GetStatisticByTrainingId(statistic.TrainingId, token.UserId)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", httputils.ErrInternalServer, err)
	}

	stat.CorrectAnswers.Words = statisticWords.CorrectAnswers.Words
	stat.WrongAnswers.Words = statisticWords.WrongAnswers.Words

	return stat, nil
}
