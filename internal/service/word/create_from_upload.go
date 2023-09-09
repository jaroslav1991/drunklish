package word

import (
	"drunklish/internal/pkg/httputils"
	"drunklish/internal/service/word/dto"
	"fmt"
	"time"
)

func (w *Word) CreateListFromUpload(upload dto.RequestFromUpload, tkn string) (*[]dto.ResponseFromCreateWord, error) {
	token, err := w.parseTokenFn(tkn)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", httputils.ErrValidation)
	}
	createdAt := time.Now()

	var words []dto.ResponseFromCreateWord

	for i := 0; i < len(upload.Words.Words); i++ {
		word, err := w.repo.CreateFromUpload(upload.Words.Words[i].Word, upload.Words.Words[i].Translate, createdAt, token.UserId)
		if err != nil {
			return nil, err
		}

		words = append(words, *word)
	}

	return &words, nil
}
