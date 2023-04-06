package word

import (
	"drunklish/internal/pkg/httputils"
	"drunklish/internal/service/word/dto"
	"drunklish/internal/service/word/validator"
	"fmt"
)

func (w *Word) CreateWord(word dto.CreateWordRequest) (*dto.ResponseFromCreateWor, error) {
	if checkLengthWordAndTranslate := validator.CheckLengthWordAndTranslate(word.Word, word.Translate); !checkLengthWordAndTranslate {
		return nil, fmt.Errorf("invalid length fail: %w", httputils.ErrValidation)
	}

	createdWord, err := w.repo.Create(word)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", httputils.ErrInternalServer, err)
	}
	return createdWord, nil
}
