package handlers

import (
	"drunklish/internal/service/word/dto"
)

type WordService interface {
	CreateWord(word dto.CreateWordRequest) (*dto.ResponseFromCreateWor, error)
	GetWordsByUserId(word dto.RequestForGettingWord) (*dto.ResponseWords, error)
	DeleteWordByWord(word dto.RequestForDeletingWord) (*dto.ResponseFromDeleting, error)
}
