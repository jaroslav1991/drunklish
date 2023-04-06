//go:generate mockgen -package=$GOPACKAGE -source=$GOFILE -destination=interfaces_mock.go

package word

import (
	"drunklish/internal/model"
	"drunklish/internal/service/word/dto"
	"time"
)

type Word struct {
	repo Repository
}

func NewWordService(repo Repository) *Word {
	return &Word{repo: repo}
}

type Repository interface {
	Create(word dto.CreateWordRequest) (*model.Word, error)
	GetWords(word dto.RequestForGettingWord) ([]*dto.ResponseWord, error)
	GetWordsByCreated(userId int64, createdAt time.Time) (*model.Word, error)
	DeleteWord(word dto.RequestForDeletingWord) error
}
