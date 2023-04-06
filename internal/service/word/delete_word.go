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

func (w *Word) DeleteWordByWord(word dto.RequestForDeletingWord) error {
	if err := w.repo.DeleteWord(word); err != nil {
		return fmt.Errorf("%w", httputils.ErrWordNotExist)
	}

	return nil
}
