package word

import (
	"drunklish/internal/model"
	"drunklish/internal/pkg/httputils"
	"drunklish/internal/service/word/dto"
	"fmt"
	"time"
)

func (w *Word) GetWordsByUserId(userId int64) ([]*dto.ResponseWord, error) {
	words, err := w.repo.GetWords(userId)
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
