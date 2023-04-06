package word

import (
	"drunklish/internal/model"
	"drunklish/internal/pkg/httputils"
	"drunklish/internal/service/word/dto"
	"fmt"
	"time"
)

// todo: implement validator where check user in db for user_id

func (w *Word) GetWordsByUserId(word dto.RequestForGettingWord) (*dto.ResponseWords, error) {
	words, err := w.repo.GetWords(word)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", httputils.ErrInternalServer, err)
	}

	return words, nil
}

func (w *Word) GetWordsByCreatedAt(userId int64, createdAt time.Time) (*model.Word, error) {
	words, err := w.repo.GetWordsByCreated(userId, createdAt)
	if err != nil {
		return nil, err
	}

	return words, nil
}
