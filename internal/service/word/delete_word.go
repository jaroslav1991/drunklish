package word

import (
	"drunklish/internal/pkg/httputils"
	"errors"
	"fmt"
)

var (
	ErrWord = errors.New("word not exist")
)

func (w *Word) DeleteWordByWord(word string, userId int64) error {
	if err := w.repo.DeleteWord(word, userId); err != nil {
		return fmt.Errorf("%w", httputils.ErrWordNotExist)
	}

	return nil
}
