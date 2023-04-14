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
	CheckUserInDB(userId int64) (bool, error)
	DeleteWord(word string, userId int64) (*dto.ResponseFromDeleting, error)
	Update(word, translate string, id, userId int64) (*dto.ResponseWord, error)
}
