package word

import (
	"drunklish/internal/pkg/httputils"
	"drunklish/internal/service/word/dto"
	"drunklish/internal/service/word/validator"
	"fmt"
)

func (w *Word) GetWordsByUserId(word dto.RequestForGettingWord) (*dto.ResponseWords, error) {
	token, err := w.parseTokenFn(word.Token)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", httputils.ErrValidation)
	}

	exist, err := w.repo.CheckUserInDB(token.UserId)
	if !exist {
		return nil, fmt.Errorf("user not exists: %w", httputils.ErrValidation)
	}
	if err != nil {
		return nil, fmt.Errorf("%w", httputils.ErrValidation)
	}

	words, err := w.repo.GetWords(token.UserId)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", httputils.ErrInternalServer, err)
	}

	return words, nil
}

// todo: добавить проверки на наличие FirstDate и SecondDate

func (w *Word) GetWordsByCreatedAt(period dto.RequestForGetByPeriod) (*dto.ResponseWords, error) {
	token, err := w.parseTokenFn(period.Token)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", httputils.ErrValidation)
	}

	exist, err := w.repo.CheckUserInDB(token.UserId)
	if !exist {
		return nil, fmt.Errorf("user not exists %w", httputils.ErrValidation)
	}
	if err != nil {
		return nil, fmt.Errorf("%w", httputils.ErrValidation)
	}

	periodFromCheck := validator.CheckPlacesFirstOrSecondDate(period)

	words, err := w.repo.GetWordByCreated(token.UserId, periodFromCheck.CreatedAt.FirstDate, periodFromCheck.CreatedAt.SecondDate)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", httputils.ErrInternalServer, err)
	}

	return words, nil
}
