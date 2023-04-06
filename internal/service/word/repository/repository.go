package repository

import (
	"drunklish/internal/model"
	"drunklish/internal/pkg/db"
	"drunklish/internal/service/word/dto"
	"time"
)

const (
	createWordQuery          = `insert into words (word, translate, created_at, user_id) values ($1, $2, $3, $4) returning word, translate, created_at, user_id`
	getWordsQuery            = `select word, translate from words where user_id=$1`
	getWordsByCreatedAtQuery = `select word, translate from words where user_id=$1 and created_at=$2`
	deleteWordQuery          = `delete from words where word=$1 and user_id=$2`
	selectWordQuery          = `select word, translate, user_id from words where word=$1 and user_id=$2`
)

type WordRepository struct {
	db db.DB
}

func NewWordRepository(db db.DB) *WordRepository {
	return &WordRepository{db: db}
}

func (repo *WordRepository) Create(word dto.CreateWordRequest) (*model.Word, error) {
	createdAt := time.Now()
	if err := repo.db.QueryRowx(
		createWordQuery,
		word.Word,
		word.Translate,
		createdAt,
		word.UserId,
	).Scan(
		&word.Word,
		&word.Translate,
		&word.CreatedAt,
		&word.UserId,
	); err != nil {
		return nil, err
	}

	return &model.Word{
		Word:      word.Word,
		Translate: word.Translate,
		CreatedAt: createdAt,
		UserId:    word.UserId,
	}, nil
}

func (repo *WordRepository) GetWords(wordReq dto.RequestForGettingWord) ([]*dto.ResponseWord, error) {
	var words []*dto.ResponseWord

	rows, err := repo.db.Query(getWordsQuery, wordReq.UserId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var word dto.ResponseWord
		if err := rows.Scan(&word.Word, &word.Translate); err != nil {
			return nil, err
		}

		words = append(words, &word)
	}
	return words, nil
}

func (repo *WordRepository) GetWordsByCreated(userId int64, createdAt time.Time) (*model.Word, error) {
	var word model.Word

	if err := repo.db.QueryRowx(getWordsByCreatedAtQuery, userId, createdAt).Scan(&word.Word, &word.Translate); err != nil {
		return nil, err
	}

	return &word, nil
}

func (repo *WordRepository) DeleteWord(word dto.RequestForDeletingWord) error {
	var wd model.Word

	if err := repo.db.QueryRowx(selectWordQuery, word.Word, word.UserId).Scan(&wd.Word, &wd.Translate, &wd.UserId); err != nil {
		return err
	}

	if _, err := repo.db.Exec(deleteWordQuery, word.Word, word.UserId); err != nil {
		return err
	}

	return nil
}
