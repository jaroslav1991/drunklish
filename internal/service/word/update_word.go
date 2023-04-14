package word

import (
	"drunklish/internal/pkg/httputils"
	"drunklish/internal/service/word/dto"
	"fmt"
)

func (w *Word) UpdateWord(word dto.RequestForUpdateWord) (*dto.ResponseWord, error) {
	token, err := w.parseTokenFn(word.Token)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", httputils.ErrValidation)
	}

	updatedWord, err := w.repo.Update(word.Word, word.Translate, word.Id, token.UserId)
	if err != nil {
		return nil, err
	}

	return updatedWord, nil
}
