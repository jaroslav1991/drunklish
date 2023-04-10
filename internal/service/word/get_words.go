package word

import (
	"drunklish/internal/pkg/httputils"
	"drunklish/internal/service/auth/token"
	"drunklish/internal/service/word/dto"
	"drunklish/internal/service/word/validator"
	"fmt"
)

func (w *Word) GetWordsByUserId(word dto.RequestForGettingWord) (*dto.ResponseWords, error) {
	userIdFromToken, err := token.ParseToken(word.Token)
	if err != nil {
		return nil, err
	}

	userId, err := w.repo.CheckUserInDB(userIdFromToken.UserId)
	if !userId {
		return nil, fmt.Errorf("user not exists: %w", httputils.ErrValidation)
	}
	if err != nil {
		return nil, fmt.Errorf("%w", httputils.ErrValidation)
	}

	words, err := w.repo.GetWords(userIdFromToken.UserId)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", httputils.ErrInternalServer, err)
	}

	return words, nil
}

// todo: добавить проверки на наличие FirstDate и SecondDate

func (w *Word) GetWordsByCreatedAt(period dto.RequestForGetByPeriod) (*dto.ResponseWords, error) {
	userId, err := w.repo.CheckUserInDB(period.UserId)
	if !userId {
		return nil, fmt.Errorf("user not exists %w", httputils.ErrValidation)
	}
	if err != nil {
		return nil, fmt.Errorf("%w", httputils.ErrValidation)
	}

	periodFromCheck := validator.CheckPlacesFirstOrSecondDate(period)

	words, err := w.repo.GetWordByCreated(periodFromCheck)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", httputils.ErrInternalServer, err)
	}

	return words, nil
}
