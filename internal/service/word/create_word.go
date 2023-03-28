package word

import (
	"drunklish/internal/model"
	"time"
)

const (
	createWordQuery = `insert into words (word, translate, created_at, user_id) values ($1, $2, $3, $4) returning word, translate, created_at, user_id`
)

func (w *Word) CreateWord(word *model.Word) (*model.Word, error) {
	word.CreatedAt = time.Now()

	if err := w.db.QueryRowx(
		createWordQuery,
		word.Word,
		word.Translate,
		word.CreatedAt,
		word.UserId,
	).Scan(
		&word.Word,
		&word.Translate,
		&word.CreatedAt,
		&word.UserId,
	); err != nil {
		return nil, err
	}

	return word, nil
}
