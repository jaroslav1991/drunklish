package word

import "drunklish/internal/service"

type Word struct {
	db service.DB
}

func NewWordService(db service.DB) *Word {
	return &Word{db: db}
}
