//go:generate mockgen -package=$GOPACKAGE -source=$GOFILE -destination=interfaces_mock.go

package word

import (
	"drunklish/internal/service/auth/token"
	"drunklish/internal/service/word/dto"
	"time"
)

type Word struct {
	repo         Repository
	parseTokenFn func(tokenString string) (*token.AuthClaims, error)
}

func NewWordService(repo Repository) *Word {
	return &Word{repo: repo, parseTokenFn: token.ParseToken}
}

type Repository interface {
	Create(word string, translate string, createdAt time.Time, userId int64) (*dto.ResponseFromCreateWord, error)
	GetWords(userId int64) (*dto.ResponseWords, error)
	GetWordByCreated(userId int64, firstDate, secondDate time.Time) (*dto.ResponseWords, error)
	DeleteWord(userId, id int64) (*dto.ResponseFromDeleting, error)
	Update(word, translate string, id, userId int64) (*dto.ResponseWord, error)
	CreateTraining(words, answers dto.ResponseWordList, wordsTotal, userId int64) (*dto.ResponseForTraining, error)
	GetTrainingInfoById(trainingId, userId int64) (*dto.ResponseTrainingInfo, error)
	GetStatisticByTrainingId(trainingId, userId int64) (*dto.ResponseStatistic, error)
	CreateStatisticByTrainingId(trainingId int64, correctAnswers, wrongAnswers dto.ResponseWordList, userId int64) (*dto.ResponseCreateStatistic, error)
}
