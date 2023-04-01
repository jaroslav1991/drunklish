package word

import (
	"drunklish/internal/service/word/repository"
)

type Word struct {
	repo *repository.WordRepository
}

func NewWordService(repo *repository.WordRepository) *Word {
	return &Word{repo: repo}
}
