package repository

import (
	"drunklish/internal/pkg/db"
	"drunklish/internal/service/word/dto"
	"fmt"
	"time"
)

const (
	createWordQuery      = `insert into words (word, translate, created_at, user_id) values ($1, $2, $3, $4) returning word, translate`
	getWordsQuery        = `select id, word, translate from words where user_id=$1 order by created_at`
	deleteWordQuery      = `delete from words where user_id=$1 and id=$2 returning id`
	getWordByPeriodQuery = `select id, word, translate from words where user_id=$1 and created_at>$2 and created_at<$3 order by created_at`
	updateWordQuery      = `update words  set word=$1, translate=$2 where id=$3 and user_id=$4 returning word, translate`
)

type WordRepository struct {
	db db.DB
}

func NewWordRepository(db db.DB) *WordRepository {
	return &WordRepository{db: db}
}

func (repo *WordRepository) Update(word, translate string, id, userId int64) (*dto.ResponseWord, error) {
	var response dto.ResponseWord
	if err := repo.db.QueryRowx(updateWordQuery, word, translate, id, userId).Scan(&response.Word, &response.Translate); err != nil {
		return nil, err
	}

	return &response, nil
}

func (repo *WordRepository) Create(word string, translate string, createdAt time.Time, userId int64) (*dto.ResponseFromCreateWord, error) {
	var responseWord dto.ResponseFromCreateWord
	createdAt = time.Now()
	if err := repo.db.QueryRowx(
		createWordQuery,
		word,
		translate,
		createdAt,
		userId,
	).Scan(
		&responseWord.Word,
		&responseWord.Translate,
	); err != nil {
		return nil, err
	}

	return &responseWord, nil
}

func (repo *WordRepository) GetWords(userId int64) (*dto.ResponseWords, error) {
	var words dto.ResponseWords

	rows, err := repo.db.Query(getWordsQuery, userId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var word dto.ResponseWord
		if err := rows.Scan(&word.Id, &word.Word, &word.Translate); err != nil {
			return nil, err
		}

		words.Words = append(words.Words, word)
	}
	return &words, nil
}

func (repo *WordRepository) GetWordByCreated(userId int64, firstDate, secondDate time.Time) (*dto.ResponseWords, error) {
	var words dto.ResponseWords

	rows, err := repo.db.Query(getWordByPeriodQuery, userId, firstDate, secondDate)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var word dto.ResponseWord
		if err := rows.Scan(&word.Id, &word.Word, &word.Translate); err != nil {
			return nil, err
		}

		words.Words = append(words.Words, word)
	}
	return &words, nil
}

func (repo *WordRepository) DeleteWord(userId, id int64) (*dto.ResponseFromDeleting, error) {
	if err := repo.db.QueryRowx(deleteWordQuery, userId, id).Scan(&id); err != nil {
		return nil, err
	}

	return &dto.ResponseFromDeleting{Answer: fmt.Sprintf("%d: deleting success", id)}, nil
}
