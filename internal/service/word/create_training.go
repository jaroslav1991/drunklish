package word

import (
	"drunklish/internal/pkg/httputils"
	"drunklish/internal/service/word/dto"
	"fmt"
)

func (w *Word) CreateTrainingWords(training dto.RequestForTraining) (*dto.ResponseForTraining, error) {
	token, err := w.parseTokenFn(training.Token)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", httputils.ErrValidation)
	}

	totalWords := len(training.Answers.Words)

	createdTraining, err := w.repo.CreateTraining(training.Words.Words, training.Answers.Words, int64(totalWords), token.UserId)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", httputils.ErrInternalServer, err)
	}

	return createdTraining, nil
}
