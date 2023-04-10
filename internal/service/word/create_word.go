package word

import (
	"drunklish/internal/pkg/httputils"
	"drunklish/internal/service/word/dto"
	"drunklish/internal/service/word/validator"
	"fmt"
)

func (w *Word) CreateWord(word dto.CreateWordRequest) (*dto.ResponseFromCreateWord, error) {
	if checkLengthWordAndTranslate := validator.CheckLengthWordAndTranslate(word.Word, word.Translate); !checkLengthWordAndTranslate {
		return nil, fmt.Errorf("invalid length fail: %w", httputils.ErrValidation)
	}

	token, err := w.parseTokenFn(word.Token)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", httputils.ErrValidation)
	}

	createdWord, err := w.repo.Create(word.Word, word.Translate, word.CreatedAt, token.UserId)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", httputils.ErrInternalServer, err)
	}
	return createdWord, nil
}
