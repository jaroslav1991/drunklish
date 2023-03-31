package word

import (
	"drunklish/internal/service"
	"drunklish/internal/service/word/repository"
)

type Word struct {
	db   service.DB
	repo *repository.WordRepository
}

func NewWordService(db service.DB, repo *repository.WordRepository) *Word {
	return &Word{db: db, repo: repo}
}
