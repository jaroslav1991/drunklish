package handlers

import (
	"drunklish/internal/service/word/dto"
)

type WordService interface {
	CreateWord(word dto.CreateWordRequest) (*dto.ResponseFromCreateWord, error)
	GetWordsByUserId(word dto.RequestForGettingWord) (*dto.ResponseWords, error)
	GetWordsByCreatedAt(period dto.RequestForGetByPeriod) (*dto.ResponseWords, error)
	DeleteWordByWord(word dto.RequestForDeletingWord) (*dto.ResponseFromDeleting, error)
	UpdateWord(word dto.RequestForUpdateWord) (*dto.ResponseWord, error)
}
