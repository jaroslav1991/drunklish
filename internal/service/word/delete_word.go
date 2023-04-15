package word

import (
	"drunklish/internal/pkg/httputils"
	"drunklish/internal/service/word/dto"
	"fmt"
)

func (w *Word) DeleteWordByWord(word dto.RequestForDeletingWord) (*dto.ResponseFromDeleting, error) {
	token, err := w.parseTokenFn(word.Token)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", httputils.ErrValidation)
	}

	answer, err := w.repo.DeleteWord(token.UserId, word.Id)
	if err != nil {
		return nil, fmt.Errorf("%w", httputils.ErrWordNotExist)
	}

	return answer, nil
}
