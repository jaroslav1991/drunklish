package handlers

import (
	"drunklish/internal/model"
	"drunklish/internal/service/word/dto"
)

type WordService interface {
	CreateWord(word dto.CreateWordRequest) (*model.Word, error)
	GetWordsByUserId(word dto.RequestForGettingWord) (*dto.ResponseWords, error)
	DeleteWordByWord(word dto.RequestForDeletingWord) (*dto.ResponseFromDeleting, error)
}
