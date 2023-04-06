package word

import (
	"drunklish/internal/pkg/httputils"
	"drunklish/internal/service/word/dto"
	"errors"
	"fmt"
)

var (
	ErrWord = errors.New("word not exist")
)

func (w *Word) DeleteWordByWord(word dto.RequestForDeletingWord) (*dto.ResponseFromDeleting, error) {
	answer, err := w.repo.DeleteWord(word)
	if err != nil {
		return nil, fmt.Errorf("%w", httputils.ErrWordNotExist)
	}

	return answer, nil
}
