package word

import (
	"drunklish/internal/model"
	"drunklish/internal/service/word/dto"
)

func (w *Word) CreateWord(word dto.CreateWordRequest) (*model.Word, error) {
	createdWord, err := w.repo.Create(word)
	if err != nil {
		return nil, err
	}
	return createdWord, nil
}
