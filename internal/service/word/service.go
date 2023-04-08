//go:generate mockgen -package=$GOPACKAGE -source=$GOFILE -destination=interfaces_mock.go

package word

import (
	"drunklish/internal/service/word/dto"
)

type Word struct {
	repo Repository
}

func NewWordService(repo Repository) *Word {
	return &Word{repo: repo}
}

type Repository interface {
	Create(word dto.CreateWordRequest) (*dto.ResponseFromCreateWor, error)
	GetWords(word dto.RequestForGettingWord) (*dto.ResponseWords, error)
	GetWordByCreated(period dto.RequestForGetByPeriod) (*dto.ResponseWords, error)
	CheckUserInDB(userId int64) (bool, error)
	DeleteWord(word dto.RequestForDeletingWord) (*dto.ResponseFromDeleting, error)
}
