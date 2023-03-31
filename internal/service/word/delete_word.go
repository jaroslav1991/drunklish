package word

import (
	"errors"
	"fmt"
)

const (
	deleteWordQuery = `delete from words where word=$1 and user_id=$2`
	selectWordQuery = `select word, translate, user_id from words where word=$1 and user_id=$2`
)

var (
	ErrWord = errors.New("word not exist")
)

func (w *Word) DeleteWordByWord(word string, userId int64) error {
	if err := w.repo.DeleteWord(word, userId); err != nil {
		return fmt.Errorf("%w", ErrWord)
	}

	return nil
}
