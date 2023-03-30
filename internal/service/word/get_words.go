package word

import (
	"drunklish/internal/model"
	"errors"
	"fmt"
	"time"
)

const (
	getWordsQuery            = `select word, translate from words where user_id=$1`
	getWordsByCreatedAtQuery = `select word, translate from words where user_id=$1 and created_at=$2`
)

var (
	ErrUserID = errors.New("not authorized user")
)

type ResponseWord struct {
	Word      string `json:"word"`
	Translate string `json:"translate"`
}

func (w *Word) GetWordsByUserId(userId int64) ([]*ResponseWord, error) {
	var words []*ResponseWord

	rows, err := w.db.Query(getWordsQuery, userId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var word ResponseWord
		if err := rows.Scan(&word.Word, &word.Translate); err != nil {
			return nil, err
		}

		words = append(words, &word)
	}

	if userId == 0 {
		return nil, fmt.Errorf("%w", ErrUserID)
	}

	return words, nil
}

func (w *Word) GetWordsByCreatedAt(userId int64, createdAt time.Time) (*model.Word, error) {
	var word model.Word

	if err := w.db.QueryRowx(getWordsByCreatedAtQuery, userId, createdAt).Scan(&word.Word, &word.Translate); err != nil {
		return nil, err
	}

	return &word, nil
}
