package word

import (
	"drunklish/internal/pkg/httputils"
	"drunklish/internal/service/word/dto"
	"drunklish/internal/service/word/validator"
	"fmt"
)

// todo: implement validator where check user in db for user_id

func (w *Word) GetWordsByUserId(word dto.RequestForGettingWord) (*dto.ResponseWords, error) {
	words, err := w.repo.GetWords(word)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", httputils.ErrInternalServer, err)
	}

	return words, nil
}

func (w *Word) GetWordsByCreatedAt(period dto.RequestForGetByPeriod) (*dto.ResponseWords, error) {
	periodFromCheck := validator.CheckPlacesFirstOrSecondDate(period)

	words, err := w.repo.GetWordByCreated(periodFromCheck)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", httputils.ErrInternalServer, err)
	}

	return words, nil
}
