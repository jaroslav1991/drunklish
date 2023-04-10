package word

import (
	"drunklish/internal/pkg/httputils"
	"drunklish/internal/service/word/dto"
	"fmt"
)

//var (
//	ErrWord = errors.New("word not exist")
//)

func (w *Word) DeleteWordByWord(word dto.RequestForDeletingWord) (*dto.ResponseFromDeleting, error) {
	token, err := w.parseTokenFn(word.Token)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", httputils.ErrValidation)
	}

	answer, err := w.repo.DeleteWord(word.Word, token.UserId)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", httputils.ErrWordNotExist, err)
	}

	return answer, nil
}
