package repository

import (
	"database/sql"
	"drunklish/internal/model"
	"drunklish/internal/pkg/db"
	"drunklish/internal/service/word/dto"
	"time"
)

const (
	createWordQuery      = `insert into words (word, translate, created_at, user_id) values ($1, $2, $3, $4) returning word, translate`
	getWordsQuery        = `select w.id, w.word, w.translate from words w join users u on w.user_id = u.id where w.user_id=$1 order by created_at`
	deleteWordQuery      = `delete from words where word=$1 and user_id=$2`
	selectWordQuery      = `select word, translate, user_id from words where word=$1 and user_id=$2`
	getWordByPeriodQuery = `select w.id, w.word, w.translate from words w join users u on w.user_id = u.id where w.user_id=$1 and w.created_at>$2 and w.created_at<$3 order by created_at`
	getUserQuery         = `select user_id from words where user_id=$1`
	updateWordQuery      = `update words w set word=$1, translate=$2 from users u where w.id=$3 and w.user_id=$4 and w.user_id=u.id returning w.word, w.translate`
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

func (repo *WordRepository) CheckUserInDB(userId int64) (bool, error) {
	err := repo.db.QueryRowx(getUserQuery, userId).Scan(&userId)
	if err != nil && err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return true, nil
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

func (repo *WordRepository) DeleteWord(word string, userId int64) (*dto.ResponseFromDeleting, error) {
	var wd model.Word

	if err := repo.db.QueryRowx(selectWordQuery, word, userId).Scan(&wd.Word, &wd.Translate, &wd.UserId); err != nil {
		return nil, err
	}

	if _, err := repo.db.Exec(deleteWordQuery, word, userId); err != nil {
		return nil, err
	}

	return &dto.ResponseFromDeleting{Answer: "deleting success"}, nil
}
