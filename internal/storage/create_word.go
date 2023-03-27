package storage

import "time"

const (
	createWordQuery = `insert into words (word, translate, created_at, user_id) values ($1, $2, $3, $4) returning word, translate, created_at, user_id`
)

func (s *Storage) CreateWord(word *Word) (*Word, error) {
	word.CreatedAt = time.Now()

	if err := s.DB.QueryRowx(
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
