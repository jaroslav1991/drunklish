package word

import (
	"drunklish/internal/model"
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
	var wd model.Word

	if err := w.db.QueryRowx(selectWordQuery, word, userId).Scan(&wd.Word, &wd.Translate, &wd.UserId); err != nil {
		return fmt.Errorf("%w", ErrWord)
	}

	if _, err := w.db.Exec(deleteWordQuery, word, userId); err != nil {
		return err
	}

	return nil
}
