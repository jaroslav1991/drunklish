package repository

import (
	"drunklish/internal/pkg/db"
	"drunklish/internal/service/word/dto"
	"fmt"
	"time"
)

const (
	createWordQuery      = `insert into words (word, translate, created_at, user_id) values ($1, $2, $3, $4) returning word, translate`
	getWordsQuery        = `select id, word, translate from words where user_id=$1 order by created_at desc`
	deleteWordQuery      = `delete from words where user_id=$1 and id=$2 returning id`
	getWordByPeriodQuery = `select id, word, translate from words where user_id=$1 and created_at>$2 and created_at<$3 order by created_at desc`
	updateWordQuery      = `update words  set word=$1, translate=$2 where id=$3 and user_id=$4 returning word, translate`
	createTrainingQuery  = `insert into training (words, answers, words_total, user_id) values ($1, $2, $3, $4) returning id, words`
	getTrainingInfoQuery = `select words, answers from training where id=$1 and user_id=$2`
	getStatisticQuery    = `select correct_answers, wrong_answers from statistic where training_id=$1 and user_id=$2`
	createStatisticQuery = `insert into statistic (training_id, correct_answers, wrong_answers, user_id) values ($1, $2, $3, $4) returning id`
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

func (repo *WordRepository) CreateTraining(words, answers dto.ResponseWordList, wordsTotal, userId int64) (*dto.ResponseForTraining, error) {
	var responseTraining dto.ResponseForTraining

	rows, err := repo.db.Query(createTrainingQuery, words, answers, wordsTotal, userId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&responseTraining.Id, &responseTraining.Words); err != nil {
			return nil, err
		}
	}

	return &responseTraining, nil
}

func (repo *WordRepository) GetTrainingInfoById(trainingId, userId int64) (*dto.ResponseTrainingInfo, error) {
	var trainingInfo dto.ResponseTrainingInfo

	rows, err := repo.db.Query(getTrainingInfoQuery, trainingId, userId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&trainingInfo.Words, &trainingInfo.Answers); err != nil {
			return nil, err
		}
	}

	return &trainingInfo, nil
}

func (repo *WordRepository) GetStatisticByTrainingId(trainingId, userId int64) (*dto.ResponseStatistic, error) {
	var responseStatistic dto.ResponseStatistic

	rows, err := repo.db.Query(getStatisticQuery, trainingId, userId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&responseStatistic.CorrectAnswers.Words, &responseStatistic.WrongAnswers.Words); err != nil {
			return nil, err
		}
	}

	return &responseStatistic, nil
}

func (repo *WordRepository) CreateStatisticByTrainingId(
	trainingId int64,
	correctAnswers,
	wrongAnswers dto.ResponseWordList,
	userId int64,
) (*dto.ResponseCreateStatistic, error) {
	var statistic dto.ResponseCreateStatistic

	if err := repo.db.QueryRowx(createStatisticQuery, trainingId, correctAnswers, wrongAnswers, userId).Scan(&statistic.Id); err != nil {
		return nil, err
	}

	return &statistic, nil
}

func (repo *WordRepository) CreateFromUpload(words, translates string, createdAt time.Time, userId int64) (*dto.ResponseFromCreateWord, error) {
	var response dto.ResponseFromCreateWord
	createdAt = time.Now()

	if err := repo.db.QueryRowx(createWordQuery, words, translates, createdAt, userId).Scan(&response.Word, &response.Translate); err != nil {
		return nil, err
	}

	return &response, nil
}
